package register

import (
	"errors"
	"fmt"
	"github.com/fpawel/wasmhello/internal/ui/http"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"strings"
)

func New() *Compo {
	return &Compo{}
}

type Compo struct {
	app.Compo
	inputUser, pass string
	error           *error
}

func (x *Compo) Render() app.UI {
	hasResult := x.error != nil
	resultOk := hasResult && *x.error == nil
	return app.Div().Body(
		app.Form().Class("row g-3").
			Style("margin-top", "100px").
			Style("max-width", "600px").
			Style("padding", "30px").
			Body(
				app.H4().Class("form-label").Text("Register your mining account"),
				app.Div().
					Body(
						app.Label().For("inputAccountName").Text("user name"),
						app.Input().
							Type("text").ID("inputAccountName").
							Class("form-control").
							Placeholder("The name of the account to register").
							OnChange(func(ctx app.Context, e app.Event) {
								fmt.Println(ctx.JSSrc().Get("value"))
								x.inputUser = ctx.JSSrc().Get("value").String()
							}),
					),
				app.If(x.pass == "",
					app.Div().Class("d-flex flex-row-reverse").Body(
						app.Button().
							Class("btn btn-primary").ID("button-addon2").
							Type("button").
							Body(app.Text("Register")).
							OnClick(x.registerAccount),
					),
				),
				app.If(hasResult, app.If(resultOk, x.ok()).Else(x.err())),
			),
	)
}

func (x *Compo) ok() app.UI {
	return app.Div().
		Class("alert alert-success").
		Body(
			app.P().Body(
				app.I().Class("fas fa-thumbs-up").Style("margin-right", "5px"),
				app.Text("user"),
				app.B().Text(x.inputUser).Style("margin", "5px"),
				app.Text("registered successfully"),
			),
			app.P().Body(
				app.I().Class("fas fa-thumbs-up").Style("margin-right", "5px"),
				app.Text("password"),
				app.B().Text(x.pass).Style("margin", "5px"),
			),
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

func (x *Compo) registerAccount(_ app.Context, _ app.Event) {
	defer x.Update()

	x.inputUser = strings.TrimSpace(x.inputUser)
	if x.inputUser == "" {
		x.error = new(error)
		*x.error = errors.New("please enter something")
		return
	}

	var body struct {
		Account string `json:"account"`
	}
	body.Account = x.inputUser
	resp, err := http.C.R().
		EnableTrace().
		SetBody(&body).
		Post("/api/register")
	if err != nil {
		fmt.Println(err)
		return
	}
	x.pass = string(resp.Body())
	x.error = &err
}
