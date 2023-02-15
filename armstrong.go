package main 

import "fmt"

func main () {

	var mn, temp , remain, res int 
	fmt.Print("Enter the arm numbers: ")
	fmt.Scanln(&mn)
	
	temp = mn 
	
	for temp>0{
	    remain = temp % 10
	    res += (remain * remain * remain) 
		temp = temp / 10
	
	}
	if res ==mn && mn<=9{
	fmt.Println("Its a armstrong Number")
	
	}else{
	fmt.Println("its not a armstrong Number")

    }	
}
