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
	header      XMFileHeader
	patterns    []Pattern
	Instruments []Instrument
}

func ParseFile(filename string) (*Song, error) {
	f, err := os.Open(filename)

	if err != nil {
		return nil, err
	}

	defer f.Close()

	header, err := parseHeader(f)
	if err != nil {
		return nil, err
	}

	// HeaderSize is calculated from place where this info is stored in header
	f.Seek(int64(header.HeaderSize+headerSizeOffset), 0)

	patterns := make([]Pattern, header.NoPatterns)
	for p := uint16(0); p < header.NoPatterns; p++ {
		pattern, err := parsePattern(f, header)

		if err != nil {
			return nil, err
		}

		patterns[p] = *pattern
	}

	instruments := make([]Instrument, header.NoInstruments)
	for i := uint16(0); i < header.NoInstruments; i++ {
		instrument, err := parseInstrument(f)

		if err != nil {
			return nil, err
		}

		instruments[i] = *instrument
	}

	return &Song{*header, patterns, instruments}, nil
}
