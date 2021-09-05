package bs

import (
	"fmt"
	"github.com/maxence-charriere/go-app/v8/pkg/app"
)

func NavCollapseButton(id string) app.UI {
	return app.Raw(fmt.Sprintf(`
	<button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#%s" aria-controls="%s" aria-expanded="false" aria-label="Toggle navigation">
      <span class="navbar-toggler-icon"></span>
    </button>`, id, id))
}
