# Go Wasm Ball Simulator

A simple simulator of a ball bouncing around your screen that is affected by gravity, restitution, and friction. It can be dragged and flung around via mouse.

## Structure

```
bin/
    [server binary]
serve/
    bin/
        [wasm binary]
        [wasm_exec.js]
    default.html
    index.html
    stles.css
    wasm-init.js
server/
    main.go
wasm/
    models/
        ball.go
    window/
        window.go
    main.go
```

`bin/` stores the server binary.

`serve/` is the root directory that files will be served from.

`serve/bin/` stores the Wasm binary and the `wasm_exec.js` file.

`serve/index.html` is the landing page. Displays a loader until Wasm is ready.

`serve/default.html` is the html that will be swapped out as soon as web assembly initializes.

`serve/wasm-init.js` is the JavaScript glue needed to connect to Go's Wasm.

`server/` contains the source code for the file server.

`wasm/` contains the source code for the web assembly portion of this project.

`wasm/window/` contains a set of helper functions for interacting with the DOM.

## Compilation and Running

Compile the server:

```bash
go build -o "./bin/server" "./server"
```

Compile the Wasm.

```bash
GOARCH=wasm GOOS=js go build -o "serve/bin/out.wasm" "./wasm"
```

Run the serve to download `wasm_exec.js` into `serve/bin` and to begin serving files. Set your own value for `port`

```bash
PORT={port} "./bin/server"
```

The server will begin downloading the latest `wasm_exec.js` and once complete it will be ready to serve the project files.

You can now access the project at `http://localhost:{port}`.

### NOTE

Go's Wasm port is experimental and is constantly changing. One thing I have noticed is that the version of the `wasm_exec.js` file being downloaded today (10/25/18) seems to be broken. Upon checking my local storage I found an older working copy, but I can see some changs have been made to that file since then, so it is actively being worked on. If this doesn't work today, wait a while and they will update the exec file.