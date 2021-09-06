BUILD_DIR=./build
APP_BINARY=$(BUILD_DIR)/app
SRC_DIR=./cmd/app
MAIN_SRC=$(SRC_DIR)/main.go

clean:
	rm -fr $(BUILD_DIR) || true
	mkdir $(BUILD_DIR)
	cp -r ./web $(BUILD_DIR)

build-client: clean
	GOARCH=wasm GOOS=js GOARCH=wasm GOOS=js $(GOCMD) build -o $(BUILD_DIR)/web/app.wasm $(MAIN_SRC)

build-server: build-client
	$(GOCMD) build -o $(APP_BINARY) $(MAIN_SRC)

run: build-server
	cd $(BUILD_DIR)	&& ./app $(APP_PORT)