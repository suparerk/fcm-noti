package main

import (
	"context"
	"fcm-noti/config"
	"fmt"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
)

func main() {

	app, _, _ := config.SetupFirebase()
	sendToToken(app)
}

func sendToToken(app *firebase.App) {
	ctx := context.Background()
	client, err := app.Messaging(ctx)
	if err != nil {
		log.Fatalf("error getting Messaging client: %v\n", err)
	}

	registrationToken := "dGccz7pp3t0A0Ko3ew8vzR:APA91bE9obLBwtDSJDUCE9p4tL0NKCoBRvvJ0AwCBmM_0ibq3YcAuw6xH2191ri7lW1UbfwU4CndcQhpY7vyAyLYJxhIEl3ZDs0tEc6orsHE2UsQoP8yKhr4XW8wwwcV3h2jApXnhbdn"

	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title: "Hi there BBB",
			Body:  "Hello from VB!!",
		},
		Token: registrationToken,
	}

	response, err := client.Send(ctx, message)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Successfully sent message:", response)
}
