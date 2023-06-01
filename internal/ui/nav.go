package ui

import (
	"github.com/fpawel/wasmhello/internal/components/bs"
	"github.com/fpawel/wasmhello/internal/js"
	"github.com/fpawel/wasmhello/internal/ui/route"
	"github.com/fpawel/wasmhello/internal/ui/state"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

func navBarItem(itemRoute route.R, text string, icon string) app.UI {
	appRoute := route.Base()
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
		(appRoute == "#" || appRoute == "") && itemRoute == route.Home {
		a = a.Aria("current", "page").Class("active")
	}

	return app.Div().Class("nav-item").Body(a)
}

type navBarCompo struct {
	user string
	app.Compo
}

func newNavBar(user string) app.UI {
	return &navBarCompo{user: user}
}

func (x *navBarCompo) OnMount(ctx app.Context) {
	ctx.ObserveState(state.User).Value(&x.user)
}

func (x *navBarCompo) Render() app.UI {
	return app.Nav().Class("container container-fluid navbar fixed-top navbar-expand-lg navbar-dark bg-dark bg-gradient").
		Body(
			app.A().Class("navbar-brand").Href("/").Body(
				app.Img().Src("/web/img/logo.svg").Alt("").Width(30).Height(24),
				app.Text("Title"),
			),
			bs.NavCollapseButton(idNavbarNavAltMarkup),
			app.Div().Class("collapse navbar-collapse").ID(idNavbarNavAltMarkup).Body(
				app.Div().Class("navbar-nav").Body(
					navBarItem(route.Home, "Miners", ""),
				),
				app.Div().Class("navbar-nav").Body(
					navBarItem("#doc/readme.md", "About", ""),
				),
				app.Div().Class("nav-item flex-fill").Body(),

				app.If(
					x.user != "",
					app.Div().Class("nav-item").
						Style("margin-right", "10px").
						Body(
							app.Span().Class("navbar-text fs-5 fw-bold").Text("Logged in"),
						),
					navBarItem(route.Profile, "Profile", "fas fa-sign-in-alt"),
					app.Div().Class("navbar-nav").Body(
						navBarItem(route.Logout, "Logout", "fas fa-user"),
					),
				).Else(
					navBarItem(route.RegAcc, "Register", "fas fa-registered"),
					navBarItem(route.Login, "Login", "fas fa-sign-in-alt"),
				),
			),
		)
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
