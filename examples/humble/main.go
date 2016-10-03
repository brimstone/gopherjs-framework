package main

//go:generate temple build client/templates/templates client/templates/templates.go --partials client/templates/partials
//go:generate go get -v github.com/brimstone/gopherjs-asset
//go:generate gopherjs-asset

import (
	"fmt"
	"net/http"

	"github.com/brimstone/gopherjs-framework/server"
)

func main() {
	http.Handle("/assets/", server.AssetHandler(assets))
	http.HandleFunc("/", server.ReadFile(assets, "/assets/index.html"))

	fmt.Println("ready to serve")
	http.ListenAndServe(":8081", nil)
}
