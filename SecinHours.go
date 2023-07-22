package main

import "fmt"

func main() {
	const secondsInMinute = 60
	const minutesInHour = 60
	const secInHour = minutesInHour * secondsInMinute

	fmt.Println("number of sec in an hour", secInHour)
}
