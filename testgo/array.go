package main
import (
	"fmt"
)
func sum(arr []int) []int {
	//var max_len int
	var new []int
	for i, _ := range arr {
		cur_sum := 0
		for j := i; j < len(arr); j++ {
			cur_sum = cur_sum + arr[j]
			if cur_sum == 0 {
				new = append(new, (j-i)+1)
				//max_len = max(max_len, (j-i)+1)
			}
		}
	}
	return new
}
//	func max(x1 int, x2 int) int {
//		if x1 > x2 {
//			return x1
//		}
//		return x2
//	}
func main() {
	arr := []int{2, 2, 1, -3, 4, 3, 1, -8, 6, -2, -1}
	fmt.Println(sum(arr))
}
