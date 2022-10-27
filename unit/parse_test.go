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
		rangeType *RangeType[beerproto.DiastaticPowerUnitType]
		want      *RangeType[beerproto.DiastaticPowerUnitType]

		options []func(*Option[beerproto.DiastaticPowerUnitType])
	}{
		{
			value: "245 wk min",
			options: []func(*Option[beerproto.DiastaticPowerUnitType]){
				WithSplitter[beerproto.DiastaticPowerUnitType]([]string{"–", "-"}),
				WithUnit(beerproto.DiastaticPowerUnitType_WK),
				WithFormatter[beerproto.DiastaticPowerUnitType](lxstrconv.NewDecimalFormat(language.German)),
				WithMinContains[beerproto.DiastaticPowerUnitType]([]string{"min"}),
				WithMinTrim[beerproto.DiastaticPowerUnitType]([]string{"wk", "min"}),
			},
			rangeType: &RangeType[beerproto.DiastaticPowerUnitType]{},
			want: &RangeType[beerproto.DiastaticPowerUnitType]{
				Minimum: UnitType[beerproto.DiastaticPowerUnitType]{
					Value: 245,
					Unit:  beerproto.DiastaticPowerUnitType_WK,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Parse(tt.value, tt.rangeType, tt.options...)
			if !reflect.DeepEqual(tt.rangeType, tt.want) {
				t.Errorf("Parse() = %v, want %v", tt.rangeType, tt.want)
			}
		})
	}
}
