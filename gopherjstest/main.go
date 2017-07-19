package main

import (
	"fmt"
	"time"

	"github.com/gopherjs/gopherjs/js"
	"honnef.co/go/js/dom"
)

func main() {

	js.Global.Call("addEventListener", "load", func() {
		doc := dom.GetWindow().Document()
		canvas := doc.GetElementByID("canvas").(*dom.HTMLCanvasElement)
		ctx := canvas.GetContext2d()
		img := doc.GetElementByID("img_elephant").(*dom.HTMLImageElement)

		println(img)
		var x = 1
		var aniFrm func(t time.Duration)
		aniFrm = func(t time.Duration) {
			ctx.FillStyle = "black"
			ctx.FillRect(0, 0, 640, 480)
			ctx.FillStyle = "red"
			s := fmt.Sprint(t)
			ctx.FillText(s, x, 100, 300)

			ctx.Call("drawImage", img.Object, x, 100, 100, 100)

			x++
			if x > 640 {
				x = 0
			}
			dom.GetWindow().RequestAnimationFrame(aniFrm)
		}
		dom.GetWindow().RequestAnimationFrame(aniFrm)
	})
}
