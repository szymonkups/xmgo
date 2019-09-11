package xmparser

import (
	"encoding/binary"
	"os"
)

// XMFileHeader is a binary representation of XM module file header
type XMFileHeader struct {
	ID              [17]uint8
	Name            [20]uint8
	Ctrl            uint8
	TrackerName     [20]uint8
	VersionNumber   uint16
	HeaderSize      uint32
	SongLength      uint16
	RestartPosition uint16
	NoChannels      uint16
	NoPatterns      uint16
	NoInstruments   uint16
	FreqTable       uint16
	DefTempo        uint16
	DefBPM          uint16
	PatternOrder    [256]uint8
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
