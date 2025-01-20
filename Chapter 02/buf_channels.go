package main

import (
    "fmt"
)

func main() {
    // Create a buffered channel with capacity 2
    ch := make(chan int, 2)

    // Send values to the channel
    ch <- 10
    ch <- 20
    
    // Receive values from the channel
    value1 := <-ch
    value2 := <-ch
    
    fmt.Println("Received values:", value1, value2)
}
