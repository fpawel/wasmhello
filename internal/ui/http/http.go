package http

import (
	"github.com/go-resty/resty/v2"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"net/http"
)

var (
	loc = app.Window().Get("location")
	C   = resty.New().
		SetTransport(http.DefaultTransport).
		SetBaseURL(loc.Get("protocol").String() + "//" + loc.Get("host").String())
)
