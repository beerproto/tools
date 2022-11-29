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

type CrispMalt struct {
	formatter lxstrconv.NumberFormat
}

func NewCrispMalt() *CrispMalt {
	return &CrispMalt{
		formatter: lxstrconv.NewDecimalFormat(language.English),
	}
}

func (s *CrispMalt) Parse() []*fermentables.GrainType {
	grains := []*fermentables.GrainType{}

	c := colly.NewCollector()
	page := colly.NewCollector()

	page.OnHTML("body", func(e *colly.HTMLElement) {
		grain := &fermentables.GrainType{
			Country:    "GBR",
			Standard:   fermentables.GrainType_EBC,
			GrainGroup: beerproto.GrainGroup_GRAIN_GROUP_BASE,
			Producer:   "Crisp Malt",
			Name:       e.ChildAttr(".malt-post__title", "alt"),
		}

		if grain.Name == "" {
			return
		}

		maximum := e.ChildText(".malt-post__overview .overview-table__column:nth-child(2) p")
		maximum = strings.TrimLeft(strings.ToLower(maximum), "up to")
		grain.Maximum = unit.Percent(maximum, unit.WithFormatter[beerproto.PercentUnit](s.formatter)).Maximum

		e.ForEach(".uk-visible\\@s tr", func(_ int, el *colly.HTMLElement) {

			header := strings.TrimSpace(el.ChildText("td:first-child"))

			text := el.ChildText("td:nth-child(3)")
			if text == "" || text == "-" {
				text = el.ChildText("td:nth-child(2)")
			}

			switch header {
			case "Moisture":
				grain.Moisture = unit.Percent(text, unit.WithFormatter[beerproto.PercentUnit](s.formatter),
					unit.WithDefault[beerproto.PercentUnit](unit.Min),
				)
			case "Soluble Extract As is":
				grain.SolubleProtein = unit.Percent(text, unit.WithFormatter[beerproto.PercentUnit](s.formatter))
			case "Friability":
				grain.Friability = unit.Percent(text, unit.WithFormatter[beerproto.PercentUnit](s.formatter))
			case "Extract":
				grain.Yield = unit.Percent(text, unit.WithFormatter[beerproto.PercentUnit](s.formatter))
			case "Extract difference fine-coarse":
				grain.FineCoarseDifference = unit.Percent(text, unit.WithFormatter[beerproto.PercentUnit](s.formatter))
			case "Colour":
				grain.Color = unit.Color(text, unit.WithFormatter[beerproto.ColorUnit](s.formatter))
			case "TN/TP":
				grain.Protein = unit.Percent(text, unit.WithFormatter[beerproto.PercentUnit](s.formatter))
			case "SNR/KI/ST RATIO":
				grain.KolbachIndex = unit.Percent(text, unit.WithFormatter[beerproto.PercentUnit](s.formatter))
			}
		})

		grains = append(grains, grain)
	})

	c.OnHTML(".malt-slider a[href]", func(e *colly.HTMLElement) {
		page.Visit(e.Attr("href"))
	})

	c.Visit("https://crispmalt.com/our-malts/")
	//page.Visit("https://crispmalt.com/malts/finest-maris-otter-ale-malt/")

	return grains
}
