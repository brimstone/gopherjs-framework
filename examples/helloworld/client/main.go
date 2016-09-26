package main

import (
	"log"

	"github.com/gopherjs/websocket"

	"github.com/brimstone/gopherjs-framework/client"
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/jquery"
)

const (
	// INPUT holds our user's name
	INPUT = "input#name"
	// OUTPUT holds greeting
	OUTPUT = "span#greeting"
)

var jQuery = jquery.NewJQuery

func main() {

	log.Println("Hello world")

	ws, err := websocket.New(client.WebsocketAddress() + "/ws")
	if err != nil {
		return
	}

	ws.AddEventListener("open", false, func(data *js.Object) {
		log.Println("Websocket open")
	})

	ws.AddEventListener("message", false, func(ev *js.Object) {
		msg := ev.Get("data").String()
		if len(msg) > 0 {
			jQuery(OUTPUT).SetText(msg)
		} else {
			jQuery(OUTPUT).Empty()
		}
	})

	jQuery(INPUT).On(jquery.KEYUP, func(e jquery.Event) {
		name := jQuery(e.Target).Val()
		name = jquery.Trim(name)
		err := ws.Send(name)
		if err != nil {
			log.Println("Got error while writing to websocket:", err)
		}

	})

}
