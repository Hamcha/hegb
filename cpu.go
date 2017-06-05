package hegb

import (
	"fmt"
	"os"
)

// CPU is an emulator of the Z80 CPU used in the Game Boy
type CPU struct {
	Running bool
	Test    bool

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
}

// Step executes a single CPU instruction
func (c *CPU) Step() {

	// Read next instruction
	opcode := nextu8(c)

	// Set as operation to execute
	opid := instruction(opcode)

	// Handle CB prefix
	if opid == OpCBPrefix {
		// Read next opcode
		opcode = nextu8(c)
		// Offset table to CB instructions
		opid = OpCbRotateRegBLeftRot + instruction(opcode)
	}

	// Check if the operation is implemented
	fn, ok := cpuhandlers[opid]
	if !ok {
		// If not, panic!
		panic(fmt.Errorf("operation not implemented: %d", opid))
	}
	fn(c)
}

// Run starts the CPU and blocks until the CPU is done (hopefully, never)
func (c *CPU) Run() {
	c.Running = true
	for c.Running {
		//TODO Clock accurate stepping
		c.Step()
	}
}

// Flags describe each flag in the flag register
type Flags struct {
	Carry     bool
	HalfCarry bool
	AddSub    bool
	Zero      bool
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
	CPU     int
	Machine int
}

// Add adds a number of cycles to the current counters
func (c *Cycles) Add(cpu, machine int) {
	c.CPU += cpu
	c.Machine += machine
}

func (c *CPU) Dump() {
	fmt.Fprintf(os.Stderr, "Registers:\n AF %04x | BC %04x | DE %04x\n HL %04x | SP %04x | PC %04x\n", c.AF, c.BC, c.DE, c.HL, c.SP, c.PC)
}
