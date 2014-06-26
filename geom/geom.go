package geom

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

type Coord struct {
	X, Y int
}

type Rectangle struct {
	Min, Max Coord
}

func Add(lhs, rhs Coord) Coord {
	return Coord{lhs.X + rhs.X, lhs.Y + rhs.Y}
}

func (c *Coord) Add(rhs Coord) {
	c.X += rhs.X
	c.Y += rhs.Y
}

func Sub(lhs, rhs Coord) Coord {
	return Coord{lhs.X - rhs.X, lhs.Y - rhs.Y}
}

func (c *Coord) Sub(rhs Coord) {
	c.X -= rhs.X
	c.Y -= rhs.Y
}

func (r *Rectangle) Normalise() {
	if r.Min.X > r.Max.X {
		r.Min.X, r.Max.X = r.Max.X, r.Min.X
	}

	if r.Min.Y > r.Max.Y {
		r.Min.Y, r.Max.Y = r.Max.Y, r.Min.Y
	}
}

func Normalise(r Rectangle) Rectangle {
	r.Normalise()
	return r
}

func (r *Rectangle) Width() int {
	return r.Max.X - r.Min.X
}

func (r *Rectangle) Height() int {
	return r.Max.Y - r.Min.Y
}

func (r *Rectangle) Size() Coord {
	return Coord{r.Width(), r.Height()}
}

func (r *Rectangle) IsEmpty() bool {
	return (r.Min.X == r.Max.X) || (r.Min.Y == r.Max.Y)
}

func (r *Rectangle) IsNormal() bool {
	return (r.Min.X <= r.Max.X) && (r.Min.Y <= r.Max.Y)
}

func (r *Rectangle) Expand(c Coord) {
	r.Min.Sub(c)
	r.Max.Add(c)
}

func Expand(r Rectangle, c Coord) Rectangle {
	r.Expand(c)
	return r
}

func (r *Rectangle) Translate(c Coord) {
	r.Min.Add(c)
	r.Max.Add(c)
}

func Translate(r Rectangle, c Coord) Rectangle {
	r.Translate(c)
	return r
}

func PointInRectangle(r Rectangle, p Coord) bool {
	return (r.Min.X <= p.X) && (r.Min.Y <= p.Y) && (p.X < r.Max.X) && (p.Y < r.Max.Y)
}

func RectangleIntersection(r1, r2 Rectangle) (intersect Rectangle, ok bool) {
	intersect = Rectangle{
		Min: Coord{max(r1.Min.X, r2.Min.X), max(r1.Min.Y, r2.Min.Y)},
		Max: Coord{min(r1.Max.X, r2.Max.X), min(r1.Max.Y, r2.Max.Y)},
	}

	ok = intersect.IsNormal()
	return
}

func RectangleUnion(r1, r2 Rectangle) Rectangle {
	return Rectangle{
		Min: Coord{min(r1.Min.X, r2.Min.X), min(r1.Min.Y, r2.Min.Y)},
		Max: Coord{max(r1.Max.X, r2.Max.X), max(r1.Max.Y, r2.Max.Y)},
	}
}

func RectangleContains(rOuter, rInner Rectangle) bool {
	return PointInRectangle(rOuter, rInner.Min) && PointInRectangle(rOuter, rInner.Max)
}

func RectangleFromPosSize(pos, size Coord) Rectangle {
	return Rectangle{pos, Add(pos, size)}
}

func RectangleFromSize(size Coord) Rectangle {
	return Rectangle{Max: size}
}
