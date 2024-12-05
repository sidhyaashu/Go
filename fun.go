package main

import "fmt"

// Function with no parameters and no return value
func greet() {
    fmt.Println("Hello, Go!")
}

// Function with parameters
func add(a int, b int) int {
    return a + b
}

// Function with multiple return values
func divide(a int, b int) (int, int) {
    quotient := a / b
    remainder := a % b
    return quotient, remainder
}

func main() {
    // Call a function with no parameters
    greet()

    // Call a function with parameters
    sum := add(5, 10)
    fmt.Println("Sum:", sum)

    // Call a function with multiple return values
    quotient, remainder := divide(10, 3)
    fmt.Printf("Quotient: %d, Remainder: %d\n", quotient, remainder)
}
