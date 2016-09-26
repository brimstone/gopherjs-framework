package main

import (
	"log"

	"github.com/gopherjs/websocket"

	"github.com/brimstone/gopherjs-framework/client"
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

	c, err := websocket.Dial(client.WebsocketAddress() + "/ws")
	if err != nil {
		return
	}

	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := c.Read(buf)
			if err != nil {
				log.Println("Websocket connection errored", err.Error())
				return
			}
			msg := string(buf[:n])
			log.Println(msg)
			if n > 0 {
				jQuery(OUTPUT).SetText(msg)
			} else {
				jQuery(OUTPUT).Empty()
			}
		}
	}()

	toserver := make(chan string)

	go func() {
		for {
			msg := <-toserver
			_, err := c.Write([]byte(msg))
			if err != nil {
				log.Println("Websocket writing error", err.Error())
			}
		}
	}()

	jQuery(INPUT).On(jquery.KEYUP, func(e jquery.Event) {
		name := jQuery(e.Target).Val()
		name = jquery.Trim(name)
		toserver <- name

	})

}
