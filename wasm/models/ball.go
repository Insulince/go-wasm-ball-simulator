package models

import (
	"fmt"
	"syscall/js"
	"wasm-ball-simulator/wasm/window"
)

const (
	// Force of gravity. Higher = stronger gravity.
	AccelerationConstant = 2
	// Amount of momentum that should be lost to "friction" when making contact with the left or right walls. Higher = more momentum lost.
	RestitutionDamperX = 4.5
	// Amount of momentum that should be lost to "friction" when making contact with the top or bottom walls. Higher = more momentum lost.
	RestitutionDamperY = 2
)

var (
	CurrentX = float64(0)
	CurrentY = float64(0)
	OldX     = float64(0)
	OldY     = float64(0)
)

type Ball struct {
	Element   *js.Value
	Radius    float64
	X         float64
	Y         float64
	MomentumX float64
	MomentumY float64
}

func NewBall(radius float64, x float64, y float64) (ball *Ball) {
	div := window.Document().Call("createElement", "div")
	div.Set("id", "ball")
	div.Get("classList").Call("add", "ball")
	window.Body().Call("appendChild", div)
	return &Ball{Element: &div, Radius: radius, X: x, Y: y}
}

func (b *Ball) Drag(clickEvent js.Value) {
	CurrentX = clickEvent.Get("clientX").Float()
	CurrentY = window.Window().Get("innerHeight").Float() - clickEvent.Get("clientY").Float()

	b.MomentumX = 0
	b.MomentumY = 0

	b.Move(CurrentX-b.Radius, CurrentY-b.Radius)

}

func (b *Ball) Move(x float64, y float64) {
	b.X = x
	b.Y = y

	b.CorrectMotion()

	b.Element.Call("setAttribute", "style", fmt.Sprintf("width: %vpx; height: %vpx; left: %vpx; bottom: %vpx;", b.Radius*2, b.Radius*2, b.X, b.Y))
}

func (b *Ball) CorrectMotion() {
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

func (b *Ball) SetThrowMomentum() {
	b.MomentumX = CurrentX - OldX
	b.MomentumY = CurrentY - OldY

	OldX = CurrentX
	OldY = CurrentY
}
