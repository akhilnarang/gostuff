package main

import (
    "fmt"
    "sort"
    "strconv"
)

func main() {
    var slc []int = make([]int, 3)
    var in string
    fmt.Println("Please enter an integer - you can enter X to exit")
    for true {
        fmt.Print("in: ")
        fmt.Scanln(&in) 
        if in == "X" {
            fmt.Println("X input detected - exiting")
            break
        }
        s, _ := strconv.Atoi(in)
        slc = append(slc, s)
        sort.Ints(slc[:])
        fmt.Println(slc)
	}
}
