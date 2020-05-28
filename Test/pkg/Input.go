package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewScanner(os.Stdin)
	exit_char := "X"
	for reader.Scan() {
		input := reader.Text()
		if input == exit_char {
			break
		}
		fmt.Println(reader.Text())
	}
}
