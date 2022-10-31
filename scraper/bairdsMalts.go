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

type BairdsMalts struct {
	formatter lxstrconv.NumberFormat
}

func NewBairdsMalts() *BairdsMalts {
	return &BairdsMalts{
		formatter: lxstrconv.NewDecimalFormat(language.BritishEnglish),
	}
}

func (s *BairdsMalts) Parse() []*fermentables.GrainType {
	grains := []*fermentables.GrainType{}

	page := colly.NewCollector()

	page.OnHTML("body", func(e *colly.HTMLElement) {
		grain := &fermentables.GrainType{
			Country:    "GBR",
			Standard:   fermentables.GrainType_EBC,
			GrainGroup: beerproto.GrainGroup_BASE,
			Producer:   "Bairds Malt",
			Name:       e.ChildText("h1"),
		}

		e.ForEach("div.et_pb_tab_1 p", func(_ int, el *colly.HTMLElement) {

			index := strings.Index(el.Text, ":")
			if index < 0 {
				return
			}

			text := strings.ToLower(strings.TrimSpace(el.Text[index+1:]))

			switch strings.ToLower(strings.TrimSpace(el.Text[:index])) {
			case "moisture":
				grain.Moisture = unit.Percent(el.Text,
					unit.WithFormatter[beerproto.PercentType_PercentUnitType](s.formatter))
			case "extract (0.7/0.2mm), dry":
				grain.Yield = unit.Percent(el.Text,
					unit.WithFormatter[beerproto.PercentType_PercentUnitType](s.formatter))
			case "colour (ebc/srm units)":
				grain.Color = unit.Color(el.Text,
					unit.WithFormatter[beerproto.ColorUnitType](s.formatter))
			case "total nitrogen/protein, dry":
				grain.TotalNitrogen = unit.Percent(el.Text,
					unit.WithFormatter[beerproto.PercentType_PercentUnitType](s.formatter))
			case "snr / ki/ st ratio":
				grain.KolbachIndex = unit.Percent(el.Text,
					unit.WithFormatter[beerproto.PercentType_PercentUnitType](s.formatter))
			case "diastatic power":
				grain.DiastaticPower = unit.DiastaticPower(text,
					unit.WithUnit(beerproto.DiastaticPowerUnitType_LINTNER),
					unit.WithFormatter[beerproto.DiastaticPowerUnitType](s.formatter))
			}
		})

		e.ForEach(".et_pb_gutters2 p", func(_ int, el *colly.HTMLElement) {
			index := strings.Index(el.Text, ":")
			if index < 0 {
				return
			}

			switch strings.ToLower(strings.TrimSpace(el.Text[:index])) {
			case "inclusion rate":

				right := strings.ToLower(strings.TrimSpace(el.Text[index+1:]))

				max := strings.TrimRight(strings.TrimSpace(strings.TrimLeft(right, "up to")), "%")
				if max, err := s.formatter.ParseFloat(max); err == nil {
					maximum := &beerproto.PercentType{
						Value: max,
						Unit:  beerproto.PercentType_PERCENT_SIGN,
					}
					grain.Maximum = maximum
				}
			case "suggested use":

				right := strings.TrimSpace(el.Text[index+1:])
				grain.Notes = right

			}
		})

		grains = append(grains, grain)
	})

	page.Visit("https://www.bairds-malt.co.uk/the-1823-heritage-collection/heritage-distilling-malt/")

	return grains
}

// func (s *BairdsMalts) time(value string, index int) *beerproto.TimeRangeType {
// 	right := strings.ToLower(strings.TrimSpace(value[index+1:]))
// 	dash := strings.Index(right, "-")

// 	if dash > 0 {
// 		right = strings.TrimSpace(right[dash+1:])
// 	}
// 	max := strings.TrimRight(strings.TrimLeft(strings.TrimLeft(right, "max"), "approx"), "min")
// 	time := &beerproto.TimeRangeType{}
// 	if max, err := s.formater.ParseInt(max); err == nil {
// 		time.Maximum = &beerproto.TimeType{
// 			Value: max,
// 			Unit:  beerproto.TimeType_MIN,
// 		}
// 	}

// 	return time
// }

// func (s *BairdsMalts) viscosity(value string, index int) *beerproto.ViscosityRangeType {
// 	right := strings.ToLower(strings.TrimSpace(value[index+1:]))

// 	viscosity := &beerproto.ViscosityRangeType{}

// 	min := strings.TrimRight(strings.TrimLeft(strings.TrimLeft(right, "mix"), "."), "cp")
// 	if min, err := s.formater.ParseFloat(min); err == nil {
// 		viscosity.Minimum = &beerproto.ViscosityType{
// 			Value: min,
// 			Unit:  beerproto.ViscosityUnitType_CP,
// 		}
// 	}

//		return viscosity
//	}
// func (s *BairdsMalts) diastaticPower(value string, index int) *beerproto.DiastaticPowerRangeType {
// 	right := strings.ToLower(strings.TrimSpace(value[index+1:]))

// 	diastaticPower := &beerproto.DiastaticPowerRangeType{}

// 	min := strings.TrimRight(strings.TrimSpace(strings.TrimRight(right, "min")), "wk")
// 	if min, err := s.formater.ParseFloat(min); err == nil {
// 		diastaticPower.Minimum = &beerproto.DiastaticPowerType{
// 			Value: min,
// 			Unit:  beerproto.DiastaticPowerUnitType_WK,
// 		}
// 	}

// 	return diastaticPower
// }

func (s *BairdsMalts) concentration(value string, index int) *beerproto.ConcentrationRangeType {
	right := strings.ToLower(strings.TrimSpace(value[index+1:]))

	concentration := &beerproto.ConcentrationRangeType{}

	min := strings.TrimRight(strings.TrimLeft(strings.TrimLeft(right, "max"), "max"), "ppm")
	if min, err := s.formatter.ParseFloat(min); err == nil {
		concentration.Minimum = &beerproto.ConcentrationType{
			Value: min,
			Unit:  beerproto.ConcentrationUnitType_PPM,
		}
	}

	return concentration
}

func (s *BairdsMalts) color(value string, index int) (color *beerproto.ColorRangeType) {
	color = &beerproto.ColorRangeType{}

	right := strings.ToLower(strings.TrimSpace(value[index+1:]))
	indexA := strings.Index(right, "–")
	splitMin := right[:indexA]
	splitMax := right[indexA+3:]

	min := strings.TrimRight(strings.TrimLeft(splitMin, "min"), "°ebc")
	if min, err := s.formatter.ParseFloat(min); err == nil {
		color.Minimum = &beerproto.ColorType{
			Value: min,
			Unit:  beerproto.ColorUnitType_EBC,
		}
	}

	max := strings.TrimRight(strings.TrimLeft(splitMax, "max"), "°ebc")
	if max, err := s.formatter.ParseFloat(max); err == nil {
		color.Maximum = &beerproto.ColorType{
			Value: max,
			Unit:  beerproto.ColorUnitType_EBC,
		}
	}

	return
}

func (s *BairdsMalts) percent(value string, index int) (percent *beerproto.PercentRangeType) {
	percent = &beerproto.PercentRangeType{}
	right := strings.ToLower(strings.TrimSpace(value[index+1:]))

	if strings.HasSuffix(right, "max") {
		max := strings.TrimRight(strings.TrimSpace(strings.TrimRight(right, "max")), "%")
		if max, err := s.formatter.ParseFloat(max); err == nil {
			percent.Maximum = &beerproto.PercentType{
				Value: max,
				Unit:  beerproto.PercentType_PERCENT_SIGN,
			}
		}
	}

	if strings.HasSuffix(right, "min") {
		min := strings.TrimRight(strings.TrimSpace(strings.TrimRight(right, "min")), "%")
		if min, err := s.formatter.ParseFloat(min); err == nil {
			percent.Minimum = &beerproto.PercentType{
				Value: min,
				Unit:  beerproto.PercentType_PERCENT_SIGN,
			}
		}
	}

	return
}
