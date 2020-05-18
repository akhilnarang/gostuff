package main

import "fmt"
import "time"

func main() {
    start := time.Now()
    fmt.Println("Today is", time.Now())
    fmt.Println("The above line took", time.Now().Sub(start), "to execute")
    fmt.Println(time.Since(start))
}
