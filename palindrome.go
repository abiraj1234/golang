package main 



import "fmt"

func main() {
	
	var pan, temp, remain int
	
	fmt.Print("Enter the Number: ")
	fmt.Scan(&pan)
	
	temp = pan
	
	res:= 0
	
	for{
	    remain = temp % 10
		res = (res*10) + remain
		temp= temp /10
		
		if (temp==0){
			break
		}
    }	
	
		
	if res == pan{
		fmt.Print("The given Number is palindrome")
	}else{
		fmt.Print("The given Number is Not palindrome")
	
    }	
}		
	