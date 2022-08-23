package scraper

import (
	"strings"

	beerproto "github.com/beerproto/beerproto_go"
	"github.com/beerproto/beerproto_go/fermentables"
	colly "github.com/gocolly/colly/v2"
	"golang.org/x/text/language"
	"tawesoft.co.uk/go/lxstrconv"
)

type AtelierDoMalte struct {
	formater lxstrconv.NumberFormat
}

func NewAtelierDoMalte() *AtelierDoMalte {
	return &AtelierDoMalte{
		formater: lxstrconv.NewDecimalFormat(language.Portuguese),
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
				grain.Yield = s.percent(el.Text, index)
			case "diferença moagem fina e grossa":
				grain.CoarseGrind = s.percent(el.Text, index)
			case "tempo de sacarificação":
				grain.Saccharification = s.time(el.Text, index)
			case "umidade":
				grain.Moisture = s.percent(el.Text, index)
			case "proteína solúvel":
				grain.SolubleProtein = s.percent(el.Text, index)
			case "indice kolback":
				grain.KolbachIndex = s.percent(el.Text, index)
			case "fan (free amino acids)":
				grain.Fan = s.concentration(el.Text, index)
			case "nitrogenio solúvel":
				grain.SolubleNitrogen = s.concentration(el.Text, index)
			case "viscosidade do mosto":
				grain.Viscosity = s.viscosity(el.Text, index)
			case "friabilidade":
				grain.Friability = s.percent(el.Text, index)
			case "poder diastático":
				grain.DiastaticPower = s.diastaticPower(el.Text, index)
			case "dmsp (precursor de dms)":
				grain.DmsP = s.concentration(el.Text, index)
			case "b glucans":
				grain.BetaGlucan = s.concentration(el.Text, index)
			case "cor do mosto":
				grain.Color = s.color(el.Text, index)
			}

		})

		grains = append(grains, grain)
	})

	page.Visit("https://atelierdomalte.com.br/cardapio")

	return grains
}

func (s *AtelierDoMalte) time(value string, index int) *beerproto.TimeRangeType {
	right := strings.ToLower(strings.TrimSpace(value[index+1:]))
	dash := strings.Index(right, "-")

	if dash > 0 {
		right = strings.TrimSpace(right[dash+1:])
	}
	max := strings.TrimRight(strings.TrimLeft(strings.TrimLeft(right, "max"), "approx"), "min")
	time := &beerproto.TimeRangeType{}
	if max, err := s.formater.ParseInt(max); err == nil {
		time.Maximum = &beerproto.TimeType{
			Value: max,
			Unit:  beerproto.TimeType_MIN,
		}
	}

	return time
}

func (s *AtelierDoMalte) viscosity(value string, index int) *beerproto.ViscosityRangeType {
	right := strings.ToLower(strings.TrimSpace(value[index+1:]))

	viscosity := &beerproto.ViscosityRangeType{}

	min := strings.TrimRight(strings.TrimLeft(strings.TrimLeft(right, "máx"), "."), "cp")
	if min, err := s.formater.ParseFloat(min); err == nil {
		viscosity.Minimum = &beerproto.ViscosityType{
			Value: min,
			Unit:  beerproto.ViscosityUnitType_CP,
		}
	}

	return viscosity
}
func (s *AtelierDoMalte) diastaticPower(value string, index int) *beerproto.DiastaticPowerRangeType {
	right := strings.ToLower(strings.TrimSpace(value[index+1:]))

	diastaticPower := &beerproto.DiastaticPowerRangeType{}

	min := strings.TrimRight(strings.TrimLeft(right, "mín"), "wk")
	if min, err := s.formater.ParseFloat(min); err == nil {
		diastaticPower.Minimum = &beerproto.DiastaticPowerType{
			Value: min,
			Unit:  beerproto.DiastaticPowerUnitType_WK,
		}
	}

	return diastaticPower
}

func (s *AtelierDoMalte) concentration(value string, index int) *beerproto.ConcentrationRangeType {
	right := strings.ToLower(strings.TrimSpace(value[index+1:]))

	concentration := &beerproto.ConcentrationRangeType{}

	min := strings.TrimRight(strings.TrimLeft(strings.TrimLeft(right, "max"), "máx"), "ppm")
	if min, err := s.formater.ParseFloat(min); err == nil {
		concentration.Minimum = &beerproto.ConcentrationType{
			Value: min,
			Unit:  beerproto.ConcentrationUnitType_PPM,
		}
	}

	return concentration
}

func (s *AtelierDoMalte) color(value string, index int) (color *beerproto.ColorRangeType) {
	color = &beerproto.ColorRangeType{}

	right := strings.ToLower(strings.TrimSpace(value[index+1:]))
	indexA := strings.Index(right, "a")
	splitMin := right[:indexA]
	splitMax := right[indexA+1:]

	min := strings.TrimRight(strings.TrimLeft(splitMin, "min"), "ebc")
	if min, err := s.formater.ParseFloat(min); err == nil {
		color.Minimum = &beerproto.ColorType{
			Value: min,
			Unit:  beerproto.ColorUnitType_EBC,
		}
	}

	bracket := strings.Index(splitMax, "(")
	if bracket > 0 {
		splitMax = splitMax[:bracket]
		splitMax = strings.TrimSpace(splitMax)
	}

	max := strings.TrimRight(strings.TrimLeft(splitMax, "max "), "ebc")
	if max, err := s.formater.ParseFloat(max); err == nil {
		color.Maximum = &beerproto.ColorType{
			Value: max,
			Unit:  beerproto.ColorUnitType_EBC,
		}
	}

	return
}

func (s *AtelierDoMalte) percent(value string, index int) (percent *beerproto.PercentRangeType) {
	percent = &beerproto.PercentRangeType{}
	right := strings.ToLower(strings.TrimSpace(value[index+1:]))

	indexA := strings.Index(right, "a")
	splitMin := right
	if indexA > 0 {
		splitMin = right[:indexA]
		splitMax := right[indexA+1:]

		bracket := strings.Index(splitMax, "(")
		if bracket > 0 {
			splitMax = splitMax[:bracket]
			splitMax = strings.TrimSpace(splitMax)
		}

		max := strings.TrimRight(strings.TrimLeft(splitMax, "max"), "%")
		if max, err := s.formater.ParseFloat(max); err == nil {
			percent.Maximum = &beerproto.PercentType{
				Value: max,
				Unit:  beerproto.PercentType_PERCENT_SIGN,
			}
		}
	}
	min := strings.TrimRight(strings.TrimLeft(splitMin, "min"), "%")
	if min, err := s.formater.ParseFloat(min); err == nil {
		percent.Minimum = &beerproto.PercentType{
			Value: min,
			Unit:  beerproto.PercentType_PERCENT_SIGN,
		}
	}

	return
}
