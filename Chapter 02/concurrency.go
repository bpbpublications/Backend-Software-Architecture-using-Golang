package main

import (
    "fmt"
    "time"
)

func main() {
    // Launch a new goroutine
    go printMessage("Hello from goroutine!")
    
    fmt.Println("Hello from main!")

    // Launch another goroutine
    go printMessage("Hello from second goroutine!") 

    // Wait for the goroutine to finish
    time.Sleep(1 * time.Second)
}

func printMessage(message string) {
    fmt.Println(message)
}
