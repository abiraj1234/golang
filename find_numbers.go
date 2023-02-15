package main 

import "fmt"


func findNumbers(nums []int) int {
    count:=0;
    for _,i := range nums{
        j:=0
        
            for i>0{
                i=i/10

                j++
            }
            if j%2==0{
                count++
            }
        
    }
    return count
    
}

func main() {

	nums := [] int {12,654,7837,98,87,9,}
	fmt.Print(findNumbers(nums))
}	
