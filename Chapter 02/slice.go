package main

import (
    "fmt"
)

func main() {
    // Create a slice of integers
    slice := []int{10, 20, 30, 40}
    
    // Append an element to the slice
    slice = append(slice, 50)
    
    fmt.Println("Slice:", slice)
    
    // Create a sub-slice from the slice
    subSlice := slice[1:3]
    fmt.Println("Sub-slice:", subSlice)
}
