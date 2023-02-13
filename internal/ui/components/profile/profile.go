package profile

import (
	"github.com/fpawel/wasmhello/internal/ui/uinfo"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type Compo struct {
	app.Compo
}

func (x *Compo) OnMount(ctx app.Context) {
	if !uinfo.LoggedIn() {
		ctx.Navigate("/#home")
	}
}

func (x *Compo) Render() app.UI {
	return app.Div().Body(
		app.H3().Class("jumbotron").Text("Miners page not implemented"),
		app.If(
			uinfo.LoggedIn(),
			app.H4().Text("Account:"),
			app.H5().Text(uinfo.AccPass().Account),
			app.H4().Text("Password"),
			app.H5().Text(uinfo.AccPass().Password),
			app.Button().
				Class("btn btn-primary").ID("button-addon2").
				Type("button").
				Body(app.Text("Logout")).
				OnClick(func(ctx app.Context, e app.Event) {
					uinfo.Logout()
				}),
		),
	)
}
