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

	if ok := Contains(arr[0], []string{"min"}); ok {
		min := TrimAny(arr[0], []string{"min", "°ebc"})
		if min, err := formater.ParseFloat(min); err == nil {
			color.Minimum = &beerproto.ColorType{
				Value: min,
				Unit:  unit,
			}
		}

	} else if ok = Contains(arr[0], []string{"max"}); ok {
		max := TrimAny(arr[0], []string{"max", "°ebc"})
		if max, err := formater.ParseFloat(max); err == nil {
			color.Maximum = &beerproto.ColorType{
				Value: max,
				Unit:  unit,
			}
		}
	}
	if len(arr) == 1 {
		max := TrimAny(arr[0], []string{"max", "°ebc"})
		if max, err := formater.ParseFloat(max); err == nil {
			color.Maximum = &beerproto.ColorType{
				Value: max,
				Unit:  unit,
			}
		}
	}

	if len(arr) == 2 {
		min := TrimAny(arr[0], []string{"min", "°ebc"})
		if min, err := formater.ParseFloat(min); err == nil {
			color.Minimum = &beerproto.ColorType{
				Value: min,
				Unit:  unit,
			}
		}
		max := TrimAny(arr[1], []string{"max", "°ebc"})
		if max, err := formater.ParseFloat(max); err == nil {
			color.Maximum = &beerproto.ColorType{
				Value: max,
				Unit:  unit,
			}
		}

	}

	return
}

func percent(value string, formater lxstrconv.NumberFormat, unit beerproto.PercentType_PercentUnitType) (percent *beerproto.PercentRangeType) {
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

	if ok := Contains(arr[0], []string{"min"}); ok {
		min := TrimAny(arr[0], []string{"min", "%"})
		if min, err := formater.ParseFloat(min); err == nil {
			percent.Minimum = &beerproto.PercentType{
				Value: min,
				Unit:  unit,
			}
		}
	} else if ok = Contains(arr[0], []string{"max"}); ok {
		max := TrimAny(arr[0], []string{"max", "%"})
		if max, err := formater.ParseFloat(max); err == nil {
			percent.Maximum = &beerproto.PercentType{
				Value: max,
				Unit:  unit,
			}
		}

	} else {
		max := TrimAny(arr[0], []string{"%"})
		if max, err := formater.ParseFloat(max); err == nil {
			percent.Maximum = &beerproto.PercentType{
				Value: max,
				Unit:  unit,
			}
		}

	}

	if len(arr) == 2 {
		max := TrimAny(arr[1], []string{"max", "%"})
		if max, err := formater.ParseFloat(max); err == nil {
			percent.Maximum = &beerproto.PercentType{
				Value: max,
				Unit:  unit,
			}
		}

	}

	return
}

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
	match := false

	for _, l := range cutset {
		if strings.HasPrefix(s, l) {
			s = strings.TrimSpace(strings.TrimSuffix(s, l))
			match = true
		}
	}
	return match, ""
}

func EndsWith(s string, cutset []string) (bool, string) {
	s = strings.TrimSpace(s)
	match := false
	for _, l := range cutset {
		if strings.HasSuffix(s, l) {
			s = strings.TrimSpace(strings.TrimSuffix(s, l))
			match = true
		}
	}
	return match, s
}

func EndsWithArray(arr []string, cutset []string) (bool, string, []string) {
	for i, s := range arr {
		if ok, str := EndsWith(s, cutset); ok {
			return true, str, arr[i:]
		}
	}
	return false, "", arr
}

func Contains(s string, cutset []string) bool {
	s = strings.TrimSpace(s)
	words := strings.Split(s, " ")
	for _, l := range cutset {
		for _, w := range words {
			if l == w {
				return true
			}
		}
	}
	return false
}

func TrimAny(s string, cutset []string) string {
	s = strings.TrimSpace(s)

	for _, l := range cutset {
		s = strings.ReplaceAll(s, l, "")
	}

	return strings.TrimSpace(s)

}
