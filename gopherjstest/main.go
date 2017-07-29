package main

import (
	"fmt"
	"time"

	"github.com/gopherjs/gopherjs/js"
	"honnef.co/go/js/dom"
)

type Canvas struct {
	ctx     *dom.CanvasRenderingContext2D
	Animate func(t time.Duration)
}

func MakeCanvas(id string) *Canvas {
	c := Canvas{}

	doc := dom.GetWindow().Document()
	canvas := doc.GetElementByID(id).(*dom.HTMLCanvasElement)
	c.ctx = canvas.GetContext2d()
	return &c
}

func main() {

	js.Global.Call("addEventListener", "load", func() {
		doc := dom.GetWindow().Document()
		img := doc.GetElementByID("img_elephant").(*dom.HTMLImageElement)
		canvas := MakeCanvas("canvas")
		println(img)
		var x = 1
		var aniFrm func(t time.Duration)
		aniFrm = func(t time.Duration) {
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
			dom.GetWindow().RequestAnimationFrame(aniFrm)
		}
		dom.GetWindow().RequestAnimationFrame(aniFrm)
	})
}
