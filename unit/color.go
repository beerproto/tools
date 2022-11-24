package unit

import (
	beerproto "github.com/beerproto/beerproto_go"
	"github.com/beerproto/beerproto_go/fermentables"
)

func WithColorFromStandard[TUnit Unit](standard fermentables.GrainType_StandardType) OptionsFunc[TUnit] {
	return func(s *Option[TUnit]) {
		switch standard {
		case fermentables.GrainType_EBC:
			WithUnit(beerproto.ColorUnitType_EBC)
			return
		case fermentables.GrainType_ASBC:
			WithUnit(beerproto.ColorUnitType_SRM)
			return
		case fermentables.GrainType_ION:
			WithUnit(beerproto.ColorUnitType_LOVI)
			return
		}
	}
}

func Color(value string, options ...OptionsFunc[beerproto.ColorUnitType]) *beerproto.ColorRangeType {
	rangeType := &RangeType[beerproto.ColorUnitType, float64]{}

	options = append(options,
		WithMinContains[beerproto.ColorUnitType]([]string{"±5", "min"}),
		WithMinTrim[beerproto.ColorUnitType]([]string{"±5", "EBC", "min"}),
		WithMaxTrim[beerproto.ColorUnitType]([]string{"EBC", "max"}),
		WithMaxContains[beerproto.ColorUnitType]([]string{"max"}),
		WithUnit(beerproto.ColorUnitType_EBC),
		WithDefault[beerproto.ColorUnitType](Max),
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
