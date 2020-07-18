package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Name struct {
	first string 
	last string 
}

func main() {
	var path string
	fmt.Println("Enter the path to the file!")
	fmt.Print("path: ")
	fmt.Scanln(&path)
	var name[] Name
	f, _ := os.Open(path)

    rd := bufio.NewReader(f)
    for {
        line, _, err:= rd.ReadLine()
        
        if err != nil {
			if err != io.EOF {
				fmt.Println("Error occurred: ", err)
			}
            break
        }       
            tname := strings.Split(string(line)," ")
            tpname:= Name {
                tname[0],
                tname[1],
            } 
            name=append(name,tpname)
	}

	f.Close()
    for i := 0; i < len(name); i++ {
        fmt.Print(i+1, ". ")
        fmt.Print("First Name: ", name[i].first, " ")
		fmt.Print("Last Name: ", name[i].last, "\n")
    }
}	