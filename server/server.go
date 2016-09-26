package server

import (
	"io"
	"log"
	"net/http"

	"golang.org/x/net/websocket"

	"github.com/shurcooL/httpgzip"
)

// AssetHandler handles serving assets
func AssetHandler(assets http.FileSystem) http.Handler {
	return httpgzip.FileServer(assets, httpgzip.FileServerOptions{ServeError: httpgzip.Detailed})
}

// ReadFile Handles reading in a file to a http handler
func ReadFile(filesystem http.FileSystem, filename string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		reader, err := filesystem.Open(filename)
		if err == nil {
			io.Copy(w, reader)
		}
	}
}

func readWebsocket(ws *websocket.Conn, outChan chan string, errChan chan error) {
	var message string
	var err error
	for {
		// read in a message
		err = websocket.Message.Receive(ws, &message)
		if err != nil {
			errChan <- err
			return
		}
		// send it out of our channel
		outChan <- message
	}
}

// WebsocketHandler does things
func WebsocketHandler(handler func(<-chan string, chan<- string, <-chan bool)) websocket.Handler {
	return websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()
		in := make(chan string)
		out := make(chan string)
		done := make(chan bool)
		clienterr := make(chan error)
		//doneReader := make(chan bool)

		go readWebsocket(ws, in, clienterr)
		go handler(in, out, done)
		for {
			select {
			case e := <-clienterr:
				if e.Error() != "EOF" {
					log.Println("Error from websocket: ", e.Error())
				}
				done <- true
				return
			case data := <-out:
				ws.Write([]byte(data))
			}
		}

	})
}
