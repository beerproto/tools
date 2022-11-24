package scraper

import (
	"golang.org/x/text/language"
	"tawesoft.co.uk/go/lxstrconv"
)

var gladfieldmalt = NewMalting("gladfieldmalt", "NZ", lxstrconv.NewDecimalFormat(language.English),
	WithGrainGroupURL("https://www.gladfieldmalt.co.nz/our-malts/base-malts"),
	WithGrainGroupURL("https://www.gladfieldmalt.co.nz/our-malts/distill-malts"),
	WithGrainGroupURL("https://www.gladfieldmalt.co.nz/our-malts/specialty-malts"),
	WithGrainGroupURL("https://www.gladfieldmalt.co.nz/our-malts/sub-base-malt"),
	WithTitleSelector(".dmRespColsWrapper h4 span"),
	WithGrainGroupSelector(".dmBody .dmRespRow"),
	WithProductRowSelector(".site_content table tr"),
	WithProductHeaderSelector("td:first-child"),
	WithProductValueMaxSelector("td:nth-child(3)"),
	WithYield("Extract - Fine Dry min%"),
	WithMoisture("Moisture (max) %"),
	WithSaccharification("Saccharification time"),
	WithColor("Wort Colour"),
	WithTotalNitrogen("Total Nitrogen (%)"),
	WithKolbachIndex("Kolbach Index"),
	WithDiPh("pH"),
	WithDiastaticPower("Diastatic Power (WK) min."),
	WithFan("FAN (mg/l) min"),
	WithFriability("Friability (min) %"),
)
