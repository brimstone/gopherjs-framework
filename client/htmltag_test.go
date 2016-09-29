package client_test

import (
	"testing"

	"../client"
)

func TestHTMLDiv(t *testing.T) {
	tag := client.NewDiv().Text("I'm a div")
	html := tag.HTML()
	if html != "<div>I'm a div</div>" {
		t.Error("expected a div")
	}
}
