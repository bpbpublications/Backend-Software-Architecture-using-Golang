package main

import (
    "fmt"
    "sync"
)

var count int
var mu sync.Mutex

func increment() {
    mu.Lock()
    defer mu.Unlock()
    count++
}

func main() {
    var wg sync.WaitGroup
    wg.Add(3)

    // Start three goroutines
    go func() {
        increment()
        wg.Done()
    }()
    go func() {
        increment()
        wg.Done()
    }()
    go func() {
        increment()
        wg.Done()
    }()

    wg.Wait()
    fmt.Println("Final count:", count)
}
