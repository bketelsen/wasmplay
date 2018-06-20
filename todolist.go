package main

import (
	"bytes"
	"syscall/js"

	"html/template"

	"github.com/albrow/vdom"
)

type TodoList struct {
	Todos []Todo
	Component
}

func NewTodoList(el js.Value) *TodoList {
	comp := Component{
		Name:     "todolist",
		Root:     el,
		Tree:     &vdom.Tree{},
		Template: todolisttemplate,
	}
	return &TodoList{
		Todos:     make([]Todo, 1),
		Component: comp,
	}
}

var todolisttemplate = `<ul>
{{range $i, $x := $.Todos}} 
	{{$x.Render}} 
{{end}}
</ul>`

func (todoList *TodoList) Render() error {

	tmpl, err := template.New("todolist").Parse(todoList.Template)
	if err != nil {
		return err
	}
	// Execute the template with the given todo and write to a buffer
	buf := bytes.NewBuffer([]byte{})
	if err := tmpl.Execute(buf, todoList); err != nil {
		return err
	}
	// Parse the resulting html into a virtual tree
	newTree, err := vdom.Parse(buf.Bytes())
	if err != nil {
		return err
	}

	if todoList.Tree != nil {
		// Calculate the diff between this render and the last render
		//	patches, err := vdom.Diff(todo.tree, newTree)
		//	if err != nil {
		//		return err
		//	}

		// Effeciently apply changes to the actual DOM
		//		if err := patches.Patch(todo.Root); err != nil {
		//			return err
		//		}
	} else {

		todoList.Tree = newTree
	}
	// Remember the virtual DOM state for the next render to diff against
	todoList.Tree = newTree

	todoList.Root.Set("innerHTML", string(newTree.HTML()))
	return nil
}
