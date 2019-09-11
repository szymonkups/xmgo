package xmparser

import (
	"os"
	"encoding/binary"
)

// Cell is a binary representation of a cell entry in XM file
type Cell struct {
	Note        uint8
	Instrument  uint8
	Volume      uint8
	Effect      uint8
	EffectParam uint8
}

func parseCell(f *os.File) (Cell, error) {
	var b uint8
	err := binary.Read(f, binary.LittleEndian, &b)
	c := Cell{}

	if err != nil {
		return c, err
	}

	// Compression
	if b & 0b10000000 > 0 {
		// Next byte is note 
		if b & 0b00000001 > 0 {
			err = binary.Read(f, binary.LittleEndian, &(c.Note))
			if err != nil {
				return c, err
			}
		}

		// Next byte is instrument
		if b & 0b00000010 > 0 {
			err = binary.Read(f, binary.LittleEndian, &(c.Instrument))
			if err != nil {
				return c, err
			}
		}

		// Next byte is volume
		if b & 0b00000100 > 0 {
			err = binary.Read(f, binary.LittleEndian, &(c.Volume))
			if err != nil {
				return c, err
			}
		}

		// Next byte is effect
		if b & 0b00001000 > 0 {
			err = binary.Read(f, binary.LittleEndian, &(c.Effect))
			if err != nil {
				return c, err
			}
		}

		// Next byte is effect parameter
		if b & 0b00010000 > 0 {
			err = binary.Read(f, binary.LittleEndian, &(c.EffectParam))
			if err != nil {
				return c, err
			}
		}
	} else if b > 0 {
		// No compression

		// Move back so we can read whole cell again including first byte
		f.Seek(-1, 1)

		binary.Read(f, binary.LittleEndian, &c)
		if err != nil {
			return c, err
		}
	}

	return c, nil
}