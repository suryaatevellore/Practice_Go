package main

import "fmt"
import "sort"
import "strconv"


func main() {
	var s []int	
	s=make([]int,3)
	i:=0
	num:=""
	for {
		fmt.Print("\nEnter a number(pres x/X to stop) :")
		fmt.Scan(&num)
				
		if num=="X"||num=="x" {
			break 
		}
		
		n,_ :=strconv.Atoi(num)
		
		//adding
		if i>2 {
			s=append(s,n)
		} else {
			s[i]=n
		}
		
		//sorting
		sort.Ints(s[0:i+1])
		
		//Printing
		fmt.Print("Sorted slice is:-")
		for j:=0;j<=i;j++ {
			fmt.Printf("\n%d",s[j])
		}
		i=i+1
	}
}
				
		
		
