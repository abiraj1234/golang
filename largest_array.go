package main 

import "fmt"

func main () {
	var num [50] int 
	var next int 
	
	fmt.Print("Enter the number of elements: ")
	fmt.Scan(&next)
	
	for i:=0; i<next; i++ {
		fmt.Print("Enter the number: ")
		fmt.Scan(&num[i])
	}
	max := 0
	
	for j:=0; j<next; j++{
		if max < num[j]{
			max = num[j]
		}	
	}

	fmt.Print("largest number of array is : ", max)
}	