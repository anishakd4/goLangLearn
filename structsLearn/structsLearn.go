package main

import "fmt"

type messageToSend struct {
	phoneNumber string
	message string
}

func getMessageText(m messageToSend) {
	x := fmt.Sprintf("Sending message: '%s' to: %v\n", m.message, m.phoneNumber)
	fmt.Printf(x)
}

func main() {
	getMessageText(messageToSend{
		phoneNumber: "sdfsfsdfdsfs",
		message: "message",
	})
}