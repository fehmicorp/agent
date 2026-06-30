package defender

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	status, err := GetDefenderStatus()
	if err != nil {
		errPayload := map[string]string{"error": "Failed to collect Windows Defender metrics: " + err.Error()}
		jsonBytes, _ := json.MarshalIndent(errPayload, "", "  ")
		fmt.Println(string(jsonBytes))
		os.Exit(1)
	}
	jsonOutput, err := json.MarshalIndent(status, "", "  ")
	if err != nil {
		fmt.Printf("{\x22error\x22: \x22JSON Generation Error: %v\x22}\n", err)
		os.Exit(1)
	}
	fmt.Println(string(jsonOutput))
}
