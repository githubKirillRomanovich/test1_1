package main

import (
	"fmt"
)

type Human struct {
	Name string
	Age  int
}

func (h Human) SayHello() {
	fmt.Printf("Hi, my name is %s and I am %d years old.\n", h.Name, h.Age)
}

func (h Human) IsAdult() bool {
	return h.Age >= 18
}

type Action struct {
	Human
	Task string
}

func (a Action) DoSomething() {
	fmt.Printf("%s is doing task: %s\n", a.Name, a.Task)
}

func main() {
	action := Action{
		Human: Human{
			Name: "Kirill",
			Age:  33,
		},
		Task: "coding",
	}

	action.SayHello()
	fmt.Println("Is adult:", action.IsAdult())

	action.DoSomething()
}
