package profile

import (
	"github.com/fpawel/wasmhello/internal/ui/state"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type compo struct {
	app.Compo
	user string
}

func New() app.UI {
	return &compo{}
}

func (x *compo) OnMount(ctx app.Context) {
	ctx.ObserveState(state.User).Value(&x.user)
	ctx.GetState(state.User, &x.user)
}

func (x *compo) Render() app.UI {
	return app.Div().Body(
		app.H3().Class("jumbotron").Text("Profile"),
		app.If(
			x.user != "",
			app.H4().Text("user:"),
			app.H5().Text(x.user),
			app.Button().
				Class("btn btn-primary").ID("button-addon2").
				Type("button").
				Body(app.Text("Logout")).
				OnClick(
					func(ctx app.Context, e app.Event) {
						ctx.SetState(state.User, "", app.Persist)
					}),
		),
	)
}
