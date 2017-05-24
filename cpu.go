package hegb

// CPU is an emulator of the Z80 CPU used in the Game Boy
type CPU struct {
	Running bool

	MMU *MMU
	GPU *GPU

	// Registers
	AF Register
	BC Register
	DE Register
	HL Register

	PS uint16 // Stack pointer
	PC uint16 // Program counter

	Cycles
}

// Step executes a single CPU instruction
func (c *CPU) Step() {
	//TODO
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
type Register struct {
	Pair uint16
}

// Left returns the left (MSB) byte
func (r Register) Left() uint8 {
	return uint8(r.Pair >> 8)
}

// Right returns the right (LSB) byte
func (r Register) Right() uint8 {
	return uint8(r.Pair)
}

// SetLeft overwrites the left (MSB) byte
func (r *Register) SetLeft(val uint8) {
	r.Pair = (uint16(val) << 8) | (r.Pair & 0xff)
}

// SetRight overwrites the right (LSB) byte
func (r *Register) SetRight(val uint8) {
	r.Pair = (r.Pair & 0xff00) | uint16(val)
}

// Cycles represents the number of cycles an operation took
type Cycles struct {
	CPU     int
	Machine int
}

// Add adds a number of cycles to the current counters
func (c *Cycles) Add(inc Cycles) {
	c.CPU += inc.CPU
	c.Machine += inc.Machine
}
