package main

import (
    "fmt"
    "errors"
)

// Function that returns an error
func divide(a, b int) (int, error) {
    if b == 0 {
        // Return an error if the divisor is zero
        return 0, errors.New("cannot divide by zero")
    }
    return a / b, nil
}

func main() {
    // Example 1: Successful division
    result, err := divide(10, 2)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Result:", result)
    }

    // Example 2: Division by zero
    result, err = divide(10, 0)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Result:", result)
    }
}
