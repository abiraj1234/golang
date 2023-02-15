package main 

import "fmt"

func main () {

    fmt.Print("Enter the Number: ")
	 
	var x int 
	fmt.Scanln(&x)
	if x%2==0 {
		fmt.Println(x, "is the Even number")
	}else{
		fmt.Println(x, "is the odd number")
	}
}	