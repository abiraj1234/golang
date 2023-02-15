package main

import (
	"fmt"
)

func sum(arr []int) int {
	var max_len int
	for i, _ := range arr {
		cur_sum := 0
		for j := i; j < len(arr); j++ {
			cur_sum = cur_sum + arr[j]
			if cur_sum == 0 {
				max_len = maxn(max_len, (j-i)+1)
			}
		}
	}
	return max_len
}

func sumn(arr []int) int {
	k := make(map[int]int)
	var max_len int
	cur_sum := 0
	for i, _ := range arr {
		cur_sum += arr[i]
		if _, ok := k[cur_sum]; ok {
			max_len = maxn(max_len, i-k[cur_sum])
		} else {
			k[cur_sum] = i
		}
	}
	return max_len
}
func maxn(v1 int, v2 int) int {
	if v1 > v2 {
		return v1
	}
	return v2
}

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
	arr := []int{2, 2, 1, -3, 4, 3, 1, -8, 6, -2, -1}
	fmt.Println(sum(arr))
	fmt.Println(sumn(arr))
	fmt.Println(local_minima(arr))
}
