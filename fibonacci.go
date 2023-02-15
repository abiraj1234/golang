package main 

import "fmt"

func main () {
	var m , m1 ,m2, mterm int
	
	m1 = 0
	m2 = 1
	
	fmt.Print("Enter the number of elements :  ")
	fmt.Scan(&m)
	
	
	for i:=1; i<=m; i++{
		if (i==1) {
			fmt.Print(" ",m1)
			continue 
		}
		if (i==2) {
			fmt.Print(" ",m2)
			continue
		}
		mterm = m1 + m2 
		m1 = m2
		m2 =mterm
		fmt.Print(" ",mterm)
	}
}
	

	
