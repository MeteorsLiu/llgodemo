package graphviz

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)
//go:linkname Toggle C.gvToggle
func Toggle(c.Int)
// llgo:link (*LtSymlistT).NEWcontext C.gvNEWcontext
func (recv_ *LtSymlistT) NEWcontext(demand_loading c.Int) *GVCT {
	return nil
}
//go:linkname Context C.gvContext
func Context() *GVCT
// llgo:link (*LtSymlistT).ContextPlugins C.gvContextPlugins
func (recv_ *LtSymlistT) ContextPlugins(demand_loading c.Int) *GVCT {
	return nil
}
// llgo:link (*GVCT).CInfo C.gvcInfo
func (recv_ *GVCT) CInfo() **int8 {
	return nil
}
// llgo:link (*GVCT).CVersion C.gvcVersion
func (recv_ *GVCT) CVersion() *int8 {
	return nil
}
// llgo:link (*GVCT).CBuildDate C.gvcBuildDate
func (recv_ *GVCT) CBuildDate() *int8 {
	return nil
}
// llgo:link (*GVCT).ParseArgs C.gvParseArgs
func (recv_ *GVCT) ParseArgs(argc c.Int, argv **int8) c.Int {
	return 0
}
// llgo:link (*GVCT).NextInputGraph C.gvNextInputGraph
func (recv_ *GVCT) NextInputGraph() *GraphT {
	return nil
}
// llgo:link (*GVCT).PluginsGraph C.gvPluginsGraph
func (recv_ *GVCT) PluginsGraph() *GraphT {
	return nil
}
// llgo:link (*GVCT).Layout C.gvLayout
func (recv_ *GVCT) Layout(g *GraphT, engine *int8) c.Int {
	return 0
}
// llgo:link (*GVCT).LayoutJobs C.gvLayoutJobs
func (recv_ *GVCT) LayoutJobs(g *GraphT) c.Int {
	return 0
}
// llgo:link (*GraphT).LayoutDone C.gvLayoutDone
func (recv_ *GraphT) LayoutDone() c.Int {
	return 0
}
// llgo:link (*GraphT).AttachAttrs C.attach_attrs
func (recv_ *GraphT) AttachAttrs() {
}
// llgo:link (*GVCT).Render C.gvRender
func (recv_ *GVCT) Render(g *GraphT, format *int8, out *c.FILE) c.Int {
	return 0
}
// llgo:link (*GVCT).RenderFilename C.gvRenderFilename
func (recv_ *GVCT) RenderFilename(g *GraphT, format *int8, filename *int8) c.Int {
	return 0
}
// llgo:link (*GVCT).RenderContext C.gvRenderContext
func (recv_ *GVCT) RenderContext(g *GraphT, format *int8, context unsafe.Pointer) c.Int {
	return 0
}
// llgo:link (*GVCT).RenderData C.gvRenderData
func (recv_ *GVCT) RenderData(g *GraphT, format *int8, result **int8, length *c.Uint) c.Int {
	return 0
}
//go:linkname FreeRenderData C.gvFreeRenderData
func FreeRenderData(data *int8)
// llgo:link (*GVCT).RenderJobs C.gvRenderJobs
func (recv_ *GVCT) RenderJobs(g *GraphT) c.Int {
	return 0
}
// llgo:link (*GVCT).FreeLayout C.gvFreeLayout
func (recv_ *GVCT) FreeLayout(g *GraphT) c.Int {
	return 0
}
// llgo:link (*GVCT).Finalize C.gvFinalize
func (recv_ *GVCT) Finalize() {
}
// llgo:link (*GVCT).FreeContext C.gvFreeContext
func (recv_ *GVCT) FreeContext() c.Int {
	return 0
}
// llgo:link (*GVCT).PluginList C.gvPluginList
func (recv_ *GVCT) PluginList(kind *int8, sz *c.Int) **int8 {
	return nil
}
/** Add a library from your user application
 * @param gvc Graphviz context to add library to
 * @param lib library to add
 */
// llgo:link (*GVCT).AddLibrary C.gvAddLibrary
func (recv_ *GVCT) AddLibrary(lib *PluginLibraryT) {
}
/** Perform a Transitive Reduction on a graph
 * @param g  graph to be transformed.
 */
// llgo:link (*GraphT).ToolTred C.gvToolTred
func (recv_ *GraphT) ToolTred() c.Int {
	return 0
}
