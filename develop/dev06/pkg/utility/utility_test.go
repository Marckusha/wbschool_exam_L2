package utility

import (
	"testing"
	"wbschool_exam_L2/develop/dev06/pkg/root"
	"wbschool_exam_L2/develop/dev06/pkg/utility"

	"github.com/stretchr/testify/assert"
)

//cut file.txt -d: -f1
func TestFields1(t *testing.T) {
	conf := utility.Config{
		Dilimeter: ":",
		Fields:    []int{1},
	}

	strs, err := root.ReadLines("file.txt")

	if err != nil {
		t.Errorf("file not found")
		return
	}

	expectVal := []string{
		"Name",
		"Item1",
		"Item2",
		"Item3",
		"Item4",
		"Item5",
		"",
		"Item6",
	}

	util := utility.NewCutUtility(&conf, strs)
	res := util.ExecuteUtility()

	assert.Equal(t, res, expectVal, "error: incorrect expected strings")
}

//cut file.txt -d: -f1,4
func TestFields2(t *testing.T) {
	conf := utility.Config{
		Dilimeter: ":",
		Fields:    []int{1, 4},
	}

	strs, err := root.ReadLines("file.txt")

	if err != nil {
		t.Errorf("file not found")
		return
	}

	expectVal := []string{
		"Name:Ganre",
		"Item1:Ganre1",
		"Item2:Ganre2",
		"Item3:Ganre1",
		"Item4:Ganre1",
		"Item5:Ganre1",
		"",
		"Item6:Ganre2",
	}

	util := utility.NewCutUtility(&conf, strs)

	res := util.ExecuteUtility()
	assert.Equal(t, res, expectVal, "error: incorrect expected strings")
}

//cut file.txt -d: -f2,1,5,10 -s
func TestFields3(t *testing.T) {
	conf := utility.Config{
		Dilimeter:   ":",
		Fields:      []int{2, 1, 5, 10},
		IsSeparated: true,
	}

	strs, err := root.ReadLines("file.txt")

	if err != nil {
		t.Errorf("file not found")
		return
	}

	expectVal := []string{
		"Name:Company:Multiplayer",
		"Item1:Company1:Yes",
		"Item2:Company2:Yes",
		"Item3:Company3:Yes",
		"Item4:Company4:Yes",
		"Item5:Company5:Yes",
		"Item6:Company6:Yes",
	}

	util := utility.NewCutUtility(&conf, strs)

	res := util.ExecuteUtility()
	assert.Equal(t, res, expectVal, "error: incorrect expected strings")
}
