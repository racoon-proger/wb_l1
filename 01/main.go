package main

import "fmt"

// реализовал наследование путем встраивание полей структуры Human в структуру Action
type Human struct {
	Name string
}

func (h *Human) speak() {
	fmt.Printf("hello i am %s", h.Name)
}

type Action struct {
	Human
}

func main() {
	action := Action{
		Human: Human{
			Name: "Evgen",
		},
	}
	action.speak()
}
