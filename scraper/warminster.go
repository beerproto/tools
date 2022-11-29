package scraper

import (
	"golang.org/x/text/language"
	"tawesoft.co.uk/go/lxstrconv"
)

var warminster = NewMalting("Warminster Malt", "GRB", lxstrconv.NewDecimalFormat(language.BritishEnglish),
	WithBaseURL("https://www.warminster-malt.co.uk/warminster-base-malts/"),
	//WithProductSelector(".malt-post a[href]"),
	WithGrainGroupSelector(".entry-content .et_pb_section"),
	WithTitleSelector("h2"),
	WithProductRowSelector(".et_section_regular .et_pb_row .et_pb_column"),
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
