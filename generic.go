package main

import "fmt"

// Generic function that swaps two values of any type
func swap[T any](a, b T) (T, T) {
    return b, a
}

func main() {
    // Swapping two integers
    x, y := 1, 2
    x, y = swap(x, y)
    fmt.Println("Swapped integers:", x, y)

    // Swapping two strings
    str1, str2 := "Hello", "World"
    str1, str2 = swap(str1, str2)
    fmt.Println("Swapped strings:", str1, str2)
}
