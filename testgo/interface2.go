 package main

import (
	"fmt"
)

type Employee struct {
	Name string
}

type Manager struct {
	title string
}

func (m Manager) Title() string {
	return m.title
}

type Hastitle interface {
	Title() string
}

func main() {
	x := "84843"
	print(x)
	print(Employee{Name: "donald"})
	print(Manager{title: "CEO"})
}

func print(data interface{}) {
	switch val := data.(type) {
	case string:
		fmt.Printf("%s is a string\n", val)
	case int:
		fmt.Printf("%d is a interger\n", val)
	case Employee:
		fmt.Printf("%s is a Employee\n", val.Name)
	case Hastitle:
		fmt.Printf("%s is title of give type\n", val.Title())
	default:
		fmt.Printf("its wrong\n")
	}
}
