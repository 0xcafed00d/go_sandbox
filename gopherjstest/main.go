package main

import (
	"fmt"
	"time"

	"github.com/gopherjs/gopherjs/js"
	"honnef.co/go/js/dom"
)

type Canvas struct {
	ctx           *dom.CanvasRenderingContext2D
	animateFunc   func(t time.Duration)
	animateEnable bool
}

func MakeCanvas(id string) *Canvas {
	c := Canvas{}

	doc := dom.GetWindow().Document()
	canvas := doc.GetElementByID(id).(*dom.HTMLCanvasElement)
	c.ctx = canvas.GetContext2d()
	return &c
}

func (c *Canvas) SetAnimateFunc(f func(t time.Duration)) {
	c.animateFunc = f
}

func (c *Canvas) doAnimate(t time.Duration) {
	if c.animateFunc != nil && c.animateEnable {
		c.animateFunc(t)
		dom.GetWindow().RequestAnimationFrame(c.doAnimate)
	}
}

func (c *Canvas) Animate(animate bool) {
	c.animateEnable = animate
	dom.GetWindow().RequestAnimationFrame(c.doAnimate)
}

func main() {

	js.Global.Call("addEventListener", "load", func() {
		doc := dom.GetWindow().Document()
		img := doc.GetElementByID("img_elephant").(*dom.HTMLImageElement)
		canvas := MakeCanvas("canvas")
		println(img)
		var x = 1
		canvas.SetAnimateFunc(func(t time.Duration) {
			canvas.ctx.FillStyle = "black"
			canvas.ctx.FillRect(0, 0, 640, 480)
			canvas.ctx.FillStyle = "red"
			s := fmt.Sprint(t)
			canvas.ctx.FillText(s, x, 100, 300)

			canvas.ctx.Call("drawImage", img.Object, x, 100, 100, 100)

			x++
			if x > 640 {
				x = 0
			}
		})

		canvas.Animate(true)
	})
}
