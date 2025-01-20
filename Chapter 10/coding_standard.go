package main

import (
    "fmt"
    "strconv"
)
	
func createIdentity(id int) (string, error) {
    if id <= 0 {
        return "", fmt.Errorf("invalid input value for id: %d", id)
    }
    return "Name " + strconv.Itoa(id), nil
}
	
func main() {
    user, err := createIdentity(-1)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("User:", user)
}
