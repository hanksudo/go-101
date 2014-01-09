package main

import "fmt"

// Share type(x int, y int) or (x, y int)
func add(x int, y int) int {
    return x + y
}

// multi results
func swap(x, y string) (string, string) {
    return y, x
}

// Named results - will return x, y
func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return
}


func main() {
    fmt.Println(add(42, 13))

    a, b := swap("hello", "world")
    fmt.Println(a, b)

    fmt.Println(split(17))
}

