package logout

import (
	"github.com/fpawel/wasmhello/internal/ui/route"
	"github.com/fpawel/wasmhello/internal/ui/state"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

func New() app.UI {
	return &compo{}
}

type compo struct {
	app.Compo
}

func (x *compo) OnMount(ctx app.Context) {
	ctx.SetState(state.User, "", app.Persist)
	ctx.Navigate(route.Home)
}

func (x *compo) Render() app.UI {
	return app.Raw("")
}
