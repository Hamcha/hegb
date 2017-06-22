package hegb

import "fmt"
import "os"

// Gameboy is an emulated Game boy
type Gameboy struct {
	cpu     *CPU
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
	cpu := &CPU{
		rom: romdata,

		WRAMExtra: []WRAM{{}},

		Test:         options.Test,
		DumpCode:     options.DumpCode,
		UseBootstrap: options.UseBootstrap,
	}

	// If bootstrap is skipped, skip to entrypoint
	if !options.UseBootstrap {
		cpu.PC = Register(romdata.Header.Entrypoint)
	}

	return &Gameboy{cpu, options}
}

// Run starts up the emulated game boy and blocks until execution ends
func (g *Gameboy) Run() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprint(os.Stderr, "CPU panicked, dump and error message follows:\n\n")
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
