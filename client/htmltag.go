package client

type HTMLTag struct {
	tag        string
	attributes map[string]string
	text       string
	children   []HTMLTag
}

func NewDiv() HTMLTag {
	return HTMLTag{
		tag: "div",
	}
}

func (tag HTMLTag) Text(text string) HTMLTag {
	tag.text = text
	return tag
}

func (tag HTMLTag) HTML() string {
	return "<" + tag.tag + ">" + tag.text + "</" + tag.tag + ">"
}
