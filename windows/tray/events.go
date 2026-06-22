package main

import (
	"log"

	"github.com/fehmicorp/agent/windows/services/notification"
)

func onExit() {
	log.Println("tray closed")
}

func PushNotification(Id int) {
	switch Id {
	case 2:
		log.Println("Inventory scan triggered...")
		err := notification.Push(notification.Options{
			AppId:    "Fehmi Endpoint Agent",
			Title:    "Inventory Scan started",
			Message:  "Fehmi agent is compiling system specs in the background.",
			IconPath: "../config/assets/fav.ico",
		})
		if err != nil {
			log.Printf("Notification error: %v", err)
		}
	}
}
