package main

//go:generate go get -v -u github.com/brimstone/gopherjs-asset
//go:generate gopherjs-asset

import (
	"fmt"
	"net/http"

	"github.com/brimstone/gopherjs-framework/server"
)

func main() {
	http.Handle("/ws", server.WebsocketHandler(func(in <-chan string, out chan<- string, done <-chan bool) {
		fmt.Println("thing")
		out <- "Hello from server"
		for {
			select {
			case data := <-in:
				fmt.Println("Received:", data)
				if len(data) > 0 {
					out <- "Welcome to GopherJS, " + data + "!"
				} else {
					out <- ""
				}
			case <-done:
				return
			}
		}
	}))

	http.Handle("/assets/", server.AssetHandler(assets))
	http.HandleFunc("/", server.ReadFile(assets, "/assets/index.html"))

	fmt.Println("ready")
	http.ListenAndServe(":8080", nil)
}
