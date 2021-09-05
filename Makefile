GOCMD=/home/fpawel/sdk/go1.16.5/bin/go

client:
	GOARCH=wasm GOOS=js GOARCH=wasm GOOS=js $(GOCMD) build -o web/app.wasm

server:
	GOARCH=wasm GOOS=js GOARCH=wasm GOOS=js $(GOCMD) build -o web/app.wasm
	$(GOCMD) build

run: client server
	./wasmhello 8001