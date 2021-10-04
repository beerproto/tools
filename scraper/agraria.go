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
				arr := strings.Split(el.ChildText("tbody tr:first-child th"), " ")
				grain.Color = &beerproto.ColorRangeType{}
				if min, err := portuguese.ParseFloat(arr[0]); err == nil {
					grain.Color.Minimum = &beerproto.ColorType{
						Value: min,
						Unit:  beerproto.ColorUnitType_EBC,
					}
				}
				if min, err := portuguese.ParseFloat(arr[2]); err == nil {
					grain.Color.Maximum = &beerproto.ColorType{
						Value: min,
						Unit:  beerproto.ColorUnitType_EBC,
					}
				}
				if strings.TrimSpace(el.ChildText("tbody tr:nth-child(2) th")) == "Use:" {
					grain.Notes = strings.TrimSpace(el.ChildText("tbody tr:nth-child(3) th"))
				}

				if strings.TrimSpace(el.ChildText("tbody tr:nth-child(4) th")) == "Quantity:" {
					p := strings.TrimLeft(strings.ToLower(el.ChildText("tbody tr:nth-child(5) th")), "up to ")
					percent := strings.TrimRight(p, "%")
					if v, err := portuguese.ParseFloat(percent); err == nil {
						grain.Maximum = &beerproto.PercentType{
							Value: v,
							Unit:  beerproto.PercentType_PERCENT_SIGN,
						}
					}
				}
			}

		})

		e.ForEach(".conteudo table tr", func(_ int, el *colly.HTMLElement) {
			switch strings.TrimSpace(el.ChildText("th:first-child")) {
			case "Humidity":
				grain.Moisture = &beerproto.PercentRangeType{}
				if min, err := portuguese.ParseFloat(el.ChildText("td:nth-child(2)")); err == nil {
					grain.Moisture.Minimum = &beerproto.PercentType{
						Value: min,
						Unit:  beerproto.PercentType_PERCENT_SIGN,
					}
				}
				if max, err := portuguese.ParseFloat(el.ChildText("td:nth-child(3)")); err == nil {
					grain.Moisture.Maximum = &beerproto.PercentType{
						Value: max,
						Unit:  beerproto.PercentType_PERCENT_SIGN,
					}
				}
			case "Extract from fine grinding w.f.*":
				grain.FineGrind = &beerproto.PercentRangeType{}
				if min, err := portuguese.ParseFloat(el.ChildText("td:nth-child(2)")); err == nil {
					grain.FineGrind.Minimum = &beerproto.PercentType{
						Value: min,
						Unit:  beerproto.PercentType_PERCENT_SIGN,
					}
				}
				if max, err := portuguese.ParseFloat(el.ChildText("td:nth-child(3)")); err == nil {
					grain.FineGrind.Maximum = &beerproto.PercentType{
						Value: max,
						Unit:  beerproto.PercentType_PERCENT_SIGN,
					}
				}
			case "Expected yield":
				grain.Yield = &beerproto.PercentRangeType{}
				if min, err := portuguese.ParseFloat(el.ChildText("td:nth-child(2)")); err == nil {
					grain.Yield.Minimum = &beerproto.PercentType{
						Value: min,
						Unit:  beerproto.PercentType_PERCENT_SIGN,
					}
				}
				if max, err := portuguese.ParseFloat(el.ChildText("td:nth-child(3)")); err == nil {
					grain.Yield.Maximum = &beerproto.PercentType{
						Value: max,
						Unit:  beerproto.PercentType_PERCENT_SIGN,
					}
				}
			case "Saccharification time":
				grain.Saccharification = &beerproto.TimeRangeType{}
				if min, err := portuguese.ParseInt(el.ChildText("td:nth-child(2)")); err == nil {
					grain.Saccharification.Minimum = &beerproto.TimeType{
						Value: min,
						Unit:  beerproto.TimeType_MIN,
					}
				}
				if max, err := portuguese.ParseInt(el.ChildText("td:nth-child(3)")); err == nil {
					grain.Saccharification.Maximum = &beerproto.TimeType{
						Value: max,
						Unit:  beerproto.TimeType_MIN,
					}
				}
			case "Friabilitye":
				grain.Friability = &beerproto.PercentRangeType{}
				if min, err := portuguese.ParseFloat(el.ChildText("td:nth-child(2)")); err == nil {
					grain.Friability.Minimum = &beerproto.PercentType{
						Value: min,
						Unit:  beerproto.PercentType_PERCENT_SIGN,
					}
				}
				if max, err := portuguese.ParseFloat(el.ChildText("td:nth-child(3)")); err == nil {
					grain.Friability.Maximum = &beerproto.PercentType{
						Value: max,
						Unit:  beerproto.PercentType_PERCENT_SIGN,
					}
				}
			case "Beta-glucans":
				grain.BetaGlucan = &beerproto.ConcentrationRangeType{}
				if min, err := portuguese.ParseFloat(el.ChildText("td:nth-child(2)")); err == nil {
					grain.BetaGlucan.Minimum = &beerproto.ConcentrationType{
						Value: min,
						Unit:  beerproto.ConcentrationUnitType_MGL,
					}
				}
				if max, err := portuguese.ParseFloat(el.ChildText("td:nth-child(3)")); err == nil {
					grain.BetaGlucan.Maximum = &beerproto.ConcentrationType{
						Value: max,
						Unit:  beerproto.ConcentrationUnitType_MGL,
					}
				}
			case "Viscosity":
				grain.Viscosity = &beerproto.ViscosityRangeType{}
				if min, err := portuguese.ParseFloat(el.ChildText("td:nth-child(2)")); err == nil {
					grain.Viscosity.Minimum = &beerproto.ViscosityType{
						Value: min,
						Unit:  beerproto.ViscosityUnitType_MPAS,
					}
				}
				if max, err := portuguese.ParseFloat(el.ChildText("td:nth-child(3)")); err == nil {
					grain.Viscosity.Maximum = &beerproto.ViscosityType{
						Value: max,
						Unit:  beerproto.ViscosityUnitType_MPAS,
					}
				}
			case "Diastatic power":
				grain.DiastaticPower = &beerproto.DiastaticPowerRangeType{}
				if min, err := portuguese.ParseFloat(el.ChildText("td:nth-child(2)")); err == nil {
					grain.DiastaticPower.Minimum = &beerproto.DiastaticPowerType{
						Value: min,
						Unit:  beerproto.DiastaticPowerUnitType_WK,
					}
				}
				if max, err := portuguese.ParseFloat(el.ChildText("td:nth-child(3)")); err == nil {
					grain.DiastaticPower.Maximum = &beerproto.DiastaticPowerType{
						Value: max,
						Unit:  beerproto.DiastaticPowerUnitType_WK,
					}
				}
			case "Protein":
				grain.Protein = &beerproto.PercentRangeType{}
				if min, err := portuguese.ParseFloat(el.ChildText("td:nth-child(2)")); err == nil {
					grain.Protein.Minimum = &beerproto.PercentType{
						Value: min,
						Unit:  beerproto.PercentType_PERCENT_SIGN,
					}
				}
				if max, err := portuguese.ParseFloat(el.ChildText("td:nth-child(3)")); err == nil {
					grain.Protein.Maximum = &beerproto.PercentType{
						Value: max,
						Unit:  beerproto.PercentType_PERCENT_SIGN,
					}
				}
			case "FAN (Free Amino Nitrogen)":
				grain.Fan = &beerproto.ConcentrationRangeType{}
				if min, err := portuguese.ParseFloat(el.ChildText("td:nth-child(2)")); err == nil {
					grain.Fan.Minimum = &beerproto.ConcentrationType{
						Value: min,
						Unit:  beerproto.ConcentrationUnitType_MG100L,
					}
				}
				if max, err := portuguese.ParseFloat(el.ChildText("td:nth-child(3)")); err == nil {
					grain.Fan.Maximum = &beerproto.ConcentrationType{
						Value: max,
						Unit:  beerproto.ConcentrationUnitType_MG100L,
					}
				}
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
