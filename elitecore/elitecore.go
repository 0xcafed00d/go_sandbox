package elitecore

const (
	NumGalaxies         = 8
	NumGalaxiesMask     = 0x07
	NumSystemsPerGalaxy = 256
)

var (
	allDigrams        = "abouseitiletstonlonuthno..lexegezacebisousesarmaindirea.eratenberalavetiedorquanteisrion" /* Dots should be nullprint characters */
	planetNameDigrams = allDigrams[24:]

	govnames = []string{
		"Anarchy",
		"Feudal",
		"Multi-gov",
		"Dictatorship",
		"Communist",
		"Confederacy",
		"Democracy",
		"Corporate State",
	}

	econnames = []string{
		"Rich Industrial",
		"Average Industrial",
		"Poor Industrial",
		"Mainly Industrial",
		"Mainly Agricultural",
		"Rich Agricultural",
		"Average Agricultural",
		"Poor Agricultural",
	}
)

type MarketUnitType int

const (
	UNIT_T  MarketUnitType = 0
	UNIT_KG MarketUnitType = 1
	UNIT_G  MarketUnitType = 2
)

type tradegood struct {
	baseprice uint32
	gradient  int32
	basequant uint32
	maskbyte  uint32
	units     MarketUnitType
	name      string
}

var commodities = []tradegood{
	{0x13, -0x02, 0x06, 0x01, UNIT_T, "Food"},
	{0x14, -0x01, 0x0A, 0x03, UNIT_T, "Textiles"},
	{0x41, -0x03, 0x02, 0x07, UNIT_T, "Radioactives"},
	{0x28, -0x05, 0xE2, 0x1F, UNIT_T, "Slaves"},
	{0x53, -0x05, 0xFB, 0x0F, UNIT_T, "Liquor/Wines"},
	{0xC4, +0x08, 0x36, 0x03, UNIT_T, "Luxuries"},
	{0xEB, +0x1D, 0x08, 0x78, UNIT_T, "Narcotics"},
	{0x9A, +0x0E, 0x38, 0x03, UNIT_T, "Computers"},
	{0x75, +0x06, 0x28, 0x07, UNIT_T, "Machinery"},
	{0x4E, +0x01, 0x11, 0x1F, UNIT_T, "Alloys"},
	{0x7C, +0x0d, 0x1D, 0x07, UNIT_T, "Firearms"},
	{0xB0, -0x09, 0xDC, 0x3F, UNIT_T, "Furs"},
	{0x20, -0x01, 0x35, 0x03, UNIT_T, "Minerals"},
	{0x61, -0x01, 0x42, 0x07, UNIT_KG, "Gold"},
	{0xAB, -0x02, 0x37, 0x1F, UNIT_KG, "Platinum"},
	{0x2D, -0x01, 0xFA, 0x0F, UNIT_G, "Gem-Stones"},
	{0x35, +0x0F, 0xC0, 0x07, UNIT_T, "Alien Items"},
}

type SystemInfo struct {
	X              int
	Y              int
	Economy        int
	Govtype        int
	Techlev        int
	Population     int
	Productivity   int
	Radius         int
	Name           string
	Description    string
	InhabitantDesc string
}

type MarketItem struct {
	itemID   int
	price    int
	quantity int
	unit     MarketUnitType
	name     string
}

type Galaxy struct {
	Systems []SystemInfo
}

func GenerateGalaxy(num int) Galaxy {

}
