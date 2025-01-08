package graphviz

import (
	"github.com/goplus/llgo/c"
	_ "unsafe"
)

type PackMode c.Int

const (
	PackModeLUndef  PackMode = 0
	PackModeLClust  PackMode = 1
	PackModeLNode   PackMode = 2
	PackModeLGraph  PackMode = 3
	PackModeLArray  PackMode = 4
	PackModeLAspect PackMode = 5
)

type PackvalT c.Uint
type PackInfo struct {
	Aspect    float32
	Sz        c.Int
	Margin    c.Uint
	DoSplines c.Int
	Mode      PackMode
	Fixed     *c.Int
	Vals      *PackvalT
	Flags     c.Int
}
//go:linkname PutRects C.putRects
func PutRects(ng uintptr, bbs *Boxf, pinfo *PackInfo) *Pointf
//go:linkname PackRects C.packRects
func PackRects(ng uintptr, bbs *Boxf, pinfo *PackInfo) c.Int
//go:linkname PutGraphs C.putGraphs
func PutGraphs(uintptr, **AgraphT, *AgraphT, *PackInfo) *Pointf
//go:linkname PackGraphs C.packGraphs
func PackGraphs(uintptr, **AgraphT, *AgraphT, *PackInfo) c.Int
//go:linkname PackSubgraphs C.packSubgraphs
func PackSubgraphs(uintptr, **AgraphT, *AgraphT, *PackInfo) c.Int
//go:linkname PackGraph C.pack_graph
func PackGraph(ng uintptr, gs **AgraphT, root *AgraphT, fixed *c.Int) c.Int
//go:linkname ShiftGraphs C.shiftGraphs
func ShiftGraphs(__llgo_arg_0 uintptr, __llgo_arg_1 **AgraphT, __llgo_arg_2 *Pointf, __llgo_arg_3 *AgraphT, bool c.Int) c.Int
// llgo:link (*AgraphT).GetPackMode C.getPackMode
func (recv_ *AgraphT) GetPackMode(dflt PackMode) PackMode {
	return 0
}
// llgo:link (*AgraphT).GetPack C.getPack
func (recv_ *AgraphT) GetPack(not_def c.Int, dflt c.Int) c.Int {
	return 0
}
// llgo:link (*AgraphT).GetPackInfo C.getPackInfo
func (recv_ *AgraphT) GetPackInfo(dflt PackMode, dfltMargin c.Int, __llgo_arg_2 *PackInfo) PackMode {
	return 0
}
// llgo:link (*AgraphT).GetPackModeInfo C.getPackModeInfo
func (recv_ *AgraphT) GetPackModeInfo(dflt PackMode, __llgo_arg_1 *PackInfo) PackMode {
	return 0
}
//go:linkname ParsePackModeInfo C.parsePackModeInfo
func ParsePackModeInfo(p *int8, dflt PackMode, pinfo *PackInfo) PackMode
// llgo:link (*AgraphT).IsConnected C.isConnected
func (recv_ *AgraphT) IsConnected() c.Int {
	return 0
}
// llgo:link (*AgraphT).Ccomps C.ccomps
func (recv_ *AgraphT) Ccomps(*uintptr, *int8) **AgraphT {
	return nil
}
// llgo:link (*AgraphT).Cccomps C.cccomps
func (recv_ *AgraphT) Cccomps(*uintptr, *int8) **AgraphT {
	return nil
}
// llgo:link (*AgraphT).Pccomps C.pccomps
func (recv_ *AgraphT) Pccomps(*uintptr, *int8, *c.Int) **AgraphT {
	return nil
}
// llgo:link (*AgraphT).MapClust C.mapClust
func (recv_ *AgraphT) MapClust() *AgraphT {
	return nil
}
