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
		rangeType *RangeType[beerproto.DiastaticPowerUnitType, float64]
		want      *RangeType[beerproto.DiastaticPowerUnitType, float64]
		options   []OptionsFunc[beerproto.DiastaticPowerUnitType]
	}{
		{
			value: "245 wk min",
			options: []OptionsFunc[beerproto.DiastaticPowerUnitType]{
				WithFormatter[beerproto.DiastaticPowerUnitType](lxstrconv.NewDecimalFormat(language.German)),
				WithMinContains[beerproto.DiastaticPowerUnitType]([]string{"min"}),
				WithMinTrim[beerproto.DiastaticPowerUnitType]([]string{"wk", "min"}),
				WithUnit(beerproto.DiastaticPowerUnitType_WK),
			},
			rangeType: &RangeType[beerproto.DiastaticPowerUnitType, float64]{},
			want: &RangeType[beerproto.DiastaticPowerUnitType, float64]{
				Minimum: &UnitType[beerproto.DiastaticPowerUnitType, float64]{
					Value: 245,
					Unit:  beerproto.DiastaticPowerUnitType_WK,
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
