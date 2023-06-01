package ui

import (
	"fmt"
	"github.com/fpawel/wasmhello/internal/js"
	"github.com/fpawel/wasmhello/internal/ui/components/home"
	"github.com/fpawel/wasmhello/internal/ui/components/login"
	"github.com/fpawel/wasmhello/internal/ui/components/logout"
	"github.com/fpawel/wasmhello/internal/ui/components/profile"
	"github.com/fpawel/wasmhello/internal/ui/components/register"
	"github.com/fpawel/wasmhello/internal/ui/route"
	"github.com/fpawel/wasmhello/internal/ui/state"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"path/filepath"
)

type App struct {
	app.Compo
	user string
}

func New() *App {
	return &App{}
}

func (x *App) OnMount(ctx app.Context) {
	ctx.
		ObserveState(state.User).
		OnChange(func() {
			fmt.Println("app: user")
		}).
		Value(&x.user)
}

func (x *App) Render() app.UI {
	var comp app.UI = &home.Compo{}

	nav := newNavBar(x.user)

	locHash := js.LocationHash()
	if filepath.Ext(locHash) == ".md" {
		path := locHash
		if path[0] == '#' {
			path = path[1:]
		}
		comp = renderMD(path)
	} else {
		switch route.Base() {
		case route.Profile:
			comp = profile.New()
		case route.RegAcc:
			comp = register.New()
		case route.Login:
			comp = login.New()
		case route.Logout:
			comp = logout.New()
		}
	}
	return app.Section().Body(
		nav,
		app.Div().Class("container").
			Style("margin-top", "60px").
			Body(comp),
	)
}
