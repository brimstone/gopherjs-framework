package client

import (
	"github.com/go-humble/temple/temple"
	"github.com/go-humble/view"
)

type AppOptions struct {
}

type App struct {
	Tmpl *temple.Template
	view.DefaultView
}

// Render renders the App view and satisfies the view.View interface.
func (a *App) Render() error {
	if err := a.Tmpl.ExecuteEl(a.Element(), nil); err != nil {
		return err
	}
	return nil
}

// NewApp Returns a new App
func NewApp(config AppOptions) (*App, error) {
	a := new(App)
	return a, nil
}
