package home

import (
	"encoding/json"
	"fmt"
	"github.com/fpawel/wasmhello/internal/js"
	"github.com/fpawel/wasmhello/internal/server/datatype"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
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

func (l *Compo) OnNav(ctx app.Context) {
	pages := l.getPages()
	fmt.Println("click pages:", pages)

	l.minersPage = 0

	locHash := js.LocationHash()
	parts := strings.Split(locHash, "/")
	if len(parts) > 2 && parts[1] == "page" {
		l.minersPage, _ = strconv.Atoi(parts[2])
	}
	if l.minersPage < 0 {
		l.minersPage = 0
	}
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

type pageNav struct {
	page        int
	currentPage int
	first       bool
	last        bool
}

func (x pageNav) render() app.HTMLLi {
	link := app.A().Class("page-link").Href(fmt.Sprintf("#home/page/%d", x.page))
	span := app.Span()
	item := app.Li().
		Style("width", "44px").
		Class("page-item")
	currentPage := x.page == x.currentPage
	disabled := false
	pageLinkAdd := true

	if x.last {
		span = span.Text("»»")
		if currentPage {
			disabled = true
		}
	} else if x.first {
		span = span.Text("««")
		if currentPage {
			disabled = true
		}
	} else if x.page == -1 {
		span = span.Text("...")
		disabled = true
	} else {
		if currentPage {
			item = item.Class("active").Aria("current", "page")
			pageLinkAdd = false
		}
		span = app.Span().Text(x.page + 1)
	}
	if disabled {
		link = link.Aria("disabled", "true")
		item = item.Class("disabled")
	}
	if pageLinkAdd {
		link = link.Class("page-link-add")
	}

	return item.Body(link.Body(span))

}

func (l *Compo) tableNavigation() app.UI {
	pages := l.getPages()
	fmt.Println("render Pages:", pages)
	return app.Nav().
		Aria("label", "Table navigation").
		Body(
			app.Ul().
				Class("pagination justify-content-end").
				Body(
					pageNav{
						first:       true,
						page:        0,
						currentPage: l.minersPage,
					}.render(),
					app.Range(pages).Slice(func(i int) app.UI {
						return pageNav{
							page:        pages[i],
							currentPage: l.minersPage,
						}.render()
					}),
					pageNav{
						last:        true,
						page:        l.minersCount/minersRowsCount - 1,
						currentPage: l.minersPage,
					}.render(),
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
