package main

import (
	"fmt"
	"syscall/js"
	"time"
)

func main() {

	el := js.Global().Get("document").Call("getElementById", "thing")
	todolist := NewTodoList(el)

	todo := Todo{
		Title: "Wake Up",
	}
	todo2 := Todo{
		Title: "Rule the World",
	}
	todolist.Todos = []Todo{todo, todo2}
	err := todolist.Render()
	fmt.Println(err)
	fmt.Println("Sleeping for 5")
	time.Sleep(5 * time.Second)

	todo3 := Todo{
		Title: "Take a Nap",
	}

	fmt.Println("Adding another task")
	todolist.Todos = append(todolist.Todos, todo3)
	err = todolist.Render()
	fmt.Println(err)
}
