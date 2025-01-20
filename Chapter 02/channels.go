package main

import (
    "fmt"
)

func main() {
    // Create a channel for integers
    ch := make(chan int)

    // Launch a goroutine that sends a value to the channel
    go func() {
        ch <- 44 // Send value to the channel
    }()
    
    // Receive the value from the channel
    value := <-ch
    fmt.Println("Received value:", value)
}
