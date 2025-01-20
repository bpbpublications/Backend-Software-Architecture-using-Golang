package main

import (
    "fmt"
)

func main() {
    // Declare a map with string keys and integer values
    myMap := make(map[string]int)
    
    // Add key-value pairs to the map
    myMap["a"] = 1
    myMap["b"] = 2
    myMap["c"] = 3
    
    fmt.Println("Map:", myMap)
    
    // Access values using keys
    fmt.Println("Value for key 'b':", myMap["b"])
}
