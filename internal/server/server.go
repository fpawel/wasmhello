package server

import (
	"fmt"
	"github.com/fpawel/wasmhello/internal/ui"
	"github.com/maxence-charriere/go-app/v8/pkg/app"
	"log"
	"net/http"
)

func Run(port string) {
	app.Route("/", ui.New())
	app.RunWhenOnBrowser()
	http.Handle("/", &app.Handler{
		Name:        "Hello",
		Description: "An Hello World! example",
		Styles: []string{
			"/web/css/main.css",
		},
		Icon: app.Icon{
			Default: "/web/img/logo.svg",
		},
		RawHeaders: []string{bootstrapCDNs},
	})

	http.HandleFunc("/api/register", registerAccount)
	http.HandleFunc("/api/login", login)
	http.HandleFunc("/api/miners", serveMiners)
	http.HandleFunc("/api/miners/count", serveMinersCount)

	log.Printf("localhost:%s", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		log.Fatal(err)
	}
}

const (
	bootstrapCDNs = `
<link href="https://cdn.jsdelivr.net/npm/bootstrap@5.0.2/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-EVSTQN3/azprG1Anm3QDgpJLIm9Nao0Yz1ztcQTwFspd3yD65VohhpuuCOmLASjC" crossorigin="anonymous">
<script src="https://cdn.jsdelivr.net/npm/bootstrap@5.0.2/dist/js/bootstrap.bundle.min.js" integrity="sha384-MrcW6ZMFYlzcLA8Nl+NtUVF0sA7MsXsP1UyJoMp4YLEuNSfAP+JcXn/tWtIaxVXM" crossorigin="anonymous"></script>
<link rel="stylesheet" href="https://pro.fontawesome.com/releases/v5.10.0/css/all.css" integrity="sha384-AYmEC3Yw5cVb3ZcuHtOA93w35dYTsvhLPVnYs9eStHfGJvOvKxVfELGroGkvsg+p" crossorigin="anonymous"/>
<script
  src="https://code.jquery.com/jquery-3.6.0.slim.min.js"
  integrity="sha256-u7e5khyithlIdTpu22PHhENmPcRdFiHRjhAuHcs05RI="
  crossorigin="anonymous"></script>
`
)
