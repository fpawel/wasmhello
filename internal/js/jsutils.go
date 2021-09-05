package js

import "github.com/maxence-charriere/go-app/v8/pkg/app"

func LocationHash() string {
	return app.Window().Get("location").Get("hash").String()
}

func LogErr(args ...interface{}) {
	app.Window().Get("console").Call("error", args...)
}
