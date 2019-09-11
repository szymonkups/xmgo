package xmparser

import (
	"encoding/binary"
	"fmt"
	"os"
)

type Pattern struct {
	Header PatternHeader
	Rows   []Row
}

type PatternHeader struct {
	HeaderLength uint32
	PackingType  uint8
	NoRows       uint16
	DataSize     uint16
}

type Row = []Cell

// ParsePatterns parses XM file patterns starting of given offset
func parsePattern(f *os.File, fileHeader *XMFileHeader) (*Pattern, error) {
	pattern := Pattern{}

	header, err := parsePatternHeader(f)

	if err != nil {
		return nil, err
	}

	pattern.Header = *header
	pattern.Rows = make([]Row, header.NoRows)

	if header.DataSize > 0 {
		for r := uint16(0); r < header.NoRows; r++ {
			pattern.Rows[r] = make([]Cell, fileHeader.NoChannels)
			for c := uint16(0); c < fileHeader.NoChannels; c++ {
				cell, err := parseCell(f)

				if err != nil {
					return nil, err
				}

				pattern.Rows[r][c] = *cell
			}
		}
	}

	return &pattern, nil
}

func parsePatternHeader(f *os.File) (*PatternHeader, error) {
	header := PatternHeader{}
	err := binary.Read(f, binary.LittleEndian, &header)

	if err != nil {
		return nil, err
	}

	if header.PackingType != supportedPacking {
		return nil, fmt.Errorf("Wrong patter packing, expected %d but got %d", supportedPacking, header.PackingType)
	}

	return &header, nil
}
