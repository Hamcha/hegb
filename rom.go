package hegb

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strings"
)

// MemoryController is the required interface for any sort of ROM MBC
type MemoryController interface {
	Read() uint8
	Write(addr uint16, data uint8)
}

// ROM defines the header and content of a ROM file
type ROM struct {
	Header     ROMHeader
	Controller MemoryController
}

func LoadROM(data []byte) *ROM {
	headerPacked := struct {
		Entrypoint      uint32
		NintendoLogo    [0x30]byte
		Title           [0x0b]byte
		ManCode         [4]byte
		GBCFlag         GBCFlag
		NewLicenseeCode [2]byte
		SGBflag         byte
		ROMType         ROMType
		ROMSize         ROMSizeType
		RAMSize         RAMSizeType
		DestCode        DestinationCode
		OldLicenseeCode uint8
		MaskROMversion  uint8
		HeaderChecksum  uint8
		GlobalChecksum  [2]byte
	}{}
	binary.Read(bytes.NewReader(data[0x100:]), binary.BigEndian, &headerPacked)

	rom := new(ROM)
	rom.Header = ROMHeader{
		Entrypoint:      uint16(headerPacked.Entrypoint),
		NintendoLogo:    headerPacked.NintendoLogo,
		Title:           strings.Trim(string(headerPacked.Title[:]), "\000"),
		ManufacturerID:  string(headerPacked.ManCode[:]),
		GBCFlag:         headerPacked.GBCFlag,
		HasSuperGB:      headerPacked.SGBflag == 0x03,
		OldLicenseeCode: headerPacked.OldLicenseeCode,
		NewLicenseeCode: headerPacked.NewLicenseeCode,
		Type:            headerPacked.ROMType,
		ROMSize:         headerPacked.ROMSize,
		RAMSize:         headerPacked.RAMSize,
		Region:          headerPacked.DestCode,
		MaskROMVersion:  headerPacked.MaskROMversion,
	}

	//TODO Make controller

	return rom
}

func (r ROM) String() string {
	return fmt.Sprintf("ROM title: \"%s\"\nRegion: %s\nGBC support: %s\nSGB support: %v\nROM type: %s\nROM size: %s\nRAM size: %s", r.Header.Title, r.Header.Region, r.Header.GBCFlag, r.Header.HasSuperGB, r.Header.Type, r.Header.ROMSize, r.Header.RAMSize)
}

// ROMHeader contains the header fields of a ROM file
type ROMHeader struct {
	Entrypoint     uint16
	NintendoLogo   [0x30]byte
	Title          string
	ManufacturerID string
	GBCFlag
	HasSuperGB      bool
	OldLicenseeCode uint8
	NewLicenseeCode [2]byte
	Type            ROMType
	ROMSize         ROMSizeType
	RAMSize         RAMSizeType
	Region          DestinationCode
	MaskROMVersion  uint8
}

// ROM types
type ROMType uint8

const (
	ROMTypeONLY    ROMType = 0x00 // ROM Only (32kB)
	ROMTypeRAM     ROMType = 0x08 // ROM + RAM
	ROMTypeRB      ROMType = 0x09 // ROM + RAM + Battery
	ROMTypeMBC1    ROMType = 0x01 // Memory Bank Controller 1 (16/4Mbit)
	ROMTypeMBC1R   ROMType = 0x02 // MBC1 + RAM
	ROMTypeMBC1RB  ROMType = 0x03 // MBC1 + RAM + Battery
	ROMTypeMBC2    ROMType = 0x05 // Memory Bank Controller 2 (<= 2Mbit)
	ROMTypeMBC2RB  ROMType = 0x06 // MBC2 + RAM + Battery
	ROMTypeMBC3    ROMType = 0x11 // Memory Bank Controller 3 (16Mbit)
	ROMTypeMBC3R   ROMType = 0x12 // MBC3 + RAM
	ROMTypeMBC3RB  ROMType = 0x13 // MBC3 + RAM + Battery
	ROMTypeMBC3TB  ROMType = 0x0f // MBC3 + Timer + Battery
	ROMTypeMBC3TRB ROMType = 0x10 // MBC3 + Timer + RAM + Battery
	ROMTypeMBC5    ROMType = 0x19 // Memory Bank Controller 5 (64Mbit)
	ROMTypeMBC5R   ROMType = 0x1a // MBC5 + RAM
	ROMTypeMBC5RB  ROMType = 0x1b // MBC5 + RAM + Battery
	ROMTypeMBC5RMB ROMType = 0x1c // MBC5 + Rumble
	ROMTypeMBC5RR  ROMType = 0x1d // MBC5 + Rumble + RAM
	ROMTypeMBC5RRB ROMType = 0x1e // MBC5 + Rumble + RAM + Battery
	ROMTypeMBC6RB  ROMType = 0x20 // MBC6 + RAM + Battery
	ROMTypeMBC7RBA ROMType = 0x22 // MBC7 + RAM + Battery + Accelerometer
	ROMTypeMMM1    ROMType = 0x0b // no idea
	ROMTypeMMM1R   ROMType = 0x0c // MMM1 + RAM
	ROMTypeMMM1RB  ROMType = 0x0d // MMM1 + RAM + Battery
	ROMTypeCAMERA  ROMType = 0xfc // Pocket Camera
	ROMTypeTAMA5   ROMType = 0xfd // Bandai TAMA5
	ROMTypeHUC3    ROMType = 0xfe // (???)
	ROMTypeHUC1RB  ROMType = 0xff // Hudson Soft Infrared MBC1
)

func (d ROMType) String() string {
	switch d {
	case ROMTypeONLY:
		return "ROM Only (32kB)"
	case ROMTypeRAM:
		return "ROM + RAM"
	case ROMTypeRB:
		return "ROM + RAM + Battery"
	case ROMTypeMBC1:
		return "Memory Bank Controller 1 (16/4Mbit)"
	case ROMTypeMBC1R:
		return "MBC1 + RAM"
	case ROMTypeMBC1RB:
		return "MBC1 + RAM + Battery"
	case ROMTypeMBC2:
		return "Memory Bank Controller 2 (<= 2Mbit)"
	case ROMTypeMBC2RB:
		return "MBC2 + RAM + Battery"
	case ROMTypeMBC3:
		return "Memory Bank Controller 3 (16Mbit)"
	case ROMTypeMBC3R:
		return "MBC3 + RAM"
	case ROMTypeMBC3RB:
		return "MBC3 + RAM + Battery"
	case ROMTypeMBC3TB:
		return "MBC3 + Timer + Battery"
	case ROMTypeMBC3TRB:
		return "MBC3 + Timer + RAM + Battery"
	case ROMTypeMBC5:
		return "Memory Bank Controller 5 (64Mbit)"
	case ROMTypeMBC5R:
		return "MBC5 + RAM"
	case ROMTypeMBC5RB:
		return "MBC5 + RAM + Battery"
	case ROMTypeMBC5RMB:
		return "MBC5 + Rumble"
	case ROMTypeMBC5RR:
		return "MBC5 + Rumble + RAM"
	case ROMTypeMBC5RRB:
		return "MBC5 + Rumble + RAM + Battery"
	case ROMTypeMBC6RB:
		return "MBC6 + RAM + Battery"
	case ROMTypeMBC7RBA:
		return "MBC7 + RAM + Battery + Accelerometer"
	case ROMTypeMMM1:
		return "MMM1 (?)"
	case ROMTypeMMM1R:
		return "MMM1 + RAM"
	case ROMTypeMMM1RB:
		return "MMM1 + RAM + Battery"
	case ROMTypeCAMERA:
		return "Pocket Camera"
	case ROMTypeTAMA5:
		return "Bandai TAMA5"
	case ROMTypeHUC3:
		return "HUC3 (?)"
	case ROMTypeHUC1RB:
		return "Hudson Soft Infrared MBC1"
	}
	return "<unknown>"
}

// ROM sizes
type ROMSizeType uint8

const (
	ROMSize32K  ROMSizeType = 0x00 //  32kB,   2 banks
	ROMSize64K  ROMSizeType = 0x01 //  64kB,   4 banks
	ROMSize128K ROMSizeType = 0x02 // 128kB,   8 banks
	ROMSize256K ROMSizeType = 0x03 // 256kB,  16 banks
	ROMSize512K ROMSizeType = 0x04 // 512kB,  32 banks
	ROMSize1M   ROMSizeType = 0x05 //   1MB,  64 banks (only 63 used in MBC1)
	ROMSize2M   ROMSizeType = 0x06 //   2MB, 128 banks (only 125 used in MBC1)
	ROMSize4M   ROMSizeType = 0x07 //   4MB, 256 banks
	ROMSize8M   ROMSizeType = 0x08 //   8MB, 512 banks
)

func (d ROMSizeType) String() string {
	switch d {
	case ROMSize32K:
		return "32kiB, 2 banks"
	case ROMSize64K:
		return "64kiB, 4 banks"
	case ROMSize128K:
		return "128kiB, 8 banks"
	case ROMSize256K:
		return "256kiB, 16 banks"
	case ROMSize512K:
		return "512kiB, 32 banks"
	case ROMSize1M:
		return "1MiB, 64 banks"
	case ROMSize2M:
		return "2MiB, 128 banks"
	case ROMSize4M:
		return "4MiB, 256 banks"
	case ROMSize8M:
		return "8MiB, 512 banks"
	}
	return "<unknown>"
}

//! RAM Sizes
type RAMSizeType uint8

const (
	RAMSizeNONE  RAMSizeType = 0x00 //   No RAM
	RAMSize2KB   RAMSizeType = 0x01 //  2kB RAM
	RAMSize8KB   RAMSizeType = 0x02 //  8kB RAM (1 bank)
	RAMSize32KB  RAMSizeType = 0x03 // 32kB RAM (4 banks of 8kB)
	RAMSize64KB  RAMSizeType = 0x05 // 64kB RAM (8 banks of 8kB)
	RAMSize128KB RAMSizeType = 0x04 // 128kB RAM (16 banks of 8kB)
)

func (d RAMSizeType) String() string {
	switch d {
	case RAMSizeNONE:
		return "None"
	case RAMSize2KB:
		return "2kiB"
	case RAMSize8KB:
		return "8kiB (1 bank)"
	case RAMSize32KB:
		return "32kiB (4 banks of 8kiB)"
	case RAMSize64KB:
		return "64kiB (8 banks of 8kiB)"
	case RAMSize128KB:
		return "128kiB (16 banks of 8kiB)"
	}
	return "<unknown>"
}

//! Destination Code
type DestinationCode uint8

const (
	RegionJapanese    DestinationCode = 0x00 // Japanese game
	RegionNonJapanese DestinationCode = 0x01 // Non-Japanese game
)

func (d DestinationCode) String() string {
	switch d {
	case RegionJapanese:
		return "Japan-only"
	case RegionNonJapanese:
		return "Worldwide"
	}
	return "<unknown>"
}

type GBCFlag uint8

const (
	GBCSupported GBCFlag = 0x80
	GBCOnly      GBCFlag = 0xc0
)

func (d GBCFlag) String() string {
	switch d {
	case GBCSupported:
		return "Supports GBC"
	case GBCOnly:
		return "Requires GBC"
	}
	return "No GBC support"
}
