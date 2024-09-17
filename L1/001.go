package main

import "fmt"

func main() {
	person := Human{
		name: "Vanya",
	}

	person.sayHello()

	action := Action{
		Human: person,
	}

	action.sayHello()
	action.doSomeAction()
}

// Human Создаем структуру Human
type Human struct {
	name string
}

// Создаем какую-то функцию для Human
func (h *Human) sayHello() {
	fmt.Println("My name is", h.name)
}

// Action Создаем структуру Action от родительской Human
type Action struct {
	Human
	doSomething string
}

// Создаем какую-то функцию для Action
func (a *Action) doSomeAction() {
	fmt.Println("some work")
}
