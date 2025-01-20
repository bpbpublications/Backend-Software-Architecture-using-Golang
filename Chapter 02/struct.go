package main

import (
    "fmt"
)

// Define a struct
type Person struct {
    Name string
    Age  int
}

func main() {
    // Initialize a struct
    person := Person{
        Name: "Alice",
        Age:  30,
    }
    
    // Access struct fields
    fmt.Println("Name:", person.Name)
    fmt.Println("Age:", person.Age)
    
    // Modify struct fields
    person.Age = 31
    fmt.Println("Updated age:", person.Age)
}
