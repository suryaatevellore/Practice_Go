package main

import (
	"bufio"
	"fmt"
	// "io/ioutil"
	"os"
	"strings"
)

const maxLength = 20

type Name struct {
	fname string
	lname string
}

func trimLength(s string) string {
	// Ensure only the first 20 characters
	// will be considered
	r := []rune(s)
	return string(r[:maxLength])
}

func main() {
	var filename string
	struct_slice := make([]Name, 0)
	Reader := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter the filename: ")
	Reader.Scan()
	filename = Reader.Text()

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		read_name := scanner.Text()
		name_split := strings.Split(read_name, " ")
		fname, lname := trimLength(name_split[0]), trimLength(name_split[1])
		struct_slice = append(struct_slice, Name{fname, lname})
	}

	// Slice is created
	for i := 0; i < len(struct_slice); i += 1 {
		fmt.Println("First Name:", struct_slice[i].fname, ", Last Name:", struct_slice[i].lname)
	}

}
