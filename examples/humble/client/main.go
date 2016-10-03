package main

import (
	"log"

	"honnef.co/go/js/dom"

	"github.com/brimstone/gopherjs-framework/examples/humble/client/templates"
	"github.com/go-humble/temple/temple"
	"github.com/go-humble/view"
)

type App struct {
	tmpl *temple.Template
	view.DefaultView
}

var (
	appTmpl  = templates.MustGetTemplate("app")
	document = dom.GetWindow().Document()
)

func main() {
	log.Println("Client!")

	myapp := &App{}
	myapp.SetElement(document.QuerySelector("body"))
	appTmpl.ExecuteEl(myapp.Element(), nil)
	log.Println(myapp.Element().InnerHTML())
}
