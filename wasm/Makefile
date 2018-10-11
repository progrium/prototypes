.PHONY: dev

build:
	GOARCH=wasm GOOS=js go build -o static/app.wasm .

dev:
	go run dev/main.go