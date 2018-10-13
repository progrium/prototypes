.PHONY: dev

build:
	go generate ./...
	GOARCH=wasm GOOS=js go build -o static/app.wasm ./wasm

dev:
	go run dev/main.go