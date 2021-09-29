package models

type Configs struct {
	IsIgnore  bool
	IsInvert  bool
	IsFixed   bool
	IsNum     bool
	FormatOut ForamSTDOUT
	FormatPos int
}

type ForamSTDOUT int

const (
	Defualt ForamSTDOUT = iota
	A
	B
	C
)
