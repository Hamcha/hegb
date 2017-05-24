package hegb

// WRAM is "work" RAM, aka RAM in the system, not cartridge
type WRAM [4 * 1024]byte

// ZRAM is zero page RAM (aka high page RAM, because it sits at FF80)
type ZRAM [128]byte

// MMU manages access to the emulated Game boy memory
type MMU struct {
	WRAM
	WRAMExtra []WRAM
	WRAMID    uint8

	ZRAM

	rom *ROM
	cpu *CPU
	gpu *GPU
}
