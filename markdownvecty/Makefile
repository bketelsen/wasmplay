deps:
	~/gowasm/bin/go get ./...

wasm: deps
	GOROOT=~/gowasm GOARCH=wasm GOOS=js ~/gowasm/bin/go build -o example.wasm markdown.go

build-server:
	go build -o server-app ../server/server.go

run: build-server wasm
	./server-app