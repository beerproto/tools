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

type AtelierDoMalte struct {
	formatter lxstrconv.NumberFormat
}

func NewAtelierDoMalte() *AtelierDoMalte {
	return &AtelierDoMalte{
		formatter: lxstrconv.NewDecimalFormat(language.Portuguese),
	}
}

func (s *AtelierDoMalte) Parse() []*fermentables.GrainType {
	grains := []*fermentables.GrainType{}

	page := colly.NewCollector()

	page.OnHTML(".cs-widget.cs-text-widget", func(e *colly.HTMLElement) {
		grain := &fermentables.GrainType{
			Country:    "BRA",
			Standard:   fermentables.GrainType_EBC,
			GrainGroup: beerproto.GrainGroup_BASE,
			Producer:   "Atelier Do Malte",
		}

		e.ForEach("p:first-child", func(_ int, el *colly.HTMLElement) {
			grain.Name = strings.TrimSpace(el.Text)
		})

		e.ForEach("p", func(_ int, el *colly.HTMLElement) {
			index := strings.Index(el.Text, ":")
			if index < 0 {
				return
			}

			switch strings.ToLower(strings.TrimSpace(el.Text[:index])) {
			case "extrato":
				grain.Yield = unit.Percent(el.Text,
					unit.WithFormatter[beerproto.PercentType_PercentUnitType](s.formatter))
			case "diferença moagem fina e grossa":
				grain.CoarseGrind = unit.Percent(el.Text,
					unit.WithFormatter[beerproto.PercentType_PercentUnitType](s.formatter))
			case "tempo de sacarificação":
				grain.Saccharification = unit.Time(el.Text,
					unit.WithFormatter[beerproto.TimeType_TimeUnitType](s.formatter))
			case "umidade":
				grain.Moisture = unit.Percent(el.Text,
					unit.WithFormatter[beerproto.PercentType_PercentUnitType](s.formatter))
			case "proteína solúvel":
				grain.SolubleProtein = unit.Percent(el.Text,
					unit.WithFormatter[beerproto.PercentType_PercentUnitType](s.formatter))
			case "indice kolback":
				grain.KolbachIndex = unit.Percent(el.Text,
					unit.WithFormatter[beerproto.PercentType_PercentUnitType](s.formatter))
			case "fan (free amino acids)":
				grain.Fan = unit.Concentration(el.Text,
					unit.WithFormatter[beerproto.ConcentrationUnitType](s.formatter))
			case "nitrogenio solúvel":
				grain.SolubleNitrogen = unit.Concentration(el.Text,
					unit.WithFormatter[beerproto.ConcentrationUnitType](s.formatter))
			case "viscosidade do mosto":
				grain.Viscosity = unit.Viscosity(el.Text,
					unit.WithFormatter[beerproto.ViscosityUnitType](s.formatter))
			case "friabilidade":
				grain.Friability = unit.Percent(el.Text,
					unit.WithFormatter[beerproto.PercentType_PercentUnitType](s.formatter))
			case "poder diastático":
				grain.DiastaticPower = unit.DiastaticPower(el.Text, unit.WithFormatter[beerproto.DiastaticPowerUnitType](s.formatter))
			case "dmsp (precursor de dms)":
				grain.DmsP = unit.Concentration(el.Text,
					unit.WithFormatter[beerproto.ConcentrationUnitType](s.formatter))
			case "b glucans":
				grain.BetaGlucan = unit.Concentration(el.Text,
					unit.WithFormatter[beerproto.ConcentrationUnitType](s.formatter))
			case "cor do mosto":
				grain.Color = unit.Color(el.Text,
					unit.WithFormatter[beerproto.ColorUnitType](s.formatter))
			}

		})

		grains = append(grains, grain)
	})

	page.Visit("https://atelierdomalte.com.br/cardapio")

	return grains
}
