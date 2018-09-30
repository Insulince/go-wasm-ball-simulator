package main

import (
	"io/ioutil"
	"net/http"
	"syscall/js"
	"wasm-goroutines/wasm/window"
)

func init() {
	renderDefaultHtml()
}

func renderDefaultHtml() {
	defaultHtmlLocation := "/default.html"

	htmlResponse, err := http.Get(defaultHtmlLocation)
	if err != nil {
		panic(err)
	}
	defer htmlResponse.Body.Close()
	defaultHtml, err := ioutil.ReadAll(htmlResponse.Body)
	if err != nil {
		panic(err)
	}
	window.Body().Set("innerHTML", string(defaultHtml))
}

func main() {
	startButton := window.GetElementById("start-button")

	window.AddOnClick(
		startButton,
		func(value []js.Value) {
			go func() {
				// Do stuff.
			}()
		},
	)

	select {}
	println("Web Assembly stopped.")
}
