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
		options   OptionsFunc[beerproto.DiastaticPowerUnitType]
	}{
		{
			value: "245 wk min",
			options: func(opts *Option[beerproto.DiastaticPowerUnitType]) {
				opts.WithFormatter(lxstrconv.NewDecimalFormat(language.German)).WithMinContains([]string{"min"}).WithMinTrim([]string{"wk", "min"}).WithSplitter([]string{"–", "-"}).WithUnit(beerproto.DiastaticPowerUnitType_WK)
			},
			rangeType: &RangeType[beerproto.DiastaticPowerUnitType]{},
			want: &RangeType[beerproto.DiastaticPowerUnitType]{
				Minimum: &UnitType[beerproto.DiastaticPowerUnitType]{
					Value: 245,
					Unit:  beerproto.DiastaticPowerUnitType_WK,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parse(tt.value, tt.rangeType, tt.options)
			if !reflect.DeepEqual(tt.rangeType, tt.want) {
				t.Errorf("parse() = %v, want %v", tt.rangeType, tt.want)
			}
		})
	}
}
