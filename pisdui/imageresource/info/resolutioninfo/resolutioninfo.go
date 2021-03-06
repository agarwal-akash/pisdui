package resolutioninfo

import (
	"os"

	util "github.com/fabulousduck/pisdui/pisdui/util/file"
)

//Note: HorizontalResolution and VerticalResolution are byte buffers as they are fixed point integers
type Resolutioninfo struct {
	HorizontalResolution     float32
	HorizontalResolutionUnit string
	WidthResolutionUnit      string
	VerticalResolution       float32
	VerticalResolutionUnit   string
	HeightUnit               string
}

func (resolutioninfo *Resolutioninfo) GetTypeID() int {
	return 1005
}

func NewResolutionInfo() *Resolutioninfo {
	return new(Resolutioninfo)
}

func (resolutioninfo *Resolutioninfo) Parse(file *os.File) {
	resolutioninfo.HorizontalResolution = parseFixedPoint(util.ReadRawBytes(file, 4))
	resolutioninfo.HorizontalResolutionUnit = parseUnit(util.ReadBytesShort(file))
	resolutioninfo.WidthResolutionUnit = parseUnit(util.ReadBytesShort(file))
	resolutioninfo.VerticalResolution = parseFixedPoint(util.ReadRawBytes(file, 4))
	resolutioninfo.VerticalResolutionUnit = parseUnit(util.ReadBytesShort(file))
	resolutioninfo.HeightUnit = parseUnit(util.ReadBytesShort(file))
}

func parseUnit(unit uint16) string {
	opts := map[uint16]string{
		1: "inches",
		2: "cm",
		3: "points",
		4: "picas",
		5: "columns",
	}
	return opts[unit]
}

//This assumes the buffer is big endian
func parseFixedPoint(buffer []byte) float32 {
	var n float32
	f := buffer[1] | buffer[1]<<8 | buffer[2]<<16
	n = float32(f) + float32(buffer[0])/100
	return n
}
