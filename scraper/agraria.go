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

func Agraria() []*fermentables.GrainType {
	grains := []*fermentables.GrainType{}

	portuguese := lxstrconv.NewDecimalFormat(language.Portuguese)

	c := colly.NewCollector()
	page := colly.NewCollector()

	page.OnHTML("body", func(e *colly.HTMLElement) {
		grain := &fermentables.GrainType{
			Country:    "BRA",
			Standard:   fermentables.GrainType_EBC,
			GrainGroup: beerproto.GrainGroup_BASE,
			Producer:   "Agraria",
			Name:       e.ChildText("h1"),
		}

		e.ForEach(".conteudo table", func(_ int, el *colly.HTMLElement) {

			switch strings.TrimSpace(el.ChildText("thead th:first-child")) {
			case "Wort color:":
				grain.Color = unit.Color(el.ChildText("tbody tr:first-child th"))

				if strings.TrimSpace(el.ChildText("tbody tr:nth-child(2) th")) == "Use:" {
					grain.Notes = strings.TrimSpace(el.ChildText("tbody tr:nth-child(3) th"))
				}

				if strings.TrimSpace(el.ChildText("tbody tr:nth-child(4) th")) == "Quantity:" {
					p := strings.TrimLeft(strings.ToLower(el.ChildText("tbody tr:nth-child(5) th")), "up to")
					grain.Maximum = unit.Percent(p, unit.WithFormatter[beerproto.PercentType_PercentUnitType](portuguese)).Maximum
				}
			}

		})

		e.ForEach(".conteudo table tr", func(_ int, el *colly.HTMLElement) {
			switch strings.TrimSpace(el.ChildText("th:first-child")) {
			case "Humidity":
				grain.Moisture = unit.Percent(el.Text, unit.WithFormatter[beerproto.PercentType_PercentUnitType](portuguese))
			case "Extract from fine grinding w.f.*":
				grain.FineGrind = unit.Percent(el.Text, unit.WithFormatter[beerproto.PercentType_PercentUnitType](portuguese))

			case "Expected yield":
				grain.Yield = unit.Percent(el.Text, unit.WithFormatter[beerproto.PercentType_PercentUnitType](portuguese))

			case "Saccharification time":
				grain.Saccharification = unit.Time(el.Text, unit.WithFormatter[beerproto.TimeType_TimeUnitType](portuguese))

			case "Friabilitye":
				grain.Friability = unit.Percent(el.Text, unit.WithFormatter[beerproto.PercentType_PercentUnitType](portuguese))
			case "Beta-glucans":
				grain.BetaGlucan = unit.Concentration(el.Text, unit.WithUnit(beerproto.ConcentrationUnitType_MGL),
					unit.WithFormatter[beerproto.ConcentrationUnitType](portuguese))

			case "Viscosity":
				grain.Viscosity = unit.Viscosity(el.Text, unit.WithUnit(beerproto.ViscosityUnitType_MPAS),
					unit.WithFormatter[beerproto.ViscosityUnitType](portuguese))
			case "Diastatic power":
				grain.DiastaticPower = unit.DiastaticPower(el.Text,
					unit.WithFormatter[beerproto.DiastaticPowerUnitType](portuguese))

			case "Protein":
				grain.Protein = unit.Percent(el.Text, unit.WithFormatter[beerproto.PercentType_PercentUnitType](portuguese))
			case "FAN (Free Amino Nitrogen)":
				grain.Fan = unit.Concentration(el.Text, unit.WithFormatter[beerproto.ConcentrationUnitType](portuguese))
			}
		})

		grains = append(grains, grain)
	})

	c.OnHTML(".areas a[href]", func(e *colly.HTMLElement) {
		url := fmt.Sprintf("http://%s/en/%s", e.Request.URL.Host, e.Attr("href"))
		page.Visit(url)
	})

	c.Visit("http://www.agraria.com.br/en/malt/products")

	return grains
}
