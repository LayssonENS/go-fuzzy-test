package words

import (
	"errors"
	"unicode/utf8"
)

func Reverse(s string) (string, error) {
	if !utf8.ValidString(s) {
		return s, errors.New("invalid")
	}
	b := []byte(s)
	var response []byte
	for i := len(b) - 1; i >= 0; i-- {
		response = append(response, b[i])
	}
	return string(response), nil
}
