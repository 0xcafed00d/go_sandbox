package geom

func minf(x, y float64) float64 {
	if x < y {
		return x
	}
	return y
}

func maxf(x, y float64) float64 {
	if x > y {
		return x
	}
	return y
}

type Coordf struct {
	X, Y float64
}

type Rectanglef struct {
	Min, Max Coordf
}

func Addf(lhs, rhs Coordf) Coordf {
	return Coordf{lhs.X + rhs.X, lhs.Y + rhs.Y}
}

func (c *Coordf) Add(rhs Coordf) {
	c.X += rhs.X
	c.Y += rhs.Y
}

func Subf(lhs, rhs Coordf) Coordf {
	return Coordf{lhs.X - rhs.X, lhs.Y - rhs.Y}
}

func (c *Coordf) Sub(rhs Coordf) {
	c.X -= rhs.X
	c.Y -= rhs.Y
}

func (r *Rectanglef) Normalise() {
	if r.Min.X > r.Max.X {
		r.Min.X, r.Max.X = r.Max.X, r.Min.X
	}

	if r.Min.Y > r.Max.Y {
		r.Min.Y, r.Max.Y = r.Max.Y, r.Min.Y
	}
}

func Normalisef(r Rectanglef) Rectanglef {
	r.Normalise()
	return r
}

func (r *Rectanglef) Width() float64 {
	return r.Max.X - r.Min.X
}

func (r *Rectanglef) Height() float64 {
	return r.Max.Y - r.Min.Y
}

func (r *Rectanglef) Size() Coordf {
	return Coordf{r.Width(), r.Height()}
}

func (r *Rectanglef) IsEmpty() bool {
	return (r.Min.X == r.Max.X) || (r.Min.Y == r.Max.Y)
}

func (r *Rectanglef) IsNormal() bool {
	return (r.Min.X <= r.Max.X) && (r.Min.Y <= r.Max.Y)
}

func (r *Rectanglef) Expand(c Coordf) {
	r.Min.Sub(c)
	r.Max.Add(c)
}

func Expandf(r Rectanglef, c Coordf) Rectanglef {
	r.Expand(c)
	return r
}

func (r *Rectanglef) Translate(c Coordf) {
	r.Min.Add(c)
	r.Max.Add(c)
}

func Translatef(r Rectanglef, c Coordf) Rectanglef {
	r.Translate(c)
	return r
}

func PointInRectanglef(r Rectanglef, p Coordf) bool {
	return (r.Min.X <= p.X) && (r.Min.Y <= p.Y) && (p.X < r.Max.X) && (p.Y < r.Max.Y)
}

func RectanglefIntersection(r1, r2 Rectanglef) (intersect Rectanglef, ok bool) {
	intersect = Rectanglef{
		Min: Coordf{maxf(r1.Min.X, r2.Min.X), maxf(r1.Min.Y, r2.Min.Y)},
		Max: Coordf{minf(r1.Max.X, r2.Max.X), minf(r1.Max.Y, r2.Max.Y)},
	}

	ok = intersect.IsNormal()
	return
}

func RectanglefUnion(r1, r2 Rectanglef) Rectanglef {
	return Rectanglef{
		Min: Coordf{minf(r1.Min.X, r2.Min.X), minf(r1.Min.Y, r2.Min.Y)},
		Max: Coordf{maxf(r1.Max.X, r2.Max.X), maxf(r1.Max.Y, r2.Max.Y)},
	}
}

func RectanglefContains(rOuter, rInner Rectanglef) bool {
	return PointInRectanglef(rOuter, rInner.Min) && PointInRectanglef(rOuter, rInner.Max)
}

func RectanglefFromPosSize(pos, size Coordf) Rectanglef {
	return Rectanglef{pos, Addf(pos, size)}
}

func RectanglefFromSize(size Coordf) Rectanglef {
	return Rectanglef{Max: size}
}
