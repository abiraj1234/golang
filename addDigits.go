package main 

import "fmt"

func addDigits(num int) int {
	if num<10{
		return num
	}
	res := 0
	for num > 0 {
		rem:= num % 10 
		res = res + rem 
		num = num/10 
	
	}
	return addDigits(res)
}

func main () {
		
		num := 3856
		fmt.Print(addDigits(num))
}	