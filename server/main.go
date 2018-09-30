package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const ServeFrom = "./serve"
const WasmExecURL = "https://raw.githubusercontent.com/golang/go/master/misc/wasm/wasm_exec.js"
const WasmExecFile = ServeFrom + "/bin/wasm_exec.js"

func main() {
	if _, err := os.Stat(WasmExecFile); err != nil {
		log.Printf("Error locating wasm_exec.js: %v\n", err)
		log.Printf("Downloading wasm_exec.js...\n")

		out, err := os.Create(WasmExecFile)
		if err != nil {
			panic(err)
		}

		resp, err := http.Get(WasmExecURL)
		if err != nil {
			panic(err)
		}

		_, err = io.Copy(out, resp.Body)
		if err != nil {
			panic(err)
		}

		out.Close()
		resp.Body.Close()
	}

	fileServer := http.FileServer(http.Dir(ServeFrom))
	http.Handle("/", fileServer)

	fmt.Printf("Listening on port 3000...\n")
	err := http.ListenAndServe(":3000", nil)
	log.Fatalf("Serve error: %v\n", err)
}
