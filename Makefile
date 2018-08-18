wasm:
	GOARCH=wasm GOOS=js go build -o example.wasm .

build-server:
	go build -o server-app server/server.go

run: build-server wasm
	./server-app