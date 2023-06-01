package route

import (
	"github.com/fpawel/wasmhello/internal/js"
	"strings"
)

type R string

const (
	Home    = "#home"
	Profile = "#profile"
	RegAcc  = "#regacc"
	Login   = "#login"
	Logout  = "#logout"
)

func Base() R {
	locHash := js.LocationHash()
	parts := strings.Split(locHash, "/")
	if len(parts) == 0 || parts[0] == "#" {
		return Home
	}
	return R(parts[0])
}
