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

type PaulsMalt struct {
	formatter lxstrconv.NumberFormat
}

func NewPaulsMalt() *PaulsMalt {
	return &PaulsMalt{
		formatter: lxstrconv.NewDecimalFormat(language.English),
	}
}

func (s *PaulsMalt) Parse() []*fermentables.GrainType {
	grains := []*fermentables.GrainType{}

	c := colly.NewCollector()
	page := colly.NewCollector()

	page.OnHTML("body", func(e *colly.HTMLElement) {
		grain := &fermentables.GrainType{
			Country:    "GBR",
			Standard:   fermentables.GrainType_EBC,
			GrainGroup: beerproto.GrainGroup_BASE,
			Producer:   "Pauls Malt",
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
			case "IOB Colour in Wort":
				grain.Color = unit.Color(text, unit.WithFormatter[beerproto.ColorUnitType](s.formatter))
			case "Moisture":
				grain.Moisture = unit.Percent(text, unit.WithFormatter[beerproto.PercentType_PercentUnitType](s.formatter))
			case "Total Protein dry basis":
				grain.Protein = unit.Percent(text, unit.WithFormatter[beerproto.PercentType_PercentUnitType](s.formatter))
			case "Total Nitrogen, Dry":
				grain.TotalNitrogen = unit.Percent(text, unit.WithFormatter[beerproto.PercentType_PercentUnitType](s.formatter))
			case "IOB FAN in Wort":
				grain.Fan = unit.Concentration(text, unit.WithFormatter[beerproto.ConcentrationUnitType](s.formatter))
			case "EBC Wort pH":
				grain.DiPh = unit.Acidity(text, unit.WithFormatter[beerproto.AcidityUnitType](s.formatter)).Maximum
			case "Alpha Amylase dry basis":
				grain.AlphaAmylase = unit.Time(text, unit.WithFormatter[beerproto.TimeType_TimeUnitType](s.formatter)).Maximum
			case "IOB B-Glucan in Wort":
				grain.BetaGlucan = unit.Concentration(text, unit.WithUnit(beerproto.ConcentrationUnitType_MGL),
					unit.WithFormatter[beerproto.ConcentrationUnitType](s.formatter))
			case "Friability":
				grain.Friability = unit.Percent(text, unit.WithFormatter[beerproto.PercentType_PercentUnitType](s.formatter))
			case "EBC Extract 0.2mm dry basis":
				grain.Yield = unit.Percent(text, unit.WithFormatter[beerproto.PercentType_PercentUnitType](s.formatter))

			}

		})

		grains = append(grains, grain)
	})

	c.OnHTML(".section-content .col a[href]", func(e *colly.HTMLElement) {
		page.Visit(e.Attr("href"))
	})

	c.Visit("https://paulsmalt.co.uk/our-products/")
	//page.Visit("https://paulsmalt.co.uk/product/keep-flying-maris-otter/")

	return grains
}
