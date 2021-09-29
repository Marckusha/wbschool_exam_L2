package main

import (
	"fmt"
	"wbschool_exam_L2/develop/dev02/pkq/unpack"
)

/*
"a4bc2d5e" => "aaaabccddddde"
"abcd" => "abcd"
"45" => "" (некорректная строка)
"" => ""
Дополнительное задание: поддержка escape - последовательностей
qwe\4\5 => qwe45 (*)
qwe\45 => qwe44444 (*)
qwe\\5 => qwe\\\\\ (*)
*/

func main() {
	var s string = "a4bc2d5e"
	res, _ := unpack.Unpack(s)
	fmt.Println(res)
}
