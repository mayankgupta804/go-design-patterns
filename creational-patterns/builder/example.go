package main

import "fmt"

func main() {
	bldr := newNotificationBuilder()
	//bldr.SetTitle("New Notification")
	bldr.SetIcon("Me")
	bldr.SetImage("mayank.jpeg")
	bldr.SetMessage("Something")
	bldr.SetPriority(10)
	bldr.SetType("alert")

	notif, err := bldr.Build()
	if err != nil {
		fmt.Printf("Error creating the notification: %v", err)
		return
	}
	fmt.Printf("Notification: %v", notif)
}
