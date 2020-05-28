package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func getPosition(array []int) int {
	var pos int
	for i := 0; i < 3; i += 1 {
		// Return the first vacant position
		if array[i] == 0 {
			pos = i
			break
		}
	}
	return pos
}
func main() {
	var times_entered int
	numbers := make([]int, 3)
	reader := bufio.NewScanner(os.Stdin)
	exit_char := "X"
	for {
		fmt.Print("Enter a number (x/X to Stop): ")
		reader.Scan()
		input := reader.Text()
		if input == exit_char {
			break
		}
		number, _ := strconv.Atoi(input)
		if times_entered == 3 {
			numbers = append(numbers, number)
		} else {
			pos := getPosition(numbers)
			numbers[pos] = number
			times_entered += 1
		}
		sort.SliceStable(numbers, func(i, j int) bool { return numbers[i] < numbers[j] })
		fmt.Println("Sorted slice is :", numbers)
	}

}
