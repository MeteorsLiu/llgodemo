package graphviz

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

type GVCOMMONS struct {
	Info             **int8
	Cmdname          *int8
	Verbose          c.Int
	Config           c.Int
	AutoOutfileNames c.Int
	Errorfn          unsafe.Pointer
	ShowBoxes        **int8
	Lib              **int8
	ViewNum          c.Int
	Builtins         *c.Int
	DemandLoading    c.Int
}
type GVCOMMONT GVCOMMONS
