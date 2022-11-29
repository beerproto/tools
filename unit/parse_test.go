package unit

import (
	"reflect"
	"testing"

	beerproto "github.com/beerproto/beerproto_go"
	"golang.org/x/text/language"
	"tawesoft.co.uk/go/lxstrconv"
)

func TestParse(t *testing.T) {

	tests := []struct {
		name      string
		value     string
		rangeType *RangeType[beerproto.DiastaticPowerUnit, float64]
		want      *RangeType[beerproto.DiastaticPowerUnit, float64]
		options   []OptionsFunc[beerproto.DiastaticPowerUnit]
	}{
		{
			value: "245 wk min",
			options: []OptionsFunc[beerproto.DiastaticPowerUnit]{
				WithFormatter[beerproto.DiastaticPowerUnit](lxstrconv.NewDecimalFormat(language.German)),
				WithMinContains[beerproto.DiastaticPowerUnit]([]string{"min"}),
				WithMinTrim[beerproto.DiastaticPowerUnit]([]string{"wk", "min"}),
				WithUnit(beerproto.DiastaticPowerUnit_DIASTATIC_POWER_UNIT_WK),
			},
			rangeType: &RangeType[beerproto.DiastaticPowerUnit, float64]{},
			want: &RangeType[beerproto.DiastaticPowerUnit, float64]{
				Minimum: &UnitType[beerproto.DiastaticPowerUnit, float64]{
					Value: 245,
					Unit:  beerproto.DiastaticPowerUnit_DIASTATIC_POWER_UNIT_WK,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parse(tt.value, tt.rangeType, tt.options...)
			if !reflect.DeepEqual(tt.rangeType, tt.want) {
				t.Errorf("parse() = %v, want %v", tt.rangeType, tt.want)
			}
		})
	}
}
