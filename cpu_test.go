package hegb

import (
	"errors"
	"fmt"
	"os"
	"testing"
)

func TestLoadRegister(t *testing.T) {
	gb := runCode([]byte{
		0x01, 0x12, 0x34, // LD BC, 0x3412
		0x11, 0xc3, 0xff, // LD DE, 0xFFC3
		0x21, 0xaa, 0xbb, // LD HL, 0xBBAA
		0x41, // LD B, C
		0x63, // LD H, E
		0x70, // LD (HL), B
		0x7e, // LD A, (HL)
	})

	// Check values
	checkReg(t, gb, map[RegID]uint16{
		RegBC: 0x1212,
		RegDE: 0xffc3,
		RegHL: 0xc3aa,
		RegA:  0x12,
	})
	checkCycles(t, gb, Cycles{13, 60})
}

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
	checkCycles(t, gb, Cycles{12, 48})
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
	checkCycles(t, gb, Cycles{14, 56})
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
	checkCycles(t, gb, Cycles{4, 32})
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
	checkCycles(t, gb, Cycles{4, 32})
}

func TestIncrement8(t *testing.T) {
	gb := runCode([]byte{
		0x04, // INC B
		0x0c, // INC C
		0x14, // INC D
		0x1c, // INC E
		0x24, // INC D
		0x2c, // INC E
	})

	checkReg(t, gb, map[RegID]uint16{
		RegBC: 0x0101,
		RegDE: 0x0101,
		RegHL: 0x0101,
		RegSP: 0x0101,
	})
	checkCycles(t, gb, Cycles{6, 24})
}

func TestDecrement8(t *testing.T) {
	gb := runCode([]byte{
		0x05, // DEC B
		0x0d, // DEC C
		0x15, // DEC D
		0x1d, // DEC E
		0x25, // DEC H
		0x2d, // DEC L
	})

	checkReg(t, gb, map[RegID]uint16{
		RegBC: 0xffff,
		RegDE: 0xffff,
		RegHL: 0xffff,
	})
	checkCycles(t, gb, Cycles{6, 24})
}
func TestBitSet(t *testing.T) {
	gb := runCode([]byte{
		0xcb, 0xc7, // SET 0, A
		0xcb, 0xdf, // SET 3, A
		0xcb, 0xe8, // SET 5, B
	})
	checkReg(t, gb, map[RegID]uint16{
		RegA: 0x09,
		RegB: 0x20,
	})
	checkCycles(t, gb, Cycles{6, 24})
}

func TestBitReset(t *testing.T) {
	gb := runCode([]byte{
		0x3e, 0xff, // LD A, 0xFF
		0xcb, 0x8f, // RES 1, A
		0xcb, 0xaf, // RES 5, A
	})
	checkReg(t, gb, map[RegID]uint16{
		RegA: 0xdd,
	})
	checkCycles(t, gb, Cycles{6, 24})
}

func TestBitCheck(t *testing.T) {
	// Test with zero flag set
	gb := runCode([]byte{
		0x3e, 0xfd, // LD A, 0xFD
		0xcb, 0x4f, // BIT 1, A
	})
	checkReg(t, gb, map[RegID]uint16{
		RegAF: 0xfda0,
	})
	checkCycles(t, gb, Cycles{4, 16})

	// Check that zero flag is clear after this one
	gb = runCode([]byte{
		0x3e, 0xff, // LD A, 0xFD
		0xcb, 0x4f, // BIT 1, A
	})
	checkReg(t, gb, map[RegID]uint16{
		RegAF: 0xff20,
	})
}

func TestSwapNibble(t *testing.T) {
	gb := runCode([]byte{
		0x3e, 0xdf, // LD A, 0xDF
		0xcb, 0x37, // SWAP A
		// Extra: Check zero flag
		0x06, 0x00, // LD B, 0x00
		0xcb, 0x30, // SWAP B
	})
	checkReg(t, gb, map[RegID]uint16{
		RegAF: 0xfd80,
		RegB:  0,
	})
	checkCycles(t, gb, Cycles{8, 32})
}

func TestCBRotate(t *testing.T) {
	gb := runCode([]byte{
		0x3e, 0xf0, // LD A, 0xF0
		0x06, 0x0f, // LD B, 0x0F
		0x0e, 0xf0, // LD C, 0xF0
		0x16, 0x0f, // LD D, 0x0F
		0xcb, 0x07, // RLC A
		0xcb, 0x08, // RRC B
		0xcb, 0x11, // RL C
		0xcb, 0x1a, // RR D
		// Reset carry flag
		0xcb, 0x37, // SWAP A
		// Check WithCarry rotation with carry not set
		0xcb, 0x11, // RL C
	})
	checkReg(t, gb, map[RegID]uint16{
		RegAF: 0x1e10,
		RegB:  0x87,
		RegC:  0xc2,
		RegD:  0x87,
	})
	checkCycles(t, gb, Cycles{20, 80})
}

func TestCBShift(t *testing.T) {
	gb := runCode([]byte{
		0x3e, 0xf0, // LD A, 0xF0
		0x06, 0xff, // LD B, 0x0F
		0x0e, 0xff, // LD C, 0xFF
		0xcb, 0x27, // SLA A
		0xcb, 0x38, // SRL B
		0xcb, 0x29, // SRA C
	})
	checkReg(t, gb, map[RegID]uint16{
		RegAF: 0xe010,
		RegB:  0x7f,
		RegC:  0xff,
	})
	checkCycles(t, gb, Cycles{12, 48})
}

func TestRotateAcc(t *testing.T) {
	checkReg(t, runCode([]byte{
		0x3e, 0xf0, // LD A, 0xF0
		0x07, // RLCA
	}), map[RegID]uint16{
		RegAF: 0xe110,
	})

	checkReg(t, runCode([]byte{
		0x3e, 0xf0, // LD A, 0xF0
		0x0f, // RRCA
	}), map[RegID]uint16{
		RegAF: 0x7800,
	})

	checkReg(t, runCode([]byte{
		0x3e, 0xf0, // LD A, 0xF0
		0x17, // RLA
	}), map[RegID]uint16{
		RegAF: 0xe010,
	})

	gb := runCode([]byte{
		0x3e, 0xff, // LD A, 0xF0
		0x1f, // RRA
	})
	checkReg(t, gb, map[RegID]uint16{
		RegAF: 0x7f10,
	})
	checkCycles(t, gb, Cycles{3, 12})
}

// Test all instructions to check that they are all handled
func TestHandlerPresence(t *testing.T) {
	handled := 0
	unhandled := 0
	jsrows := make([]byte, 0x200)
	// Test standard instructions
	for i := OpNop; i <= OpRestart38; i++ {
		// There are some holes, skip them
		if i == 0xfc || i == 0xfd || i == OpCBPrefix || i == 0xf4 || i == 0xeb || i == 0xec || i == 0xed || i == 0xe3 || i == 0xe4 || i == 0xdd || i == 0xdb || i == 0xd3 {
			jsrows[i-OpNop] = '-'
			continue
		}
		if _, ok := cpuhandlers[i]; ok {
			handled++
			jsrows[i-OpNop] = '1'
			continue
		}
		fmt.Fprintf(os.Stderr, "%02X | %s is MISSING!\n", uint16(i), i)
		unhandled++
		jsrows[i-OpNop] = '0'
	}

	// CB prefix instructions
	for i := OpCbRotateRegBLeftRot; i <= OpCbSetDirectA7; i++ {
		if _, ok := cpuhandlers[i]; ok {
			handled++
			jsrows[0x100+i-OpCbRotateRegBLeftRot] = '1'
			continue
		}
		fmt.Fprintf(os.Stderr, "CB %02X | %s is MISSING!\n", uint16(i-OpCbRotateRegBLeftRot), i)
		unhandled++
		jsrows[0x100+i-OpCbRotateRegBLeftRot] = '0'
	}

	fmt.Fprintf(os.Stderr, "Summary: %d handled, %d missing (%.2f%% total)\n", handled, unhandled, (float32(handled) / float32(handled+unhandled) * 100))
	fmt.Fprintf(os.Stderr, "JS table code: %s\n", jsrows)
	if unhandled > 0 {
		t.Fail()
	}
}

// Test framework

func runCode(code []byte) *Gameboy {
	rom := makeTestROM(append(code, byte(OpStop)))
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
		Controller: &testController{data: code},
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
				t.Fatalf("[Register mismatch] Register %s expected to be %04x, is %04x instead", regid, val, act)
			}
		case RegA, RegF, RegB, RegC, RegD, RegE, RegH, RegL:
			act := getreg8(gb.cpu, regid)
			if act != uint8(val) {
				t.Fatalf("[Register mismatch] Register %s expected to be %02x, is %02x instead", regid, uint8(val), act)
			}
		}
	}
}

func checkCycles(t *testing.T, gb *Gameboy, cycles Cycles) {
	if gb.cpu.Cycles.CPU != cycles.CPU {
		t.Fatalf("[Cycle mismatch] Expected %d CPU cycles, got %d", cycles.CPU, gb.cpu.Cycles.CPU)
	}
	if gb.cpu.Cycles.Machine != cycles.Machine {
		t.Fatalf("[Cycle mismatch] Expected %d machine cycles, got %d", cycles.Machine, gb.cpu.Cycles.Machine)
	}
}
