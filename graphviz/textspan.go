package graphviz

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

type X_PostscriptAlias struct {
	Name          *int8
	Family        *int8
	Weight        *int8
	Stretch       *int8
	Style         *int8
	XfigCode      c.Int
	SvgFontFamily *int8
	SvgFontWeight *int8
	SvgFontStyle  *int8
}
type PostscriptAlias X_PostscriptAlias

type TextfontT struct {
	Name            *int8
	Color           *int8
	PostscriptAlias *PostscriptAlias
	Size            float64
	Flags           c.Uint
	Cnt             c.Uint
}

type TextspanT struct {
	Str               *int8
	Font              *TextfontT
	Layout            unsafe.Pointer
	FreeLayout        unsafe.Pointer
	YoffsetLayout     float64
	YoffsetCenterline float64
	Size              Pointf
	Just              int8
}
