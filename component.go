package main

import (
	"syscall/js"

	"github.com/albrow/vdom"
)

type Component struct {
	Name     string
	Root     js.Value
	Tree     *vdom.Tree
	Template string
}

type App struct {
	Root js.Value
	Tree *vdom.Tree
}
