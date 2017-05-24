package hegb

// Gameboy is an emulated Game boy
type Gameboy struct {
	mmu     *MMU
	cpu     *CPU
	gpu     *GPU
	options EmulatorOptions
}

type EmulatorOptions struct {
	SkipBootstrap bool
}

// MakeGB creates a Game Boy and loads the rom in it
func MakeGB(romdata *ROM, options EmulatorOptions) *Gameboy {
	gpu := &GPU{}
	mmu := &MMU{
		rom: romdata,
		gpu: gpu,
	}
	cpu := &CPU{
		GPU: gpu,
		MMU: mmu,
	}
	mmu.cpu = cpu

	return &Gameboy{mmu, cpu, gpu, options}
}

// Run starts up the emulated game boy and blocks until execution ends
func (g *Gameboy) Run() {
	g.cpu.Run()
}
