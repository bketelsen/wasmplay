wasm2:
	GOROOT=~/gowasm GOARCH=wasm GOOS=js ~/gowasm/bin/go build -o example.wasm markdown.go

wasm:
	GOROOT=~/gowasm GOARCH=wasm GOOS=js ~/gowasm/bin/go build -o example.wasm .

build-server:
	go build -o server-app server/server.go

run: build-server wasm
	./server-app