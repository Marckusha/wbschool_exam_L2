package main

import "wbschool_exam_L2/develop/dev05/pkg/root"

func main() {
	c := root.NewCommand()
	root.SetFlags(c)
	c.Execute()
}
