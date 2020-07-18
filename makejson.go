package main

import (
	"fmt"
	"encoding/json"
)

func main() {
	fmt.Println("Enter details!")
	data := make(map[string] string)
	var name, address string
	fmt.Print("name: ")
	fmt.Scanln(&name)
	fmt.Print("address: ")
	fmt.Scanln(&address)
	data["name"] = name
	data["address"] = address
	fmt.Println("Map data is: ", data)
	json_data, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error occurred trying to marshal the map into JSON")
	} else {
		fmt.Println("JSON encoded data is: ", json_data)
		fmt.Println("JSON as string is: ", string(json_data))
	}
}