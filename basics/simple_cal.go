package main

import "fmt"

func main() {
    var a, b float64

    fmt.Print("Enter first number: ")
    fmt.Scan(&a)

    fmt.Print("Enter second number: ")
    fmt.Scan(&b)

    fmt.Println("Sum:", a+b)
    fmt.Println("Difference:", a-b)
    fmt.Println("Product:", a*b)

    if b != 0 {
        fmt.Println("Quotient:", a/b)
    } else {
        fmt.Println("Cannot divide by zero")
    }
}

