package main

import "fmt"

func main() {
	congrats := "happy birthday"
	penniesPerText := 2.0
	averageOpenRate, displayMessage := .23, "is the avetage"
	accountAge := 2.6
	accountAgeInt := int(accountAge)
	fmt.Printf(congrats)
	fmt.Println("The type of PenniesPerText is %T\n", penniesPerText)
	fmt.Println(averageOpenRate, displayMessage)
	fmt.Println("Your account has existed for", accountAgeInt, "years")
}
