package main

import "fmt"

func main() {
    // Declare a variable
    var num int = 42

    // Declare a pointer to num
    var ptr *int = &num

    fmt.Println("Value of num:", num)             // Prints the value of num
    fmt.Println("Address of num:", &num)         // Prints the memory address of num
    fmt.Println("Value stored in ptr:", ptr)     // Prints the memory address of num (pointer value)
    fmt.Println("Value pointed to by ptr:", *ptr) // Dereferences the pointer to get the value of num

    // Modify num through the pointer
    *ptr = 100
    fmt.Println("\nAfter modifying through pointer:")
    fmt.Println("Value of num:", num)             // The value of num is updated
    fmt.Println("Value pointed to by ptr:", *ptr) // The dereferenced pointer also shows the updated value
}
