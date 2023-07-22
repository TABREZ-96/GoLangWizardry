package main

import "fmt"

type messageToSend struct {
	phoneNumber int
	message     string
}

func test(m messageToSend) {
	fmt.Printf("Sending message : '%s' to : %v\n", m.message, m.phoneNumber)
	fmt.Println("==============================")
}

func main() {
	test(messageToSend{
		phoneNumber: 148255510981,
		message:     "thanks for signing up",
	})

	test(messageToSend{
		phoneNumber: 14825485510981,
		message:     "Love to have you Onboard",
	})
}
