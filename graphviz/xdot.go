package graphviz

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

type XdotGradType c.Int

const (
	XdotGradTypeXdNone   XdotGradType = 0
	XdotGradTypeXdLinear XdotGradType = 1
	XdotGradTypeXdRadial XdotGradType = 2
)

type XdotColorStop struct {
	Frac  float32
	Color *int8
}

type XdotLinearGrad struct {
	X0     float64
	Y0     float64
	X1     float64
	Y1     float64
	NStops c.Int
	Stops  *XdotColorStop
}

type XdotRadialGrad struct {
	X0     float64
	Y0     float64
	R0     float64
	X1     float64
	Y1     float64
	R1     float64
	NStops c.Int
	Stops  *XdotColorStop
}

type XdotColor struct {
	Type XdotGradType
	U    struct {
		Ring XdotRadialGrad
	}
}
type XdotAlign c.Int

const (
	XdotAlignXdLeft   XdotAlign = 0
	XdotAlignXdCenter XdotAlign = 1
	XdotAlignXdRight  XdotAlign = 2
)

type XdotPoint struct {
	X float64
	Y float64
	Z float64
}

type XdotRect struct {
	X float64
	Y float64
	W float64
	H float64
}

type XdotPolyline struct {
	Cnt uintptr
	Pts *XdotPoint
}

type XdotText struct {
	X     float64
	Y     float64
	Align XdotAlign
	Width float64
	Text  *int8
}

type XdotImage struct {
	Pos  XdotRect
	Name *int8
}

type XdotFont struct {
	Size float64
	Name *int8
}
type XdotKind c.Int

const (
	XdotKindXdFilledEllipse   XdotKind = 0
	XdotKindXdUnfilledEllipse XdotKind = 1
	XdotKindXdFilledPolygon   XdotKind = 2
	XdotKindXdUnfilledPolygon XdotKind = 3
	XdotKindXdFilledBezier    XdotKind = 4
	XdotKindXdUnfilledBezier  XdotKind = 5
	XdotKindXdPolyline        XdotKind = 6
	XdotKindXdText            XdotKind = 7
	XdotKindXdFillColor       XdotKind = 8
	XdotKindXdPenColor        XdotKind = 9
	XdotKindXdFont            XdotKind = 10
	XdotKindXdStyle           XdotKind = 11
	XdotKindXdImage           XdotKind = 12
	XdotKindXdGradFillColor   XdotKind = 13
	XdotKindXdGradPenColor    XdotKind = 14
	XdotKindXdFontchar        XdotKind = 15
)

type XopKind c.Int

const (
	XopKindXopEllipse   XopKind = 0
	XopKindXopPolygon   XopKind = 1
	XopKindXopBezier    XopKind = 2
	XopKindXopPolyline  XopKind = 3
	XopKindXopText      XopKind = 4
	XopKindXopFillColor XopKind = 5
	XopKindXopPenColor  XopKind = 6
	XopKindXopFont      XopKind = 7
	XopKindXopStyle     XopKind = 8
	XopKindXopImage     XopKind = 9
	XopKindXopGradColor XopKind = 10
	XopKindXopFontchar  XopKind = 11
)

type X_XdotOp struct {
	Kind XdotKind
	U    struct {
		GradColor XdotColor
	}
	Drawfunc unsafe.Pointer
}
type XdotOp X_XdotOp
// llgo:type C
type DrawfuncT func(*XdotOp, c.Int)
// llgo:type C
type FreefuncT func(*XdotOp)

type Xdot struct {
	Cnt      uintptr
	Sz       uintptr
	Ops      *XdotOp
	Freefunc unsafe.Pointer
	Flags    c.Int
}

type XdotStats struct {
	Cnt          uintptr
	NEllipse     uintptr
	NPolygon     uintptr
	NPolygonPts  uintptr
	NPolyline    uintptr
	NPolylinePts uintptr
	NBezier      uintptr
	NBezierPts   uintptr
	NText        uintptr
	NFont        uintptr
	NStyle       uintptr
	NColor       uintptr
	NImage       uintptr
	NGradcolor   uintptr
	NFontchar    uintptr
}
//go:linkname ParseXDotF C.parseXDotF
func ParseXDotF(__llgo_arg_0 *int8, opfns *DrawfuncT, sz uintptr) *Xdot
// llgo:link (*Xdot).FreeXDot C.freeXDot
func (recv_ *Xdot) FreeXDot() {
}
