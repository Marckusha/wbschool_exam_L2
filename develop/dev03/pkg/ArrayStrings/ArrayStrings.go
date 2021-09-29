package ArrayStrings

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var months = []string{"jan", "feb", "mar", "apr", "may", "jun", "jul", "aug", "sep", "oct", "nov", "dec"}

//хранит полное значение строки
//и строку, по которой будет производится сортировка
type SortString struct {
	Value      string //строка
	IndexValue string //строка, по которой сортируется
}

type ArrayStrings []SortString

func NewArrayStrings(str []string) ArrayStrings {
	var arr = make([]SortString, len(str))

	for i := 0; i < len(arr); i++ {
		arr[i].Value = str[i]
		arr[i].IndexValue = str[i]
	}

	return arr
}

func (ar1 ArrayStrings) Equal(ar2 ArrayStrings) bool {
	if len(ar1) != len(ar2) {
		return false
	}

	for i := 0; i < len(ar1); i++ {
		if ar1[i].Value != ar2[i].Value {
			return false
		}
	}

	return true
}

//устанавливаем колонку, по которой производится сортировка
//одновременно инициализируем IndexValue
func (str ArrayStrings) SetSortColumn(v int) {

	if v <= 0 {
		fmt.Println("invalid number at field start")
		return
	}

	for i := 0; i < len(str); i++ {
		ar := strings.Fields(str[i].Value)
		if len(ar) == 0 {
			continue
		}

		if v > len(ar) {
			str[i].IndexValue = ar[0]
		} else {
			str[i].IndexValue = ar[v-1]
		}
	}
}

//Оставляет в массиве только уникальные значения
func (str *ArrayStrings) Unique() {
	m := make(map[string]struct{})
	var newStr []SortString

	for _, elem := range *str {
		if _, ok := m[elem.Value]; !ok {
			m[elem.Value] = struct{}{}
			newStr = append(newStr, elem)
		}
	}

	*str = newStr
}

//обрезает пробелы с концов строк
func (ar ArrayStrings) IgnoreSpace() {

	for i := 0; i < len(ar); i++ {
		var b strings.Builder
		b.Grow(len(ar[i].Value))
		for _, ch := range ar[i].Value {
			if !unicode.IsSpace(ch) {
				b.WriteRune(ch)
			}
		}
		ar[i].Value = b.String()
	}
}

func (s ArrayStrings) StandartSort(i, j int) bool {
	return s[i].IndexValue < s[j].IndexValue
}

func (s ArrayStrings) NumberSort(i, j int) bool {

	fVal, err1 := strconv.Atoi(s[i].IndexValue)
	sVal, err2 := strconv.Atoi(s[j].IndexValue)

	if err1 != nil && err2 != nil {
		return s[i].IndexValue < s[j].IndexValue
	} else if err1 != nil {
		return true
	} else if err2 != nil {
		return false
	}

	return fVal < sVal
}

func (s ArrayStrings) MonthSort(i, j int) bool {

	fVal, sVal := -1, -1
	for index, elem := range months {
		if strings.HasPrefix(strings.ToLower(s[i].IndexValue), elem) {
			fVal = index
		}
		if strings.HasPrefix(strings.ToLower(s[j].IndexValue), elem) {
			sVal = index
		}
	}

	if fVal == -1 && sVal == -1 {
		return s[i].IndexValue < s[j].IndexValue
	} else if fVal == -1 {
		return true
	} else if sVal == -1 {
		return false
	}

	return fVal <= sVal
}

func (s ArrayStrings) Reverse() {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

//проверка сортирован ли массив или нет
func (s ArrayStrings) IsSorted(typeSort func(i, j int) bool) bool {

	for i := 0; i < len(s)-1; i++ {
		fmt.Println()
		if !typeSort(i, i+1) {
			//fmt.Println("disorder:", s[i+1])
			return false
		}
	}

	return true
}

//тест вывод
func (ar ArrayStrings) Print() {
	for _, elem := range ar {
		fmt.Println(elem.Value)
	}
}

//тест вывод
func (ar ArrayStrings) PrintIndexValue() {
	fmt.Println("Index value")
	for _, elem := range ar {
		fmt.Println(elem.IndexValue)
	}
}
