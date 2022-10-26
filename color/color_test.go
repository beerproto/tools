package color

import (
	"reflect"
	"testing"

	beerproto "github.com/beerproto/beerproto_go"

	"golang.org/x/text/language"
	"tawesoft.co.uk/go/lxstrconv"
)

func Test_color(t *testing.T) {
	tests := []struct {
		value     string
		options   []func(*ColorOption)
		wantColor *beerproto.ColorRangeType
	}{
		{
			value: "3.0 - 3.5",
			options: []func(*ColorOption){
				WithUnit(beerproto.ColorUnitType_EBC),
				WithFormatter(lxstrconv.NewDecimalFormat(language.BritishEnglish)),
			},
			wantColor: &beerproto.ColorRangeType{
				Minimum: &beerproto.ColorType{
					Unit:  beerproto.ColorUnitType_EBC,
					Value: 3,
				},
				Maximum: &beerproto.ColorType{
					Unit:  beerproto.ColorUnitType_EBC,
					Value: 3.5,
				},
			},
		},
		{
			value: "3.0",
			options: []func(*ColorOption){
				WithUnit(beerproto.ColorUnitType_EBC),
				WithFormatter(lxstrconv.NewDecimalFormat(language.BritishEnglish)),
			},
			wantColor: &beerproto.ColorRangeType{
				Maximum: &beerproto.ColorType{
					Unit:  beerproto.ColorUnitType_EBC,
					Value: 3,
				},
			},
		},
		{
			value: "max 3.5",
			options: []func(*ColorOption){
				WithUnit(beerproto.ColorUnitType_EBC),
				WithFormatter(lxstrconv.NewDecimalFormat(language.BritishEnglish)),
			},
			wantColor: &beerproto.ColorRangeType{
				Maximum: &beerproto.ColorType{
					Unit:  beerproto.ColorUnitType_EBC,
					Value: 3.5,
				},
			},
		},
		{
			value: "min 3.5",
			options: []func(*ColorOption){
				WithUnit(beerproto.ColorUnitType_EBC),
				WithFormatter(lxstrconv.NewDecimalFormat(language.BritishEnglish)),
			},
			wantColor: &beerproto.ColorRangeType{
				Minimum: &beerproto.ColorType{
					Unit:  beerproto.ColorUnitType_EBC,
					Value: 3.5,
				},
			},
		},
		{
			value: "min 3.0 - max 3.5",
			options: []func(*ColorOption){
				WithUnit(beerproto.ColorUnitType_EBC),
				WithFormatter(lxstrconv.NewDecimalFormat(language.BritishEnglish)),
			},
			wantColor: &beerproto.ColorRangeType{
				Minimum: &beerproto.ColorType{
					Unit:  beerproto.ColorUnitType_EBC,
					Value: 3,
				},
				Maximum: &beerproto.ColorType{
					Unit:  beerproto.ColorUnitType_EBC,
					Value: 3.5,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.value, func(t *testing.T) {
			if gotColor := Parse(tt.value, tt.options...); !reflect.DeepEqual(gotColor, tt.wantColor) {
				t.Errorf("color() = %v, want %v", gotColor, tt.wantColor)
			}
		})
	}
}