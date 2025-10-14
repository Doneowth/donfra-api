APP=donfra-api

run:
	go run ./cmd/$(APP)

build:
	go build -o ./bin/$(APP) ./cmd/$(APP)

docker:
	docker build -t $(APP):dev .

docker-run:
	docker run --rm -p 8080:8080 \
	 -e CORS_ORIGIN=http://localhost:3000 \
	 -e PASSCODE=19930115 \
	 $(APP):dev
