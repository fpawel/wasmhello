package uinfo

import (
	"encoding/json"
	"fmt"
	"github.com/fpawel/wasmhello/internal/server/datatype"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

const (
	tokenKey = "app-hello-wasm-token"
	inputKey = "app-hello-wasm-input"
)

type Input struct {
	Login    datatype.AccPass
	Register datatype.AccPass
}

func Token() (token string) {
	return app.Window().Get("localStorage").Call("getItem", tokenKey).String()
}

func AccPass() datatype.AccPass {
	tok := Token()
	r, err := datatype.JwtParse(tok)
	if err != nil {
		fmt.Println(err)
	}
	return r
}

func LoggedIn() bool {
	fmt.Println("LoggedIn:", AccPass())
	return AccPass().Account != ""
}

func SetToken(token string) {
	app.Window().Get("localStorage").Call("setItem", tokenKey, token)
}

func RemoveToken() {
	app.Window().Get("localStorage").Call("removeItem", tokenKey)
}

func GetInput() (r Input) {
	s := app.Window().Get("localStorage").Call("getItem", inputKey).String()
	if err := json.Unmarshal([]byte(s), &r); err != nil {
		fmt.Println(err)
	}
	return r
}

func SetInput(r Input) {
	b, err := json.Marshal(r)
	if err != nil {
		fmt.Println(err)
		return
	}
	app.Window().Get("localStorage").Call("setItem", inputKey, string(b)).String()
}
