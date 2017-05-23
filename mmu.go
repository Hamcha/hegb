package hegb

type wram [4 * 1024]byte

type zram [128]byte

// MMU emulates the Game boy memory
type MMU struct {
	rom *ROM

	wram
	wramExtra []wram
	wramID    uint8

	zram zram

	gpu *GPU
}
