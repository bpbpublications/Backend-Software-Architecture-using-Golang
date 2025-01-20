package main

import (
    "fmt"
    "sync"
)

func worker(id int, wg *sync.WaitGroup) {
    defer wg.Done()
    fmt.Printf("Worker %d starting\n", id)
    // Simulate work
    fmt.Printf("Worker %d done\n", id)
}

func main() {
    var wg sync.WaitGroup

    // Launch three goroutines
    for i := 1; i <= 3; i++ {
        wg.Add(1)
        go worker(i, &wg)
    }

    wg.Wait() // Wait for all workers to finish
    fmt.Println("All workers completed")
}
