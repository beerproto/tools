package scraper

import (
	"strings"

	beerproto "github.com/beerproto/beerproto_go"
	"github.com/beerproto/beerproto_go/fermentables"
	"github.com/beerproto/tools/unit"
	colly "github.com/gocolly/colly/v2"
	"golang.org/x/text/language"
	"tawesoft.co.uk/go/lxstrconv"
)

type BelgoMalt struct {
	formatter lxstrconv.NumberFormat
}

func NewBelgoMalt() *BelgoMalt {
	return &BelgoMalt{
		formatter: lxstrconv.NewDecimalFormat(language.German),
	}
}

func (s *BelgoMalt) Parse() []*fermentables.GrainType {
	grains := []*fermentables.GrainType{}

	c := colly.NewCollector()
	page := colly.NewCollector()

	page.OnHTML("body", func(e *colly.HTMLElement) {
		grain := &fermentables.GrainType{
			Country:    "BEL",
			Standard:   fermentables.GrainType_EBC,
			GrainGroup: beerproto.GrainGroup_GRAIN_GROUP_BASE,
			Producer:   "Belgomalt",
			Name:       e.ChildText("h1.product-title"),
		}

		if grain.Name == "" {
			return
		}

		e.ForEach(".woocommerce-tabs tr", func(_ int, el *colly.HTMLElement) {
			text := el.ChildText("td:nth-child(2)")
			if text == "" {
				text = el.ChildText("td")
			}
			header := strings.TrimSpace(el.ChildText("th:first-child"))
			if header == "" {
				header = strings.TrimSpace(el.ChildText("td:first-child"))

			}
			switch header {
			case "EBC Colour":
				text = strings.Trim(strings.ToLower(text), "between")
				grain.Color = unit.Color(text, unit.WithFormatter[beerproto.ColorUnit](s.formatter))
			case "Moisture":
				grain.Moisture = unit.Percent(text, unit.WithFormatter[beerproto.PercentUnit](s.formatter))
			case "Total Protein dry basis":
				grain.Protein = unit.Percent(text, unit.WithFormatter[beerproto.PercentUnit](s.formatter))
			case "Total Nitrogen, Dry":
				grain.TotalNitrogen = unit.Percent(text, unit.WithFormatter[beerproto.PercentUnit](s.formatter))
			case "Soluble Protein":
				text = strings.Trim(strings.ToLower(text), "between")
				grain.SolubleProtein = unit.Percent(text, unit.WithFormatter[beerproto.PercentUnit](s.formatter))
			case "IOB FAN in Wort":
				grain.Fan = unit.Concentration(text, unit.WithFormatter[beerproto.ConcentrationUnit](s.formatter))
			case "Diastatic Power":
				text = strings.Trim(strings.ToLower(text), "minimum")
				grain.DiastaticPower = unit.DiastaticPower(text, unit.WithFormatter[beerproto.DiastaticPowerUnit](s.formatter))
			case "EBC Wort pH":
				text = strings.Trim(strings.ToLower(text), "between")
				grain.DiPh = unit.Acidity(text, unit.WithFormatter[beerproto.AcidityUnit](s.formatter)).Maximum
			case "Alpha Amylase dry basis":
				grain.AlphaAmylase = unit.Time(text, unit.WithFormatter[beerproto.TimeUnit](s.formatter)).Maximum
			case "EBC B-Glucan in Wort":
				grain.BetaGlucan = unit.Concentration(text, unit.WithUnit(beerproto.ConcentrationUnit_CONCENTRATION_UNIT_MGL),
					unit.WithFormatter[beerproto.ConcentrationUnit](s.formatter))
			case "Friability":
				text = strings.Trim(strings.ToLower(text), "minimum")
				grain.Friability = unit.Percent(text, unit.WithFormatter[beerproto.PercentUnit](s.formatter))
			case "EBC Extract 0.2mm dry basis":
				text = strings.Trim(strings.ToLower(text), "minimum")
				grain.Yield = unit.Percent(text, unit.WithFormatter[beerproto.PercentUnit](s.formatter))

			}

		})

		grains = append(grains, grain)
	})

	c.OnHTML(".section-content .col a[href]", func(e *colly.HTMLElement) {
		page.Visit(e.Attr("href"))
	})

	c.Visit("https://belgomalt.be/our-products/")
	//page.Visit("https://belgomalt.be/our-products/no-ox/")

	return grains
}
