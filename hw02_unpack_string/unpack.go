package hw02unpackstring

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(packed string) (string, error) {
	r := []rune(packed)
	if packed == "" {
		return "", nil
	}

	var s strings.Builder
	var isHas bool
	for i := 0; i < len(r); i++ {
		if i+1 < len(r) && unicode.IsDigit(r[i+1]) && r[i+1] == '0' {
			continue
		}

		if unicode.IsDigit(r[i]) {
			if isHas {
				count := int(r[i] - '0')
				if count == 0 {
					continue
				}
				s.WriteString(strings.Repeat(string(r[i-1]), count-1))
				isHas = false
				continue
			}
			if i == 0 || unicode.IsDigit(r[i-1]) {
				return "", ErrInvalidString
			}
			count := int(r[i] - '0')
			if count == 0 {
				continue
			}
			s.WriteString(strings.Repeat(string(r[i-1]), count-1))
			isHas = false
			continue
		}

		if r[i] == '\\' {
			if unicode.IsDigit(r[i+1]) || r[i+1] == '\\' {

				s.WriteRune(r[i+1])
				if unicode.IsDigit(r[i+1]) {
					isHas = true
				}
				i++
				continue
			}
		}

		s.WriteRune(r[i])
		isHas = false
	}
	return s.String(), nil
}
