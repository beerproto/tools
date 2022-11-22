package scraper

import (
	"fmt"
	"strings"

	beerproto "github.com/beerproto/beerproto_go"
	"github.com/beerproto/beerproto_go/fermentables"
	"github.com/beerproto/tools/unit"
	"github.com/beerproto/tools/utils"
	colly "github.com/gocolly/colly/v2"
	"tawesoft.co.uk/go/lxstrconv"
)

type MaltingOptions struct {
	baseURL               *string
	titleSelector         *string
	productSelector       *string
	productRowSelector    *string
	productHeaderSelector *string
	productHeaderTrim     *[]string
	productValueSelector  *string
	productValueTrim      *[]string
	standard              *fermentables.GrainType_StandardType
	grainGroup            *beerproto.GrainGroup

	moisture      *string
	yield         *string
	color         *string
	protein       *string
	maximum       *string
	friability    *string
	betaGlucan    *string
	alphaAmylase  *string
	diPh          *string
	fan           *string
	totalNitrogen *string
}

type MaltingOptionsFunc func(opts *MaltingOptions)

func WithProductValueTrim(trim []string) MaltingOptionsFunc {
	return func(s *MaltingOptions) {
		if s.productValueTrim == nil {
			s.productValueTrim = &trim
		}
	}
}

func WithProductHeaderTrim(trim []string) MaltingOptionsFunc {
	return func(s *MaltingOptions) {
		if s.productHeaderTrim == nil {
			s.productHeaderTrim = &trim
		}
	}
}

func WithTotalNitrogen(header string) MaltingOptionsFunc {
	return func(s *MaltingOptions) {
		if s.totalNitrogen == nil {
			s.totalNitrogen = &header
		}
	}
}

func WithFan(header string) MaltingOptionsFunc {
	return func(s *MaltingOptions) {
		if s.fan == nil {
			s.fan = &header
		}
	}
}

func WithDiPh(header string) MaltingOptionsFunc {
	return func(s *MaltingOptions) {
		if s.diPh == nil {
			s.diPh = &header
		}
	}
}

func WithAlphaAmylase(header string) MaltingOptionsFunc {
	return func(s *MaltingOptions) {
		if s.alphaAmylase == nil {
			s.alphaAmylase = &header
		}
	}
}

func WithBetaGlucan(header string) MaltingOptionsFunc {
	return func(s *MaltingOptions) {
		if s.betaGlucan == nil {
			s.betaGlucan = &header
		}
	}
}

func WithFriability(header string) MaltingOptionsFunc {
	return func(s *MaltingOptions) {
		if s.friability == nil {
			s.friability = &header
		}
	}
}

func WithMaximum(header string) MaltingOptionsFunc {
	return func(s *MaltingOptions) {
		if s.maximum == nil {
			s.maximum = &header
		}
	}
}

func WithProtein(header string) MaltingOptionsFunc {
	return func(s *MaltingOptions) {
		if s.protein == nil {
			s.protein = &header
		}
	}
}

func WithYield(header string) MaltingOptionsFunc {
	return func(s *MaltingOptions) {
		if s.yield == nil {
			s.yield = &header
		}
	}
}

func WithMoisture(header string) MaltingOptionsFunc {
	return func(s *MaltingOptions) {
		if s.moisture == nil {
			s.moisture = &header
		}
	}
}

func WithProductValueSelector(selector string) MaltingOptionsFunc {
	return func(s *MaltingOptions) {
		if s.productValueSelector == nil {
			s.productValueSelector = &selector
		}
	}
}

func WithProductHeaderSelector(selector string) MaltingOptionsFunc {
	return func(s *MaltingOptions) {
		if s.productHeaderSelector == nil {
			s.productHeaderSelector = &selector
		}
	}
}

func WithProductRowSelector(selector string) MaltingOptionsFunc {
	return func(s *MaltingOptions) {
		if s.productRowSelector == nil {
			s.productRowSelector = &selector
		}
	}
}

func WithBaseURL(url string) MaltingOptionsFunc {
	return func(s *MaltingOptions) {
		if s.baseURL == nil {
			s.baseURL = &url
		}
	}
}

func WithTitleSelector(selector string) MaltingOptionsFunc {
	return func(s *MaltingOptions) {
		if s.titleSelector == nil {
			s.titleSelector = &selector
		}
	}
}

func WithProductSelector(selector string) MaltingOptionsFunc {
	return func(s *MaltingOptions) {
		if s.productSelector == nil {
			s.productSelector = &selector
		}
	}
}

func WithStandard(standard fermentables.GrainType_StandardType) MaltingOptionsFunc {
	return func(s *MaltingOptions) {
		if s.standard == nil {
			s.standard = &standard
		}
	}
}

func WithGrainGroup(grainGroup beerproto.GrainGroup) MaltingOptionsFunc {
	return func(s *MaltingOptions) {
		if s.grainGroup == nil {
			s.grainGroup = &grainGroup
		}
	}
}

type Malting struct {
	formatter lxstrconv.NumberFormat
	country   string
	producer  string
	options   MaltingOptions
}

func NewMalting(producer string, country string, formatter lxstrconv.NumberFormat, options ...MaltingOptionsFunc) *Malting {
	options = append(options,
		WithGrainGroup(beerproto.GrainGroup_BASE),
		WithStandard(fermentables.GrainType_EBC),
		WithTitleSelector("h1"),
	)
	opts := &MaltingOptions{}
	for _, opt := range options {
		opt(opts)
	}

	return &Malting{
		formatter: formatter,
		producer:  producer,
		country:   country,
	}
}

func (s *Malting) Parse() []*fermentables.GrainType {
	grains := []*fermentables.GrainType{}

	c := colly.NewCollector()

	page := colly.NewCollector()

	page.OnHTML("body", func(e *colly.HTMLElement) {
		grain := &fermentables.GrainType{
			Country:    s.country,
			Standard:   *s.options.standard,
			GrainGroup: *s.options.grainGroup,
			Producer:   s.producer,
		}

		if s.options.titleSelector != nil {
			grain.Name = e.ChildText(*s.options.titleSelector)
		}

		if grain.Name == "" {
			return
		}

		e.ForEach(*s.options.productRowSelector, func(_ int, el *colly.HTMLElement) {

			text := strings.TrimSpace(el.ChildText(*s.options.productValueSelector))
			header := strings.TrimSpace(el.ChildText(*s.options.productHeaderSelector))

			if s.options.productValueTrim != nil {
				text = utils.TrimAny(header, *s.options.productValueTrim)
			}

			if s.options.productHeaderTrim != nil {
				header = utils.TrimAny(header, *s.options.productHeaderTrim)
			}

			if s.options.moisture != nil {
				if header == *s.options.moisture {
					grain.Moisture = unit.Percent(text,
						unit.WithFormatter[beerproto.PercentType_PercentUnitType](s.formatter))

				}
			}

			if s.options.yield != nil {
				if header == *s.options.yield {
					grain.Moisture = unit.Percent(text,
						unit.WithFormatter[beerproto.PercentType_PercentUnitType](s.formatter))

				}
			}

			if s.options.color != nil {
				if header == *s.options.color {
					grain.Color = unit.Color(text,
						unit.WithFormatter[beerproto.ColorUnitType](s.formatter),
						unit.WithUnit(beerproto.ColorUnitType_SRM))
				}
			}

			if s.options.protein != nil {
				if header == *s.options.protein {
					grain.Protein = unit.Percent(text, unit.WithFormatter[beerproto.PercentType_PercentUnitType](s.formatter))
				}
			}

			if s.options.maximum != nil {
				if header == *s.options.maximum {
					grain.Maximum = unit.Percent(text, unit.WithFormatter[beerproto.PercentType_PercentUnitType](s.formatter)).Maximum
				}
			}

			if s.options.friability != nil {
				if header == *s.options.friability {
					grain.Friability = unit.Percent(text, unit.WithFormatter[beerproto.PercentType_PercentUnitType](s.formatter))
				}
			}

			if s.options.betaGlucan != nil {
				if header == *s.options.betaGlucan {
					grain.BetaGlucan = unit.Concentration(text, unit.WithUnit(beerproto.ConcentrationUnitType_MGL),
						unit.WithFormatter[beerproto.ConcentrationUnitType](s.formatter))
				}
			}

			if s.options.alphaAmylase != nil {
				if header == *s.options.alphaAmylase {
					grain.AlphaAmylase = unit.Time(text, unit.WithFormatter[beerproto.TimeType_TimeUnitType](s.formatter)).Maximum
				}
			}

			if s.options.diPh != nil {
				if header == *s.options.diPh {
					grain.DiPh = unit.Acidity(text, unit.WithFormatter[beerproto.AcidityUnitType](s.formatter)).Maximum
				}
			}

			if s.options.fan != nil {
				if header == *s.options.fan {
					grain.Fan = unit.Concentration(text, unit.WithFormatter[beerproto.ConcentrationUnitType](s.formatter))
				}
			}

			if s.options.totalNitrogen != nil {
				if header == *s.options.totalNitrogen {
					grain.TotalNitrogen = unit.Percent(text, unit.WithFormatter[beerproto.PercentType_PercentUnitType](s.formatter))
				}
			}

		})

		grains = append(grains, grain)
	})

	c.OnHTML(*s.options.productSelector, func(e *colly.HTMLElement) {
		url := fmt.Sprintf("%s://%s%s", e.Request.URL.Scheme, e.Request.URL.Host, e.Attr("href"))
		page.Visit(url)
	})

	c.Visit(*s.options.baseURL)

	return grains
}
