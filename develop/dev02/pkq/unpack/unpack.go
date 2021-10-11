package unpack

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

//Unpack return string "a4bc2d5e" => "aaaabccddddde"
func Unpack(str string) (result string, err error) {

	if str == "" {
		return
	}

	buffer := ""
	r := []rune(str)

	for i := 0; i < len(r); i++ {

		if buffer != "" {
			if unicode.IsLetter(r[i]) {
				result += buffer
				buffer = string(r[i])
			} else if string(r[i]) == "\\" && i+1 < len(r) {
				result += buffer
				buffer = string(r[i+1])
				i++
			} else if unicode.IsDigit(r[i]) {
				num := ""
				for ; i < len(r) && unicode.IsDigit(r[i]); i++ {
					num += string(r[i])
				}
				i--

				mult, err := strconv.Atoi(num)
				if err != nil {
					fmt.Fprintln(os.Stderr, "incorrect number", err)
					return "", err
				}
				result += strings.Repeat(buffer, mult)
				buffer = ""
			}
		} else {
			if unicode.IsLetter(r[i]) {
				buffer = string(r[i])
			} else if string(r[i]) == "\\" && i+1 < len(r) {
				buffer = string(r[i+1])
				i++
			} else {
				return "", fmt.Errorf("error")
			}
		}

		if i == len(r)-1 && buffer == string(r[i]) {
			result += buffer
		}
	}

	return
}
