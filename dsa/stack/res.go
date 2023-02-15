package main

import (
	"fmt"
	"strings"
)

func main() {
	res := make(map[string]int)
	arr := []string{"hi", "happy", "new", "year"}
	for _, xn := range arr {
		res[strings.ToUpper(xn)] = len(xn)
	}
	fmt.Println(res)
}
