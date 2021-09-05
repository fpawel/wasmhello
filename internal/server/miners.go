package server

import (
	"encoding/json"
	"fmt"
	"github.com/fpawel/wasmhello/internal/server/datatype"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

var miners []datatype.Miner

func init() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 1000; i++ {
		m := datatype.Miner{
			Name:               randStringRunes(10),
			TerraBytes:         fmt.Sprintf("%d", rand.Intn(100)),
			Commit:             randStringRunes(20),
			Factor:             math.Round(rand.Float64() * 1000. / 1000),
			ContributorsNumber: rand.Intn(10),
		}
		miners = append(miners, m)
	}
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func getMiners(q url.Values) ([]datatype.Miner, error) {
	fromStr := q.Get("from")
	if len(fromStr) == 0 {
		fromStr = "0"
	}
	toStr := q.Get("to")
	if len(toStr) == 0 {
		toStr = fmt.Sprintf("%d", len(miners))
	}
	from, err := strconv.Atoi(fromStr)
	if err != nil {
		return nil, fmt.Errorf("'from': %w", err)
	}
	to, err := strconv.Atoi(toStr)
	if err != nil {
		return nil, fmt.Errorf("'to': %w", err)
	}
	if from < 0 || to < 0 || to > len(miners) || from > to {
		return nil, fmt.Errorf("invalid range: 'from' %s 'to' %s", fromStr, toStr)
	}
	return miners[from:to], nil
}

func serveMinersCount(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Unsupported", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	if _, err := w.Write([]byte(strconv.Itoa(len(miners)))); err != nil {
		log.Println("ERROR:", err)
	}
}

func serveMiners(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Unsupported", http.StatusBadRequest)
		return
	}
	// Read body
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	miners, err := getMiners(r.URL.Query())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	b, err = json.Marshal(miners)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(b); err != nil {
		log.Println("ERROR:", err)
	}
}
