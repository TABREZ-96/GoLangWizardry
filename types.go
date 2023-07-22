package main

import "fmt"

func main() {
	//intilializing variables
	var smsSendingLimit int
	var costPerSMS float64
	var hasPermision bool
	var username string

	fmt.Printf(
		"%v %f %v %q\n",
		smsSendingLimit,
		costPerSMS,
		hasPermision,
		username,
	)

}
