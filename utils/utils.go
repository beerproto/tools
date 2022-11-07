package utils

import (
	"strings"
)

func ToLower(cutset []string) []string {
	arr := []string{}
	for _, l := range cutset {
		arr = append(arr, strings.ToLower(l))
	}
	return arr
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
