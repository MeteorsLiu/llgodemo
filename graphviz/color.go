package graphviz

import (
	"github.com/goplus/llgo/c"
	_ "unsafe"
)

type HsvrgbacolorT struct {
	Name *int8
	H    int8
	S    int8
	V    int8
	R    int8
	G    int8
	B    int8
	A    int8
}
type ColorTypeT c.Int

const (
	ColorTypeTHSVADOUBLE  ColorTypeT = 0
	ColorTypeTRGBABYTE    ColorTypeT = 1
	ColorTypeTRGBAWORD    ColorTypeT = 2
	ColorTypeTRGBADOUBLE  ColorTypeT = 3
	ColorTypeTCOLORSTRING ColorTypeT = 4
	ColorTypeTCOLORINDEX  ColorTypeT = 5
)

type ColorS struct {
	U struct {
		RGBA [4]float64
	}
	Type ColorTypeT
}
type ColorT ColorS
