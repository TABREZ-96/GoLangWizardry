//Go has built-in support for concurrency through goroutines and channels. Goroutines are lightweight threads that allow you to execute functions concurrently. Channels facilitate communication and synchronization between goroutines.
package main

import (
    "fmt"
    "time"
)

func printNumbers() {
    for i := 1; i <= 5; i++ {
        fmt.Println(i)
        time.Sleep(500 * time.Millisecond)
    }
}

func printLetters() {
    for char := 'a'; char <= 'e'; char++ {
        fmt.Println(string(char))
        time.Sleep(500 * time.Millisecond)
    }
}

func main() {
    // Execute functions concurrently using goroutines
    go printNumbers()
    go printLetters()

    // Sleep for some time to allow goroutines to execute
    time.Sleep(3 * time.Second)
}
