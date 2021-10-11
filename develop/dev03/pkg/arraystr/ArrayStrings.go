package arraystr

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var months = []string{"jan", "feb", "mar", "apr", "may", "jun", "jul", "aug", "sep", "oct", "nov", "dec"}

//SortString хранит полное значение строки
//и строку-колонку, по которой будет производится сортировка
type SortString struct {
	Value      string //полная строка
	IndexValue string //строка-колонка, по которой сортируется
}

//ArrayStrings main arrays
type ArrayStrings []SortString

//NewArrayStrings create main arrays
func NewArrayStrings(str []string) ArrayStrings {
	var arr = make([]SortString, len(str))

	for i := 0; i < len(arr); i++ {
		arr[i].Value = str[i]
		arr[i].IndexValue = str[i]
	}

	return arr
}

//Equal сравнение двух массивов строк
func (arstr ArrayStrings) Equal(ar ArrayStrings) bool {
	if len(arstr) != len(ar) {
		return false
	}

	for i := 0; i < len(arstr); i++ {
		if arstr[i].Value != ar[i].Value {
			return false
		}
	}

	return true
}

//SetSortColumn устанавливаем колонку, по которой производится сортировка
//одновременно инициализируем IndexValue
func (arstr ArrayStrings) SetSortColumn(v int) {

	if v <= 0 {
		fmt.Println("invalid number at field start")
		return
	}

	for i := 0; i < len(arstr); i++ {
		ar := strings.Fields(arstr[i].Value)
		if len(ar) == 0 {
			continue
		}

		if v > len(ar) {
			arstr[i].IndexValue = ar[0]
		} else {
			arstr[i].IndexValue = ar[v-1]
		}
	}
}

//Unique Оставляет в массиве только уникальные значения
func (arstr *ArrayStrings) Unique() {
	m := make(map[string]struct{})
	var newStr []SortString

	for _, elem := range *arstr {
		if _, ok := m[elem.Value]; !ok {
			m[elem.Value] = struct{}{}
			newStr = append(newStr, elem)
		}
	}

	*arstr = newStr
}

//IgnoreSpace обрезает пробелы с концов строк
func (arstr ArrayStrings) IgnoreSpace() {

	for i := 0; i < len(arstr); i++ {
		var b strings.Builder
		b.Grow(len(arstr[i].Value))
		for _, ch := range arstr[i].Value {
			if !unicode.IsSpace(ch) {
				b.WriteRune(ch)
			}
		}
		arstr[i].Value = b.String()
	}
}

//StandartSort ...
func (arstr ArrayStrings) StandartSort(i, j int) bool {
	return arstr[i].IndexValue < arstr[j].IndexValue
}

//NumberSort ...
func (arstr ArrayStrings) NumberSort(i, j int) bool {

	fVal, err1 := strconv.Atoi(arstr[i].IndexValue)
	sVal, err2 := strconv.Atoi(arstr[j].IndexValue)

	if err1 != nil && err2 != nil {
		return arstr[i].IndexValue < arstr[j].IndexValue
	} else if err1 != nil {
		return true
	} else if err2 != nil {
		return false
	}

	return fVal < sVal
}

//MonthSort ...
func (arstr ArrayStrings) MonthSort(i, j int) bool {

	fVal, sVal := -1, -1
	for index, elem := range months {
		if strings.HasPrefix(strings.ToLower(arstr[i].IndexValue), elem) {
			fVal = index
		}
		if strings.HasPrefix(strings.ToLower(arstr[j].IndexValue), elem) {
			sVal = index
		}
	}

	if fVal == -1 && sVal == -1 {
		return arstr[i].IndexValue < arstr[j].IndexValue
	} else if fVal == -1 {
		return true
	} else if sVal == -1 {
		return false
	}

	return fVal <= sVal
}

//Reverse ...
func (arstr ArrayStrings) Reverse() {
	for i, j := 0, len(arstr)-1; i < j; i, j = i+1, j-1 {
		arstr[i], arstr[j] = arstr[j], arstr[i]
	}
}

//IsSorted проверка сортирован ли массив или нет
func (arstr ArrayStrings) IsSorted(typeSort func(i, j int) bool) bool {

	for i := 0; i < len(arstr)-1; i++ {
		if !typeSort(i, i+1) {
			return false
		}
	}

	return true
}
