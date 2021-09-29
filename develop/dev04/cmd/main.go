package main

import (
	"fmt"
	"wbschool_exam_L2/develop/dev04/pkg/anagram"
)

func main() {
	str := []string{"пятка", "яткап", "листок", "конверт", "пятак", "слиток", "столик", "тяпка", "каптер", "паркет", "отец"}
	m := anagram.GetMapAnagram(str)
	fmt.Println(m)
}
