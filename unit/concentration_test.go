package unit

import (
	"reflect"
	"testing"

	beerproto "github.com/beerproto/beerproto_go"
	"golang.org/x/text/language"
	"tawesoft.co.uk/go/lxstrconv"
)

func TestConcentration(t *testing.T) {
	tests := []struct {
		name                   string
		value                  string
		options                []OptionsFunc[beerproto.ConcentrationUnitType]
		wantConcentrationRange *beerproto.ConcentrationRangeType
	}{
		{
			value: "630 -730",
			options: []OptionsFunc[beerproto.ConcentrationUnitType]{
				WithFormatter[beerproto.ConcentrationUnitType](lxstrconv.NewDecimalFormat(language.BritishEnglish)),
				WithUnit(beerproto.ConcentrationUnitType_MGL),
			},
			wantConcentrationRange: &beerproto.ConcentrationRangeType{
				Minimum: &beerproto.ConcentrationType{
					Value: 630,
					Unit:  beerproto.ConcentrationUnitType_MGL,
				},
				Maximum: &beerproto.ConcentrationType{
					Value: 730,
					Unit:  beerproto.ConcentrationUnitType_MGL,
				},
			},
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDiastaticPower := Concentration(tt.value, tt.options...); !reflect.DeepEqual(gotDiastaticPower, tt.wantConcentrationRange) {
				t.Errorf("Concentration() = %v, want %v", gotDiastaticPower, tt.wantConcentrationRange)
			}
		})
	}
}
