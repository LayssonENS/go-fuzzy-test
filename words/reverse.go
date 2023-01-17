package words

import (
	"errors"
	"unicode/utf8"
)

func Reverse(s string) (string, error) {
	var response string
	r := []rune(s)

	if !utf8.ValidString(s) {
		return s, errors.New("invalid")
	}
	for i := len(r) - 1; i >= 0; i-- {
		response = response + string(r[i])
	}

	return response, nil
}
