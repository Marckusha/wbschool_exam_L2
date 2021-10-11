package models

//Configs consist args for grep utility
type Configs struct {
	IsIgnore  bool
	IsInvert  bool
	IsFixed   bool
	IsNum     bool
	FormatOut ForamSTDOUT
	FormatPos int
}

//ForamSTDOUT type algorithm for grep
type ForamSTDOUT int

//Defualt ...
const (
	Defualt ForamSTDOUT = iota
	A
	B
	C
)
