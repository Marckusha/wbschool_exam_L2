package anagram

import (
	"testing"
	"wbschool_exam_L2/develop/dev04/pkg/anagram"

	"github.com/stretchr/testify/assert"
)

func TestOk1(t *testing.T) {
	testData := []string{"пятка", "яткап", "листок", "конверт", "пятак", "слиток", "столик", "тяпка", "каптер", "паркет", "отец"}
	expectData := map[string][]string{
		"каптер": {"паркет"},
		"листок": {"слиток", "столик"},
		"пятка":  {"пятак", "тяпка", "яткап"},
	}

	m := anagram.GetMapAnagram(testData)

	assert.Equal(t, m, expectData, "not equal values")
}

func TestOk2(t *testing.T) {
	testData := []string{}
	expectData := map[string][]string{}

	m := anagram.GetMapAnagram(testData)

	assert.Equal(t, m, expectData, "not equal values")
}

func TestOk3(t *testing.T) {
	testData := []string{"арбуз", "зураб", "листок", "бузар", "слиток", "тяпка"}
	expectData := map[string][]string{
		"арбуз":  {"бузар", "зураб"},
		"листок": {"слиток"},
	}

	m := anagram.GetMapAnagram(testData)

	assert.Equal(t, m, expectData, "not equal values")
}

func TestOk4(t *testing.T) {
	testData := []string{"арбуз"}
	expectData := map[string][]string{}

	m := anagram.GetMapAnagram(testData)

	assert.Equal(t, m, expectData, "not equal values")
}
