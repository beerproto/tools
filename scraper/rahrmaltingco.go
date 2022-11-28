package scraper

import (
	"golang.org/x/text/language"
	"tawesoft.co.uk/go/lxstrconv"
)

var rahrmaltingco = NewMalting("Rahr Malting Co", "USA", lxstrconv.NewDecimalFormat(language.AmericanEnglish),
	WithBaseURL("https://rahrmaltingco.com/products-2?pagesize=32"),

	WithProductSelector(".product-grid .product-title a[href]"),
	WithTitleSelector("h1"),
	WithProductRowSelector("table tr"),
	WithProductHeaderSelector("td:first-child"),
	WithProductValueSelector("td:nth-child(2)"),
	WithColor("Color °L"),
	WithMoisture("Moisture % Max"),
	WithProtein("Protein Total"),
	WithDiastaticPower("Diastatic Power"),
	WithMaximum("Usage Rate"),
	WithFineGrind("Extract FG Min"),
)
