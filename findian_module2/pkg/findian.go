package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// var example string
	fmt.Printf("Enter a string (no case restrictions): ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		line := scanner.Text()
		example_nocase := strings.ToLower(line)
		if strings.HasPrefix(example_nocase, "i") && strings.HasSuffix(example_nocase, "n") && strings.Contains(example_nocase, "a") {
			fmt.Println("Found!")
		} else {
			fmt.Println("Not Found!")
		}
	}

}
