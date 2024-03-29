package scraper

import (
	"fmt"
	"strings"

	beerproto "github.com/beerproto/beerproto_go"
	"github.com/beerproto/beerproto_go/fermentables"
	"github.com/beerproto/tools/unit"
	colly "github.com/gocolly/colly/v2"
	"golang.org/x/text/language"
	"tawesoft.co.uk/go/lxstrconv"
)

type Bestmalz struct {
	formatter lxstrconv.NumberFormat
}

func NewBestmalz() *Bestmalz {
	return &Bestmalz{
		formatter: lxstrconv.NewDecimalFormat(language.German),
	}
}

func (s *Bestmalz) Parse() []*fermentables.GrainType {
	grains := []*fermentables.GrainType{}

	c := colly.NewCollector()
	page := colly.NewCollector()

	page.OnHTML("body", func(e *colly.HTMLElement) {
		grain := &fermentables.GrainType{
			Country:    "DEU",
			Standard:   fermentables.GrainType_EBC,
			GrainGroup: beerproto.GrainGroup_GRAIN_GROUP_BASE,
			Producer:   "Bestmalz",
			Name:       e.ChildText("h1"),
		}

		e.ForEach("table tbody tr", func(_ int, el *colly.HTMLElement) {

			spec := el.DOM.Children().First().Text()
			text := el.DOM.Children().First().NextAll().Text()

			switch strings.ToLower(strings.TrimSpace(spec)) {
			case "moisture content":
				grain.Moisture = unit.Percent(text,
					unit.WithFormatter[beerproto.PercentUnit](s.formatter))
			case "fine-coarse difference ebc":
				grain.CoarseGrind = unit.Percent(text,
					unit.WithFormatter[beerproto.PercentUnit](s.formatter))
			case "viscosity (8.6%)":
				grain.Viscosity = unit.Viscosity(text,
					unit.WithFormatter[beerproto.ViscosityUnit](s.formatter))
			case "friability":
				grain.Friability = unit.Percent(text,
					unit.WithFormatter[beerproto.PercentUnit](s.formatter))
			case "protein, dry basis":
				grain.Protein = unit.Percent(text,
					unit.WithFormatter[beerproto.PercentUnit](s.formatter))
			case "soluble nitrogen":
				grain.SolubleNitrogen = unit.Concentration(text,
					unit.WithFormatter[beerproto.ConcentrationUnit](s.formatter))
			case "kolbach index":
				grain.KolbachIndex = unit.Percent(text,
					unit.WithFormatter[beerproto.PercentUnit](s.formatter))
			case "wort color":
				grain.Color = unit.Color(text,
					unit.WithFormatter[beerproto.ColorUnit](s.formatter))
			case "wort ph":
				unit.Acidity(text,
					unit.WithFormatter[beerproto.AcidityUnit](s.formatter))
			case "diastatic power":
				grain.DiastaticPower = unit.DiastaticPower(text, unit.WithFormatter[beerproto.DiastaticPowerUnit](s.formatter))
			}
		})
	})

	c.OnHTML(".fusion-portfolio-post div a[href]", func(e *colly.HTMLElement) {
		url := fmt.Sprintf(e.Attr("href"))
		page.Visit(url)
	})

	c.Visit("https://bestmalz.de/en/our-malts/base-malts/")

	return grains
}
