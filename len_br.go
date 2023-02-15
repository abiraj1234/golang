package main 

import "fmt"


func main () {

	var len, br float64
	
	var area float64 
	
	fmt.Print("Enter the length of rectangle: ")
	fmt.Scanln(&len)
	
	fmt.Print("Enter the breadth of rectangle: ")
	fmt.Scanln(&br)
	
	area = len * br 
	fmt.Println(float64(area))

	
	
	
}
	