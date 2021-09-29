package unpack

import (
	"testing"
	"wbschool_exam_L2/develop/dev02/pkq/unpack"

	"github.com/stretchr/testify/assert"
)

func TestUnpuck(t *testing.T) {

	testCases := []struct {
		str    string
		expStr string
	}{
		{
			str:    "a4bc2d5e",
			expStr: "aaaabccddddde",
		},
		{
			str:    "abcd",
			expStr: "abcd",
		},
		{
			str:    "",
			expStr: "",
		},
	}

	for _, test := range testCases {
		result, _ := unpack.Unpack(test.str)
		assert.Equal(t, result, test.expStr, "error")
	}
}

func TestUnpuckForError(t *testing.T) {
	var data = "45"
	_, err := unpack.Unpack(data)

	if err == nil {
		t.Errorf("test for OK Failed - error")
	}
}
