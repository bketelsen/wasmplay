package main

import (
	dom "github.com/gowasm/go-js-dom"
	"github.com/gowasm/vdom"
)

type Component struct {
	Name     string
	Root     dom.Element
	Tree     *vdom.Tree
	Template string
}

type App struct {
	Root dom.Element
	Tree *vdom.Tree
}
