package ui

import (
	"fmt"
	"github.com/fpawel/wasmhello/internal/js"
	"github.com/fpawel/wasmhello/internal/ui/components/home"
	"github.com/fpawel/wasmhello/internal/ui/components/login"
	"github.com/fpawel/wasmhello/internal/ui/components/profile"
	"github.com/fpawel/wasmhello/internal/ui/components/regacc"
	"github.com/fpawel/wasmhello/internal/ui/uinfo"
	"github.com/maxence-charriere/go-app/v8/pkg/app"
	"path/filepath"
	"strings"
)

type App struct {
	app.Compo
	loggedIn bool
}

func New() *App {
	return &App{}
}

func getBaseRoute() Route {
	locHash := js.LocationHash()
	parts := strings.Split(locHash, "/")
	if len(parts) == 0 || parts[0] == "#" {
		return RouteHome
	}
	return Route(parts[0])
}

func (x *App) OnMount(ctx app.Context) {
	uinfo.AppUpdate = x.Update
}

func (x *App) OnDismount(ctx app.Context) {
	uinfo.AppUpdate = func() {
		fmt.Println("AppUpdate: dismounted")
	}
}

func (x *App) Render() app.UI {
	var comp app.UI = &home.Compo{}

	locHash := js.LocationHash()
	if filepath.Ext(locHash) == ".md" {
		path := locHash
		if path[0] == '#' {
			path = path[1:]
		}
		comp = renderMD(path)
	} else {
		switch getBaseRoute() {
		case RouteProfile:
			comp = &profile.Compo{}
		case RouteRegAcc:
			comp = regacc.New()
		case RouteLogin:
			comp = login.New()
		}
	}
	return app.Section().Body(
		&navBar{},
		app.Div().Class("container").
			Style("margin-top", "60px").
			Body(comp),
	)
}

func (x *App) OnNav(ctx app.Context) {
	fmt.Println("OnNav")
	//x.Update()
}
