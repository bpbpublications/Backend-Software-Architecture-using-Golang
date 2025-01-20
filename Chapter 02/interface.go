package main

import (
    "fmt"
)


type Shape interface {
    Area() float64
    Perimeter() float64
}

type Circle struct {
    radius float64
}

// Circle implements the Shape interface
func (c Circle) Area() float64 {
    return 3.14 * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
    return 2 * 3.14 * c.radius
}


func printShapeDetails(s Shape) {
    fmt.Printf("Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
}

func printType(value interface{}) {
    switch v := value.(type) {
    case int:
        fmt.Println("This is an integer:", v)
    case string:
        fmt.Println("This is a string:", v)
    default:
        fmt.Println("Unknown type")
    }
}

func main() {
    c := Circle{radius: 5}

    // The function works with any type that implements the Shape interface
    printShapeDetails(c)

    printType(42)          // Output: This is an integer: 42
    printType("Hello")     // Output: This is a string: Hello
    printType(3.14)        // Output: Unknown type
}
