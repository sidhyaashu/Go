package main

import "fmt"

func main() {
    // Array
    arr := [5]int{1, 2, 3, 4, 5}
    fmt.Println("Array:", arr)

    // Slice
    slice := []int{10, 20, 30, 40, 50}
    fmt.Println("Slice:", slice)

    // Map
    myMap := map[string]int{
        "Alice": 25,
        "Bob":   30,
        "Carol": 35,
    }
    fmt.Println("Map:", myMap)

    // for-range loop with Array
    fmt.Println("\nIterating over Array:")
    for index, value := range arr {
        fmt.Printf("Index: %d, Value: %d\n", index, value)
    }

    // for-range loop with Slice
    fmt.Println("\nIterating over Slice:")
    for index, value := range slice {
        fmt.Printf("Index: %d, Value: %d\n", index, value)
    }

    // for-range loop with Map
    fmt.Println("\nIterating over Map:")
    for key, value := range myMap {
        fmt.Printf("Key: %s, Value: %d\n", key, value)
    }
}
