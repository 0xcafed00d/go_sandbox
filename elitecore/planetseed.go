package elitecore

type PlanetSeed struct {
	w0, w1, w2 uint16
}

func NewPlanetSeed(w0, w1, w2 uint16) PlanetSeed {
	return PlanetSeed{w0, w1, w2}
}

/* rotate 8 bit number leftwards */
func rotatel(x uint16) uint16 {
	temp := x & 128
	return (2 * (x & 127)) + (temp >> 7)
}

func twist(x uint16) uint16 {
	return (256 * rotatel(x>>8)) + rotatel(x&255)
}

/* Apply to base seed; once for galaxy 2  */
/* twice for galaxy 3, etc. */
/* Eighth application gives galaxy 1 again*/

func nextGalaxy(s PlanetSeed) PlanetSeed {
	return PlanetSeed{twist(s.w0), twist(s.w1), twist(s.w2)}
}

func tweakseed(s PlanetSeed) PlanetSeed {
	temp := s.w0 + s.w1 + s.w2
	return PlanetSeed{s.w1, s.w2, temp}
}
