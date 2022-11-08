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

type LesMaltiers struct {
	formatter lxstrconv.NumberFormat
}

func NewLesMaltiers() *LesMaltiers {
	return &LesMaltiers{
		formatter: lxstrconv.NewDecimalFormat(language.English),
	}
}

func (s *LesMaltiers) Parse() []*fermentables.GrainType {
	grains := []*fermentables.GrainType{}

	c := colly.NewCollector()
	page := colly.NewCollector()

	page.OnHTML("body", func(e *colly.HTMLElement) {
		grain := &fermentables.GrainType{
			Country:    "FR",
			Standard:   fermentables.GrainType_EBC,
			GrainGroup: beerproto.GrainGroup_BASE,
			Producer:   "Les Maltiers",
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
			case "Couleur":
				grain.Color = unit.Color(text, unit.WithFormatter[beerproto.ColorUnitType](s.formatter))
			case "EBC Colour":
				grain.Color = unit.Color(text, unit.WithFormatter[beerproto.ColorUnitType](s.formatter))
			case "Moisture":
				grain.Moisture = unit.Percent(text, unit.WithFormatter[beerproto.PercentType_PercentUnitType](s.formatter))
			case "Protéines totales":
				grain.Protein = unit.Percent(text, unit.WithFormatter[beerproto.PercentType_PercentUnitType](s.formatter))
			case "Protéines solubles":
				grain.SolubleProtein = unit.Percent(text, unit.WithFormatter[beerproto.PercentType_PercentUnitType](s.formatter))
			case "EBC FAN in Wort":
				grain.Fan = unit.Concentration(text, unit.WithFormatter[beerproto.ConcentrationUnitType](s.formatter))
			case "EBC Wort Viscosity":
				grain.Viscosity = unit.Viscosity(text, unit.WithFormatter[beerproto.ViscosityUnitType](s.formatter))
			case "Pouvoir diastasique":
				grain.DiastaticPower = unit.DiastaticPower(text, unit.WithFormatter[beerproto.DiastaticPowerUnitType](s.formatter))
			case "pH:":
				grain.DiPh = unit.Acidity(text, unit.WithFormatter[beerproto.AcidityUnitType](s.formatter)).Maximum
			case "Alpha Amylase dry basis":
				grain.AlphaAmylase = unit.Time(text, unit.WithFormatter[beerproto.TimeType_TimeUnitType](s.formatter)).Maximum
			case "B glucans":
				grain.BetaGlucan = unit.Concentration(text, unit.WithUnit(beerproto.ConcentrationUnitType_MGL),
					unit.WithFormatter[beerproto.ConcentrationUnitType](s.formatter))
			case "Friabilité":
				grain.Friability = unit.Percent(text, unit.WithFormatter[beerproto.PercentType_PercentUnitType](s.formatter))
			case "Extrait":
				grain.Yield = unit.Percent(text, unit.WithFormatter[beerproto.PercentType_PercentUnitType](s.formatter))
			case "EBC Extract 0.2mm dry basis":
				grain.FineGrind = unit.Percent(text, unit.WithFormatter[beerproto.PercentType_PercentUnitType](s.formatter))

			}

		})

		grains = append(grains, grain)
	})

	c.OnHTML(".section-content .col a[href]", func(e *colly.HTMLElement) {
		page.Visit(e.Attr("href"))
	})

	c.Visit("https://www.lesmaltiers.fr/shop/")

	return grains
}
