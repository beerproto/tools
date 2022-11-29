package scraper

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"strings"

	beerproto "github.com/beerproto/beerproto_go"
	"github.com/beerproto/beerproto_go/fermentables"
	"github.com/beerproto/tools/unit"
	colly "github.com/gocolly/colly/v2"
	"golang.org/x/text/language"
	"tawesoft.co.uk/go/lxstrconv"
)

type MouterijDingemans struct {
	formatter lxstrconv.NumberFormat
}

func NewMouterijDingemans() *MouterijDingemans {
	return &MouterijDingemans{
		formatter: lxstrconv.NewDecimalFormat(language.English),
	}
}

func (s *MouterijDingemans) Parse() []*fermentables.GrainType {
	grains := []*fermentables.GrainType{}

	c := colly.NewCollector()
	c.WithTransport(&http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	})
	page := colly.NewCollector()
	page.WithTransport(&http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	})

	page.OnHTML("html", func(e *colly.HTMLElement) {
		grain := &fermentables.GrainType{
			Country:    "BEL",
			Standard:   fermentables.GrainType_EBC,
			GrainGroup: beerproto.GrainGroup_GRAIN_GROUP_BASE,
			Producer:   "Mouterij Dingemans",
		}

		grain.Name = e.ChildText("h1")

		e.ForEach(".horizontal-tabs-panes .field", func(_ int, el *colly.HTMLElement) {

			header := strings.TrimRight(strings.TrimSpace(el.ChildText(".field-label")), ":")

			text := el.ChildText(".field-items")

			switch header {
			case "Extract fine D.M.":
				grain.Yield = unit.Percent(text, unit.WithFormatter[beerproto.PercentUnit](s.formatter))
			case "Moisture":
				grain.Moisture = unit.Percent(text, unit.WithFormatter[beerproto.PercentUnit](s.formatter),
					unit.WithDefault[beerproto.PercentUnit](unit.Min),
				)
			case "Colour":
				grain.Color = unit.Color(text, unit.WithFormatter[beerproto.ColorUnit](s.formatter))
			case "Total protein":
				grain.Protein = unit.Percent(text, unit.WithFormatter[beerproto.PercentUnit](s.formatter))
			case "pH":
				grain.DiPh = unit.Acidity(text, unit.WithFormatter[beerproto.AcidityUnit](s.formatter)).Maximum
			}
		})

		grains = append(grains, grain)
	})

	c.OnHTML(".portfolio-list li a", func(e *colly.HTMLElement) {
		url := fmt.Sprintf("https://%s%s", e.Request.URL.Host, e.Attr("href"))
		page.Visit(url)
	})

	c.Visit("https://www.dingemansmout.be/malts")
	//page.Visit("https://www.dingemansmout.be/malt/organic-pale-ale-md")

	return grains
}
