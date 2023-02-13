package regacc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/fpawel/wasmhello/internal/ui/uinfo"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"io"
	"net/http"
	"strings"
)

func New() *Compo {
	return &Compo{}
}

type Compo struct {
	app.Compo
	input string
	error *error
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
						app.Label().For("inputAccountName").Text("Account name"),
						app.Input().
							Type("text").ID("inputAccountName").
							Class("form-control").
							Placeholder("The name of the account to register").
							OnChange(func(ctx app.Context, e app.Event) {
								fmt.Println(ctx.JSSrc().Get("value"))
								x.input = ctx.JSSrc().Get("value").String()
							}),
					),
				app.Div().Class("d-flex flex-row-reverse").
					Body(
						app.Button().
							Class("btn btn-primary").ID("button-addon2").
							Type("button").
							Body(app.Text("Register")).
							OnClick(x.registerAccount),
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
			app.B().Text(x.input).Style("margin", "5px"),
			app.Text("registered successfully"),
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

func (x *Compo) registerAccount(ctx app.Context, e app.Event) {
	defer x.Update()

	x.input = strings.TrimSpace(x.input)
	if len(x.input) == 0 {
		x.error = new(error)
		*x.error = errors.New("Please enter something")
		return
	}
	loc := app.Window().Get("location")
	host := loc.Get("protocol").String() + "//" + loc.Get("host").String() + "/api/register"
	body, _ := json.Marshal(map[string]interface{}{
		"account": x.input,
	})
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
	b, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		*x.error = errors.New(string(b))
		return
	}
	r := uinfo.GetInput()
	r.Register.Account = x.input
	r.Register.Password = string(b)
	uinfo.SetInput(r)
}
