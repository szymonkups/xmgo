package xmparser

import (
	"encoding/binary"
	"os"
)

type Instrument struct {
	Header  InstrumentHeader
	Info    InstrumentInfo
	Samples []Sample
}

type InstrumentHeader struct {
	HeaderSize uint32
	Name       [22]uint8
	Type       uint8
	NoSamples  uint16
}

type InstrumentInfo struct {
	SampleHeaderSize uint32
	SamplePerNote    [96]uint8
	VolPoints        [48]uint8
	PanPoints        [48]uint8
	NoVolPoints      uint8
	NoPanPoints      uint8
	VolSustainPoint  uint8
	VolLoopStart     uint8
	VolLoopEnd       uint8
	PanSustainPoint  uint8
	PanLoopStart     uint8
	PanLoopEnd       uint8
	VolType          uint8
	PanType          uint8
	VibratoType      uint8
	VibratoSweep     uint8
	VibratoDepth     uint8
	VibratoRate      uint8
	VolFadeout       uint16
	Reserved         uint16
}

func parseInstrument(f *os.File) (*Instrument, error) {
	startPosition, err := f.Seek(0, os.SEEK_CUR)
	if err != nil {
		return nil, err
	}

	header := InstrumentHeader{}
	err = binary.Read(f, binary.LittleEndian, &header)

	if err != nil {
		return nil, err
	}

	// If there are samples - read additional info
	info := InstrumentInfo{}
	if header.NoSamples > 0 {
		err := binary.Read(f, binary.LittleEndian, &info)

		if err != nil {
			return nil, err
		}
	}

	// Skip instrument header - it might be bigger than data already read
	_, err = f.Seek(startPosition+int64(header.HeaderSize), os.SEEK_SET)
	if err != nil {
		return nil, err
	}

	samples := make([]Sample, header.NoSamples)
	if header.NoSamples > 0 {
		_, err = parseSamples(f, header.NoSamples, info.SampleHeaderSize)
		if err != nil {
			return nil, err
		}
	}

	return &Instrument{header, info, samples}, nil
}
