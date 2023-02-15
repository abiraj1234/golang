package main 


import "fmt"

func main () {
	
	var rev, remain int 
	
	fmt.Print("Enter the Number of reverse = ")
	fmt.Scan(&rev)
	
	reverse := 0 
	
	for rev > 0{
		remain = rev % 10
		reverse = reverse*10 + remain 
		rev = rev/10
	}
		  
    fmt.Println("The Reverse of the Given Number =", reverse)		  
}	