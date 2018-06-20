package main

import (
	"bytes"
	"fmt"

	"html/template"

	"github.com/albrow/vdom"
)

type Todo struct {
	Title     string
	Completed bool
	//Root      js.Value
	tree *vdom.Tree
}

var todotemplate = `<li class="todo-list-item {{ if .Completed }} completed {{ end }}">
	<input class="toggle" type="checkbox" {{ if .Completed }} checked {{ end }}>
	<label class="todo-label">{{ .Title }}</label>
	<button class="destroy"></button>
	<input class="edit" onfocus="this.value = this.value;" value="{{ .Title }}">
</li>`

func (todo *Todo) Render() string {

	tmpl, err := template.New("todo").Parse(todotemplate)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	// Execute the template with the given todo and write to a buffer
	buf := bytes.NewBuffer([]byte{})
	if err := tmpl.Execute(buf, todo); err != nil {
		fmt.Println(err)
		return ""
	}
	// Parse the resulting html into a virtual tree
	newTree, err := vdom.Parse(buf.Bytes())
	if err != nil {
		fmt.Println(err)
		return ""
	}
	if todo.tree != nil {
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

		todo.tree = newTree
	}
	// Remember the virtual DOM state for the next render to diff against
	todo.tree = newTree

	//todo.Root.Set("innerHTML", string(newTree.HTML()))
	return string(newTree.HTML())
}
