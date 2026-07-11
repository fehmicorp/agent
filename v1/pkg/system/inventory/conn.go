package inventory

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/fehmicorp/agent/v1/pkg/db/sqlite"
)

var store *sqlite.SQLiteStore

type QueueItem struct {
	ID          int64          `json:"id"`
	DeviceID    string         `json:"deviceId"`
	Fingerprint string         `json:"fingerprint"`
	CollectedAt time.Time      `json:"collectedAt"`
	Status      string         `json:"status"`
	Inventory   *InventoryInfo `json:"inventory"`
}

func InitialSQLite(dir string) error {

	store = sqlite.NewSQLiteStore("agent.db", "inventory")

	_, err := store.Init(dir, []sqlite.TblQuery{
		{
			Key:        "id",
			Type:       "INTEGER",
			Preference: "PRIMARY KEY AUTOINCREMENT",
		},
		{
			Key:  "device_id",
			Type: "TEXT",
		},
		{
			Key:  "fingerprint",
			Type: "TEXT",
		},
		{
			Key:  "collected_at",
			Type: "DATETIME",
		},
		{
			Key:  "status",
			Type: "TEXT",
		},
		{
			Key:  "payload",
			Type: "TEXT",
		},
	})

	return err
}

func SaveScan(inv *InventoryInfo) error {

	if store == nil {
		return fmt.Errorf("inventory sqlite not initialized")
	}

	data, err := json.Marshal(inv)
	if err != nil {
		return err
	}

	return store.InsertDynamic(map[string]any{
		"device_id":    inv.System.DeviceID,
		"fingerprint":  inv.System.Fingerprint,
		"collected_at": time.Now().UTC(),
		"status":       "pending",
		"payload":      string(data),
	})
}

func GetPending(limit int) ([]QueueItem, error) {

	if store == nil {
		return nil, fmt.Errorf("inventory sqlite not initialized")
	}

	if limit <= 0 {
		limit = 10
	}
	rows, err := store.Query(`
		SELECT
			id,
			device_id,
			fingerprint,
			collected_at,
			status,
			payload
		FROM inventory
		WHERE status = ?
		ORDER BY id ASC
		LIMIT ?;
	`, "pending", limit)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var result []QueueItem

	for rows.Next() {

		var item QueueItem

		var payload string

		err := rows.Scan(
			&item.ID,
			&item.DeviceID,
			&item.Fingerprint,
			&item.CollectedAt,
			&item.Status,
			&payload,
		)

		if err != nil {
			return nil, err
		}

		var inv InventoryInfo

		if err := json.Unmarshal([]byte(payload), &inv); err != nil {
			continue
		}

		item.Inventory = &inv

		result = append(result, item)
	}

	return result, rows.Err()
}

func MarkSynced(ids ...int64) error {

	if len(ids) == 0 {
		return nil
	}

	for _, id := range ids {

		if err := store.UpdateDynamic(
			map[string]any{
				"status": "synced",
			},
			"id=?",
			id,
		); err != nil {
			return err
		}
	}

	return nil
}

func DeleteSynced() error {

	return store.DeleteDynamic(
		"status=?",
		"synced",
	)
}
