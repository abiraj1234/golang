package main


import "fmt"

func main () {

	var minimum, maximum int 
	
	fmt.Print("Enter the minimum: ")
	fmt.Scan(&minimum)
	
	fmt.Print("Enter the maximum: ")
	fmt.Scan(&maximum)
	
	
	print("the prime number is " )
	for i:= minimum; i<=maximum; i++ {
	    p:=0
		//print(minimum)
		for j:=2;j<minimum;j++{
		  if i%j==0{
		     //fmt.Print(j)
		      p+=1
			  break
		   }
		
		
		}
		if p==0{
		  fmt.Println(i)
		
		}
		}
		
		
		
		
}