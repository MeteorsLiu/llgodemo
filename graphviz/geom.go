package graphviz

import (
	"github.com/goplus/llgo/c"
	_ "unsafe"
)

type Point struct {
	X c.Int
	Y c.Int
}

type PointfS struct {
	X float64
	Y float64
}
type Pointf PointfS

type Linef struct {
	P Pointf
	M float64
}

type Box struct {
	LL Point
	UR Point
}

type Boxf struct {
	LL Pointf
	UR Pointf
}
