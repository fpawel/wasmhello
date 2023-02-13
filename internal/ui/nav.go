package ui

import (
	"github.com/fpawel/wasmhello/internal/components/bs"
	"github.com/fpawel/wasmhello/internal/js"
	"github.com/fpawel/wasmhello/internal/ui/uinfo"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type Route string

const (
	RouteHome    = "#home"
	RouteProfile = "#profile"
	RouteRegAcc  = "#regacc"
	RouteLogin   = "#login"
)

func navBarItem(itemRoute Route, text string, icon string) app.UI {

	appRoute := getBaseRoute()
	a := app.A().Class("btn btn-outline-primary").Href(string(itemRoute)).
		Style("margin-right", "20px").Body(
		app.If(
			len(icon) > 0,
			app.I().Class(icon).Style("margin-right", "5px"),
		),
		app.Text(text),
	)
	if string(itemRoute) == js.LocationHash() ||
		itemRoute == appRoute ||
		(appRoute == "#" || appRoute == "") && itemRoute == RouteHome {
		a = a.Aria("current", "page").Class("active")
	}

	return app.Div().Class("nav-item").Body(a)
}

type navBar struct {
	app.Compo
}

func (x *navBar) Render() app.UI {
	return app.Nav().Class("container container-fluid navbar fixed-top navbar-expand-lg navbar-dark bg-dark bg-gradient").
		Body(
			app.A().Class("navbar-brand").Href("/").Body(
				app.Img().Src("/web/img/logo.svg").Alt("").Width(30).Height(24),
				app.Text("Title"),
			),
			bs.NavCollapseButton(idNavbarNavAltMarkup),
			app.Div().Class("collapse navbar-collapse").ID(idNavbarNavAltMarkup).Body(
				app.Div().Class("navbar-nav").Body(
					navBarItem(RouteHome, "Miners", ""),
				),
				app.Div().Class("navbar-nav").Body(
					navBarItem("#doc/readme.md", "About", ""),
				),
				app.Div().Class("nav-item flex-fill").Body(),

				app.If(
					uinfo.LoggedIn(),
					app.Div().Class("nav-item").
						Style("margin-right", "10px").
						Body(
							app.Span().Class("navbar-text fs-5 fw-bold").Text("Logged in"),
						),
					navBarItem(RouteProfile, "Profile", "fas fa-sign-in-alt"),
				),

				app.Div().Class("nav-item").Body(
					app.Button().Class("btn btn-outline-primary nav-item dropdown-toggle").
						ID(idNavbarDropdown).
						DataSet("bs-toggle", "dropdown").
						Aria("expanded", "false").
						Type("button").
						Body(icon("fas fa-user"), app.Text("Account")),
					dropdownNavMenu())))
}

func dropdownNavMenu() app.UI {
	return app.Ul().Class("dropdown-menu").
		Class("dropdown-menu-end").
		Aria("labelledby", idNavbarDropdown).Body(
		dropdownItem(RouteRegAcc, "Register", "fas fa-registered"),
		// app.Li().Class("dropdown-divider"),
		app.If(
			uinfo.LoggedIn(),
			app.Li().Body(app.A().Class("dropdown-item").Href("#").
				Text("Logout").OnClick(func(ctx app.Context, e app.Event) {
				uinfo.Logout()
			})),
		).Else(
			dropdownItem(RouteLogin, "Login", "fas fa-sign-in-alt"),
		),
	)
}

func dropdownItem(route Route, text, iconClass string) app.UI {
	return app.Li().Body(app.A().Class("dropdown-item").Href(string(route)).
		Body(
			icon(iconClass),
			app.Text(text)))

}

func icon(class string) app.UI {
	return app.I().Class(class).Style("margin-right", "5px")
}

//func dropdownMenuItemShowModal(text string) app.UI  {
//	return app.A().
//		Class("dropdown-item").
//		Href(jsutils.LocationHash()).
//		DataSet("bs-toggle", "modal").
//		DataSet("bs-target", "#"+idAppModalDialog).Body(app.Text(text))
//}

const (
	idAppModalDialog     = "appModalDialog"
	idNavbarNavAltMarkup = "navbarNavAltMarkup"
	idNavbarDropdown     = "navbarDropdown"
)
