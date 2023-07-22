package main

import "fmt"

func main() {
	messageLen := 10
	maxMessageLen := 20
	fmt.Println("Trying to send message of lenght:", messageLen)

	if messageLen <= maxMessageLen {
		fmt.Println("Messgae sent")
	} else {
		fmt.Println("Message not sent")
	}
}
