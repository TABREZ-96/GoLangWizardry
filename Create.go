package main

import "fmt"

func main() {
	const name = "Tabrez Sayed"
	const openRate = 30.5

	msg := fmt.Sprintf("hi %s,your open rate is %.1f percent", name, openRate)

	fmt.Print(msg)

}
