package graphviz

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

type DeviceEngineS struct {
	Initialize unsafe.Pointer
	Format     unsafe.Pointer
	Finalize   unsafe.Pointer
}
type DeviceEngineT DeviceEngineS

type FormatterEngineS struct {
	Unused [8]uint8
}
type FormatterEngineT FormatterEngineS

type RenderEngineS struct {
	BeginJob     unsafe.Pointer
	EndJob       unsafe.Pointer
	BeginGraph   unsafe.Pointer
	EndGraph     unsafe.Pointer
	BeginLayer   unsafe.Pointer
	EndLayer     unsafe.Pointer
	BeginPage    unsafe.Pointer
	EndPage      unsafe.Pointer
	BeginCluster unsafe.Pointer
	EndCluster   unsafe.Pointer
	BeginNodes   unsafe.Pointer
	EndNodes     unsafe.Pointer
	BeginEdges   unsafe.Pointer
	EndEdges     unsafe.Pointer
	BeginNode    unsafe.Pointer
	EndNode      unsafe.Pointer
	BeginEdge    unsafe.Pointer
	EndEdge      unsafe.Pointer
	BeginAnchor  unsafe.Pointer
	EndAnchor    unsafe.Pointer
	BeginLabel   unsafe.Pointer
	EndLabel     unsafe.Pointer
	Textspan     unsafe.Pointer
	ResolveColor unsafe.Pointer
	Ellipse      unsafe.Pointer
	Polygon      unsafe.Pointer
	Beziercurve  unsafe.Pointer
	Polyline     unsafe.Pointer
	Comment      unsafe.Pointer
	LibraryShape unsafe.Pointer
}
type RenderEngineT RenderEngineS

type LayoutEngineS struct {
	Layout  unsafe.Pointer
	Cleanup unsafe.Pointer
}
type LayoutEngineT LayoutEngineS

type TextlayoutEngineS struct {
	Bool func(*c.Int) func(*TextspanT, **int8) c.Int
}
type TextlayoutEngineT TextlayoutEngineS

type LoadimageEngineS struct {
	Loadimage unsafe.Pointer
}
type LoadimageEngineT LoadimageEngineS
type PenType c.Int

const (
	PenTypePENNONE   PenType = 0
	PenTypePENDASHED PenType = 1
	PenTypePENDOTTED PenType = 2
	PenTypePENSOLID  PenType = 3
)

type FillType c.Int

const (
	FillTypeFILLNONE   FillType = 0
	FillTypeFILLSOLID  FillType = 1
	FillTypeFILLLINEAR FillType = 2
	FillTypeFILLRADIAL FillType = 3
)

type FontType c.Int

const (
	FontTypeFONTREGULAR FontType = 0
	FontTypeFONTBOLD    FontType = 1
	FontTypeFONTITALIC  FontType = 2
)

type LabelType c.Int

const (
	LabelTypeLABELPLAIN LabelType = 0
	LabelTypeLABELHTML  LabelType = 1
)

type RenderFeaturesT struct {
	Flags         c.Int
	DefaultPad    float64
	Knowncolors   **int8
	SzKnowncolors c.Int
	ColorType     ColorTypeT
}
type DeviceFeaturesT struct {
	Flags           c.Int
	DefaultMargin   c.Int
	DefaultPagesize c.Int
	DefaultDpi      c.Int
}

type PluginActiveDeviceS struct {
	Engine   *DeviceEngineT
	Id       c.Int
	Features *DeviceFeaturesT
	Type     *int8
}
type PluginActiveDeviceT PluginActiveDeviceS

type PluginActiveRenderS struct {
	Engine   *RenderEngineT
	Id       c.Int
	Features *RenderFeaturesT
	Type     *int8
}
type PluginActiveRenderT PluginActiveRenderS

type PluginActiveLoadimageT struct {
	Engine *LoadimageEngineT
	Id     c.Int
	Type   *int8
}

type DeviceCallbacksS struct {
	Refresh       unsafe.Pointer
	ButtonPress   unsafe.Pointer
	ButtonRelease unsafe.Pointer
	Motion        unsafe.Pointer
	Modify        unsafe.Pointer
	Del           unsafe.Pointer
	Read          unsafe.Pointer
	Layout        unsafe.Pointer
	Render        unsafe.Pointer
}
type DeviceCallbacksT DeviceCallbacksS
// llgo:type C
type EventKeyCallbackT func(*c.Int) c.Int

type EventKeyBindingS struct {
	Keystring *int8
	Callback  unsafe.Pointer
}
type EventKeyBindingT EventKeyBindingS
type MapShapeT c.Int

const (
	MapShapeTMAPRECTANGLE MapShapeT = 0
	MapShapeTMAPCIRCLE    MapShapeT = 1
	MapShapeTMAPPOLYGON   MapShapeT = 2
)

type ObjType c.Int

const (
	ObjTypeROOTGRAPHOBJTYPE ObjType = 0
	ObjTypeCLUSTEROBJTYPE   ObjType = 1
	ObjTypeNODEOBJTYPE      ObjType = 2
	ObjTypeEDGEOBJTYPE      ObjType = 3
)

type EmitStateT c.Int

const (
	EmitStateTEMITGDRAW  EmitStateT = 0
	EmitStateTEMITCDRAW  EmitStateT = 1
	EmitStateTEMITTDRAW  EmitStateT = 2
	EmitStateTEMITHDRAW  EmitStateT = 3
	EmitStateTEMITGLABEL EmitStateT = 4
	EmitStateTEMITCLABEL EmitStateT = 5
	EmitStateTEMITTLABEL EmitStateT = 6
	EmitStateTEMITHLABEL EmitStateT = 7
	EmitStateTEMITNDRAW  EmitStateT = 8
	EmitStateTEMITEDRAW  EmitStateT = 9
	EmitStateTEMITNLABEL EmitStateT = 10
	EmitStateTEMITELABEL EmitStateT = 11
)

type ObjStateS struct {
	Parent *ObjStateT
	Type   ObjType
	U      struct {
		G *c.Int
	}
	EmitState            EmitStateT
	Pencolor             ColorT
	Fillcolor            ColorT
	Stopcolor            ColorT
	GradientAngle        c.Int
	GradientFrac         float64
	Pen                  PenType
	Fill                 FillType
	Penwidth             float64
	Rawstyle             **int8
	Z                    float64
	TailZ                float64
	HeadZ                float64
	Label                *int8
	Xlabel               *int8
	Taillabel            *int8
	Headlabel            *int8
	Url                  *int8
	Id                   *int8
	Labelurl             *int8
	Tailurl              *int8
	Headurl              *int8
	Tooltip              *int8
	Labeltooltip         *int8
	Tailtooltip          *int8
	Headtooltip          *int8
	Target               *int8
	Labeltarget          *int8
	Tailtarget           *int8
	Headtarget           *int8
	ExplicitTooltip      c.Uint
	ExplicitTailtooltip  c.Uint
	ExplicitHeadtooltip  c.Uint
	ExplicitLabeltooltip c.Uint
	ExplicitTailtarget   c.Uint
	ExplicitHeadtarget   c.Uint
	ExplicitEdgetarget   c.Uint
	ExplicitTailurl      c.Uint
	ExplicitHeadurl      c.Uint
	Labeledgealigned     c.Uint
	UrlMapShape          MapShapeT
	UrlMapN              uintptr
	UrlMapP              *c.Int
	UrlBsplinemapPolyN   uintptr
	UrlBsplinemapN       *uintptr
	UrlBsplinemapP       *c.Int
	TailendurlMapN       c.Int
	TailendurlMapP       *c.Int
	HeadendurlMapN       c.Int
	HeadendurlMapP       *c.Int
}
type ObjStateT ObjStateS
