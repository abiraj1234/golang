package main

import "fmt"

func local_minima(arr []int) int {
	res := arr[0]
	for i, val := range arr {
		if i != 0 {
			if val < res {
				res = val
			}
		}
	}
	return res
}

func main() {
	arr := []int{2, 2, 1, -3, 4, 3, 1, -8, 6, -2, -1, 3, -3}

	fmt.Println(local_minima(arr))
}
