package main

import (
	"github.com/russross/blackfriday/v2"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"testing"
)

func TestMain1(t *testing.T) {
	b, err := ioutil.ReadFile("/home/fpawel/GOPATH/src/github.com/fpawel/wasmhello/readme.md")
	require.NoError(t, err)
	output := blackfriday.Run(b)
	err = ioutil.WriteFile("/home/fpawel/GOPATH/src/github.com/fpawel/wasmhello/build/example.html", output, 0644)
	require.NoError(t, err)
}
