package graphviz

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

type Prwr struct {
	Unused [8]uint8
}
// llgo:type C
type Pruserfn func(*int8) c.Int

type Prbinding struct {
	Name *int8
	Fn   unsafe.Pointer
}

type Propts struct {
	Ingraphs   **AgraphT
	NOutgraphs uintptr
	Outgraphs  **AgraphT
	Out        Prwr
	Err        Prwr
	Flags      c.Int
	Bindings   *Prbinding
}
