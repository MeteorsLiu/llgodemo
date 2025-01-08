package graphviz

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

type ImagetypeT c.Int

const (
	ImagetypeTFTNULL ImagetypeT = 0
	ImagetypeTFTBMP  ImagetypeT = 1
	ImagetypeTFTGIF  ImagetypeT = 2
	ImagetypeTFTPNG  ImagetypeT = 3
	ImagetypeTFTJPEG ImagetypeT = 4
	ImagetypeTFTPDF  ImagetypeT = 5
	ImagetypeTFTPS   ImagetypeT = 6
	ImagetypeTFTEPS  ImagetypeT = 7
	ImagetypeTFTSVG  ImagetypeT = 8
	ImagetypeTFTXML  ImagetypeT = 9
	ImagetypeTFTRIFF ImagetypeT = 10
	ImagetypeTFTWEBP ImagetypeT = 11
	ImagetypeTFTICO  ImagetypeT = 12
	ImagetypeTFTTIFF ImagetypeT = 13
)

type ImagescaleT c.Int

const (
	ImagescaleTIMAGESCALEFALSE  ImagescaleT = 0
	ImagescaleTIMAGESCALETRUE   ImagescaleT = 1
	ImagescaleTIMAGESCALEWIDTH  ImagescaleT = 2
	ImagescaleTIMAGESCALEHEIGHT ImagescaleT = 3
	ImagescaleTIMAGESCALEBOTH   ImagescaleT = 4
)

type ImageposT c.Int

const (
	ImageposTIMAGEPOSTOPLEFT      ImageposT = 0
	ImageposTIMAGEPOSTOPCENTER    ImageposT = 1
	ImageposTIMAGEPOSTOPRIGHT     ImageposT = 2
	ImageposTIMAGEPOSMIDDLELEFT   ImageposT = 3
	ImageposTIMAGEPOSMIDDLECENTER ImageposT = 4
	ImageposTIMAGEPOSMIDDLERIGHT  ImageposT = 5
	ImageposTIMAGEPOSBOTTOMLEFT   ImageposT = 6
	ImageposTIMAGEPOSBOTTOMCENTER ImageposT = 7
	ImageposTIMAGEPOSBOTTOMRIGHT  ImageposT = 8
)

type UsershapeS struct {
	Link       DtlinkT
	Name       *int8
	MacroId    c.Int
	MustInline c.Int
	Nocache    c.Int
	F          *c.FILE
	Type       ImagetypeT
	Stringtype *int8
	X          float64
	Y          float64
	W          float64
	H          float64
	Dpi        c.Int
	Data       unsafe.Pointer
	Datasize   uintptr
	Datafree   unsafe.Pointer
}
type UsershapeT UsershapeS
