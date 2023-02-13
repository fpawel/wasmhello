package ui

import (
	"fmt"
	"github.com/fpawel/wasmhello/internal/js"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/russross/blackfriday/v2"
	"io"
	"net/http"
	"path"
	"strings"
)

func renderMD(Path string) app.UI {
	u := js.URL()
	u.Path = path.Join("web", Path)
	req, err := http.NewRequest("GET", u.String(), nil)
	req.Header.Add("js.fetch:mode", "cors")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return renderError("Internal error")
	}
	defer resp.Body.Close()

	response, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return renderError("Internal error")
	}
	if resp.StatusCode != http.StatusOK {
		return renderError(string(response))
	}

	sb := new(strings.Builder)
	sb.WriteString("<section>")
	sb.Write(blackfriday.Run(response))
	sb.WriteString("</section>")
	return app.Raw(sb.String())
}

func renderError(err interface{}) app.UI {
	return app.Div().
		Class("alert alert-danger").
		Body(
			app.I().Class("fas fa-exclamation-triangle").Style("margin-right", "5px"),
			app.Text(err))
}
