package server

import (
	"encoding/json"
	"fmt"
	"github.com/fpawel/wasmhello/internal/server/datatype"
	"github.com/google/uuid"
	"io"
	"log"
	"net/http"
	"sync"
)

var (
	muAccounts sync.Mutex
	accounts   = map[string]userData{
		"1": {"1", "1"},
	}
)

type userData struct {
	Password string
	Token    string
}

func registerAccount(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Unsupported", http.StatusBadRequest)
		return
	}
	// Read body
	b, err := io.ReadAll(r.Body)
	defer func() {
		_ = r.Body.Close()
	}()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var input struct {
		Account string `json:"account"`
	}
	if err := json.Unmarshal(b, &input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	muAccounts.Lock()
	defer muAccounts.Unlock()
	if _, f := accounts[input.Account]; f {
		http.Error(w, "user already exists", http.StatusBadRequest)
		return
	}
	pass := uuid.New().String()
	accounts[input.Account] = userData{
		Password: pass,
	}

	fmt.Println("register:", input.Account, pass)

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	if _, err := w.Write([]byte(pass)); err != nil {
		log.Println("ERROR:", err)
	}

}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Unsupported", http.StatusBadRequest)
		return
	}
	// Read body
	b, err := io.ReadAll(r.Body)
	defer func() {
		_ = r.Body.Close()
	}()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var input datatype.UserPass
	if err := json.Unmarshal(b, &input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	muAccounts.Lock()
	defer muAccounts.Unlock()
	u, f := accounts[input.User]
	if !f {
		http.Error(w, "user doesn't exists", http.StatusBadRequest)
		return
	}
	if u.Password != input.Pass {
		http.Error(w, "Access denied", http.StatusForbidden)
		return
	}

	tok, err := datatype.JwtTokenize(input)
	if err != nil {
		http.Error(w, "Access denied", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	if _, err := w.Write([]byte(tok)); err != nil {
		log.Println("ERROR:", err)
	}
}
