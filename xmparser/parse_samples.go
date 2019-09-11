package xmparser

import (
	"encoding/binary"
	"os"
)

type Sample struct {
	Header SampleHeader
}

type SampleHeader struct {
	DataLength uint32
	LoopStart  uint32
	LoopLength uint32
	Vol        uint8
	Finetune   int8
	LoopType   uint8
	Pan        uint8
	NoteNumber int8
	Reserved   uint8
	Name       [22]uint8
}

func parseSamples(f *os.File, noSamples uint16, headerSize uint32) ([]Sample, error) {
	samples := make([]Sample, noSamples)

	// Save current file position so we can skip headers
	startPos, err := f.Seek(0, os.SEEK_CUR)

	if err != nil {
		return nil, err
	}

	// Parse all sample headers
	for h := uint16(0); h < noSamples; h++ {
		header := SampleHeader{}

		err = binary.Read(f, binary.LittleEndian, &header)

		if err != nil {
			return nil, err
		}

		samples[h].Header = header

		// Skip header size after each read - there might be more undocumented data than we read
		startPos, err = f.Seek(startPos+int64(headerSize), os.SEEK_SET)

		if err != nil {
			return nil, err
		}
	}

	// Parse all sample data
	for d := uint16(0); d < noSamples; d++ {
		header := samples[d].Header
		// TODO: parse data

		// Skip by data length from header
		startPos += int64(header.DataLength)
		_, err = f.Seek(startPos, os.SEEK_SET)

		if err != nil {
			return nil, err
		}
	}

	return samples, nil
}
