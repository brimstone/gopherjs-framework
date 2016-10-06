package main

import (
	"log"

	"honnef.co/go/js/dom"

	"github.com/brimstone/gopherjs-framework/client"
	"github.com/brimstone/gopherjs-framework/examples/humble/client/templates"
	"github.com/go-humble/router"
)

var (
	document = dom.GetWindow().Document()
)

func main() {
	log.Println("Client!")

	appView := &client.App{
		Tmpl: templates.MustGetTemplate("app"),
	}
	appView.SetElement(document.QuerySelector("body"))

	// Create a new Router object
	r := router.New()
	// Use HandleFunc to add routes.
	r.HandleFunc("/", func(_ *router.Context) {
		if err := appView.Render(); err != nil {
			panic(err)
		}
	})
	// You must call Start in order to start listening for changes
	// in the url and trigger the appropriate handler function.
	r.Start()

}
