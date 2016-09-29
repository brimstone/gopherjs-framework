package client_test

import (
	"testing"

	"github.com/brimstone/gopherjs-framework/client"
	"github.com/stretchr/testify/assert"
)

func TestHTMLDiv(t *testing.T) {
	tag := client.NewDiv().Text("I'm a div")
	assert.Equal(t, tag.HTML(), "<div>I'm a div</div>")
}

func TestHTMLAddChild(t *testing.T) {
	tag := client.NewDiv()
	child := client.NewDiv()
	tag.AddChild(child)
	assert.Equal(t, "<div><div></div></div>", tag.HTML())
}
