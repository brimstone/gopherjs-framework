package client

// HTMLFragment is the smallest unit of html
type HTMLFragment interface {
	Text(string) HTMLFragment
	HTML() string
}
