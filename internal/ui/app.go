package ui

import "github.com/maxence-charriere/go-app/v8/pkg/app"

type App struct {
	app.Compo
	route route
}

func (x *App) routableContent() app.UI {
	switch x.route {
	case routeHome:
		return NewFoodList()
	case routeFeatures:
		return app.H3().Body(app.Text("Features"))
	case routePricing:
		return app.H3().Body(app.Text("Pricing"))
	default:
		panic(x.route)
	}
}

func (x *App) navigate(route route)  {
	x.route = route
	x.Update()
}

func (x *App) Render() app.UI {
	return app.Div().Class("container").Body(
		app.H1().Text("Sample web application"),
		&Navbar{
			OnNavigate: x.navigate,
		},
		x.routableContent(),
	)
}
