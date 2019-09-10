package xmparser

import (
	"encoding/binary"
	"os"
)

const (
	idOffset         = 0
	idText           = "Extended Module: "
	versionOffset    = 58
	supportedVersion = 0x104
)

// XMFileHeader is a binary representation of XM module file header
type XMFileHeader struct {
	ID              [17]byte
	Name            [20]byte
	Ctrl            byte
	TrackerName     [20]byte
	VersionNumber   int16
	HeaderSize      int32
	SongLength      int16
	RestartPosition int16
	NoChannels      int16
	NoPatterns      int16
	NoInstruments   int16
	FreqTable       int16
	DefTempo        int16
	DefBPM          int16
	PatternOrder    [256]byte
}

// ParseHeader parses a provided file to XMLFileHeader
func ParseHeader(f *os.File) (XMFileHeader, error) {
	f.Seek(0, 0)

	header := XMFileHeader{}
	err := binary.Read(f, binary.LittleEndian, &header)

	if err != nil {
		return header, err
	}

	return header, nil
}
