package diastaticPower

import (
	beerproto "github.com/beerproto/beerproto_go"
	"github.com/beerproto/tools/unit"
	"golang.org/x/text/language"
	"tawesoft.co.uk/go/lxstrconv"
)

type DiastaticPowerOption struct {
	formatter lxstrconv.NumberFormat
	unit      beerproto.DiastaticPowerUnitType
}

func WithFormatter(formatter lxstrconv.NumberFormat) func(*DiastaticPowerOption) {
	return func(s *DiastaticPowerOption) {
		s.formatter = formatter
	}
}

func WithUnit(unit beerproto.DiastaticPowerUnitType) func(*DiastaticPowerOption) {
	return func(s *DiastaticPowerOption) {
		s.unit = unit
	}
}

func Parse(value string, options ...func(*DiastaticPowerOption)) (diastaticPower *beerproto.DiastaticPowerRangeType) {
	rangeType := &unit.RangeType[beerproto.DiastaticPowerUnitType]{}

	opts := []func(*unit.Option[beerproto.DiastaticPowerUnitType]){
		unit.WithSplitter[beerproto.DiastaticPowerUnitType]([]string{"–", "-"}),
		unit.WithUnit(beerproto.DiastaticPowerUnitType_WK),
		unit.WithFormatter[beerproto.DiastaticPowerUnitType](lxstrconv.NewDecimalFormat(language.German)),
		unit.WithMinContains[beerproto.DiastaticPowerUnitType]([]string{"min"}),
		unit.WithMinTrim[beerproto.DiastaticPowerUnitType]([]string{"wk", "min"}),
	}

	unit.Parse(value, rangeType, opts...)

	return nil
}
