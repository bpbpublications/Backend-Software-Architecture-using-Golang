package main

import (
    "fmt"
    "sync"
)

func main() {
    var wg sync.WaitGroup

    wg.Add(2) // Add two goroutines to the wait group

    go func() {
        defer wg.Done()
        fmt.Println("Goroutine 1")
    }()

    go func() {
        defer wg.Done()
        fmt.Println("Goroutine 2")
    }()

    wg.Wait() // Wait for all goroutines to finish
    fmt.Println("All goroutines finished")
}
