package scraper

import (
	"fmt"
	"strings"

	beerproto "github.com/beerproto/beerproto_go"
	"github.com/beerproto/beerproto_go/fermentables"
	colly "github.com/gocolly/colly/v2"
	"golang.org/x/text/language"
	"tawesoft.co.uk/go/lxstrconv"
)

type Bestmalz struct {
	formater lxstrconv.NumberFormat
}

func NewBestmalz() *Bestmalz {
	return &Bestmalz{
		formater: lxstrconv.NewDecimalFormat(language.German),
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
			GrainGroup: beerproto.GrainGroup_BASE,
			Producer:   "Bestmalz",
			Name:       e.ChildText("h1"),
		}

		e.ForEach("table tbody tr", func(_ int, el *colly.HTMLElement) {

			spec := el.DOM.Children().First().Text()
			text := el.DOM.Children().First().NextAll().Text()

			switch strings.ToLower(strings.TrimSpace(spec)) {
			case "moisture content":
				grain.Moisture = percent(text, s.formater, beerproto.PercentType_PERCENT_SIGN)
			case "extract fine grind, dry basis":
				//grain.Yield = percent(el.Text, s.formater)
			case "fine-coarse difference ebc":
				grain.CoarseGrind = percent(text, s.formater, beerproto.PercentType_PERCENT_SIGN)
			case "viscosity (8.6%)":
				grain.Viscosity = viscosity(text, s.formater)
			case "friability":
				grain.Friability = percent(text, s.formater, beerproto.PercentType_PERCENT_SIGN)
			case "glassiness":
			//	grain.gl = s.diastaticPower(el.Text, index)
			case "protein, dry basis":
				grain.Protein = percent(text, s.formater, beerproto.PercentType_PERCENT_SIGN)
			case "soluble nitrogen":
				grain.SolubleNitrogen = concentration(text, s.formater, beerproto.ConcentrationUnitType_MG100L)
			case "kolbach index":
				grain.KolbachIndex = percent(text, s.formater, beerproto.PercentType_PERCENT_SIGN)
			case "wort color":
				grain.Color = color(text, s.formater, beerproto.ColorUnitType_EBC)
			case "wort ph":
				grain.DiPh = acidity(text, s.formater)
			case "grading > 2.5mm":
			//	grain.gr = s.diastaticPower(el.Text, index)
			case "diastatic power":
				grain.DiastaticPower = diastaticPower(text, s.formater, beerproto.DiastaticPowerUnitType_WK)
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
