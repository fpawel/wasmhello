package home

import (
	"encoding/json"
	"fmt"
	"github.com/fpawel/wasmhello/internal/js"
	"github.com/fpawel/wasmhello/internal/server/datatype"
	"github.com/maxence-charriere/go-app/v8/pkg/app"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

type Compo struct {
	app.Compo
	miners      []datatype.Miner
	minersCount int
	minersPage  int
}

const minersRowsCount = 40

func (l *Compo) fetchMiners(ctx app.Context) {
	fmt.Println("requested miners")
	minersCount, err := fetchMinersCount()
	if err != nil {
		js.LogErr("fetchMinersCount:", err)
		return
	}
	to := minersRowsCount * (l.minersPage + 1)
	if minersCount < to {
		to = minersCount
	}

	from := minersRowsCount * l.minersPage

	miners, err := fetchMiners(from, to)
	if err != nil {
		fmt.Println("fetchMiners:", err)
		return
	}
	fmt.Println("got miners", from, to, len(miners))
	l.miners = miners
	l.minersCount = minersCount
	ctx.Async(func() {
		ctx.Dispatch(func(ctx app.Context) {
			l.Update()
		})
	})
}

func (l *Compo) OnMount(ctx app.Context) {
	l.fetchMiners(ctx)
}

func (l *Compo) getPages() []int {

	n := l.minersPage

	count := l.minersCount / minersRowsCount

	var xs []int
	if count < 7 {
		for i := 0; i < count; i++ {
			xs = append(xs, i)
		}
	} else if n < 3 {
		for i := 0; i < 6; i++ {
			xs = append(xs, i)
		}
		xs = append(xs, -1, -1, count-1)
	} else if count-n < 4 {
		xs = append(xs, 0, -1, -1)
		for i := count - 6; i < count; i++ {
			xs = append(xs, i)
		}
	} else {
		xs = append(xs, 0, -1)
		for i := 1; i < count-1; i++ {
			v := n - i
			if v < 0 {
				v *= -1
			}
			if i == 0 || i == count-1 || v < 3 {
				xs = append(xs, i)
			}
		}
		xs = append(xs, -1, count-1)
	}
	return xs
}

func (l *Compo) tableNavigation() app.UI {

	locHash := js.LocationHash()
	pages := l.getPages()
	fmt.Println("render Pages:", pages)
	return app.Nav().
		Aria("label", "Table navigation").
		Body(
			app.Ul().
				Class("pagination justify-content-end").
				Body(
					app.Li().Style("width", "44px").
						Class("page-item").
						Body(
							app.A().
								Class("page-link page-link-add").Aria("label", "Previous").
								Href(locHash).
								Body(
									app.Span().Aria("hidden", "true").Text("««"),
								),
						).OnClick(func(ctx app.Context, e app.Event) {
						l.minersPage = 0
						l.fetchMiners(ctx)
					}),
					app.Range(pages).Slice(func(i int) app.UI {
						link := app.A().
							Class("page-link").
							Href(fmt.Sprintf("#home/page/%d", pages[i])).
							Body(
								app.Span().Text(pages[i]),
							)
						if pages[i] == -1 {
							link = link.Aria("disabled", "true").Text("...")
							link = link.Class("page-link-add")
						}
						ret := app.Li().Style("width", "44px").
							Class("page-item").
							Aria("current", "page")
						if pages[i] == -1 {
							link = link.Class("page-link-add")
							ret = ret.Class("disabled")
						} else if pages[i] == l.minersPage {
							ret = ret.Class("active")
						} else {
							link = link.Class("page-link-add")
							ret = ret.OnClick(func(ctx app.Context, e app.Event) {
								fmt.Println("click pages:", l.getPages())
								l.minersPage = l.getPages()[i]
								l.fetchMiners(ctx)
							})
						}

						ret = ret.Body(link)

						return ret
					}),
					app.Li().Style("width", "44px").
						Class("page-item").
						Body(
							app.A().
								Class("page-link page-link-add").Aria("label", "Next").
								Href(locHash).
								Body(
									app.Span().Aria("hidden", "true").Text("»»"),
								),
						).OnClick(func(ctx app.Context, e app.Event) {
						l.minersPage = l.minersCount/minersRowsCount - 1
						l.fetchMiners(ctx)
					}),
				))
}

func (l *Compo) Render() app.UI {
	if len(l.miners) == 0 {
		return app.Div()
	}
	return app.Div().Body(
		app.Div().Class("jumbotron").Body(
			app.H3().Text("Miners page not implemented"),
		),
		app.Table().
			Class("table dark-theme").Style("padding", "0.5rem 0.5rem").
			Body(
				app.THead().
					Body(
						app.Tr().Style("background", "#36393f").
							Style("z-index", "1").
							Style("position", "sticky").
							Style("top", "50px").Body(
							app.Td().ColSpan(5).
								Body(l.tableNavigation()),
						),
						app.Tr().Body(
							app.Th().Scope("col").Text("name"),
							app.Th().Scope("col").Text("terra-bytes"),
							app.Th().Scope("col").Text("commit"),
							app.Th().Scope("col").Text("factor"),
							app.Th().Scope("col").Text("contributors"),
						),
					),
				app.TBody().Body(
					app.Range(l.miners).Slice(func(i int) app.UI {
						m := l.miners[i]
						return app.Tr().Body(
							app.Th().Scope("row").Text(m.Name),
							app.Td().Text(m.TerraBytes),
							app.Td().Text(m.Commit),
							app.Td().Text(m.Factor),
							app.Td().Text(m.ContributorsNumber),
						)
					}),
				),
			),
	)
}

func fetchMiners(from, to int) ([]datatype.Miner, error) {
	loc := app.Window().Get("location")

	u := url.URL{
		Host: loc.Get("host").String(),
		Path: "/api/miners",
	}
	q := u.Query()
	q.Set("from", strconv.Itoa(from))
	q.Set("to", strconv.Itoa(to))
	u.RawQuery = q.Encode()

	fmt.Println(u.String())

	req, err := http.NewRequest("GET", u.String(), nil)
	req.Header.Add("js.fetch:mode", "cors")
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result []datatype.Miner
	if err := json.Unmarshal(b, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func fetchMinersCount() (int, error) {
	loc := app.Window().Get("location")

	u := url.URL{
		Host: loc.Get("host").String(),
		Path: "/api/miners/count",
	}

	req, err := http.NewRequest("GET", u.String(), nil)
	req.Header.Add("js.fetch:mode", "cors")
	if err != nil {
		return 0, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	result, err := strconv.Atoi(string(b))
	if err != nil {
		return 0, err
	}
	return result, nil
}
