package hegb

import (
	"errors"
	"fmt"
)

// MemoryController is the required interface for any sort of ROM MBC
type MemoryController interface {
	Read(addr uint16) (uint8, error)
	Write(addr uint16, data uint8) error
}

type rombank [16 * 1024]byte
type rambank [8 * 1024]byte

// No MBC (just 32k)
type mbc0 struct {
	HasRAM     bool
	HasBattery bool

	romtype ROMType
	ramsize RAMSizeType

	rombanks []rombank
	rambanks []rambank
}

func loadMBC0(rominfo ROMHeader, data []byte) (*mbc0, error) {
	romtype := rominfo.Type

	rombanks, rambanks, err := loadBanks(rominfo, data)
	if err != nil {
		return nil, err
	}

	mbc := &mbc0{
		HasRAM:     romtype == ROMTypeRAM || romtype == ROMTypeRB,
		HasBattery: romtype == ROMTypeRB,
		romtype:    romtype,
		ramsize:    rominfo.RAMSize,
		rombanks:   rombanks,
		rambanks:   rambanks,
	}

	return mbc, nil
}

func (m *mbc0) Read(addr uint16) (uint8, error) {
	if addr < 0x4000 {
		return m.rombanks[0][addr], nil
	}
	if addr < 0x8000 {
		return m.rombanks[1][addr-0x4000], nil
	}
	if addr < 0xa000 {
		return 0, errors.New("trying to access VRAM in ROM")
	}
	// External RAM (on cartridge)
	if m.HasRAM {
		if addr < 0xc000 {
			return m.rambanks[0][addr-0xa000], nil
		}
	} else {
		return 0, errors.New("trying to access inexistent cartridge RAM")
	}
	return 0, errors.New("out of bound ROM read")
}

func (m *mbc0) Write(addr uint16, data uint8) error {
	// No MBC0 games write to cartridge (or should), ignore any write
	return nil
}

func loadBanks(rominfo ROMHeader, data []byte) ([]rombank, []rambank, error) {
	isMBC1 := rominfo.Type == ROMTypeMBC1 || rominfo.Type == ROMTypeMBC1R || rominfo.Type == ROMTypeMBC1RB
	var romcount int
	switch rominfo.ROMSize {
	case ROMSize32K:
		romcount = 2
	case ROMSize64K:
		romcount = 4
	case ROMSize128K:
		romcount = 8
	case ROMSize256K:
		romcount = 16
	case ROMSize512K:
		romcount = 32
	case ROMSize1M:
		if isMBC1 {
			romcount = 63
		} else {
			romcount = 64
		}
	case ROMSize2M:
		if isMBC1 {
			romcount = 125
		} else {
			romcount = 128
		}
	case ROMSize4M:
		romcount = 256
	case ROMSize8M:
		romcount = 512
	default:
		return nil, nil, fmt.Errorf("unsupported ROM size flag: %2x", rominfo.ROMSize)
	}

	rombanks := make([]rombank, romcount)
	if len(data) < romcount*16*1024 {
		return nil, nil, fmt.Errorf("ROM size should be %d bytes, but only %d could be found", romcount*16*1024, len(data))
	}
	for bankidx := 0; bankidx < romcount; bankidx++ {
		copy(rombanks[bankidx][:], data[bankidx*16*1024:(bankidx+1)*16*1024])
	}

	var ramcount uint8
	switch rominfo.RAMSize {
	case RAMSizeNONE:
		ramcount = 0
	case RAMSize2KB, RAMSize8KB:
		ramcount = 1
	case RAMSize32KB:
		ramcount = 4
	case RAMSize64KB:
		ramcount = 8
	case RAMSize128KB:
		ramcount = 16
	default:
		return nil, nil, fmt.Errorf("unsupported RAM size flag: %2x", rominfo.RAMSize)
	}
	rambanks := make([]rambank, ramcount, ramcount)

	return rombanks, rambanks, nil
}
