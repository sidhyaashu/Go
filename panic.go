package main

import "fmt"

// Function that causes a panic
func causePanic() {
    panic("Something went wrong!")
}

func main() {
    fmt.Println("Program started")

    // Calling the function that causes panic
    causePanic()

    fmt.Println("This will not be executed")
}
