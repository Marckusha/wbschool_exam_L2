package anagram

import (
	"bytes"
	"sort"
	"strings"
)

//принимает массив строк и возвращает мапу с анаграммами согласно тз
func GetMapAnagram(str []string) map[string][]string {

	an := make(map[string]string) // мапа, в которой хрянятся массивы анаграмм ключей
	result := make(map[string][]string)

	for _, elem := range str {
		sortString := strSort(elem)

		if value, ok := an[sortString]; ok && value != elem {
			result[value] = append(result[value], elem)
		} else if value != elem {
			an[sortString] = elem
		}
	}

	for _, val := range result {
		sort.Strings(val)
	}

	return result
}

//функция для сортировки символов в строке
func strSort(str string) string {
	var (
		res       bytes.Buffer
		firstChar rune = 'а' //cyrillic 'а'
		countChar int  = 33
		counter        = make([]int, countChar)
	)

	str = strings.ToLower(str)

	for _, elem := range str {
		index := rune(elem) - firstChar
		if index >= 0 && index < rune(countChar) {
			counter[index] += 1
		} else {
			//return error
		}
	}

	for i := 0; i < countChar; i++ {
		ch := string(rune(i) + firstChar)
		res.WriteString(strings.Repeat(ch, counter[i]))
	}

	return res.String()
}
