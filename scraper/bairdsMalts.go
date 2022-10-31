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
