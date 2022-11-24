package scraper

import (
	"golang.org/x/text/language"
	"tawesoft.co.uk/go/lxstrconv"
)

var greatwesternmalting = NewMalting("Great Western Malting", "USA", lxstrconv.NewDecimalFormat(language.AmericanEnglish),
	WithGrainGroupURL("https://www.greatwesternmalting.com/gwm/base-malt/"),
	WithGrainGroupURL("https://www.greatwesternmalting.com/gwm/crystal-malt/"),
	WithGrainGroupURL("https://www.greatwesternmalting.com/gwm/identity-preserved-malt/"),
	WithGrainGroupURL("https://www.greatwesternmalting.com/gwm/organic-malt/"),
	WithGrainGroupURL("https://www.greatwesternmalting.com/gwm/specialty-malt/"),
	WithGrainGroupURL("https://www.greatwesternmalting.com/gwm/specialty-grain/"),
	WithGrainGroupSelector(".entry-content .one_third a"),
	WithTitleSelector("h1"),
	WithProductRowSelector(".malt-detail"),
	WithProductHeaderSelector("h2"),
	WithProductValueSelector("p"),
	WithColor("ASBC COLOR"),
	WithMoisture("MOISTURE (%)"),
	WithProtein("PROTEIN (%)"),
	WithMaximum("RATE"),
)
