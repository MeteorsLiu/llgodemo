package graphviz

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

type GraphT AgraphS
type NodeT AgnodeS
type EdgeT AgedgeS
type AttrsymT AgsymS

type HtmllabelT struct {
	Unused [8]uint8
}

type Port struct {
	P           Pointf
	Theta       float64
	Bp          *Boxf
	Defined     c.Int
	Constrained c.Int
	Clip        c.Int
	Dyna        c.Int
	Order       int8
	Side        int8
	Name        *int8
}
type SplineInfo struct {
	Bool       func(*c.Int) func(*EdgeT) c.Int
	IgnoreSwap c.Int
	IsOrtho    c.Int
}

type PathendT struct {
	Nb       Boxf
	Np       Pointf
	Sidemask c.Int
	Boxn     c.Int
	Boxes    [20]Boxf
}

type Path struct {
	Start Port
	End   Port
	Nbox  uintptr
	Boxes *Boxf
	Data  unsafe.Pointer
}

type Bezier struct {
	List  *Pointf
	Size  uintptr
	Sflag uint32
	Eflag uint32
	Sp    Pointf
	Ep    Pointf
}

type Splines struct {
	List *Bezier
	Size uintptr
	Bb   Boxf
}

type TextlabelT struct {
	Text      *int8
	Fontname  *int8
	Fontcolor *int8
	Charset   c.Int
	Fontsize  float64
	Dimen     Pointf
	Space     Pointf
	Pos       Pointf
	U         struct {
		Txt struct {
			Span   *TextspanT
			Nspans uintptr
		}
	}
	Valign int8
	Set    c.Int
	Html   c.Int
}
type GraphvizPolygonStyleT struct {
	Filled     c.Int
	Radial     c.Int
	Rounded    c.Int
	Diagonals  c.Int
	Auxlabels  c.Int
	Invisible  c.Int
	Striped    c.Int
	Dotted     c.Int
	Dashed     c.Int
	Wedged     c.Int
	Underline  c.Int
	Fixedshape c.Int
	Shape      c.Uint
}

type PolygonT struct {
	Regular     c.Int
	Peripheries uintptr
	Sides       uintptr
	Orientation float64
	Distortion  float64
	Skew        float64
	Option      GraphvizPolygonStyleT
	Vertices    *Pointf
}

type InsideT struct {
	S struct {
		N        *NodeT
		Bp       *Boxf
		Lastn    *NodeT
		Radius   float64
		LastPoly *PolygonT
		Last     uintptr
		Outp     uintptr
		Scalex   float64
		Scaley   float64
		BoxURx   float64
		BoxURy   float64
	}
}

type StrokeT struct {
	Nvertices uintptr
	Vertices  *Pointf
}

type ShapeFunctions struct {
	Initfn unsafe.Pointer
	Freefn unsafe.Pointer
	Portfn unsafe.Pointer
	Bool   func(*c.Int) func(*InsideT, Pointf) c.Int
	Pboxfn unsafe.Pointer
	Codefn unsafe.Pointer
}
type ShapeKind c.Int

const (
	ShapeKindSHUNSET  ShapeKind = 0
	ShapeKindSHPOLY   ShapeKind = 1
	ShapeKindSHRECORD ShapeKind = 2
	ShapeKindSHPOINT  ShapeKind = 3
	ShapeKindSHEPSF   ShapeKind = 4
)

type ShapeDesc struct {
	Name      *int8
	Fns       *ShapeFunctions
	Polygon   *PolygonT
	Usershape c.Int
}

type AdjmatrixT struct {
	Unused [8]uint8
}

type RankT struct {
	N         c.Int
	V         **NodeT
	An        c.Int
	Av        **NodeT
	Ht1       float64
	Ht2       float64
	Pht1      float64
	Pht2      float64
	Candidate c.Int
	Valid     c.Int
	CacheNc   c.Int
	Flat      *AdjmatrixT
}
type RatioT c.Int

const (
	RatioTRNONE     RatioT = 0
	RatioTRVALUE    RatioT = 1
	RatioTRFILL     RatioT = 2
	RatioTRCOMPRESS RatioT = 3
	RatioTRAUTO     RatioT = 4
	RatioTREXPAND   RatioT = 5
)

type LayoutT struct {
	Quantum   float64
	Scale     float64
	Ratio     float64
	Dpi       float64
	Margin    Pointf
	Page      Pointf
	Size      Pointf
	Filled    c.Int
	Landscape c.Int
	Centered  c.Int
	RatioKind RatioT
	Xdots     unsafe.Pointer
	Id        *int8
}

type FieldT struct {
	Size  Pointf
	B     Boxf
	NFlds c.Int
	Lp    *TextlabelT
	Fld   **FieldT
	Id    *int8
	LR    int8
	Sides int8
}

type NlistT struct {
	List **NodeT
	Size uintptr
}

type Elist struct {
	List **EdgeT
	Size uintptr
}
type FontnameKind c.Int

const (
	FontnameKindNATIVEFONTS FontnameKind = 0
	FontnameKindPSFONTS     FontnameKind = 1
	FontnameKindSVGFONTS    FontnameKind = 2
)
/// @addtogroup cgraph_graph
/// @{
type AgraphinfoT struct {
	Hdr          AgrecT
	Drawing      *LayoutT
	Label        *TextlabelT
	Bb           Boxf
	Border       [4]Pointf
	GuiState     int8
	HasLabels    int8
	HasImages    c.Int
	Charset      int8
	Rankdir      c.Int
	Ht1          float64
	Ht2          float64
	Flags        uint16
	Alg          unsafe.Pointer
	Gvc          *GVCT
	Cleanup      unsafe.Pointer
	NeatoNlist   **NodeT
	Move         c.Int
	Dist         **float64
	Spring       **float64
	SumT         **float64
	T            ***float64
	Ndim         uint16
	Odim         uint16
	NCluster     c.Int
	Clust        **GraphT
	Dotroot      *GraphT
	Nlist        *NodeT
	Rank         *RankT
	Parent       *GraphT
	Level        c.Int
	Minrep       *NodeT
	Maxrep       *NodeT
	Comp         NlistT
	Minset       *NodeT
	Maxset       *NodeT
	Minrank      c.Int
	Maxrank      c.Int
	HasFlatEdges c.Int
	Showboxes    int8
	Fontnames    FontnameKind
	Nodesep      c.Int
	Ranksep      c.Int
	Ln           *NodeT
	Rn           *NodeT
	Leader       *NodeT
	Rankleader   **NodeT
	Expanded     c.Int
	Installed    int8
	SetType      int8
	LabelPos     int8
	ExactRanksep c.Int
}
/// @addtogroup cgraph_node
/// @{
type AgnodeinfoT struct {
	Hdr           AgrecT
	Shape         *ShapeDesc
	ShapeInfo     unsafe.Pointer
	Coord         Pointf
	Width         float64
	Height        float64
	Bb            Boxf
	Ht            float64
	Lw            float64
	Rw            float64
	OutlineWidth  float64
	OutlineHeight float64
	Label         *TextlabelT
	Xlabel        *TextlabelT
	Alg           unsafe.Pointer
	State         int8
	GuiState      int8
	Clustnode     c.Int
	Pinned        int8
	Id            c.Int
	Heapindex     c.Int
	Hops          c.Int
	Pos           *float64
	Dist          float64
	Showboxes     int8
	HasPort       c.Int
	Rep           *NodeT
	Set           *NodeT
	NodeType      int8
	Mark          uintptr
	Onstack       int8
	Ranktype      int8
	WeightClass   int8
	Next          *NodeT
	Prev          *NodeT
	In            Elist
	Out           Elist
	FlatOut       Elist
	FlatIn        Elist
	Other         Elist
	Clust         *GraphT
	UFSize        c.Int
	UFParent      *NodeT
	Rank          c.Int
	Order         c.Int
	Mval          float64
	SaveIn        Elist
	SaveOut       Elist
	TreeIn        Elist
	TreeOut       Elist
	Par           *EdgeT
	Low           c.Int
	Lim           c.Int
	Priority      c.Int
	Pad           [1]float64
}
/// @addtogroup cgraph_edge
/// @{
type AgedgeinfoT struct {
	Hdr         AgrecT
	Spl         *Splines
	TailPort    Port
	HeadPort    Port
	Label       *TextlabelT
	HeadLabel   *TextlabelT
	TailLabel   *TextlabelT
	Xlabel      *TextlabelT
	EdgeType    int8
	Compound    int8
	Adjacent    int8
	LabelOntop  int8
	GuiState    int8
	ToOrig      *EdgeT
	Alg         unsafe.Pointer
	Factor      float64
	Dist        float64
	Path        PpolylineT
	Showboxes   int8
	ConcOppFlag c.Int
	Xpenalty    int16
	Weight      c.Int
	Cutvalue    c.Int
	TreeIndex   c.Int
	Count       int16
	Minlen      c.Int
	ToVirt      *EdgeT
}

type LayoutFeaturesT struct {
	Flags c.Int
}
