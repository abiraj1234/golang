package main 

import "fmt"

func twoSum(nums []int, target int) []int {
    for i := 0; i< len(nums); i++ {
        for j := i+1; j<len(nums); j++{
             if nums[i]+nums[j] == target{
                 return []int{nums[i],nums[j]}
        }
    
        }
    }
    return []int{}
}

func main () {
	nums := []int {2,7,11,5}
	fmt.Println(twoSum(nums,12))
	
}
	