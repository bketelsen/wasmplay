// +build wasm
package main

import (
	"fmt"

	dom "github.com/gowasm/go-js-dom"

	"time"
)

func main() {

	window := dom.GetWindow()
	fmt.Println(window)
	doc := window.Document()
	fmt.Println(doc)
	el := doc.GetElementByID("thing")
	fmt.Println(el)
	fmt.Println(el.InnerHTML())
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

	fmt.Println("Sleeping for 5")
	time.Sleep(5 * time.Second)
	todolist.Todos = todolist.Todos[0:1]
	err = todolist.Render()
	fmt.Println(err)
}
