package hegb

var bootstrap = []byte{
	0x31, 0xFE, 0xFF, 0xAF, 0x21, 0xFF, 0x9F, 0x32, 0xCB, 0x7C, 0x20, 0xFB, 0x21,
	0x26, 0xFF, 0x0E, 0x11, 0x3E, 0x80, 0x32, 0xE2, 0x0C, 0x3E, 0xF3, 0xE2, 0x32,
	0x3E, 0x77, 0x77, 0x3E, 0xFC, 0xE0, 0x47, 0x11, 0x04, 0x01, 0x21, 0x10, 0x80,
	0x1A, 0xCD, 0x95, 0x00, 0xCD, 0x96, 0x00, 0x13, 0x7B, 0xFE, 0x34, 0x20, 0xF3,
	0x11, 0xD8, 0x00, 0x06, 0x08, 0x1A, 0x13, 0x22, 0x23, 0x05, 0x20, 0xF9, 0x3E,
	0x19, 0xEA, 0x10, 0x99, 0x21, 0x2F, 0x99, 0x0E, 0x0C, 0x3D, 0x28, 0x08, 0x32,
	0x0D, 0x20, 0xF9, 0x2E, 0x0F, 0x18, 0xF3, 0x67, 0x3E, 0x64, 0x57, 0xE0, 0x42,
	0x3E, 0x91, 0xE0, 0x40, 0x04, 0x1E, 0x02, 0x0E, 0x0C, 0xF0, 0x44, 0xFE, 0x90,
	0x20, 0xFA, 0x0D, 0x20, 0xF7, 0x1D, 0x20, 0xF2, 0x0E, 0x13, 0x24, 0x7C, 0x1E,
	0x83, 0xFE, 0x62, 0x28, 0x06, 0x1E, 0xC1, 0xFE, 0x64, 0x20, 0x06, 0x7B, 0xE2,
	0x0C, 0x3E, 0x87, 0xE2, 0xF0, 0x42, 0x90, 0xE0, 0x42, 0x15, 0x20, 0xD2, 0x05,
	0x20, 0x4F, 0x16, 0x20, 0x18, 0xCB, 0x4F, 0x06, 0x04, 0xC5, 0xCB, 0x11, 0x17,
	0xC1, 0xCB, 0x11, 0x17, 0x05, 0x20, 0xF5, 0x22, 0x23, 0x22, 0x23, 0xC9, 0xCE,
	0xED, 0x66, 0x66, 0xCC, 0x0D, 0x00, 0x0B, 0x03, 0x73, 0x00, 0x83, 0x00, 0x0C,
	0x00, 0x0D, 0x00, 0x08, 0x11, 0x1F, 0x88, 0x89, 0x00, 0x0E, 0xDC, 0xCC, 0x6E,
	0xE6, 0xDD, 0xDD, 0xD9, 0x99, 0xBB, 0xBB, 0x67, 0x63, 0x6E, 0x0E, 0xEC, 0xCC,
	0xDD, 0xDC, 0x99, 0x9F, 0xBB, 0xB9, 0x33, 0x3E, 0x3C, 0x42, 0xB9, 0xA5, 0xB9,
	0xA5, 0x42, 0x3C, 0x21, 0x04, 0x01, 0x11, 0xA8, 0x00, 0x1A, 0x13, 0xBE, 0x20,
	0xFE, 0x23, 0x7D, 0xFE, 0x34, 0x20, 0xF5, 0x06, 0x19, 0x78, 0x86, 0x23, 0x05,
	0x20, 0xFB, 0x86, 0x20, 0xFE, 0x3E, 0x01, 0xE0, 0x50,
}

// WRAM is "work" RAM, aka RAM in the system, not cartridge
type WRAM [4 * 1024]byte

// ZRAM is zero page RAM (aka high page RAM, because it sits at FF80)
type ZRAM [128]byte

// MMU manages access to the emulated Game boy memory
type MMU struct {
	WRAM      WRAM
	WRAMExtra []WRAM
	WRAMID    uint8

	UseBootstrap bool

	InterruptMask uint8

	ZRAM ZRAM

	rom *ROM
	gpu *GPU
}

func (m *MMU) Read(addr uint16) uint8 {
	// 0000 - 00ff => Bootstrap, if enabled
	if m.UseBootstrap && addr < 0x100 {
		return bootstrap[addr]
	}
	// 0000 - 7fff => ROM banks
	if addr < 0x8000 {
		dat, err := m.rom.Controller.Read(addr)
		assert("ROM read error", err)
		return dat
	}
	// 8000 - 9fff => VRAM bank (switchable in GBC)
	if addr < 0xa000 {
		return m.gpu.vram[m.gpu.vramID][addr-0x8000]
	}
	// a000 - bfff => External RAM (switchable)
	if addr < 0xc000 {
		dat, err := m.rom.Controller.Read(addr)
		assert("ROM read error", err)
		return dat
	}
	// c000 - cfff => Work RAM fixed bank
	if addr < 0xd000 {
		return m.WRAM[addr-0xc000]
	}
	// d000 - dfff => Switchable Work RAM bank
	if addr < 0xe000 {
		return m.WRAMExtra[m.WRAMID][addr-0xd000]
	}
	// e000 - fdff => Mirror of c000 - ddff
	if addr < 0xfe00 {
		return m.Read(addr - 0x2000)
	}
	// fe00 - fe9f => Sprite attribute table
	if addr < 0xfea0 {
		panic("todo")
	}
	// fea0 - feff => Not usable
	if addr < 0xff00 {
		return 0
	}
	// ff00 - ff7f => I/O Registers
	if addr < 0xff80 {
		panic("todo")
	}
	// ff80 - fffe => High RAM (HRAM)
	if addr < 0xffff {
		return m.ZRAM[addr-0xff80]
	}
	// ffff => Interrupt mask
	return m.InterruptMask
}

func (m *MMU) Write(addr uint16, value uint8) {
	// 0000 - 7fff => ROM banks (usually non writable)
	if addr < 0x8000 {
		err := m.rom.Controller.Write(addr, value)
		assert("ROM write error", err)
		return
	}
	// 8000 - 9fff => VRAM bank (switchable in GBC)
	if addr < 0xa000 {
		m.gpu.vram[m.gpu.vramID][addr-0x8000] = value
		return
	}
	// a000 - bfff => External RAM (switchable)
	if addr < 0xc000 {
		err := m.rom.Controller.Write(addr, value)
		assert("ROM write error", err)
		return
	}
	// c000 - cfff => Work RAM fixed bank
	if addr < 0xd000 {
		m.WRAM[addr-0xc000] = value
		return
	}
	// d000 - dfff => Switchable Work RAM bank
	if addr < 0xe000 {
		m.WRAMExtra[m.WRAMID][addr-0xd000] = value
		return
	}
	// e000 - fdff => Mirror of c000 - ddff (not writable)
	if addr < 0xfe00 {
		return
	}
	// fe00 - fe9f => Sprite attribute table
	if addr < 0xfea0 {
		panic("todo")
	}
	// fea0 - feff => Not usable
	if addr < 0xff00 {
		return
	}
	// ff00 - ff7f => I/O Registers
	if addr < 0xff80 {
		panic("todo")
	}
	// ff80 - fffe => High RAM (HRAM)
	if addr < 0xffff {
		m.ZRAM[addr-0xff80] = value
		return
	}
	// ffff => Interrupt mask
	m.InterruptMask = value
}
