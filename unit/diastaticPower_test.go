package unit

import (
	"reflect"
	"testing"

	beerproto "github.com/beerproto/beerproto_go"
	"golang.org/x/text/language"
	"tawesoft.co.uk/go/lxstrconv"
)

func TestDiastaticPower(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		options []OptionsFunc[beerproto.DiastaticPowerUnitType]

		wantDiastaticPower *beerproto.DiastaticPowerRangeType
	}{
		{
			value: "> 250 WK ",
			options: []OptionsFunc[beerproto.DiastaticPowerUnitType]{
				WithFormatter[beerproto.DiastaticPowerUnitType](lxstrconv.NewDecimalFormat(language.German)),
			},
			wantDiastaticPower: &beerproto.DiastaticPowerRangeType{
				Minimum: &beerproto.DiastaticPowerType{
					Value: 250,
					Unit:  beerproto.DiastaticPowerUnitType_WK,
				},
			},
		},
		{
			value: "245 wk min",
			options: []OptionsFunc[beerproto.DiastaticPowerUnitType]{
				WithFormatter[beerproto.DiastaticPowerUnitType](lxstrconv.NewDecimalFormat(language.German)),
			},
			wantDiastaticPower: &beerproto.DiastaticPowerRangeType{
				Minimum: &beerproto.DiastaticPowerType{
					Value: 245,
					Unit:  beerproto.DiastaticPowerUnitType_WK,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDiastaticPower := DiastaticPower(tt.value, tt.options...); !reflect.DeepEqual(gotDiastaticPower, tt.wantDiastaticPower) {
				t.Errorf("DiastaticPower() = %v, want %v", gotDiastaticPower, tt.wantDiastaticPower)
			}
		})
	}
}
