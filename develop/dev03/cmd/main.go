package main

import (
	"fmt"
	"wbschool_exam_L2/develop/dev03/pkg/root"
)

func main() {
	//os.Args = append(os.Args, "file1.txt", "-c")

	c := root.NewCommand()
	root.SetFlags(c)
	c.Execute()
	fmt.Print(root.TestString)
}
