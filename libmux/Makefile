
build:
	go build -o libmux.dylib -buildmode=c-shared ./libmux.go
	cd node && ./node_modules/.bin/tsc --lib es2015 libmux.ts