package main

import (
    "fmt"
    "sync"
)

var data int
var rwMutex sync.RWMutex

func readData() {
    rwMutex.RLock()
    defer rwMutex.RUnlock()
    fmt.Println("Read data:", data)
}

func writeData(newData int) {
    rwMutex.Lock()
    defer rwMutex.Unlock()
    data = newData
    fmt.Println("Wrote data:", data)
}

func main() {
    var wg sync.WaitGroup
    wg.Add(2)

    go func() {
        writeData(42)
        wg.Done()
    }()

    go func() {
        readData()
        wg.Done()
    }()

    wg.Wait()
}
