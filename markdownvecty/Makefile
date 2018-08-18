deps:
	go get ./...

wasm: deps
	GOARCH=wasm GOOS=js go build -o example.wasm markdown.go

build-server:
	go build -o server-app ../server/server.go

run: build-server wasm
	./server-app