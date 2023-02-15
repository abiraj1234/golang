package main 


import "fmt"

func main (){
	
	fmt.Print("Enter the Elements: ")
	var nums int
	fmt.Scanln(&nums)
	
	var user_input [70]int 
	val := 0
	var average int 
	for i:=0; i<nums; i++ {
		fmt.Print("Enter the numbers: ")
		fmt.Scanln(&user_input[i])
		val = val + user_input[i]
	}
	average = val/nums
	fmt.Println("average of numbers ", average)
}		