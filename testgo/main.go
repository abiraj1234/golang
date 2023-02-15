package main

import "fmt"

func main() {
	a := sortArray([]int{-1, 0, -1, -1, 1, 0, 1, 0})
	fmt.Println("arrau", a)
}

func sortArray(nums []int) []int {

	for i := 1; i < len(nums); i++ {
		if i != 0 && nums[i] < nums[i-1] {
			nums[i], nums[i-1] = nums[i-1], nums[i]
			i -= 2
		}
	}

	fmt.Println()
	return nums
}
