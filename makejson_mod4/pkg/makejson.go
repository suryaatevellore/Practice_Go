package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	// input
	var m map[string]string
	m = make(map[string]string)
	reader := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter a name: ")
	reader.Scan()
	m["name"] = reader.Text()
	fmt.Print("Enter an address: ")
	reader.Scan()
	m["address"] = reader.Text()
	map_json, _ := json.Marshal(m)
	fmt.Println("JSON-ified map looks like this:")
	fmt.Print(string(map_json))
}
