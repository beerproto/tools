package unit

import (
	beerproto "github.com/beerproto/beerproto_go"
)

func Concentration(value string, options ...OptionsFunc[beerproto.ConcentrationUnit]) *beerproto.ConcentrationRangeType {
	rangeType := &RangeType[beerproto.ConcentrationUnit, float64]{}

	options = append(options,
		WithMinContains[beerproto.ConcentrationUnit]([]string{">", "min"}),
		WithMinTrim[beerproto.ConcentrationUnit]([]string{"mg/l", ">", "min"}),
		WithMaxTrim[beerproto.ConcentrationUnit]([]string{"mg/l", "<", "max"}),
		WithMaxContains[beerproto.ConcentrationUnit]([]string{"<", "max"}),
		WithUnit(beerproto.ConcentrationUnit_CONCENTRATION_UNIT_MG100L),
		WithDefault[beerproto.ConcentrationUnit](Max),
	)

	parse(value, rangeType, options...)

	concentration := &beerproto.ConcentrationRangeType{}

	if rangeType.Maximum != nil {
		concentration.Maximum = &beerproto.ConcentrationType{
			Unit:  rangeType.Maximum.Unit,
			Value: rangeType.Maximum.Value,
		}
	}

	if rangeType.Minimum != nil {
		concentration.Minimum = &beerproto.ConcentrationType{
			Unit:  rangeType.Minimum.Unit,
			Value: rangeType.Minimum.Value,
		}
	}

	return concentration
}
