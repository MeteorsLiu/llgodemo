package graphviz

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

type IDTYPE uint64

type AgtagS struct {
	Objtype c.Uint
	Mtflock c.Uint
	Attrwf  c.Uint
	Seq     c.Uint
	Id      IDTYPE
}
type AgtagT AgtagS

type AgobjS struct {
	Tag  AgtagT
	Data *AgrecT
}
type AgobjT AgobjS

type AgraphS struct {
	Base    AgobjT
	Desc    AgdescT
	SeqLink DtlinkT
	IdLink  DtlinkT
	NSeq    *DictT
	NId     *GraphvizNodeSet
	ESeq    *DictT
	EId     *DictT
	GSeq    *DictT
	GId     *DictT
	Parent  *AgraphT
	Root    *AgraphT
	Clos    *AgclosT
}
type AgraphT AgraphS

type AgdescS struct {
	Directed  c.Uint
	Strict    c.Uint
	NoLoop    c.Uint
	Maingraph c.Uint
	NoWrite   c.Uint
	HasAttrs  c.Uint
	HasCmpnd  c.Uint
}
type AgdescT AgdescS

type AgdstateS struct {
	Id unsafe.Pointer
}
type AgdstateT AgdstateS

type AgclosS struct {
	Disc         AgdiscT
	State        AgdstateT
	Strdict      *DictT
	Seq          [3]uint64
	Cb           *AgcbstackT
	LookupByName [3]*DictT
	LookupById   [3]*DictT
}
type AgclosT AgclosS

type AgnodeS struct {
	Base    AgobjT
	Root    *AgraphT
	Mainsub AgsubnodeT
}
type AgnodeT AgnodeS

type AgsubnodeS struct {
	SeqLink DtlinkT
	IdLink  DtlinkT
	Node    *AgnodeT
	InId    *DtlinkT
	OutId   *DtlinkT
	InSeq   *DtlinkT
	OutSeq  *DtlinkT
}
type AgsubnodeT AgsubnodeS

type AgedgeS struct {
	Base    AgobjT
	IdLink  DtlinkT
	SeqLink DtlinkT
	Node    *AgnodeT
}
type AgedgeT AgedgeS

type AgedgepairS struct {
	Out AgedgeT
	In  AgedgeT
}
type AgedgepairT AgedgepairS

type AgiddiscS struct {
	Open       unsafe.Pointer
	Map        unsafe.Pointer
	Alloc      unsafe.Pointer
	Free       unsafe.Pointer
	Print      unsafe.Pointer
	Close      unsafe.Pointer
	Idregister unsafe.Pointer
}
type AgiddiscT AgiddiscS

type AgiodiscS struct {
	Afread unsafe.Pointer
	Putstr unsafe.Pointer
	Flush  unsafe.Pointer
}
type AgiodiscT AgiodiscS

type AgdiscS struct {
	Id *AgiddiscT
	Io *AgiodiscT
}
type AgdiscT AgdiscS

type AgcbdiscS struct {
	Graph struct {
		Ins unsafe.Pointer
		Mod unsafe.Pointer
		Del unsafe.Pointer
	}
	Node struct {
		Ins unsafe.Pointer
		Mod unsafe.Pointer
		Del unsafe.Pointer
	}
	Edge struct {
		Ins unsafe.Pointer
		Mod unsafe.Pointer
		Del unsafe.Pointer
	}
}
type AgcbdiscT AgcbdiscS

type AgcbstackS struct {
	F     *AgcbdiscT
	State unsafe.Pointer
	Prev  *AgcbstackT
}
type AgcbstackT AgcbstackS

type AgsymS struct {
	Link   DtlinkT
	Name   *int8
	Defval *int8
	Id     c.Int
	Kind   int8
	Fixed  int8
	Print  int8
}
type AgsymT AgsymS

type AgattrS struct {
	H    AgrecT
	Dict *DictT
	Str  **int8
}
type AgattrT AgattrS

type AgdatadictS struct {
	H    AgrecT
	Dict struct {
		N *DictT
		E *DictT
		G *DictT
	}
}
type AgdatadictT AgdatadictS

type AgrecS struct {
	Name *int8
	Next *AgrecT
}
type AgrecT AgrecS

const (
	AGRAPH    c.Int = 0
	AGNODE    c.Int = 1
	AGEDGE    c.Int = 2
	AGOUTEDGE c.Int = 2
	AGINEDGE  c.Int = 3
)
// llgo:type C
type AgobjfnT func(*AgraphT, *AgobjT, unsafe.Pointer)
// llgo:type C
type AgobjupdfnT func(*AgraphT, *AgobjT, unsafe.Pointer, *AgsymT)
// llgo:link (*AgraphT).Agpushdisc C.agpushdisc
func (recv_ *AgraphT) Agpushdisc(disc *AgcbdiscT, state unsafe.Pointer) {
}
// llgo:link (*AgraphT).Agpopdisc C.agpopdisc
func (recv_ *AgraphT) Agpopdisc(disc *AgcbdiscT) c.Int {
	return 0
}
/// opaque type; the definition of this is internal to Graphviz
type GraphvizNodeSet struct {
	Unused [8]uint8
}
//go:linkname Agopen C.agopen
func Agopen(name *int8, desc AgdescT, disc *AgdiscT) *AgraphT
// llgo:link (*AgraphT).Agclose C.agclose
func (recv_ *AgraphT) Agclose() c.Int {
	return 0
}
//go:linkname Agread C.agread
func Agread(chan_ unsafe.Pointer, disc *AgdiscT) *AgraphT
//go:linkname Agmemread C.agmemread
func Agmemread(cp *int8) *AgraphT
// llgo:link (*AgraphT).Agmemconcat C.agmemconcat
func (recv_ *AgraphT) Agmemconcat(cp *int8) *AgraphT {
	return nil
}
//go:linkname Agreadline C.agreadline
func Agreadline(c.Int)
//go:linkname Agsetfile C.agsetfile
func Agsetfile(*int8)
// llgo:link (*AgraphT).Agconcat C.agconcat
func (recv_ *AgraphT) Agconcat(chan_ unsafe.Pointer, disc *AgdiscT) *AgraphT {
	return nil
}
// llgo:link (*AgraphT).Agwrite C.agwrite
func (recv_ *AgraphT) Agwrite(chan_ unsafe.Pointer) c.Int {
	return 0
}
// llgo:link (*AgraphT).Agisdirected C.agisdirected
func (recv_ *AgraphT) Agisdirected() c.Int {
	return 0
}
// llgo:link (*AgraphT).Agisundirected C.agisundirected
func (recv_ *AgraphT) Agisundirected() c.Int {
	return 0
}
// llgo:link (*AgraphT).Agisstrict C.agisstrict
func (recv_ *AgraphT) Agisstrict() c.Int {
	return 0
}
// llgo:link (*AgraphT).Agissimple C.agissimple
func (recv_ *AgraphT) Agissimple() c.Int {
	return 0
}
/// @addtogroup cgraph_node
/// @{
// llgo:link (*AgraphT).Agnode C.agnode
func (recv_ *AgraphT) Agnode(name *int8, createflag c.Int) *AgnodeT {
	return nil
}
// llgo:link (*AgraphT).Agidnode C.agidnode
func (recv_ *AgraphT) Agidnode(id IDTYPE, createflag c.Int) *AgnodeT {
	return nil
}
// llgo:link (*AgraphT).Agsubnode C.agsubnode
func (recv_ *AgraphT) Agsubnode(n *AgnodeT, createflag c.Int) *AgnodeT {
	return nil
}
// llgo:link (*AgraphT).Agfstnode C.agfstnode
func (recv_ *AgraphT) Agfstnode() *AgnodeT {
	return nil
}
// llgo:link (*AgraphT).Agnxtnode C.agnxtnode
func (recv_ *AgraphT) Agnxtnode(n *AgnodeT) *AgnodeT {
	return nil
}
// llgo:link (*AgraphT).Aglstnode C.aglstnode
func (recv_ *AgraphT) Aglstnode() *AgnodeT {
	return nil
}
// llgo:link (*AgraphT).Agprvnode C.agprvnode
func (recv_ *AgraphT) Agprvnode(n *AgnodeT) *AgnodeT {
	return nil
}
// llgo:link (*AgraphT).Agsubrep C.agsubrep
func (recv_ *AgraphT) Agsubrep(n *AgnodeT) *AgsubnodeT {
	return nil
}
// llgo:link (*AgnodeT).Agnodebefore C.agnodebefore
func (recv_ *AgnodeT) Agnodebefore(v *AgnodeT) c.Int {
	return 0
}
// llgo:link (*AgraphT).Agdelnode C.agdelnode
func (recv_ *AgraphT) Agdelnode(arg_n *AgnodeT) c.Int {
	return 0
}
// llgo:link (*AgnodeT).AgrelabelNode C.agrelabel_node
func (recv_ *AgnodeT) AgrelabelNode(newname *int8) c.Int {
	return 0
}
/** @defgroup cgraph_edge edges
 *  @ingroup cgraph_object
 *
 * An abstract edge has two endpoint nodes called tail and head
 * where all outedges of the same node have it as the tail
 * value and similarly all inedges have it as the head.
 * In an undirected graph, head and tail are interchangeable.
 * If a graph has multi-edges between the same pair of nodes,
 * the edge's string name behaves as a secondary key.
 *
 * Note that an abstract edge has two distinct concrete
 * representations: as an in-edge and as an out-edge.
 * In particular, the pointer as an out-edge is different
 * from the pointer as an in-edge.
 * The function @ref ageqedge canonicalizes the pointers before
 * doing a comparison and so can be used to test edge equality.
 * The sense of an edge can be flipped using @ref agopp.
 *
 * @{
 */
// llgo:link (*AgraphT).Agedge C.agedge
func (recv_ *AgraphT) Agedge(t *AgnodeT, h *AgnodeT, name *int8, createflag c.Int) *AgedgeT {
	return nil
}
// llgo:link (*AgraphT).Agidedge C.agidedge
func (recv_ *AgraphT) Agidedge(t *AgnodeT, h *AgnodeT, id IDTYPE, createflag c.Int) *AgedgeT {
	return nil
}
// llgo:link (*AgraphT).Agsubedge C.agsubedge
func (recv_ *AgraphT) Agsubedge(e *AgedgeT, createflag c.Int) *AgedgeT {
	return nil
}
// llgo:link (*AgraphT).Agfstin C.agfstin
func (recv_ *AgraphT) Agfstin(n *AgnodeT) *AgedgeT {
	return nil
}
// llgo:link (*AgraphT).Agnxtin C.agnxtin
func (recv_ *AgraphT) Agnxtin(e *AgedgeT) *AgedgeT {
	return nil
}
// llgo:link (*AgraphT).Agfstout C.agfstout
func (recv_ *AgraphT) Agfstout(n *AgnodeT) *AgedgeT {
	return nil
}
// llgo:link (*AgraphT).Agnxtout C.agnxtout
func (recv_ *AgraphT) Agnxtout(e *AgedgeT) *AgedgeT {
	return nil
}
// llgo:link (*AgraphT).Agfstedge C.agfstedge
func (recv_ *AgraphT) Agfstedge(n *AgnodeT) *AgedgeT {
	return nil
}
// llgo:link (*AgraphT).Agnxtedge C.agnxtedge
func (recv_ *AgraphT) Agnxtedge(e *AgedgeT, n *AgnodeT) *AgedgeT {
	return nil
}
// llgo:link (*AgraphT).Agdeledge C.agdeledge
func (recv_ *AgraphT) Agdeledge(arg_e *AgedgeT) c.Int {
	return 0
}
/// @addtogroup cgraph_object
/// @{
//go:linkname Agraphof C.agraphof
func Agraphof(obj unsafe.Pointer) *AgraphT
//go:linkname Agroot C.agroot
func Agroot(obj unsafe.Pointer) *AgraphT
// llgo:link (*AgraphT).Agcontains C.agcontains
func (recv_ *AgraphT) Agcontains(obj unsafe.Pointer) c.Int {
	return 0
}
//go:linkname Agnameof C.agnameof
func Agnameof(unsafe.Pointer) *int8
// llgo:link (*AgraphT).Agdelete C.agdelete
func (recv_ *AgraphT) Agdelete(obj unsafe.Pointer) c.Int {
	return 0
}
//go:linkname Agobjkind C.agobjkind
func Agobjkind(obj unsafe.Pointer) c.Int
/** @defgroup cgraph_string string utilities
 *  @brief reference-counted strings
 *  @ingroup cgraph_misc
 *
 *  Storage management of strings as reference-counted strings.
 *  The caller does not need to dynamically allocate storage.
 *
 * All uses of cgraph strings need to be freed using @ref agstrfree
 * in order to correctly maintain the reference count.
 *
 * @ref agcanonStr returns a pointer to a version of the input string
 * canonicalized for output for later re-parsing.
 * This includes quoting special characters and keywords.
 * It uses its own internal buffer, so the value will be lost on
 * the next call to @ref agcanonStr.
 * @ref agcanon is identical with @ref agcanonStr
 * except it can be used with any character string.
 * The second argument indicates whether or not the string
 * should be canonicalized as an HTML-like string.
 *
 * @{
 */
// llgo:link (*AgraphT).Agstrdup C.agstrdup
func (recv_ *AgraphT) Agstrdup(*int8) *int8 {
	return nil
}
/// creating one if necessary
// llgo:link (*AgraphT).AgstrdupHtml C.agstrdup_html
func (recv_ *AgraphT) AgstrdupHtml(*int8) *int8 {
	return nil
}
///
//go:linkname Aghtmlstr C.aghtmlstr
func Aghtmlstr(*int8) c.Int
///
// llgo:link (*AgraphT).Agstrbind C.agstrbind
func (recv_ *AgraphT) Agstrbind(*int8) *int8 {
	return nil
}
// llgo:link (*AgraphT).Agstrfree C.agstrfree
func (recv_ *AgraphT) Agstrfree(*int8) c.Int {
	return 0
}
//go:linkname Agcanon C.agcanon
func Agcanon(str *int8, html c.Int) *int8
//go:linkname Agstrcanon C.agstrcanon
func Agstrcanon(*int8, *int8) *int8
//go:linkname AgcanonStr C.agcanonStr
func AgcanonStr(str *int8) *int8
// llgo:link (*AgraphT).Agattr C.agattr
func (recv_ *AgraphT) Agattr(kind c.Int, name *int8, value *int8) *AgsymT {
	return nil
}
//go:linkname Agattrsym C.agattrsym
func Agattrsym(obj unsafe.Pointer, name *int8) *AgsymT
// llgo:link (*AgraphT).Agnxtattr C.agnxtattr
func (recv_ *AgraphT) Agnxtattr(kind c.Int, attr *AgsymT) *AgsymT {
	return nil
}
/// @param attr	if `NULL` the function returns the first attribute
/// @returns the next one in succession or `NULL` at the end of the list.
//go:linkname Agcopyattr C.agcopyattr
func Agcopyattr(oldobj unsafe.Pointer, newobj unsafe.Pointer) c.Int
/// @addtogroup cgraph_rec
/// @{
//go:linkname Agbindrec C.agbindrec
func Agbindrec(obj unsafe.Pointer, name *int8, recsize c.Uint, move_to_front c.Int) unsafe.Pointer
/// @param recsize if 0, the call to @ref agbindrec is simply a lookup
/// @returns pointer to `Agrec_t` and user data
//go:linkname Aggetrec C.aggetrec
func Aggetrec(obj unsafe.Pointer, name *int8, move_to_front c.Int) *AgrecT
//go:linkname Agdelrec C.agdelrec
func Agdelrec(obj unsafe.Pointer, name *int8) c.Int
// llgo:link (*AgraphT).Aginit C.aginit
func (recv_ *AgraphT) Aginit(kind c.Int, rec_name *int8, rec_size c.Int, move_to_front c.Int) {
}
// llgo:link (*AgraphT).Agclean C.agclean
func (recv_ *AgraphT) Agclean(kind c.Int, rec_name *int8) {
}
/// @}
//go:linkname Agget C.agget
func Agget(obj unsafe.Pointer, name *int8) *int8
//go:linkname Agxget C.agxget
func Agxget(obj unsafe.Pointer, sym *AgsymT) *int8
//go:linkname Agset C.agset
func Agset(obj unsafe.Pointer, name *int8, value *int8) c.Int
//go:linkname Agxset C.agxset
func Agxset(obj unsafe.Pointer, sym *AgsymT, value *int8) c.Int
//go:linkname Agsafeset C.agsafeset
func Agsafeset(obj unsafe.Pointer, name *int8, value *int8, def *int8) c.Int
/** @defgroup cgraph_subgraph subgraphs
 *  @ingroup cgraph_graph
 *
 * A "main" or "root" graph defines a namespace for a collection of
 * graph objects (subgraphs, nodes, edges) and their attributes.
 * Objects may be named by unique strings or by integer IDs.
 *
 * @ref agsubg finds or creates a subgraph by name.
 *
 * @ref agidsubg allows a programmer to specify the subgraph by a unique integer
 * ID.
 *
 * A new subgraph is initially empty and is of the same kind as its parent.
 * Nested subgraph trees may be created.
 * A subgraph's name is only interpreted relative to its parent.
 *
 * A program can scan subgraphs under a given graph
 * using @ref agfstsubg and @ref agnxtsubg.
 *
 * A subgraph is deleted with @ref agdelsubg (or @ref agclose).
 *
 * The @ref agparent function returns the immediate parent graph of a subgraph,
 * or itself if the graph is already a root graph.
 *
 * @{
 */
// llgo:link (*AgraphT).Agsubg C.agsubg
func (recv_ *AgraphT) Agsubg(name *int8, cflag c.Int) *AgraphT {
	return nil
}
// llgo:link (*AgraphT).Agidsubg C.agidsubg
func (recv_ *AgraphT) Agidsubg(id IDTYPE, cflag c.Int) *AgraphT {
	return nil
}
// llgo:link (*AgraphT).Agfstsubg C.agfstsubg
func (recv_ *AgraphT) Agfstsubg() *AgraphT {
	return nil
}
// llgo:link (*AgraphT).Agnxtsubg C.agnxtsubg
func (recv_ *AgraphT) Agnxtsubg() *AgraphT {
	return nil
}
// llgo:link (*AgraphT).Agparent C.agparent
func (recv_ *AgraphT) Agparent() *AgraphT {
	return nil
}
// llgo:link (*AgraphT).Agdelsubg C.agdelsubg
func (recv_ *AgraphT) Agdelsubg(sub *AgraphT) c.Int {
	return 0
}
/** @defgroup card set cardinality
 *
 * By default, nodes are stored in ordered sets for
 * efficient random access to insert, find, and delete nodes.
 *
 * @ref agnnodes, @ref agnedges, and @ref agnsubg return the
 * sizes of node, edge and subgraph sets of a graph.
 *
 * The function @ref agdegree returns the size of a node’s edge set,
 * and takes flags to select in-edges, out-edges, or both.
 *
 * The function @ref agcountuniqedges returns
 * the size of a node’s edge set, and takes flags
 * to select in-edges, out-edges, or both.
 * Unlike @ref agdegree, each loop is only counted once.
 *
 * @{
 */
// llgo:link (*AgraphT).Agnnodes C.agnnodes
func (recv_ *AgraphT) Agnnodes() c.Int {
	return 0
}
// llgo:link (*AgraphT).Agnedges C.agnedges
func (recv_ *AgraphT) Agnedges() c.Int {
	return 0
}
// llgo:link (*AgraphT).Agnsubg C.agnsubg
func (recv_ *AgraphT) Agnsubg() c.Int {
	return 0
}
// llgo:link (*AgraphT).Agdegree C.agdegree
func (recv_ *AgraphT) Agdegree(n *AgnodeT, in c.Int, out c.Int) c.Int {
	return 0
}
// llgo:link (*AgraphT).Agcountuniqedges C.agcountuniqedges
func (recv_ *AgraphT) Agcountuniqedges(n *AgnodeT, in c.Int, out c.Int) c.Int {
	return 0
}
/// @defgroup cgmem memory
/// @{
// llgo:link (*AgraphT).Agalloc C.agalloc
func (recv_ *AgraphT) Agalloc(size uintptr) unsafe.Pointer {
	return nil
}
// llgo:link (*AgraphT).Agrealloc C.agrealloc
func (recv_ *AgraphT) Agrealloc(ptr unsafe.Pointer, oldsize uintptr, size uintptr) unsafe.Pointer {
	return nil
}
// llgo:link (*AgraphT).Agfree C.agfree
func (recv_ *AgraphT) Agfree(ptr unsafe.Pointer) {
}
// llgo:link (*AgraphT).Aginternalmapclearlocalnames C.aginternalmapclearlocalnames
func (recv_ *AgraphT) Aginternalmapclearlocalnames() {
}

type AgerrlevelT c.Int

const (
	AgerrlevelTAGWARN AgerrlevelT = 0
	AgerrlevelTAGERR  AgerrlevelT = 1
	AgerrlevelTAGMAX  AgerrlevelT = 2
	AgerrlevelTAGPREV AgerrlevelT = 3
)
// llgo:type C
type Agusererrf func(*int8) c.Int
// llgo:link AgerrlevelT.Agseterr C.agseterr
func (recv_ AgerrlevelT) Agseterr() AgerrlevelT {
	return 0
}
//go:linkname Aglasterr C.aglasterr
func Aglasterr() *int8
// llgo:link AgerrlevelT.Agerr C.agerr
func (recv_ AgerrlevelT) Agerr(fmt *int8, __llgo_va_list ...interface{}) c.Int {
	return 0
}
//go:linkname Agerrorf C.agerrorf
func Agerrorf(fmt *int8, __llgo_va_list ...interface{})
//go:linkname Agwarningf C.agwarningf
func Agwarningf(fmt *int8, __llgo_va_list ...interface{})
//go:linkname Agerrors C.agerrors
func Agerrors() c.Int
//go:linkname Agreseterrors C.agreseterrors
func Agreseterrors() c.Int
//go:linkname Agseterrf C.agseterrf
func Agseterrf(Agusererrf) Agusererrf

type GraphvizAcyclicOptionsT struct {
	OutFile *c.FILE
	DoWrite c.Int
	Verbose c.Int
}
/// programmatic access to `acyclic`
///
/// See `man acyclic` for an explanation of the `acyclic` tool.
///
/// \param g Graph to operate on
/// \param opts Options to control acyclic algorithm
/// \param num_rev [inout] Running total of reversed edges
/// \return True if a cycle was found, indicating failure
// llgo:link (*AgraphT).GraphvizAcyclic C.graphviz_acyclic
func (recv_ *AgraphT) GraphvizAcyclic(opts *GraphvizAcyclicOptionsT, num_rev *uintptr) c.Int {
	return 0
}

type GraphvizTredOptionsT struct {
	Verbose           c.Int
	PrintRemovedEdges c.Int
	Out               *c.FILE
	Err               *c.FILE
}
/// @brief programmatic access to `tred` -
/// [transitive reduction](https://en.wikipedia.org/wiki/Transitive_reduction)
///
/// See `man tred` for an explanation of the `tred` tool.
///
/// \param g Graph to operate on
/// \param opts Options to control tred algorithm
// llgo:link (*AgraphT).GraphvizTred C.graphviz_tred
func (recv_ *AgraphT) GraphvizTred(opts *GraphvizTredOptionsT) {
}

type GraphvizUnflattenOptionsT struct {
	DoFans     c.Int
	MaxMinlen  c.Int
	ChainLimit c.Int
}
/// programmatic access to `unflatten`
///
/// See `man unflatten` for an explanation of the `unflatten` tool.
///
/// \param g Graph to operate on
/// \param opts Options to control unflattening
// llgo:link (*AgraphT).GraphvizUnflatten C.graphviz_unflatten
func (recv_ *AgraphT) GraphvizUnflatten(opts *GraphvizUnflattenOptionsT) {
}
/** add to a graph any edges with both endpoints within that graph
 *
 * If `edgeset` is given as `NULL`, edges from the root graph of `g` will be
 * considered. In this case if `g` itself is the root graph, this call is a
 * no-op.
 *
 * If `g` is a connected component, the edges added will be all edges attached
 * to any node in `g`.
 *
 * \param g Graph to add edges to
 * \param edgeset Graph whose edges to consider
 * \return Number of edges added
 */
// llgo:link (*AgraphT).GraphvizNodeInduce C.graphviz_node_induce
func (recv_ *AgraphT) GraphvizNodeInduce(edgeset *AgraphT) uintptr {
	return 0
}
