package main

import "fmt"

func main() {
	messagesFromDoris := []string{
		"you are doing good",
		"i am learning go",
		"i am working on go",
		"i am starting my journey to go",
	}
	numMessages := float64(len(messagesFromDoris))
	costPerMessage := .02

	totalCost := costPerMessage * numMessages

	fmt.Printf("Doris spent %.2f on text messgae today", totalCost)
}
