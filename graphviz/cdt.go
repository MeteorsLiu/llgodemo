package graphviz

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

type X_DtlinkS struct {
	Right *DtlinkT
	Hl    struct {
		X_Left *DtlinkT
	}
}
type DtlinkT X_DtlinkS

type X_DtholdS struct {
	Hdr DtlinkT
	Obj unsafe.Pointer
}
type DtholdT X_DtholdS

type X_DtdiscS struct {
	Key     c.Int
	Size    c.Int
	Link    c.Int
	Makef   unsafe.Pointer
	Freef   unsafe.Pointer
	Comparf unsafe.Pointer
}
type DtdiscT X_DtdiscS

type X_DtmethodS struct {
	Searchf unsafe.Pointer
	Type    c.Int
}
type DtmethodT X_DtmethodS

type X_DtdataS struct {
	Type c.Int
	Here *DtlinkT
	Hh   struct {
		X_Htab **DtlinkT
	}
	Ntab c.Int
	Size c.Int
	Loop c.Int
}
type DtdataT X_DtdataS

type X_DtS struct {
	Searchf unsafe.Pointer
	Disc    *DtdiscT
	Data    *DtdataT
	Meth    *DtmethodT
	Nview   c.Int
	View    *DtT
	Walk    *DtT
	User    unsafe.Pointer
}
type DtT X_DtS
type DictT X_DtS

type X_DtstatS struct {
	DtMeth  c.Int
	DtSize  c.Int
	DtN     uintptr
	DtMax   uintptr
	DtCount *uintptr
}
type DtstatT X_DtstatS
// llgo:type C
type DtsearchF func(*DtT, unsafe.Pointer, c.Int) unsafe.Pointer
// llgo:type C
type DtmakeF func(unsafe.Pointer, *DtdiscT) unsafe.Pointer
// llgo:type C
type DtfreeF func(unsafe.Pointer)
// llgo:type C
type DtcomparF func(unsafe.Pointer, unsafe.Pointer) c.Int
// llgo:link (*DtdiscT).Dtopen C.dtopen
func (recv_ *DtdiscT) Dtopen(*DtmethodT) *DtT {
	return nil
}
// llgo:link (*DtT).Dtclose C.dtclose
func (recv_ *DtT) Dtclose() c.Int {
	return 0
}
// llgo:link (*DtT).Dtview C.dtview
func (recv_ *DtT) Dtview(*DtT) *DtT {
	return nil
}
// llgo:link (*DtT).Dtdisc C.dtdisc
func (recv_ *DtT) Dtdisc(*DtdiscT) *DtdiscT {
	return nil
}
// llgo:link (*DtT).Dtmethod C.dtmethod
func (recv_ *DtT) Dtmethod(*DtmethodT) *DtmethodT {
	return nil
}
// llgo:link (*DtT).Dtflatten C.dtflatten
func (recv_ *DtT) Dtflatten() *DtlinkT {
	return nil
}
// llgo:link (*DtT).Dtextract C.dtextract
func (recv_ *DtT) Dtextract() *DtlinkT {
	return nil
}
// llgo:link (*DtT).Dtrestore C.dtrestore
func (recv_ *DtT) Dtrestore(*DtlinkT) c.Int {
	return 0
}
// llgo:link (*DtT).Dtwalk C.dtwalk
func (recv_ *DtT) Dtwalk(func(unsafe.Pointer, unsafe.Pointer) c.Int, unsafe.Pointer) c.Int {
	return 0
}
// llgo:link (*DtT).Dtrenew C.dtrenew
func (recv_ *DtT) Dtrenew(unsafe.Pointer) unsafe.Pointer {
	return nil
}
// llgo:link (*DtT).Dtsize C.dtsize
func (recv_ *DtT) Dtsize() c.Int {
	return 0
}
// llgo:link (*DtT).Dtstat C.dtstat
func (recv_ *DtT) Dtstat(*DtstatT, c.Int) c.Int {
	return 0
}
//go:linkname Dtstrhash C.dtstrhash
func Dtstrhash(unsafe.Pointer, c.Int) c.Uint
