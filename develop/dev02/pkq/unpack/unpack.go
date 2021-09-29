package unpack

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func Unpack(str string) (string, error) {
	var result string

	if str == "" {
		return "", nil
	}

	r := []rune(str)

	for i := 0; i < len(r); i++ {
		if unicode.IsLetter(r[i]) {
			if i+1 < len(r) && unicode.IsLetter(r[i+1]) || i == len(r)-1 {
				result += string(r[i])
			} else if i+1 < len(r) && unicode.IsDigit(r[i+1]) {
				char := string(r[i])
				mult, err := strconv.Atoi(string(r[i+1]))
				if err == nil {
				}
				result += strings.Repeat(char, mult)
				i++
			}
		} else if string(r[i]) == "\\" {

		} else {
			return "", fmt.Errorf("error")
		}
	}

	return result, nil
}
