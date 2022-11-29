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

type GambrinusMalting struct {
	formatter lxstrconv.NumberFormat
}

func NewGambrinusMalting() *GambrinusMalting {
	return &GambrinusMalting{
		formatter: lxstrconv.NewDecimalFormat(language.AmericanEnglish),
	}
}

func (s *GambrinusMalting) Parse() []*fermentables.GrainType {
	grains := []*fermentables.GrainType{}

	c := colly.NewCollector()

	page := colly.NewCollector()

	page.OnHTML("body", func(e *colly.HTMLElement) {
		grain := &fermentables.GrainType{
			Country:    "USA",
			Standard:   fermentables.GrainType_ASBC,
			GrainGroup: beerproto.GrainGroup_GRAIN_GROUP_BASE,
			Producer:   "Gambrinus Malting",
			Name:       e.ChildText("h1"),
		}

		e.ForEach(".full-description tr", func(_ int, el *colly.HTMLElement) {

			text := el.ChildText("td:nth-child(2)")
			header := strings.TrimSpace(el.ChildText("td:first-child"))

			switch header {
			case "Moisture % Max":
				grain.Moisture = unit.Percent(text,
					unit.WithFormatter[beerproto.PercentUnit](s.formatter))
			case "Extract FG Min":
				grain.Yield = unit.Percent(text,
					unit.WithFormatter[beerproto.PercentUnit](s.formatter))
			case "Color SRM":
				grain.Color = unit.Color(text,
					unit.WithFormatter[beerproto.ColorUnit](s.formatter),
					unit.WithUnit(beerproto.ColorUnit_COLOR_UNIT_SRM))
			case "Protein Total":
				grain.Protein = unit.Percent(text, unit.WithFormatter[beerproto.PercentUnit](s.formatter))
			case "Usage Rate":
				text = strings.TrimLeft(text, "Up to")
				grain.Maximum = unit.Percent(text, unit.WithFormatter[beerproto.PercentUnit](s.formatter)).Maximum

			}
		})

		grains = append(grains, grain)
	})

	c.OnHTML(".product-grid a[href]", func(e *colly.HTMLElement) {
		url := fmt.Sprintf("http://%s%s", e.Request.URL.Host, e.Attr("href"))
		page.Visit(url)
	})

	c.Visit("https://gambrinusmalting.com/malt")

	return grains
}
