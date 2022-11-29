package unit

import (
	beerproto "github.com/beerproto/beerproto_go"
	"github.com/beerproto/beerproto_go/fermentables"
)

func WithColorFromStandard[TUnit Unit](standard fermentables.GrainType_StandardType) OptionsFunc[TUnit] {
	return func(s *Option[TUnit]) {
		switch standard {
		case fermentables.GrainType_EBC:
			WithUnit(beerproto.ColorUnit_COLOR_UNIT_EBC)
			return
		case fermentables.GrainType_ASBC:
			WithUnit(beerproto.ColorUnit_COLOR_UNIT_SRM)
			return
		case fermentables.GrainType_ION:
			WithUnit(beerproto.ColorUnit_COLOR_UNIT_LOVI)
			return
		}
	}
}

func Color(value string, options ...OptionsFunc[beerproto.ColorUnit]) *beerproto.ColorRangeType {
	rangeType := &RangeType[beerproto.ColorUnit, float64]{}

	options = append(options,
		WithMinContains[beerproto.ColorUnit]([]string{"±5", "min"}),
		WithMinTrim[beerproto.ColorUnit]([]string{"±5", "EBC", "min"}),
		WithMaxTrim[beerproto.ColorUnit]([]string{"EBC", "max"}),
		WithMaxContains[beerproto.ColorUnit]([]string{"max"}),
		WithUnit(beerproto.ColorUnit_COLOR_UNIT_EBC),
		WithDefault[beerproto.ColorUnit](Max),
	)

	parse(value, rangeType, options...)

	color := &beerproto.ColorRangeType{}

	if rangeType.Maximum != nil {
		color.Maximum = &beerproto.ColorType{
			Unit:  rangeType.Maximum.Unit,
			Value: rangeType.Maximum.Value,
		}
	}

	if rangeType.Minimum != nil {
		color.Minimum = &beerproto.ColorType{
			Unit:  rangeType.Minimum.Unit,
			Value: rangeType.Minimum.Value,
		}
	}

	return color
}
