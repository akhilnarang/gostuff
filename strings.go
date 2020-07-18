package main

import "fmt"
import "strings"

func main() {
    var s string
    fmt.Print("Please enter a string:\n")
    fmt.Scanln(&s)
    l := len(s) - 1
    if strings.Contains(s, "a") && strings.Contains("Ii", string(s[0])) && strings.Contains("Nn", string(s[l])) {
        fmt.Print("Found!")
    } else  {
        fmt.Print("Not Found!")
    }
}
