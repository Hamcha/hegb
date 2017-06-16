package hegb

import (
	"encoding/binary"
	"fmt"
	"os"
	"strings"
)

// CPU is an emulator of the Z80 CPU used in the Game Boy
type CPU struct {
	Running  bool
	Test     bool
	DumpCode bool

	MMU *MMU
	GPU *GPU

	// Registers
	AF Register
	BC Register
	DE Register
	HL Register

	SP Register // Stack pointer
	PC Register // Program counter

	Cycles Cycles

	InterruptEnable bool

	curInstruction instruction
	curOpcodePos   uint16 // Mostly for debug purposes
}

func (c *CPU) decode() {
	// Read next instruction
	opcode := nextu8(c)

	// Set as operation to execute
	c.curInstruction = instruction(opcode)

	// Handle CB prefix
	if c.curInstruction == OpCBPrefix {
		// Read next opcode
		opcode = nextu8(c)
		// Offset table to CB instructions
		c.curInstruction = OpCbRotateRegBLeftRot + instruction(opcode)
	}
}

// Step executes a single CPU instruction
func (c *CPU) Step() {
	// Save next opcode original position
	c.curOpcodePos = uint16(c.PC)

	// Decode instruction
	c.decode()

	// Check if the operation is implemented
	fn, ok := cpuhandlers[c.curInstruction]
	if !ok {
		// If not, panic!
		panic(fmt.Errorf("operation not implemented: [%02X] %s", uint8(c.curInstruction), c.curInstruction))
	}

	if c.DumpCode {
		fmt.Fprintf(os.Stderr, "| %04x | %s |\n", uint16(c.PC)-1, c.printInstruction(c.curInstruction))
	}
	fn(c)
}

// Run starts the CPU and blocks until the CPU is done (hopefully, never)
func (c *CPU) Run() {
	c.Running = true
	c.InterruptEnable = false
	c.SP = 0xfffe
	for c.Running {
		//TODO Clock accurate stepping
		c.Step()
	}
}

// Flags describe each flag in the flag register
type Flags struct {
	Carry     bool
	AddSub    bool
	HalfCarry bool
	Zero      bool
}

func (f Flags) String() string {
	str := []rune("_ _ _ _")
	if f.Zero {
		str[0] = 'Z'
	}
	if f.AddSub {
		str[2] = 'N'
	}
	if f.HalfCarry {
		str[4] = 'H'
	}
	if f.Carry {
		str[6] = 'C'
	}
	return string(str)
}

// Flags return the current flag register in a nice struct
func (c CPU) Flags() Flags {
	flagbyte := c.AF.Right()
	return Flags{
		Carry:     flagbyte&0x10 == 0x10,
		HalfCarry: flagbyte&0x20 == 0x20,
		AddSub:    flagbyte&0x40 == 0x40,
		Zero:      flagbyte&0x80 == 0x80,
	}
}

// SetFlags sets the flags to a specific value
func (c *CPU) SetFlags(newflags Flags) {
	var flagbyte uint8
	if newflags.Carry {
		flagbyte |= 0x10
	}
	if newflags.HalfCarry {
		flagbyte |= 0x20
	}
	if newflags.AddSub {
		flagbyte |= 0x40
	}
	if newflags.Zero {
		flagbyte |= 0x80
	}
	c.AF.SetRight(flagbyte)
}

// Register represents a single register pair (16bit register that can be accessed as two 2 8bit registers)
type Register uint16

// Left returns the left (MSB) byte
func (r Register) Left() uint8 {
	return uint8(r >> 8)
}

// Right returns the right (LSB) byte
func (r Register) Right() uint8 {
	return uint8(r)
}

// SetLeft overwrites the left (MSB) byte
func (r *Register) SetLeft(val uint8) {
	*r = Register((uint16(val) << 8) | (uint16(*r) & 0xff))
}

// SetRight overwrites the right (LSB) byte
func (r *Register) SetRight(val uint8) {
	*r = Register((uint16(*r) & 0xff00) | uint16(val))
}

// Cycles represents the number of cycles an operation took
type Cycles struct {
	Machine int
	CPU     int
}

// Add adds a number of cycles to the current counters
func (c *Cycles) Add(machine, cpu int) {
	c.CPU += cpu
	c.Machine += machine
}

func (c *CPU) Dump() {
	// Print current instruction (re-decode to fix PC position)
	c.PC = Register(c.curOpcodePos)
	c.decode()
	fmt.Fprintf(os.Stderr, "Instruction: %s\n", c.printInstruction(c.curInstruction))
	// Print registers
	fmt.Fprintf(os.Stderr, "  Registers: AF %04x | BC %04x | DE %04x | HL %04x | SP %04x | PC %04x\n", c.AF, c.BC, c.DE, c.HL, c.SP, c.PC)
	// Print flags individually
	fmt.Fprintf(os.Stderr, "      Flags: %s\n", c.Flags())
	fmt.Fprintln(os.Stderr)
}

func (c *CPU) printInstruction(i instruction) string {
	const REGPOS = 20
	const FLAGPOS = 60
	str := i.String()
	// Replace parameters with their actual values
	if strings.Index(str, "d8") > 0 {
		val := c.MMU.Read(uint16(c.PC))
		str = strings.Replace(str, "d8", fmt.Sprintf("$%02x (%d)", val, val), 1)
	}
	if strings.Index(str, "d16") > 0 {
		val := binary.LittleEndian.Uint16([]byte{c.MMU.Read(uint16(c.PC)), c.MMU.Read(uint16(c.PC) + 1)})
		str = strings.Replace(str, "d16", fmt.Sprintf("$%04x (%d)", val, val), 1)
	}
	if strings.Index(str, "r8") > 0 {
		val := int8(c.MMU.Read(uint16(c.PC)))
		str = strings.Replace(str, "r8", fmt.Sprintf("%d", val), 1)
	}
	if strings.Index(str, "a8") > 0 {
		val := c.MMU.Read(uint16(c.PC))
		str = strings.Replace(str, "a8", fmt.Sprintf("$FF%02x", val), 1)
	}
	if strings.Index(str, "a16") > 0 {
		val := binary.LittleEndian.Uint16([]byte{c.MMU.Read(uint16(c.PC)), c.MMU.Read(uint16(c.PC) + 1)})
		str = strings.Replace(str, "a16", fmt.Sprintf("$%04x", val), 1)
	}
	// Add padding
	str += strings.Repeat(" ", REGPOS-len(str)) + "| "
	// Extra print: registers
	regs := i.Registers()
	if len(regs) > 0 {
		for _, reg := range regs {
			switch reg {
			case RegA, RegB, RegC, RegD, RegE, RegH, RegL:
				str += fmt.Sprintf("%s $%02x ", reg, getreg8(c, reg))
			case RegHLInd, RegBCInd, RegDEInd:
				orig := reg.Unref()
				str += fmt.Sprintf("%s $%04x %s $%02x ", orig, uint16(*reg16(c, orig)), reg, getreg8(c, reg))
			case RegCInd:
				str += fmt.Sprintf("C %02x (C) $%02x ", c.BC.Right(), getreg8(c, reg))
			case RegAF, RegBC, RegDE, RegHL:
				str += fmt.Sprintf("%s $%04x ", reg, uint16(*reg16(c, reg)))
			}
		}
	}
	// Add more padding
	str += strings.Repeat(" ", FLAGPOS-len(str)) + "| "
	// Extra print: flags
	str += fmt.Sprintf("%s", c.Flags())

	return str
}
