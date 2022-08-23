package scraper

import (
	"strings"

	beerproto "github.com/beerproto/beerproto_go"
	"tawesoft.co.uk/go/lxstrconv"
)

func acidity(value string, formater lxstrconv.NumberFormat) *beerproto.AcidityType {
	right := strings.ToLower(strings.TrimSpace(value))

	min := TrimRight(TrimLeft(right, []string{"min"}), []string{})
	if min, err := formater.ParseFloat(min); err == nil {
		acidity := &beerproto.AcidityType{
			Value: min,
			Unit:  beerproto.AcidityUnitType_PH,
		}
		return acidity

	}
	return nil

}

func time(value string, formater lxstrconv.NumberFormat) *beerproto.TimeRangeType {
	right := strings.ToLower(strings.TrimSpace(value))

	time := &beerproto.TimeRangeType{}

	min := TrimRight(TrimLeft(right, []string{"max"}), []string{""})
	if min, err := formater.ParseInt(min); err == nil {
		time.Minimum = &beerproto.TimeType{
			Value: min,
			Unit:  beerproto.TimeType_MIN,
		}
	}

	return time
}

func diastaticPower(value string, formater lxstrconv.NumberFormat) *beerproto.DiastaticPowerRangeType {
	right := strings.ToLower(strings.TrimSpace(value))

	time := &beerproto.DiastaticPowerRangeType{}

	min := TrimRight(TrimLeft(right, []string{"min"}), []string{""})
	if min, err := formater.ParseFloat(min); err == nil {
		time.Minimum = &beerproto.DiastaticPowerType{
			Value: min,
			Unit:  beerproto.DiastaticPowerUnitType_WK,
		}
	}

	return time
}

func concentration(value string, formater lxstrconv.NumberFormat, unit beerproto.ConcentrationUnitType) (concentration *beerproto.ConcentrationRangeType) {
	concentration = &beerproto.ConcentrationRangeType{}

	value = strings.ToLower(strings.TrimSpace(value))
	length := len(value)
	if length == 0 {
		return
	}

	arr := Split(value, []string{"–", "-"})

	if len(arr) == 0 {
		arr = []string{value}
	}

	ok := false
	min := ""
	max := ""

	if ok, min, arr = StartsWithArray(arr, []string{"min"}); ok {
		if min, err := formater.ParseFloat(min); err == nil {
			concentration.Minimum = &beerproto.ConcentrationType{
				Value: min,
				Unit:  unit,
			}
		}
	}

	if ok, max, arr = StartsWithArray(arr, []string{"max"}); ok {
		if max, err := formater.ParseFloat(max); err == nil {
			concentration.Maximum = &beerproto.ConcentrationType{
				Value: max,
				Unit:  unit,
			}

			return
		}
	}

	if len(arr) == 2 {
		min = arr[0]
		if min, err := formater.ParseFloat(min); err == nil {
			concentration.Minimum = &beerproto.ConcentrationType{
				Value: min,
				Unit:  unit,
			}
		}
		arr = arr[1:]
	}
	if len(arr) == 1 {
		max = arr[0]
		if max, err := formater.ParseFloat(max); err == nil {
			concentration.Maximum = &beerproto.ConcentrationType{
				Value: max,
				Unit:  unit,
			}
		}
	}
	return

	return concentration
}

func color(value string, formater lxstrconv.NumberFormat, unit beerproto.ColorUnitType) (color *beerproto.ColorRangeType) {
	color = &beerproto.ColorRangeType{}

	value = strings.ToLower(strings.TrimSpace(value))
	length := len(value)
	if length == 0 {
		return
	}

	arr := Split(value, []string{"–", "-"})

	if len(arr) == 0 {
		arr = []string{value}
	}

	ok := false
	min := ""
	max := ""

	if ok, min, arr = StartsWithArray(arr, []string{"min"}); ok {
		if min, err := formater.ParseFloat(min); err == nil {
			color.Minimum = &beerproto.ColorType{
				Value: min,
				Unit:  unit,
			}
		}
	}

	if ok, max, arr = StartsWithArray(arr, []string{"max"}); ok {
		if max, err := formater.ParseFloat(max); err == nil {
			color.Maximum = &beerproto.ColorType{
				Value: max,
				Unit:  unit,
			}

			return
		}
	}

	if ok, min, arr = EndsWithArray(arr, []string{"°ebc"}); ok {
		if min, err := formater.ParseFloat(min); err == nil {
			color.Minimum = &beerproto.ColorType{
				Value: min,
				Unit:  beerproto.ColorUnitType_EBC,
			}
		}
	}

	if ok, min, arr = EndsWithArray(arr, []string{"°ebc"}); ok {
		if max, err := formater.ParseFloat(min); err == nil {
			color.Maximum = &beerproto.ColorType{
				Value: max,
				Unit:  beerproto.ColorUnitType_EBC,
			}

			return
		}
	}

	if len(arr) == 2 {
		min = arr[0]
		if min, err := formater.ParseFloat(min); err == nil {
			color.Minimum = &beerproto.ColorType{
				Value: min,
				Unit:  unit,
			}
		}
		arr = arr[1:]
	}
	if len(arr) == 1 {
		max = arr[0]
		if max, err := formater.ParseFloat(max); err == nil {
			color.Maximum = &beerproto.ColorType{
				Value: max,
				Unit:  unit,
			}
		}
	}
	return
}

func percent(value string, formater lxstrconv.NumberFormat) (percent *beerproto.PercentRangeType) {
	percent = &beerproto.PercentRangeType{}

	value = strings.ToLower(strings.TrimSpace(value))
	length := len(value)
	if length == 0 {
		return
	}

	arr := Split(value, []string{"–", "-"})

	if len(arr) == 0 {
		arr = []string{value}
	}

	ok := false
	min := ""
	max := ""

	if ok, min, arr = StartsWithArray(arr, []string{"min"}); ok {
		if min, err := formater.ParseFloat(min); err == nil {
			percent.Minimum = &beerproto.PercentType{
				Value: min,
				Unit:  beerproto.PercentType_PERCENT_SIGN,
			}
		}
	}

	if ok, max, arr = StartsWithArray(arr, []string{"max"}); ok {
		if max, err := formater.ParseFloat(max); err == nil {
			percent.Maximum = &beerproto.PercentType{
				Value: max,
				Unit:  beerproto.PercentType_PERCENT_SIGN,
			}

			return
		}
	}

	if len(arr) == 2 {
		min = arr[0]
		if min, err := formater.ParseFloat(min); err == nil {
			percent.Minimum = &beerproto.PercentType{
				Value: min,
				Unit:  beerproto.PercentType_PERCENT_SIGN,
			}
		}
		arr = arr[1:]
	}
	if len(arr) == 1 {
		max = arr[0]
		_, max = StartsWith(max, []string{"%"})
		if max, err := formater.ParseFloat(max); err == nil {
			percent.Maximum = &beerproto.PercentType{
				Value: max,
				Unit:  beerproto.PercentType_PERCENT_SIGN,
			}
		}
	}
	return
}

// func percent(value string, formater lxstrconv.NumberFormat) (percent *beerproto.PercentRangeType) {
// 	percent = &beerproto.PercentRangeType{}
// 	right := strings.ToLower(strings.TrimSpace(value))

// 	if len(right) == 0 {
// 		return
// 	}

// 	if strings.HasSuffix(right, "max") || strings.HasPrefix(right, "max") {
// 		max := TrimRight(TrimLeft(right, []string{"max"}), []string{"%"})
// 		if max, err := formater.ParseFloat(max); err == nil {
// 			percent.Maximum = &beerproto.PercentType{
// 				Value: max,
// 				Unit:  beerproto.PercentType_PERCENT_SIGN,
// 			}
// 		}
// 	}

// 	if strings.HasSuffix(right, "min") || strings.HasPrefix(right, "min") {
// 		min := TrimRight(TrimLeft(right, []string{"min"}), []string{"%"})
// 		if min, err := formater.ParseFloat(min); err == nil {
// 			percent.Minimum = &beerproto.PercentType{
// 				Value: min,
// 				Unit:  beerproto.PercentType_PERCENT_SIGN,
// 			}
// 		}
// 	}

// 	return
// }

func viscosity(value string, formater lxstrconv.NumberFormat) *beerproto.ViscosityRangeType {
	right := strings.ToLower(strings.TrimSpace(value))

	viscosity := &beerproto.ViscosityRangeType{}

	min := TrimRight(TrimLeft(right, []string{"máx", "max", "."}), []string{"cp"})
	if min, err := formater.ParseFloat(min); err == nil {
		viscosity.Minimum = &beerproto.ViscosityType{
			Value: min,
			Unit:  beerproto.ViscosityUnitType_CP,
		}
	}

	return viscosity
}

func TrimLeft(s string, cutset []string) string {
	for _, l := range cutset {
		s = strings.TrimLeft(s, l)
	}
	return s
}

func TrimRight(s string, cutset []string) string {
	for _, l := range cutset {
		s = strings.TrimRight(s, l)
	}
	return s
}

func TrimSpace(arr []string) []string {
	result := []string{}
	for _, l := range arr {
		result = append(result, strings.TrimSpace(l))
	}
	return result
}

func Index(s string, cutset []string) int {
	for _, l := range cutset {
		i := strings.Index(s, l)
		if i > 0 {
			return i
		}
	}
	return len(s)
}

func Split(s string, cutset []string) []string {
	arr := []string{}
	for _, l := range cutset {
		if strings.Contains(s, l) {
			arr = append(arr, strings.Split(s, l)...)
		}
	}
	return TrimSpace(arr)
}

func StartsWithArray(arr []string, cutset []string) (bool, string, []string) {
	for i, s := range arr {
		if ok, str := StartsWith(s, cutset); ok {
			return true, str, arr[i:]
		}
	}
	return false, "", arr
}

func StartsWith(s string, cutset []string) (bool, string) {
	s = strings.TrimSpace(s)
	for _, l := range cutset {
		if strings.HasPrefix(s, l) {
			return true, strings.TrimPrefix(s, l)
		}
	}
	return false, ""
}

func EndsWith(s string, cutset []string) (bool, string) {
	s = strings.TrimSpace(s)
	for _, l := range cutset {
		if strings.HasSuffix(s, l) {
			return true, strings.TrimSuffix(s, l)
		}
	}
	return false, ""
}

func EndsWithArray(arr []string, cutset []string) (bool, string, []string) {
	for i, s := range arr {
		if ok, str := EndsWith(s, cutset); ok {
			return true, str, arr[i:]
		}
	}
	return false, "", arr
}
