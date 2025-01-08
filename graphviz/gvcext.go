package graphviz

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

type ApiT c.Int

const (
	ApiTAPIRender     ApiT = 0
	ApiTAPILayout     ApiT = 1
	ApiTAPITextlayout ApiT = 2
	ApiTAPIDevice     ApiT = 3
	ApiTAPILoadimage  ApiT = 4
)

type GVJS struct {
	Gvc                 *c.Int
	Next                *c.Int
	NextActive          *c.Int
	Common              *GVCOMMONT
	Obj                 *ObjStateT
	InputFilename       *int8
	GraphIndex          c.Int
	LayoutType          *int8
	OutputFilename      *int8
	OutputFile          *c.Int
	OutputData          *int8
	OutputDataAllocated c.Uint
	OutputDataPosition  c.Uint
	OutputLangname      *int8
	OutputLang          c.Int
	Render              PluginActiveRenderT
	Device              PluginActiveDeviceT
	Loadimage           PluginActiveLoadimageT
	Callbacks           *DeviceCallbacksT
	DeviceDpi           c.Int
	DeviceSetsDpi       c.Int
	Display             unsafe.Pointer
	Screen              c.Int
	Context             unsafe.Pointer
	ExternalContext     c.Int
	Imagedata           *int8
	Flags               c.Int
	NumLayers           c.Int
	LayerNum            c.Int
	PagesArraySize      c.Int
	PagesArrayFirst     c.Int
	PagesArrayMajor     c.Int
	PagesArrayMinor     c.Int
	PagesArrayElem      c.Int
	NumPages            c.Int
	Bb                  c.Int
	Pad                 c.Int
	Clip                c.Int
	PageBox             c.Int
	PageSize            c.Int
	Focus               c.Int
	Zoom                float64
	Rotation            c.Int
	View                c.Int
	CanvasBox           c.Int
	Margin              c.Int
	Dpi                 c.Int
	Width               c.Uint
	Height              c.Uint
	PageBoundingBox     c.Int
	BoundingBox         c.Int
	Scale               c.Int
	Translation         c.Int
	Devscale            c.Int
	FitMode             c.Int
	NeedsRefresh        c.Int
	Click               c.Int
	HasGrown            c.Int
	HasBeenRendered     c.Int
	Button              int8
	Pointer             c.Int
	Oldpointer          c.Int
	CurrentObj          unsafe.Pointer
	SelectedObj         unsafe.Pointer
	ActiveTooltip       *int8
	SelectedHref        *int8
	Window              unsafe.Pointer
	Keybindings         *EventKeyBindingT
	Numkeys             uintptr
	Keycodes            unsafe.Pointer
}
type GVJT GVJS

type GVGS struct {
	Unused [8]uint8
}
type GVGT GVGS

type GVCS struct {
	Unused [8]uint8
}
type GVCT GVCS

type LtSymlistT struct {
	Name    *int8
	Address unsafe.Pointer
}

type PluginAvailableS struct {
	Unused [8]uint8
}
type PluginAvailableT PluginAvailableS
