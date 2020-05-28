package main
import (
	"fmt"
)

func main(){
	var num float32
	fmt.Printf("Enter a number to truncate: ")
	fmt.Scan(&num)
	fmt.Println("Truncated number is",int32(num))
}
