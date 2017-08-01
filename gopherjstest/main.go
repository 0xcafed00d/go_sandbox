package main

import (
	"fmt"
	"time"

	"github.com/gopherjs/gopherjs/js"
	"github.com/simulatedsimian/go-js-dom"
	"github.com/simulatedsimian/go_sandbox/gopherjstest/canvas2d"
)

func main() {

	js.Global.Call("addEventListener", "load", func() {
		doc := dom.GetWindow().Document()
		img := doc.GetElementByID("img_elephant").(*dom.HTMLImageElement)
		canvas := canvas2d.MakeCanvas("canvas")
		println(img)
		var x = 1.0
		canvas.SetAnimateFunc(func(t time.Duration) {
			canvas.Ctx.Rotate(float64(x) / 1000.0)
			canvas.Ctx.FillStyle = "black"
			canvas.Ctx.FillRect(0, 0, 640, 480)
			canvas.Ctx.FillStyle = "red"
			s := fmt.Sprint(t)
			canvas.Ctx.FillText(s, int(x), 100, 300)

			canvas.Ctx.DrawImageSection(img.Object, 0, 0, 100, 100, x, 100, 100, 100)

			x++
			if x > 640 {
				x = 0
			}
		})

		canvas.Animate(true)
	})
}
