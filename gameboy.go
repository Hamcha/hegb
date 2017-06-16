package hegb

import "fmt"
import "os"

// Gameboy is an emulated Game boy
type Gameboy struct {
	mmu     *MMU
	cpu     *CPU
	gpu     *GPU
	options EmulatorOptions
}

// EmulatorOptions specifies extra options for changing how the Game boy emulator runs
type EmulatorOptions struct {
	UseBootstrap bool
	Test         bool
	DumpCode     bool
}

// MakeGB creates a Game Boy and loads the rom in it
func MakeGB(romdata *ROM, options EmulatorOptions) *Gameboy {
	gpu := &GPU{}
	mmu := &MMU{
		rom:          romdata,
		gpu:          gpu,
		WRAMExtra:    []WRAM{{}},
		UseBootstrap: options.UseBootstrap,
	}
	cpu := &CPU{
		GPU:      gpu,
		MMU:      mmu,
		Test:     options.Test,
		DumpCode: options.DumpCode,
	}
	mmu.cpu = cpu

	// If bootstrap is skipped, skip to entrypoint
	if !options.UseBootstrap {
		cpu.PC = Register(romdata.Header.Entrypoint)
	}

	return &Gameboy{mmu, cpu, gpu, options}
}

// Run starts up the emulated game boy and blocks until execution ends
func (g *Gameboy) Run() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintln(os.Stderr, "CPU panicked, dump and error message follows:\n")
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
