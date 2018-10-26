package main

import (
	"io/ioutil"
	"net/http"
	"syscall/js"
	"time"
	"wasm-ball-simulator/wasm/models"
	"wasm-ball-simulator/wasm/window"
)

var dragging = false

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
	window.AddOnMouseDown(
		window.Document(),
		func(values []js.Value) {
			go func() {
				if !dragging {
					dragging = true
				} else {
					panic("Cannot start dragging, already was dragging!")
				}
			}()
		},
	)
	window.AddOnMouseUp(
		window.Document(),
		func(values []js.Value) {
			go func() {
				if dragging {
					dragging = false
				} else {
					panic("Cannot stop dragging, already was stopped!")
				}
			}()
		},
	)

	ball := models.NewBall(25, 300, 300)

	window.AddOnMouseMove(
		window.Document(),
		func(values []js.Value) {
			go func() {
				if dragging {
					ball.Drag(values[0])
				}
			}()
		},
	)

	go eventLoop(ball)

	select {}
	println("Web Assembly stopped.")
}

func eventLoop(ball *models.Ball) {
	for {
		if false { // This is where an exit condition would be set, but there is none right now.
			break
		}

		if !dragging {
			ball.ApplyGravity()

			ball.Move(
				ball.X+ball.MomentumX,
				ball.Y+ball.MomentumY,
			)
		} else {
			ball.SetThrowMomentum()
		}

		time.Sleep(10 * time.Millisecond)
	}
}
