package login

import (
	"errors"
	"fmt"
	"github.com/fpawel/wasmhello/internal/server/datatype"
	"github.com/fpawel/wasmhello/internal/ui/http"
	"github.com/fpawel/wasmhello/internal/ui/route"
	"github.com/fpawel/wasmhello/internal/ui/state"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"strings"
)

func New() *Compo {
	return &Compo{}
}

type Compo struct {
	app.Compo
	input datatype.UserPass
	error *error
}

func (x *Compo) Render() app.UI {
	hasResult := x.error != nil
	resultOk := hasResult && *x.error == nil
	return app.Div().Body(
		app.Form().
			Style("margin-top", "100px").
			Style("max-width", "600px").
			Style("padding", "30px").Class("row g-3").
			Body(
				app.H4().Class("form-label").Text("Sign in with your credentials"),
				app.Div().
					Body(
						app.Label().For("inputAccountName").Text("user name"),
						app.Input().
							ID("inputAccountName").
							Type("text").
							Class("form-control").
							Placeholder("The name of the account").
							Value(x.input.User).
							OnChange(func(ctx app.Context, e app.Event) {
								x.input.User = ctx.JSSrc().Get("value").String()
							}),
					),
				app.Div().
					Body(
						app.Label().For("inputPassword").Text("Pass"),
						app.Input().
							ID("inputPassword").
							Type("password").
							Class("form-control").
							Placeholder("Pass").
							Value(x.input.Pass).
							OnChange(func(ctx app.Context, e app.Event) {
								x.input.Pass = ctx.JSSrc().Get("value").String()
							}),
					),
				app.Div().Class("d-flex flex-row-reverse").
					Body(
						app.Button().
							Class("btn btn-primary").ID("button-addon2").
							Type("button").
							Body(app.Text("Login")).
							OnClick(x.login),
					),
				app.If(hasResult, app.If(resultOk, x.ok()).Else(x.err())),
			),
	)
}

func (x *Compo) ok() app.UI {
	return app.Div().
		Class("alert alert-success").
		Body(
			app.I().Class("fas fa-thumbs-up").Style("margin-right", "5px"),
			app.Text("user"),
			app.B().Text(x.input.User).Style("margin", "5px"),
			app.Text("logged in successfully"),
		)
}

func (x *Compo) err() app.UI {
	if x.error == nil || *x.error == nil {
		return nil
	}
	return app.Div().
		Class("alert alert-danger").
		Body(
			app.I().Class("fas fa-exclamation-triangle").Style("margin-right", "5px"),
			app.Text((*x.error).Error()))
}

func (x *Compo) login(ctx app.Context, e app.Event) {
	x.input.User = strings.TrimSpace(x.input.User)
	x.input.Pass = strings.TrimSpace(x.input.Pass)

	if len(x.input.User) == 0 {
		x.error = new(error)
		*x.error = errors.New("please enter account")
		return
	}

	if len(x.input.Pass) == 0 {
		x.error = new(error)
		*x.error = errors.New("please enter password")
		return
	}
	r, err := http.C.R().SetBody(x.input).
		//SetHeader("js.fetch:mode", "cors").
		Post("/api/login")
	if err != nil {
		x.error = &err
		return
	}
	if r.StatusCode() != 200 {
		err := fmt.Errorf("%s: %s", r.Body(), r.Status())
		x.error = &err
		return
	}

	v, err := datatype.JwtParse(r.String())
	if err != nil {
		x.error = &err
		return
	}

	ctx.SetState(state.User, v.User, app.Persist)
	ctx.Navigate(route.Home)
}
