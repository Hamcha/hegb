package main

// CPU is an emulator of the Z80 CPU used in the Game Boy
type CPU struct {
	// Registers
	AF register
	BC register
	DE register
	HL register

	Running bool
}

type flags struct {
	Carry     bool
	HalfCarry bool
	AddSub    bool
	Zero      bool
}

// flags return the current flag register in a nice struct
func (c CPU) flags() flags {
	flagbyte := c.AF.Right()
	return flags{
		Carry:     flagbyte&0x10 == 0x10,
		HalfCarry: flagbyte&0x20 == 0x20,
		AddSub:    flagbyte&0x40 == 0x40,
		Zero:      flagbyte&0x80 == 0x80,
	}
}

func (c *CPU) setFlags(newflags flags) {
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

// register represents a single register pair (16bit register that can be accessed as two 2 8bit registers)
type register struct {
	Pair uint16
}

// Left returns the left (MSB) byte
func (r register) Left() uint8 {
	return uint8(r.Pair >> 8)
}

// Right returns the right (LSB) byte
func (r register) Right() uint8 {
	return uint8(r.Pair)
}

// SetLeft overwrites the left (MSB) byte
func (r *register) SetLeft(val uint8) {
	r.Pair = (uint16(val) << 8) | (r.Pair & 0xff)
}

// SetRight overwrites the right (LSB) byte
func (r *register) SetRight(val uint8) {
	r.Pair = (r.Pair & 0xff00) | uint16(val)
}
