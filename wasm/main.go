package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"syscall/js"
	"time"
	"wasm-test/wasm/window"
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

var dragging = false

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

	ball := newBall(window.GetElementById("ball"), 25)

	window.AddOnMouseMove(
		window.Document(),
		func(values []js.Value) {
			go func() {
				if dragging {
					currentX = values[0].Get("clientX").Float()
					currentY = window.Window().Get("innerHeight").Float() - values[0].Get("clientY").Float()

					ball.MomentumX = 0
					ball.MomentumY = 0

					ball.move(currentX-ball.Radius, currentY-ball.Radius)
				}
			}()
		},
	)

	go eventLoop(ball)

	select {}
	println("Web Assembly stopped.")
}

const AccelerationConstant = 1
const RestitutionDamperX = 2.5
const RestitutionDamperY = 4

type Ball struct {
	Element   *js.Value
	Radius    float64
	X         float64
	Y         float64
	MomentumX float64
	MomentumY float64
}

func newBall(ballElement js.Value, radius float64) (ball *Ball) {
	return &Ball{Element: &ballElement, Radius: radius, X: 30, Y: 30}
}

func (b *Ball) move(x float64, y float64) {
	b.X = x
	b.Y = y

	b.correctMotion()

	b.Element.Call("setAttribute", "style", fmt.Sprintf("width: %vpx; height: %vpx; left: %vpx; bottom: %vpx;", b.Radius*2, b.Radius*2, b.X, b.Y))
}

func (b *Ball) correctMotion() {
	innerWidth := window.Window().Get("innerWidth").Float() - b.Radius*2
	innerHeight := window.Window().Get("innerHeight").Float() - b.Radius*2

	if b.X < 0 {
		b.X = 0
		b.MomentumX *= -1
		if b.MomentumX > 0 {
			b.MomentumX -= RestitutionDamperX
		} else if b.MomentumX < 0 {
			b.MomentumX += RestitutionDamperX
		}
	} else if b.X > innerWidth {
		b.X = innerWidth
		b.MomentumX *= -1
		if b.MomentumX > 0 {
			b.MomentumX -= RestitutionDamperX
		} else if b.MomentumX < 0 {
			b.MomentumX += RestitutionDamperX
		}
	}

	if b.Y < 0 {
		b.Y = 0
		b.MomentumY *= -1

		if b.MomentumY > 0 {
			b.MomentumY -= RestitutionDamperY
		} else if b.MomentumY < 0 {
			b.MomentumY += RestitutionDamperY
		}
	} else if b.Y > innerHeight {
		b.Y = innerHeight
		b.MomentumY *= -1

		if b.MomentumY > 0 {
			b.MomentumY -= RestitutionDamperY
		} else if b.MomentumY < 0 {
			b.MomentumY += RestitutionDamperY
		}
	}
}

func (b *Ball) ApplyGravity() {
	b.MomentumY -= AccelerationConstant
}

var currentX = float64(0)
var currentY = float64(0)
var oldX = float64(0)
var oldY = float64(0)

func (b *Ball) SetThrowMomentum() {
	b.MomentumX = currentX - oldX
	b.MomentumY = currentY - oldY

	oldX = currentX
	oldY = currentY
}

func eventLoop(ball *Ball) {
	shouldContinue := true

	for shouldContinue {
		if false {
			shouldContinue = false
		}

		if !dragging {
			ball.ApplyGravity()

			ball.move(
				ball.X+ball.MomentumX,
				ball.Y+ball.MomentumY,
			)
		} else {
			ball.SetThrowMomentum()
		}

		time.Sleep(10 * time.Millisecond)
	}
}
