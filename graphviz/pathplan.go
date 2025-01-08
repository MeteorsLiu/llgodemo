package graphviz

import (
	"github.com/goplus/llgo/c"
	_ "unsafe"
)
// llgo:link (*PpolyT).Pshortestpath C.Pshortestpath
func (recv_ *PpolyT) Pshortestpath(endpoints *PpointT, output_route *PpolylineT) c.Int {
	return 0
}
// llgo:link (*PedgeT).Proutespline C.Proutespline
func (recv_ *PedgeT) Proutespline(n_barriers uintptr, input_route PpolylineT, endpoint_slopes *PvectorT, output_route *PpolylineT) c.Int {
	return 0
}
// llgo:link PpolylineT.MakePolyline C.make_polyline
func (recv_ PpolylineT) MakePolyline(sline *PpolylineT) {
}
