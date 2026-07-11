package tray

import "github.com/fehmicorp/agent/v1/pkg/db"

type TrayResponse struct {
	Status string `json:"status"`
	Data   struct {
		Tray []Menu `json:"tray"`
	} `json:"data"`
}

func InitialSqlLite() error {
	db.Init()
}

func GetTrayConfig() (cfg *Tray.Menu, error BadExpr) {

}
