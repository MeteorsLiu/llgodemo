package graphviz

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)
/// ingroup plugin_api
type PluginInstalledT struct {
	Id       c.Int
	Type     *int8
	Quality  c.Int
	Engine   unsafe.Pointer
	Features unsafe.Pointer
}

type PluginApiT struct {
	Api   ApiT
	Types *PluginInstalledT
}

type PluginLibraryT struct {
	Packagename *int8
	Apis        *PluginApiT
}
