package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func main() {
	var jsonInput []byte
	var err error

	if len(os.Args) > 1 {
		jsonInput = []byte(os.Args[1])
	} else {
		jsonInput, err = io.ReadAll(os.Stdin)
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}
	}

	var data interface{}
	err = json.Unmarshal(jsonInput, &data)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	formattedJSON, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		fmt.Println("Error formatting JSON:", err)
		return
	}

	fmt.Println(string(formattedJSON))
}
