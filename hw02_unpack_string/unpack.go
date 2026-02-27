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
	for i := 0; i < len(r); i++ {

		if i+1 < len(r) && unicode.IsDigit(r[i+1]) && r[i+1] == '0' {
			continue
		}

		if unicode.IsDigit(r[i]) {
			if i == 0 || unicode.IsDigit(r[i-1]) {
				return "", ErrInvalidString
			}
			count := int(r[i] - '0')
			if count == 0 {
				continue
			}

			s.WriteString(strings.Repeat(string(r[i-1]), count-1))
			continue
		}

		s.WriteRune(r[i])
	}
	return s.String(), nil
}
