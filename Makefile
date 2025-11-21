BIN_DIR=./bin
BIN_FILE=$(BIN_DIR)/$(APP)

.PHONY: run build docker docker-run clean

APP=donfra-api

run:
	go run ./cmd/$(APP)

build:
	mkdir -p $(BIN_DIR)
	go build -o $(BIN_FILE) ./cmd/$(APP)

docker:
	docker build -t $(APP):dev .

docker-run:
	docker run --rm -p 8080:8080 \
	 -e CORS_ORIGIN=http://localhost:3000 \
	 -e PASSCODE=19930115 \
	 $(APP):dev

clean:
	rm -f $(BIN_FILE)
	go clean -cache -testcache
