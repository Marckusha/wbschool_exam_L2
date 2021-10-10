package root

import (
	"embed"
	"os"
	"strconv"
	"strings"
	"testing"
	"wbschool_exam_L2/develop/dev03/pkg/root"

	"github.com/stretchr/testify/assert"
)

//go:embed tests/*
var data embed.FS

func TestOk(t *testing.T) {

	args := make([]string, 0)

	countTests := 7
	for i := 1; i <= countTests; i++ {

		sNum := strconv.Itoa(i)
		pathTestFile := "tests/testFile" + sNum + ".txt"
		pathExpectFile := "tests/expectFile" + sNum + ".txt"
		pathArgs := "tests/args" + sNum + ".txt"

		expectFile, _ := data.ReadFile(pathExpectFile)

		argsData, _ := data.ReadFile(pathArgs)
		argsFile := strings.Split(string(argsData), "\n")

		args = []string{}
		copy(args, os.Args)
		args = append(args, pathTestFile)
		args = append(args, argsFile...)

		c := root.NewCommand()
		root.SetFlags(c)
		c.SetArgs(args)
		c.Execute()

		assert.Equal(t, root.TestString, string(expectFile), "incorrect result. Test № "+sNum)
	}
}
