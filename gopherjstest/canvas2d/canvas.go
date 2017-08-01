package canvas2d

import (
	"time"

	"github.com/simulatedsimian/go-js-dom"
)

type Canvas struct {
	Ctx           *dom.CanvasRenderingContext2D
	animateFunc   func(t time.Duration)
	animateEnable bool
}

func MakeCanvas(id string) *Canvas {
	c := Canvas{}

	doc := dom.GetWindow().Document()
	canvas := doc.GetElementByID(id).(*dom.HTMLCanvasElement)
	c.Ctx = canvas.GetContext2d()
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
