package main

import "fmt"

type queue struct {
	array []int
}

func (q *queue) enqueue(val int) {
	q.array = append(q.array, val)
}

func (q *queue) size() int {
	return len(q.array)
}

func (q *queue) dqueue() {
	if q.size() == 0 {
		return
	}
	fmt.Println("sirapana alu nee", q.array[0])
	q.array = q.array[1:q.size()]
}

func (q *queue) print() {
	if q.size() == 0 {
		fmt.Println("onumila")
		return
	}
	for i := 0; i < q.size(); i++ {
		fmt.Printf("%d ", q.array[i])
	}
	fmt.Println()
}

func main() {
	que := queue{}
	que.enqueue(100)
	que.enqueue(200)
	que.enqueue(300)
	que.enqueue(400)
	que.print()
	que.dqueue()
	que.print()

}
