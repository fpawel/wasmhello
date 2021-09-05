package home

import "github.com/maxence-charriere/go-app/v8/pkg/app"

type Compo struct {
	app.Compo
}

func (l *Compo) Render() app.UI {
	return app.Div().Class("jumbotron").Body(
		app.H3().Text("Miners page not implemented"),
	)
}
