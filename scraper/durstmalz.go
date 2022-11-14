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

type DurstMalz struct {
	formatter lxstrconv.NumberFormat
}

func NewDurstMalz() *DurstMalz {
	return &DurstMalz{
		formatter: lxstrconv.NewDecimalFormat(language.German),
	}
}

func (s *DurstMalz) Parse() []*fermentables.GrainType {
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
			Country:    "DEU",
			Standard:   fermentables.GrainType_EBC,
			GrainGroup: beerproto.GrainGroup_BASE,
			Producer:   "Durst Malz",
		}

		title := e.ChildText("title")
		grain.Name = strings.TrimSpace(strings.Split(title, ":")[1])
		if grain.Name == "" {
			return
		}

		e.ForEach(".contenttable tr", func(_ int, el *colly.HTMLElement) {

			header := strings.TrimSpace(el.ChildText("td:first-child"))

			text := el.ChildText("td:nth-child(2)")

			switch header {
			case "Wassergehalt":
				grain.Moisture = unit.Percent(text, unit.WithFormatter[beerproto.PercentType_PercentUnitType](s.formatter),
					unit.WithDefault[beerproto.PercentType_PercentUnitType](unit.Min),
				)
			case "Feinschrotextrakt":
				grain.Yield = unit.Percent(text, unit.WithFormatter[beerproto.PercentType_PercentUnitType](s.formatter))
			case "Extraktdifferenz":
				grain.FineCoarseDifference = unit.Percent(text, unit.WithFormatter[beerproto.PercentType_PercentUnitType](s.formatter))
			case "Malzfarbe":
				grain.Color = unit.Color(text, unit.WithFormatter[beerproto.ColorUnitType](s.formatter))
			case "Eiweißgehalt":
				grain.Protein = unit.Percent(text, unit.WithFormatter[beerproto.PercentType_PercentUnitType](s.formatter))
			case "Eiweißlösung":
				grain.SolubleProtein = unit.Percent(text, unit.WithFormatter[beerproto.PercentType_PercentUnitType](s.formatter))
			case "Viskosität":
				grain.Viscosity = unit.Viscosity(text, unit.WithFormatter[beerproto.ViscosityUnitType](s.formatter))
			case "Friabilimeter mehlig":
				grain.Friability = unit.Percent(text, unit.WithFormatter[beerproto.PercentType_PercentUnitType](s.formatter))
			case "FAN":
				grain.Fan = unit.Concentration(text, unit.WithFormatter[beerproto.ConcentrationUnitType](s.formatter))
			case "Diastatische Kraft":
				grain.DiastaticPower = unit.DiastaticPower(text, unit.WithFormatter[beerproto.DiastaticPowerUnitType](s.formatter))
			}
		})

		grains = append(grains, grain)
	})

	c.OnHTML(".activeFirstLayer td a", func(e *colly.HTMLElement) {
		url := fmt.Sprintf("http://%s/%s", e.Request.URL.Host, e.Attr("href"))
		page.Visit(url)
	})

	c.Visit("http://www.durst-malz.com/malz.html")
	//page.Visit("https://www.castlemalting.com/CastleMaltingMaltSpecification.asp?Command=SpecificationShow&SpecificationID=278&CropYear=2022&Language=English&FileType=HTML")

	return grains
}
