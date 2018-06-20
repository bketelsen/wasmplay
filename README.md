# WASM Go Web "Framework"

## Running
* Install a WASM enabled version of Go (tip is good)
* Makefile assumes this version of go is at ~/gowasm
* Build server/server.go
```
    go build -o server-app server/server.go
```
* Run server-app
* navigate to http://127.0.0.1:3000

## Unfinished things

* Render() interface signature - string, string + error, error?
* Events
* Callbacks
* Is it worth embedding component?

