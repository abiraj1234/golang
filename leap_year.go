package main 


import "fmt"


func main () {
	
	var year int 
	
	fmt.Print("Enter the year to check for leap = ")
	fmt.Scan(&year)
	
	if year%400 == 0 || (year%4 == 0 && year%100 !=0){	
		fmt.Println(year, "its  a Leap year")
	}else{
		fmt.Println(year, "its non Leap year")
    }
}
	