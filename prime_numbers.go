package main

import "fmt"

func main() {
    var n int
    fmt.Print("Enter a number: ")
    fmt.Scan(&n)
    count := 0
    for i := 2; i < n; i++ {
        if n % i == 0 {
            count++
        }
    }
    fmt.Print(n)
    if count == 0 {
        fmt.Println(" is a prime number!")
    } else {
        fmt.Println(" is not a prime number!")
    }
}
