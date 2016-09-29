package client

// HTMLTag is a struct that implements HTMLFragment
type HTMLTag struct {
	tag        string
	attributes map[string]string
	text       string
	children   []HTMLTag
}

// NewDiv creates a new HTMLTag of type div
func NewDiv() *HTMLTag {
	return &HTMLTag{
		tag: "div",
	}
}

// Text assigns a string to the text attribute of the tag
func (tag *HTMLTag) Text(text string) *HTMLTag {
	tag.text = text
	return tag
}

// HTML returns the fully rendered html of the tag, including all children
func (tag HTMLTag) HTML() string {
	html := "<" + tag.tag
	html += ">"
	html += tag.text
	for _, child := range tag.children {
		html += child.HTML()
	}
	html += "</" + tag.tag + ">"
	return html
}

// AddChild adds an HTMLTag as a child to the tag
func (tag *HTMLTag) AddChild(child *HTMLTag) *HTMLTag {
	tag.children = append(tag.children, *child)
	return tag
}
