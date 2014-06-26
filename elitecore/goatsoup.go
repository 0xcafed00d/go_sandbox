package elitecore

import (
	"bytes"
)

type desc_choice [5]string

var desc_list = []desc_choice{
	/* 81 */ {"fabled", "notable", "well known", "famous", "noted"},
	/* 82 */ {"very", "mildly", "most", "reasonably", ""},
	/* 83 */ {"ancient", "\x95", "great", "vast", "pink"},
	/* 84 */ {"\x9E \x9D plantations", "mountains", "\x9C", "\x94 forests", "oceans"},
	/* 85 */ {"shyness", "silliness", "mating traditions", "loathing of \x86", "love for \x86"},
	/* 86 */ {"food blenders", "tourists", "poetry", "discos", "\x8E"},
	/* 87 */ {"talking tree", "crab", "bat", "lobster", "\xB2"},
	/* 88 */ {"beset", "plagued", "ravaged", "cursed", "scourged"},
	/* 89 */ {"\x96 civil war", "\x9B \x98 \x99s", "a \x9B disease", "\x96 earthquakes", "\x96 solar activity"},
	/* 8A */ {"its \x83 \x84", "the \xB1 \x98 \x99", "its inhabitants' \x9A \x85", "\xA1", "its \x8D \x8E"},
	/* 8B */ {"juice", "brandy", "water", "brew", "gargle blasters"},
	/* 8C */ {"\xB2", "\xB1 \x99", "\xB1 \xB2", "\xB1 \x9B", "\x9B \xB2"},
	/* 8D */ {"fabulous", "exotic", "hoopy", "unusual", "exciting"},
	/* 8E */ {"cuisine", "night life", "casinos", "sitcoms", " \xA1 "},
	/* 8F */ {"\xB0", "The planet \xB0", "The world \xB0", "This planet", "This world"},
	/* 90 */ {"n unremarkable", " boring", " dull", " tedious", " revolting"},
	/* 91 */ {"planet", "world", "place", "little planet", "dump"},
	/* 92 */ {"wasp", "moth", "grub", "ant", "\xB2"},
	/* 93 */ {"poet", "arts graduate", "yak", "snail", "slug"},
	/* 94 */ {"tropical", "dense", "rain", "impenetrable", "exuberant"},
	/* 95 */ {"funny", "weird", "unusual", "strange", "peculiar"},
	/* 96 */ {"frequent", "occasional", "unpredictable", "dreadful", "deadly"},
	/* 97 */ {"\x82 \x81 for \x8A", "\x82 \x81 for \x8A and \x8A", "\x88 by \x89", "\x82 \x81 for \x8A but \x88 by \x89", "a\x90 \x91"},
	/* 98 */ {"\x9B", "mountain", "edible", "tree", "spotted"},
	/* 99 */ {"\x9F", "\xA0", "\x87oid", "\x93", "\x92"},
	/* 9A */ {"ancient", "exceptional", "eccentric", "ingrained", "\x95"},
	/* 9B */ {"killer", "deadly", "evil", "lethal", "vicious"},
	/* 9C */ {"parking meters", "dust clouds", "ice bergs", "rock formations", "volcanoes"},
	/* 9D */ {"plant", "tulip", "banana", "corn", "\xB2weed"},
	/* 9E */ {"\xB2", "\xB1 \xB2", "\xB1 \x9B", "inhabitant", "\xB1 \xB2"},
	/* 9F */ {"shrew", "beast", "bison", "snake", "wolf"},
	/* A0 */ {"leopard", "cat", "monkey", "goat", "fish"},
	/* A1 */ {"\x8C \x8B", "\xB1 \x9F \xA2", "its \x8D \xA0 \xA2", "\xA3 \xA4", "\x8C \x8B"},
	/* A2 */ {"meat", "cutlet", "steak", "burgers", "soup"},
	/* A3 */ {"ice", "mud", "Zero-G", "vacuum", "\xB1 ultra"},
	/* A4 */ {"hockey", "cricket", "karate", "polo", "tennis"},
	/* B0 = <planet name> 	B1 = <planet name>ian  B2 = <random name> */
}

func random_name(rng *RNGSeed, digrams string) string {
	var output []byte

	len := int(rng.GenRnd() & 3)

	for n := 0; n <= len; n++ {
		x := int(rng.GenRnd() & 0x3e)

		if digrams[x] != '.' {
			output = append(output, digrams[x])
		}

		if digrams[x+1] != '.' {
			output = append(output, digrams[x+1])
		}
	}

	output = bytes.ToTitle(output)

	return string(output)
}

func ianise_planet(planetName string) string {
	lastCharIndex := len(planetName) - 1
	if (planetName[lastCharIndex] == 'i') || (planetName[lastCharIndex] == 'e') {
		planetName = planetName[:lastCharIndex-1]
	}

	return planetName + "ian"
}

func b2i(b bool) uint32 {
	if b {
		return 1
	} else {
		return 0
	}
}

func goat_soup(source, planetName, digrams string, output *bytes.Buffer, rng *RNGSeed) {

	for len(source) > 0 {
		c := uint32((source[0]) & 0xff)

		if (c & 0x80) == 0 {
			output.WriteByte(byte(c))
		} else {
			if c <= 0xA4 {
				rnd := rng.GenRnd()
				newsource := desc_list[c-0x81][b2i(rnd >= 0x33)+b2i(rnd >= 0x66)+b2i(rnd >= 0x99)+b2i(rnd >= 0xCC)]
				goat_soup(newsource, planetName, digrams, output, rng)
			} else {
				switch c {
				case 0xB0: /* planet name */
					output.WriteString(planetName)
					break

				case 0xB1: /* <planet name>ian */
					output.WriteString(ianise_planet(planetName))
					break

				case 0xB2: /* random name */
					output.WriteString(random_name(rng, digrams))
					break

				default:
					return
				}
			}
		}
	}
}

func generateGoatSoupString(seed *RNGSeed, planetName, digrams string) string {
	var output bytes.Buffer
	goat_soup("\x8F is \x97.", planetName, digrams, &output, seed)
	return output.String()
}

var inhabitant_desc1 = []string{"Large ", "Fierce ", "Small "}
var inhabitant_desc2 = []string{"Green ", "Red ", "Yellow ", "Blue ", "Black ", "Harmless "}
var inhabitant_desc3 = []string{"Slimy ", "Bug-Eyed ", "Horned ", "Bony ", "Fat ", "Furry "}
var inhabitant_desc4 = []string{"Rodents", "Frogs", "Lizards", "Lobsters", "Birds", "Humanoids", "Felines", "Insects"}

func generateInhabitantDescription(seed PlanetSeed) string {
	var output string

	//	a := uint8(seed.w0 & 0xff)
	b := uint8((seed.w0 >> 8) & 0xff)
	//	c := uint8(seed.w1 & 0xff)
	d := uint8((seed.w1 >> 8) & 0xff)
	e := uint8(seed.w2 & 0xff)
	f := uint8((seed.w2 >> 8) & 0xff)

	if e < 128 {
		return "Human Colonials"
	}

	inhab := (f / 4) & 7
	if inhab < 3 {
		output += inhabitant_desc1[inhab]
	}

	inhab = f / 32
	if inhab < 6 {
		output += inhabitant_desc2[inhab]
	}

	inhab = (d ^ b) & 7
	if inhab < 6 {
		output += inhabitant_desc3[inhab]
	}

	inhab = (inhab + (f & 3)) & 7
	output += inhabitant_desc4[inhab]

	return output
}
