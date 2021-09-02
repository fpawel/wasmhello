package ui

import (
	"github.com/maxence-charriere/go-app/v8/pkg/app"
)

type route int

const (
	routeHome route = iota
	routeFeatures
	routePricing
)

type Navbar struct {
	app.Compo
	OnNavigate func(route)
}

type NavContent struct {
	Name string
	Content app.UI
}

func (x *Navbar) brand() app.UI {
	return app.A().Class("navbar-brand").Href("#").Body(app.Text("Navbar"))
}

func (x *Navbar) navigate(r route) func(){
	return func() {
		x.OnNavigate(r)
	}
}

func (x *Navbar) Render() app.UI {
	return app.Nav().Class("navbar navbar-expand-lg navbar-dark bg-dark").Body(
		app.Div().Class("container-fluid").Body(
			x.brand(),
			app.Ul().Class("navbar-nav").Body(
				navItem{ref: "#", text: "Home", classLi: "active", ariaCurrentPage: true, onNavigate: x.navigate(routeHome)}.render(),
				navItem{ref: "#", text: "Features", onNavigate: x.navigate(routeFeatures)}.render(),
				navItem{ref: "#", text: "Pricing", onNavigate: x.navigate(routePricing)}.render(),
			),
		),
	)
}

type navItem struct {
	classLi, classA string
	text string
	ref string
	ariaCurrentPage bool
	onNavigate func()
}

func (x navItem) render() app.HTMLLi {
	a := app.A().Class("nav-link" + x.classA).Href(x.ref).Body(app.Text(x.text))
	if x.ariaCurrentPage {
		a = a.Aria("current", "page")
	}
	return app.Li().Class("nav-item"+ x.classLi).Body(a).
		OnClick(func(ctx app.Context, e app.Event) {
			x.onNavigate()
		})
}

