package hegb

import "fmt"

// Gameboy is an emulated Game boy
type Gameboy struct {
	mmu     *MMU
	cpu     *CPU
	gpu     *GPU
	options EmulatorOptions
}

// EmulatorOptions specifies extra options for changing how the Game boy emulator runs
type EmulatorOptions struct {
	SkipBootstrap bool
	Test          bool
}

// MakeGB creates a Game Boy and loads the rom in it
func MakeGB(romdata *ROM, options EmulatorOptions) *Gameboy {
	gpu := &GPU{}
	mmu := &MMU{
		rom:          romdata,
		gpu:          gpu,
		UseBootstrap: !options.SkipBootstrap,
	}
	cpu := &CPU{
		GPU:  gpu,
		MMU:  mmu,
		Test: options.Test,
	}
	mmu.cpu = cpu

	// If bootstrap is skipped, skip to entrypoint
	if options.SkipBootstrap {
		cpu.PC = Register(romdata.Header.Entrypoint)
	}

	return &Gameboy{mmu, cpu, gpu, options}
}

// Run starts up the emulated game boy and blocks until execution ends
func (g *Gameboy) Run() {
	defer func() {
		if r := recover(); r != nil {
			g.dump()
			panic(r)
		}
	}()
	g.cpu.Run()
}

func (g *Gameboy) dump() {
	g.cpu.Dump()
}

func assert(typ string, err error) {
	if err != nil {
		panic(fmt.Errorf("[%s] %s", typ, err))
	}
}
