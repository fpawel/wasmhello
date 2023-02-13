package login

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/fpawel/wasmhello/internal/server/datatype"
	"github.com/fpawel/wasmhello/internal/ui/uinfo"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"io"
	"net/http"
	"strings"
)

func New() *Compo {
	x := &Compo{
		input: uinfo.AccPass(),
	}
	if len(x.input.Account) == 0 {
		x.input = uinfo.GetInput().Login
	}
	if len(x.input.Account) == 0 {
		x.input = uinfo.GetInput().Register
	}
	return x
}

type Compo struct {
	app.Compo
	input datatype.AccPass
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
						app.Label().For("inputAccountName").Text("Account name"),
						app.Input().
							ID("inputAccountName").
							Type("text").
							Class("form-control").
							Placeholder("The name of the account").
							Value(x.input.Account).
							OnChange(func(ctx app.Context, e app.Event) {
								fmt.Println(ctx.JSSrc().Get("value"))
								x.input.Account = ctx.JSSrc().Get("value").String()
							}),
					),
				app.Div().
					Body(
						app.Label().For("inputPassword").Text("Password"),
						app.Input().
							ID("inputPassword").
							Type("password").
							Class("form-control").
							Placeholder("Password").
							Value(x.input.Password).
							OnChange(func(ctx app.Context, e app.Event) {
								fmt.Println(ctx.JSSrc().Get("value"))
								x.input.Password = ctx.JSSrc().Get("value").String()
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
			app.Text("Account"),
			app.B().Text(x.input.Account).Style("margin", "5px"),
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

	defer uinfo.AppUpdate()

	x.input.Account = strings.TrimSpace(x.input.Account)
	x.input.Password = strings.TrimSpace(x.input.Password)

	if len(x.input.Account) == 0 {
		x.error = new(error)
		*x.error = errors.New("Please enter account")
		return
	}

	if len(x.input.Password) == 0 {
		x.error = new(error)
		*x.error = errors.New("Please enter password")
		return
	}

	loc := app.Window().Get("location")
	host := loc.Get("protocol").String() + "//" + loc.Get("host").String() + "/api/login"
	body, _ := json.Marshal(x.input)
	req, err := http.NewRequest("POST", host, bytes.NewBuffer(body))
	req.Header.Add("js.fetch:mode", "cors")
	if err != nil {
		fmt.Println(err)
		return
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	x.error = new(error)
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		x.error = &err
	}
	response := string(b)

	if resp.StatusCode != 200 {
		*x.error = errors.New(response)
	} else {
		uinfo.SetToken(response)
	}

}
