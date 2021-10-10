package main

import (
	"wbschool_exam_L2/develop/dev06/pkg/root"
)

func main() {
	comm := root.NewCommand()
	root.SetFlags(comm)
	comm.Execute()
}
