package grep

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"wbschool_exam_L2/develop/dev05/pkg/grep"
	"wbschool_exam_L2/develop/dev05/pkg/models"
	"wbschool_exam_L2/develop/dev05/pkg/root"
)

//console: myGrep o file.txt -A2 -n
func TestAfter(t *testing.T) {
	str, err := root.ReadLines("file.txt")

	if err != nil {
		t.Errorf("file not found")
		return
	}

	rightStr := []string{
		"1:Hello",
		"2:ass",
		"3:df",
		"11:bool sh",
		"12:a",
		"13:d",
	}

	//input data
	config := models.Configs{}
	config.FormatOut = models.A
	config.IsNum = true
	config.FormatPos = 2
	searchStr := "o"

	gs := grep.NewGrepString(config, str)
	gs.SearchString(searchStr)
	resStr := gs.GetStringResult()
	assert.Equal(t, resStr, rightStr, "error: incorrect expected strings")
}

//console: myGrep ar file.txt -B2 -n
func TestBefore(t *testing.T) {
	str, err := root.ReadLines("file.txt")

	if err != nil {
		t.Errorf("file not found")
		return
	}

	rightStr := []string{
		"3:df",
		"4:d",
		"5:har",
		"8:sa",
		"9:df",
		"10:arbuz",
	}

	//input data
	config := models.Configs{}
	config.FormatOut = models.B
	config.IsNum = true
	config.FormatPos = 2
	searchStr := "ar"

	gs := grep.NewGrepString(config, str)
	gs.SearchString(searchStr)
	resStr := gs.GetStringResult()
	assert.Equal(t, resStr, rightStr, "error: incorrect expected strings")
}

//console: myGrep h file.txt -C1 -n -i
func TestContext(t *testing.T) {
	str, err := root.ReadLines("file.txt")

	if err != nil {
		t.Errorf("file not found")
		return
	}

	rightStr := []string{
		"1:Hello",
		"2:ass",
		"4:d",
		"5:har",
		"6:wrld",
		"10:arbuz",
		"11:bool sh",
		"12:a",
	}

	//input data
	config := models.Configs{}
	config.FormatOut = models.C
	config.IsNum = true
	config.FormatPos = 1
	config.IsIgnore = true
	searchStr := "h"

	gs := grep.NewGrepString(config, str)
	gs.SearchString(searchStr)
	resStr := gs.GetStringResult()
	assert.Equal(t, resStr, rightStr, "error: incorrect expected strings")
}

//console: myGrep d file.txt -i -v -c
func TestCount(t *testing.T) {
	str, err := root.ReadLines("file.txt")

	if err != nil {
		t.Errorf("file not found")
		return
	}

	rightCount := 8

	//input data
	config := models.Configs{}
	config.IsIgnore = true
	config.IsInvert = true
	searchStr := "d"

	gs := grep.NewGrepString(config, str)
	resCount := gs.SearchString(searchStr)
	assert.Equal(t, resCount, rightCount, "error: incorrect expected strings")
}

//console: myGrep df file.txt -F -n
func TestFixed(t *testing.T) {
	str, err := root.ReadLines("file.txt")

	if err != nil {
		t.Errorf("file not found")
		return
	}

	rightStr := []string{
		"3:df",
		"9:df",
	}

	//input data
	config := models.Configs{}
	config.IsFixed = true
	config.IsNum = true
	searchStr := "df"

	gs := grep.NewGrepString(config, str)
	gs.SearchString(searchStr)
	resStr := gs.GetStringResult()
	assert.Equal(t, resStr, rightStr, "error: incorrect expected strings")
}
