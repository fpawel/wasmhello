package main

import (
	"fmt"
	"github.com/fpawel/wasmhello/internal/ui"
	"log"
	"net/http"

	"github.com/maxence-charriere/go-app/v8/pkg/app"
)

func main() {
	app.Route("/", &ui.App{})
	app.RunWhenOnBrowser()
	http.Handle("/", &app.Handler{
		Name:        "Hello",
		Description: "An Hello World! example",
		RawHeaders: []string{
			ui.HeaderBootstrapCdn,
		},
	})

	const port = 8001

	log.Printf("localhost:%d", port)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		log.Fatal(err)
	}
}
