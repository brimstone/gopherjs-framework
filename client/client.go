package client

import "github.com/gopherjs/gopherjs/js"

// Websocketaddress returns the same host, ssl
func WebsocketAddress() string {
	location := js.Global.Get("window").Get("location")
	protocol := location.Get("protocol").String()
	host := location.Get("host").String()

	if protocol == "http:" {
		protocol = "ws:"
	} else {
		protocol = "wss:"
	}

	return protocol + "//" + host
}
