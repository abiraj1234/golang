package main

import (
	"fmt"
	"time"
)

func main() {
	go delayedPrint("first")
	go delayedPrint("second")
	go delayedPrint("thrid")
	time.Sleep(1 * time.Second)
}

func delayedPrint(str string) {
	for i := 0; i < 5; i++ {
		fmt.Printf("%s, %d,\n", str, i)
		time.Sleep(100 * time.Millisecond)
	}
}
