package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	var count int
	var symbol rune
	var shielding bool
	var res strings.Builder
	for k, r := range s {
		if string(r) == `\` && !shielding {
			shielding = true
			continue
		}

		if unicode.IsDigit(r) && !shielding {
			if symbol == 0 {
				return "", ErrInvalidString
			}
			count, _ = strconv.Atoi(string(r))
			res.WriteString(strings.Repeat(string(symbol), count))
			symbol = 0
			continue
		}

		if shielding && !(unicode.IsDigit(r) || string(r) == `\`) {
			return "", ErrInvalidString
		}

		shielding = false
		if symbol != 0 {
			res.WriteRune(symbol)
		}

		symbol = r
		if utf8.RuneCountInString(s)-1 == k {
			res.WriteRune(r)
		}
	}

	return res.String(), nil
}
