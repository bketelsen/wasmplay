package main

import (
	"fmt"
	"reflect"
	"syscall/js"
)

func main() {

	el := js.Global.Get("document").Call("getElementById", "thing")

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

}
