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
		options []OptionsFunc[beerproto.DiastaticPowerUnit]

		wantDiastaticPower *beerproto.DiastaticPowerRangeType
	}{
		{
			value: "200 Wk",
			options: []OptionsFunc[beerproto.DiastaticPowerUnit]{
				WithFormatter[beerproto.DiastaticPowerUnit](lxstrconv.NewDecimalFormat(language.German)),
			},
			wantDiastaticPower: &beerproto.DiastaticPowerRangeType{
				Minimum: &beerproto.DiastaticPowerType{
					Value: 200,
					Unit:  beerproto.DiastaticPowerUnit_DIASTATIC_POWER_UNIT_WK,
				},
			},
		},
		{
			value: "> 250 WK ",
			options: []OptionsFunc[beerproto.DiastaticPowerUnit]{
				WithFormatter[beerproto.DiastaticPowerUnit](lxstrconv.NewDecimalFormat(language.German)),
			},
			wantDiastaticPower: &beerproto.DiastaticPowerRangeType{
				Minimum: &beerproto.DiastaticPowerType{
					Value: 250,
					Unit:  beerproto.DiastaticPowerUnit_DIASTATIC_POWER_UNIT_WK,
				},
			},
		},
		{
			value: "245 wk min",
			options: []OptionsFunc[beerproto.DiastaticPowerUnit]{
				WithFormatter[beerproto.DiastaticPowerUnit](lxstrconv.NewDecimalFormat(language.German)),
			},
			wantDiastaticPower: &beerproto.DiastaticPowerRangeType{
				Minimum: &beerproto.DiastaticPowerType{
					Value: 245,
					Unit:  beerproto.DiastaticPowerUnit_DIASTATIC_POWER_UNIT_WK,
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
