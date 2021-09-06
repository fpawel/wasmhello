package js

import (
	"github.com/maxence-charriere/go-app/v8/pkg/app"
	"net/url"
)

func LocationHash() string {
	return app.Window().Get("location").Get("hash").String()
}

func LogErr(args ...interface{}) {
	app.Window().Get("console").Call("error", args...)
}

func URL() url.URL {
	return url.URL{
		Host: app.Window().Get("location").Get("host").String(),
	}
}
