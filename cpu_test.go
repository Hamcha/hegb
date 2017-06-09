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

func TestStoreHighA(t *testing.T) {
	gb := runCode([]byte{
		0x3e, 0x12, // LD A, 0x12
		0xe0, 0xc2, // LDH 0xC2, A
		0x3e, 0xff, // LD A, 0xFF
		0xf0, 0xc2, // LDH A, 0xC2
	})
	checkReg(t, gb, map[RegID]uint16{
		RegA: 0x12,
	})
	checkCycles(t, gb, Cycles{8, 40})
}

func TestStoreASP(t *testing.T) {
	gb := runCode([]byte{
		0x3e, 0x12, // LD A, 0x12
		0xea, 0xfd, 0xff, // LD 0xfffd, A
		0xea, 0xfc, 0xff, // LD 0xfffc, A
		0x31, 0xfc, 0xff, // LD SP, 0xfffc
		0xc1,             // POP BC
		0x08, 0xfc, 0xff, // LD 0xfffc, SP
		0x31, 0xfc, 0xff, // LD SP, 0xfffc
		0xd1, // POP DE
	})
	checkReg(t, gb, map[RegID]uint16{
		RegA:  0x12,
		RegBC: 0x1212,
		RegDE: 0xfffe,
		RegSP: 0xfffe,
	})
	checkCycles(t, gb, Cycles{19, 108})
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
		RegSP: 0xffff,
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
		RegSP: 0xfffd,
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
		RegSP: 0xfffe,
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

func TestInvertA(t *testing.T) {
	gb := runCode([]byte{
		0x3e, 0xf0, // LD A, 0xF0
		0x2f, // CPL
	})
	checkReg(t, gb, map[RegID]uint16{
		RegA: 0x0f,
	})
	checkCycles(t, gb, Cycles{3, 12})
}

func TestSetInvertCarry(t *testing.T) {
	gb := runCode([]byte{
		0x3e, 0x80, // LD A, 0x80
		0x37, // SCF
		0x17, // RLA (A should be 1, CF should be 1)
		0x3f, // CCF
	})
	checkReg(t, gb, map[RegID]uint16{
		RegAF: 0x0100,
	})
	checkCycles(t, gb, Cycles{5, 20})
}

func TestLoadRegister16(t *testing.T) {
	gb := runCode([]byte{
		0x21, 0x34, 0x12, // LD HL, 0x1234
		0xf9, // LD SP, HL
	})
	checkReg(t, gb, map[RegID]uint16{
		RegSP: 0x1234,
	})
	checkCycles(t, gb, Cycles{4, 20})
}

func TestLoadHighMemC(t *testing.T) {
	gb := runCode([]byte{
		0x0e, 0x81, // LD C, 0x81
		0x3e, 0x10, // LD A, 0x10
		0xe2,       // LD (C), A
		0x3e, 0xff, // LD A, 0xff
		0xf2, // LD A, (C)
	})
	checkReg(t, gb, map[RegID]uint16{
		RegA: 0x10,
	})
	checkCycles(t, gb, Cycles{10, 40})
}

func TestAnd(t *testing.T) {
	gb := runCode([]byte{
		0x0e, 0x81, // LD C, 0x81
		0x3e, 0x10, // LD A, 0x10
		0xa1,       // AND C
		0xe6, 0x01, // AND 0x01
	})
	checkReg(t, gb, map[RegID]uint16{
		RegAF: 0x00a0,
	})
	checkCycles(t, gb, Cycles{7, 28})
}

func TestOr(t *testing.T) {
	gb := runCode([]byte{
		0x0e, 0x81, // LD C, 0x80
		0x3e, 0x10, // LD A, 0x10
		0xb1,       // OR C
		0xf6, 0x01, // OR 0x01
	})
	checkReg(t, gb, map[RegID]uint16{
		RegAF: 0x9100,
	})
	checkCycles(t, gb, Cycles{7, 28})
}

func TestXor(t *testing.T) {
	gb := runCode([]byte{
		0x0e, 0x81, // LD C, 0x81
		0x3e, 0x10, // LD A, 0x10
		0xa9,       // XOR C
		0xee, 0x91, // XOR 0x91
	})
	checkReg(t, gb, map[RegID]uint16{
		RegAF: 0x0080,
	})
	checkCycles(t, gb, Cycles{7, 28})
}

func TestPushPop(t *testing.T) {
	gb := runCode([]byte{
		0x01, 0x34, 0x12, // LD BC, 0x1234
		0xc5,             // PUSH BC
		0x01, 0x11, 0x11, // LD BC, 0x1111
		0xc1, // POP BC
	})
	checkReg(t, gb, map[RegID]uint16{
		RegBC: 0x1234,
		RegSP: 0xfffe,
	})
	checkCycles(t, gb, Cycles{8, 52})
}

func TestRestart(t *testing.T) {
	gb := runCode([]byte{
		0x01, 0xff, 0xfe, // LD BC, 0xfeff
		0xcf,             // RST 0x08
		0x01, 0x34, 0x12, // LD BC, 0x1234 (dummy, should be skipped)
		0xc5, // PUSH BC (dummy, should be skipped)
		0xd1, // POP DE
	})
	checkReg(t, gb, map[RegID]uint16{
		RegBC: 0xfeff,
		RegDE: 0x0004,
		RegSP: 0xfffe,
	})
	checkCycles(t, gb, Cycles{5, 40})
}

func TestAbsoluteJump(t *testing.T) {
	gb := runCode([]byte{
		0x01, 0xff, 0xfe, // LD BC, 0xfeff
		0xd2, 0x09, 0x00, // JP NC, 0x0009
		0x01, 0x34, 0x12, // LD BC, 0x1234 (dummy, should be skipped)
		0xca, 0xff, 0xff, // JP Z, 0xffff (should be skipped)
	})
	checkReg(t, gb, map[RegID]uint16{
		RegBC: 0xfeff,
	})
	checkCycles(t, gb, Cycles{9, 40})
}

func TestJumpHL(t *testing.T) {
	gb := runCode([]byte{
		0x01, 0xff, 0xfe, // LD BC, 0xfeff
		0x21, 0x0a, 0x00, // LD HL, 0xfeff
		0xe9,             // JP (HL)
		0x01, 0x34, 0x12, // LD BC, 0x1234 (dummy, should be skipped)
		0x00, // NOP
	})
	checkReg(t, gb, map[RegID]uint16{
		RegBC: 0xfeff,
	})
	checkCycles(t, gb, Cycles{8, 32})
}

func TestRelativeJump(t *testing.T) {
	gb := runCode([]byte{
		0x01, 0xff, 0xfe, // LD BC, 0xfeff
		0x30, 0x04, // JR NC, 0x0009
		0x01, 0x34, 0x12, // LD BC, 0x1234 (dummy, should be skipped)
		0x38, 0xff, // JR Z, 0xff (should be skipped)
	})
	checkReg(t, gb, map[RegID]uint16{
		RegBC: 0xfeff,
	})
	checkCycles(t, gb, Cycles{7, 32})
}

func TestAdd8(t *testing.T) {
	gb := runCode([]byte{
		0x3e, 0xfa, // LD A, 0xfa
		0x01, 0x0f, 0x0f, // LD BC, 0x0f0f
		0x80, // ADD A, B
		0x6f, // LD L, A
		0x89, // ADC A, C
	})
	checkReg(t, gb, map[RegID]uint16{
		RegAF: 0x1900,
		RegL:  0x09,
	})
	checkCycles(t, gb, Cycles{8, 32})
}

func TestSub8(t *testing.T) {
	gb := runCode([]byte{
		0x3e, 0xfa, // LD A, 0xfa
		0x01, 0x0f, 0x0f, // LD BC, 0x0f0f
		0x90, // SUB A, B
		0x6f, // LD L, A
		0x99, // SBC A, C
	})
	checkReg(t, gb, map[RegID]uint16{
		RegAF: 0xdc60,
		RegL:  0xeb,
	})
	checkCycles(t, gb, Cycles{8, 32})
}

func TestAddHL(t *testing.T) {
	gb := runCode([]byte{
		0x21, 0xfa, 0xff, // LD HL, 0xfffa
		0x01, 0x0f, 0x00, // LD BC, 0x000f
		0x09, // ADD HL, BC
	})
	checkReg(t, gb, map[RegID]uint16{
		RegAF: 0x0010,
		RegHL: 0x0009,
	})
	checkCycles(t, gb, Cycles{7, 32})
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
			Entrypoint: 0,
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
