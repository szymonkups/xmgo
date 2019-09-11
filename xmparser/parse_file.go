package xmparser

import (
	"os"
)

const (
	idOffset         = 0
	idText           = "Extended Module: "
	headerSizeOffset = 60
	versionOffset    = 58
	supportedVersion = 0x104
	supportedPacking = 0
)

type Song struct {
	header   XMFileHeader
	patterns []Pattern
}

func ParseFile(filename string) (Song, error) {
	xm := Song{}
	f, err := os.Open(filename)

	if err != nil {
		return xm, err
	}

	defer f.Close()

	header, err := ParseHeader(f)
	if err != nil {
		return xm, err
	}

	xm.header = header
	xm.patterns = make([]Pattern, header.NoPatterns)

	for i := uint16(0); i < header.NoPatterns; i++ {
		pattern, err := ParsePattern(f, &header)

		if err != nil {
			return xm, err
		}

		xm.patterns[i] = pattern
	}

	return xm, nil
}
