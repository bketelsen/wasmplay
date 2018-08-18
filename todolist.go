package main

import (
	"bytes"
	"fmt"

	"html/template"

	dom "github.com/gowasm/go-js-dom"
	"github.com/gowasm/vdom"
)

type TodoList struct {
	Todos []Todo
	Component
}

func NewTodoList(el dom.Element) *TodoList {
	comp := Component{
		Name:     "todolist",
		Root:     el,
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
	buf := new(bytes.Buffer)
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
		patches, err := vdom.Diff(todoList.Tree, newTree)
		if err != nil {
			return err
		}

		// Effeciently apply changes to the actual DOM
		if err := patches.Patch(todoList.Root); err != nil {
			return err
		}
	} else {

		todoList.Tree = newTree
	}
	// Remember the virtual DOM state for the next render to diff against
	todoList.Tree = newTree
	enumChildren(0, newTree.Children)
	todoList.Root.SetInnerHTML(string(newTree.HTML()))
	return nil
}

func enumChildren(level int, children []vdom.Node) {
	for i, child := range children {
		fmt.Printf("%d Child %d is %v\n", level, i, child)
		enumChildren(level+1, child.Children())
	}
	return
}
