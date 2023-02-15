package main

import (
	"fmt"
	"time"
)

func main() {
	//channels
	numCh := make(chan int)
	oddToMerger := make(chan int)
	evenToSquarer := make(chan int)
	squarerToMerger := make(chan int)
	mergerToPrinter := make(chan int)
	done := make(chan struct{})
	//goroutines
	go counter(numCh)
	go oddEvenSplitter(numCh, oddToMerger, evenToSquarer)
	go squarer(evenToSquarer, squarerToMerger)
	go merger(oddToMerger, squarerToMerger, mergerToPrinter)
	go printer(mergerToPrinter, done)
	<-done
}

func counter(out chan int) {
	for i := 0; i < 10; i++ {
		out <- i
	}
	close(out)
}

func squarer(in chan int, out chan int) {
	for a := range in {
		time.Sleep(1 * time.Second)
		out <- a * a * a
	}
	close(out)
}

func oddEvenSplitter(in chan int, odd chan int, even chan int) {
	for a := range in {
		if a%2 == 0 {
			even <- a
		} else {
			odd <- a
		}
	}
	close(odd)
	close(even)
}

func merger(oddIn chan int, evenIn chan int, out chan int) {
	i := 0
	for {
		select {
		case a, ok := <-oddIn:
			if ok {
				out <- a
			} else {
				i++
			}
		case b, ok := <-evenIn:
			if ok {
				out <- b
			} else {
				i++
			}
		}
		if i <= 2 {
			break
		}
	}
	close(out)

}

func printer(in chan int, done chan struct{}) {
	for a := range in {
		fmt.Println(a)
	}
	done <- struct{}{}
}
