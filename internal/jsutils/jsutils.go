package jsutils

import "github.com/maxence-charriere/go-app/v8/pkg/app"

func LocationHash() string {
	return app.Window().Get("location").Get("hash").String()
}
