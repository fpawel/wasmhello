GOCMD=/home/fpawel/sdk/go1.16.5/bin/go

build-client:
	GOARCH=wasm GOOS=js GOARCH=wasm GOOS=js $(GOCMD) build -o web/app.wasm

build:
	GOARCH=wasm GOOS=js GOARCH=wasm GOOS=js $(GOCMD) build -o web/app.wasm
	$(GOCMD) build

run: build
	./wasmhello