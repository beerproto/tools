package scraper

import (
	"golang.org/x/text/language"
	"tawesoft.co.uk/go/lxstrconv"
)

var simpsonsmalt = NewMalting("Simpsons Malt", "GRB", lxstrconv.NewDecimalFormat(language.BritishEnglish),
	WithBaseURL("https://www.simpsonsmalt.co.uk/our-malts/"),
	WithProductSelector(".malt-post a[href]"),
	WithTitleSelector("h1.malt__title"),
	WithProductRowSelector("#tab-3 table.spec tr"),
	WithProductHeaderSelector("td.spec__label"),
	WithProductValueMinSelector("td:nth-child(2)"),
	WithProductValueMaxSelector("td:nth-child(3)"),
	WithMoisture("Moisture %"),
	WithColor("Colour ° EBC"),
	WithYield("Extract % d.b"),
	WithProtein("Protein Total"),
	WithTotalNitrogen("Total Soluble Nitrogen %"),
	WithKolbachIndex("Kolbach Index"),
	WithDiastaticPower("Diastatic Power DPWK"),
)
