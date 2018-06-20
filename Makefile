build-server:
	go build -o server-app server/server.go

run: build-server
	./server-app

wasm:
	GOROOT=~/gowasm GOARCH=wasm GOOS=js ~/gowasm/bin/go build -o example.wasm .
	