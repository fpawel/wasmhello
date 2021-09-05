package dlg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/maxence-charriere/go-app/v8/pkg/app"
	"net/http"
)

type RegisterAccount struct {
	app.Compo
	id    string
	input string
}

func New(id string) *RegisterAccount {
	return &RegisterAccount{id: id}
}

func (x *RegisterAccount) Render() app.UI {
	return app.Div().Class("modal").ID(x.id).
		TabIndex(-1).
		Aria("labelledby", x.id+"Label").
		Aria("hidden", "true").Body(
		app.Div().Class("modal-dialog modal-dialog-centered").Body(
			app.Div().Class("modal-content").
				Body(
					x.renderHeader(),
					x.renderBody(),
					x.renderFooter(),
				),
		),
	)
}

func (x *RegisterAccount) renderHeader() app.UI {
	return app.Div().Class("modal-header").Body(
		app.H5().Class("modal-title").ID(x.id+"Label").Body(app.Text("Register account")),
		app.Button().Type("button").
			Class("btn-close").
			DataSet("bs-dismiss", "modal").
			Aria("label", "Close").
			ID(x.id+"Close"),
	)
}

func (x *RegisterAccount) renderBody() app.UI {
	return app.Div().Class("modal-body input-group mb-3").Body(
		app.Input().
			Type("text").
			Class("form-control").
			Placeholder("The name of the account").
			Aria("label", "Register account").
			OnChange(func(ctx app.Context, e app.Event) {
				x.input = ctx.JSSrc.Get("value").String()
			}),
	)
}

func (x *RegisterAccount) renderFooter() app.UI {
	return app.Div().Class("modal-footer").Body(
		app.Button().Type("button").
			Class("btn btn-primary").
			Body(app.Text("Register")).
			OnClick(
				func(ctx app.Context, e app.Event) {
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

					app.Window().GetElementByID(x.id + "Close").Call("click")
				},
			),
	)
}
