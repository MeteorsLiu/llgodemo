package graphviz

import _ "unsafe"

type PxyT struct {
	X float64
	Y float64
}
type PpointT PxyT
type PvectorT PxyT

type PpolyT struct {
	Ps *PpointT
	Pn uintptr
}
type PpolylineT PpolyT

type PedgeT struct {
	A PpointT
	B PpointT
}

type VconfigS struct {
	Unused [8]uint8
}
type VconfigT VconfigS
// llgo:link (*PpolylineT).FreePath C.freePath
func (recv_ *PpolylineT) FreePath() {
}
