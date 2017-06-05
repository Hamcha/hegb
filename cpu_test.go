package hegb

import (
	"errors"
	"testing"
)

func TestLoadImmediate16(t *testing.T) {
	// Set registers to random known values
	gb := runCode([]byte{
		0x01, 0x12, 0x34, // LD BC, 0x3412
		0x11, 0x23, 0xff, // LD DE, 0xFF23
		0x21, 0xaa, 0xbb, // LD HL, 0xBBAA
		0x31, 0x00, 0x01, // LD SP, 0x0100
	})

	// Check values
	checkReg(t, gb, map[RegID]uint16{
		RegBC: 0x3412,
		RegDE: 0xff23,
		RegHL: 0xbbaa,
		RegSP: 0x0100,
	})
}

func TestLoadImmediate8(t *testing.T) {
	// Set registers to random known values
	gb := runCode([]byte{
		0x06, 0x11, // LD B, 0x11
		0x0e, 0x22, // LD C, 0x22
		0x16, 0x33, // LD D, 0x33
		0x1e, 0x44, // LD E, 0x44
		0x26, 0x55, // LD H, 0x55
		0x2e, 0x66, // LD L, 0x66
		0x3e, 0x77, // LD A, 0x77
	})

	// Check values
	checkReg(t, gb, map[RegID]uint16{
		RegBC: 0x1122,
		RegDE: 0x3344,
		RegHL: 0x5566,
		RegA:  0x77,
	})
}

func TestIncrement16(t *testing.T) {
	gb := runCode([]byte{
		0x03, // INC BC
		0x13, // INC DE
		0x23, // INC HL
		0x33, // INC SP
	})

	checkReg(t, gb, map[RegID]uint16{
		RegBC: 1,
		RegDE: 1,
		RegHL: 1,
		RegSP: 1,
	})
}

func TestDecrement16(t *testing.T) {
	gb := runCode([]byte{
		0x0b, // DEC BC
		0x1b, // DEC DE
		0x2b, // DEC HL
		0x3b, // DEC SP
	})

	checkReg(t, gb, map[RegID]uint16{
		RegBC: 0xffff,
		RegDE: 0xffff,
		RegHL: 0xffff,
		RegSP: 0xffff,
	})
}

// Test framework

func runCode(code []byte) *Gameboy {
	rom := makeTestROM(code)
	gb := MakeGB(rom, EmulatorOptions{
		SkipBootstrap: true,
		Test:          true,
	})
	gb.Run()
	return gb
}

func makeTestROM(code []byte) *ROM {
	return &ROM{
		Header: ROMHeader{
			Entrypoint: 0x200,
			Title:      "TEST",
			Type:       ROMTypeONLY,
			ROMSize:    ROMSize32K,
			RAMSize:    RAMSizeNONE,
		},
		Controller: &testController{data: append(code, byte(OpStop))},
	}
}

type testController struct {
	data []byte
}

func (t *testController) Read(addr uint16) (uint8, error) {
	if int(addr) >= len(t.data) {
		return 0, errors.New("out of bound read")
	}
	return uint8(t.data[addr]), nil
}

func (t *testController) Write(addr uint16, data uint8) error {
	if int(addr) >= len(t.data) {
		return errors.New("out of bound write")
	}
	t.data[addr] = data
	return nil
}

func checkReg(t *testing.T, gb *Gameboy, vals map[RegID]uint16) {
	for regid, val := range vals {
		switch regid {
		case RegAF, RegBC, RegDE, RegHL, RegSP:
			act := uint16(*reg16(gb.cpu, regid))
			if act != val {
				t.Fatalf("Register %s expected to be %04x, is %04x instead", regid, val, act)
			}
		case RegA, RegF, RegB, RegC, RegD, RegE, RegH, RegL:
			act := getreg8(gb.cpu, regid)
			if act != uint8(val) {
				t.Fatalf("Register %s expected to be %02x, is %02x instead", regid, uint8(val), act)
			}
		}
	}
}
