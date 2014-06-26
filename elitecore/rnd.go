package elitecore

type RNGSeed struct {
	a, b, c, d uint8
}

func NewRNGSeed(a, b, c, d uint8) RNGSeed {
	return RNGSeed{a, b, c, d}
}

func (s *RNGSeed) GenRnd() uint32 {

	var a, x uint32

	x = uint32((s.a * 2) & 0xFF)
	a = x + uint32(s.c)
	if s.a > 127 {
		a++
	}

	s.a = uint8(a & 0xFF)

	s.c = uint8(x)
	a = a / 256 /* a = any carry left from above */
	x = uint32(s.b)
	a = (a + x + uint32(s.d)) & 0xFF
	s.b = uint8(a)
	s.d = uint8(x)
	return a
}
