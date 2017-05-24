package hegb

type vram [8 * 1024]byte

// GPU emulates the graphics layer of a Game boy
type GPU struct {
	vram   [2]vram // 1 on GB, 2 on GBC
	vramID uint8
}
