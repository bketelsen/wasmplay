# WASM Go Web "Framework"

This is a proof of concept, compiling Go to web assembly and manipulating the DOM.  It's missing nearly *everything* that would make a framework useful.  I welcome any ideas about how to turn this into something useful.

## Running
* Install a WASM enabled version of Go (tip is good)
* Makefile assumes this version of go is at ~/gowasm
* run `make run` from a shell to build the server, the wasm output, and start the server.
* navigate to http://127.0.0.1:3000

## Unfinished things

* finish port of honnef.co/js/dom -> https://github.com/bketelsen/go-js-dom 
* Render() interface signature - string, string + error, error?
* vdom diff/patch
* Events
* Callbacks
* Is it worth embedding component?


## Bonus Round

```
    make wasm2 && ./setup-app
```
A vecty app running in web assembly.  Requires my fork of vecty: https://github.com/bketelsen/vecty

## License
MIT License

Copyright (c) 2018 Brian Ketelsen

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.


