package unit

import (
	beerproto "github.com/beerproto/beerproto_go"
)

func Concentration(value string, options ...OptionsFunc[beerproto.ConcentrationUnitType]) *beerproto.ConcentrationRangeType {
	rangeType := &RangeType[beerproto.ConcentrationUnitType, float64]{}

	options = append(options,
		WithMinContains[beerproto.ConcentrationUnitType]([]string{">", "min"}),
		WithMinTrim[beerproto.ConcentrationUnitType]([]string{"mg/l", ">", "min"}),
		WithMaxTrim[beerproto.ConcentrationUnitType]([]string{"mg/l", "<", "max"}),
		WithMaxContains[beerproto.ConcentrationUnitType]([]string{"<", "max"}),
		WithUnit(beerproto.ConcentrationUnitType_MG100L),
		WithDefault[beerproto.ConcentrationUnitType](Max),
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
