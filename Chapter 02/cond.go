package main

import (
    "fmt"
    "sync"
)

var cond = sync.NewCond(&sync.Mutex{})
var counter int
var wg sync.WaitGroup

func increment() {
    defer wg.Done()
    cond.L.Lock()
    defer cond.L.Unlock()
    counter++
    fmt.Println("Incremented counter:", counter)
    cond.Signal() // Signal one waiting goroutine
}

func waitForIncrement() {
    defer wg.Done()
    cond.L.Lock()
    defer cond.L.Unlock()
    for counter == 0 {
        cond.Wait() // Wait for counter to be incremented
    }
    fmt.Println("Counter is now:", counter)
}

func main() {
    wg.Add(1)
    go waitForIncrement()
    wg.Add(1)
    go increment()
    
    // Allow some time for goroutines to run
    wg.Wait()
}
