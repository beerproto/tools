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

type CastleMalting struct {
	formatter lxstrconv.NumberFormat
}

func NewCastleMalting() *CastleMalting {
	return &CastleMalting{
		formatter: lxstrconv.NewDecimalFormat(language.English),
	}
}

func (s *CastleMalting) Parse() []*fermentables.GrainType {
	grains := []*fermentables.GrainType{}

	c := colly.NewCollector()
	page := colly.NewCollector()

	page.OnHTML("body", func(e *colly.HTMLElement) {
		grain := &fermentables.GrainType{
			Country:    "BEL",
			Standard:   fermentables.GrainType_EBC,
			GrainGroup: beerproto.GrainGroup_BASE,
			Producer:   "Castle Malting",
			Name:       e.ChildText(".MaltName"),
		}

		if grain.Name == "" {
			return
		}

		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {

			header := strings.TrimSpace(el.ChildText("td:first-child"))

			min := el.ChildText("td:nth-child(3)")
			max := el.ChildText("td:nth-child(4)")

			text := ""
			if min != "" && max != "" {
				text = min + " and " + max
			} else {
				text = min + max
			}

			switch header {
			case "Moisture":
				grain.Moisture = unit.Percent(text, unit.WithFormatter[beerproto.PercentType_PercentUnitType](s.formatter),
					unit.WithDefault[beerproto.PercentType_PercentUnitType](unit.Min),
				)
			case "Extract (dry basis)":
				grain.Yield = unit.Percent(text, unit.WithFormatter[beerproto.PercentType_PercentUnitType](s.formatter))
			case "Extract difference fine-coarse":
				grain.FineCoarseDifference = unit.Percent(text, unit.WithFormatter[beerproto.PercentType_PercentUnitType](s.formatter))
			case "Wort color":
				if strings.Contains(text, "(") {
					index := strings.Index(text, "(")
					text = text[:index]
				}
				grain.Color = unit.Color(text, unit.WithFormatter[beerproto.ColorUnitType](s.formatter))
			case "Total protein":
				grain.Protein = unit.Percent(text, unit.WithFormatter[beerproto.PercentType_PercentUnitType](s.formatter))
			case "Soluble protein":
				grain.SolubleProtein = unit.Percent(text, unit.WithFormatter[beerproto.PercentType_PercentUnitType](s.formatter))
			case "Kolbach index":
				grain.KolbachIndex = unit.Percent(text, unit.WithFormatter[beerproto.PercentType_PercentUnitType](s.formatter))
			case "Viscosity":
				grain.Viscosity = unit.Viscosity(text, unit.WithFormatter[beerproto.ViscosityUnitType](s.formatter))
			case "Beta glucans":
				grain.BetaGlucan = unit.Concentration(text, unit.WithUnit(beerproto.ConcentrationUnitType_MGL),
					unit.WithFormatter[beerproto.ConcentrationUnitType](s.formatter))
			case "pH":
				grain.DiPh = unit.Acidity(text, unit.WithFormatter[beerproto.AcidityUnitType](s.formatter)).Maximum
			case "Diastatic Power":
				grain.DiastaticPower = unit.DiastaticPower(text, unit.WithFormatter[beerproto.DiastaticPowerUnitType](s.formatter))
			case "Friability":
				grain.Friability = unit.Percent(text, unit.WithFormatter[beerproto.PercentType_PercentUnitType](s.formatter))
			case "Saccharification":
				grain.Saccharification = unit.Time(text, unit.WithFormatter[beerproto.TimeType_TimeUnitType](s.formatter))
			}
		})

		grains = append(grains, grain)
	})

	c.OnHTML(".ProductItem .ProductActions a[href]", func(e *colly.HTMLElement) {

		if strings.Contains(strings.ToLower(e.Attr("title")), "html") {
			url := fmt.Sprintf("http://%s/%s", e.Request.URL.Host, e.Attr("href"))

			page.Visit(url)
		}
	})

	c.Visit("https://www.castlemalting.com/CastleMaltingMalts.asp?P=105&Language=English")
	//page.Visit("https://www.castlemalting.com/CastleMaltingMaltSpecification.asp?Command=SpecificationShow&SpecificationID=278&CropYear=2022&Language=English&FileType=HTML")

	return grains
}
