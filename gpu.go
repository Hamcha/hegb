package hegb

type vram [8 * 1024]byte

type GPU struct {
	vram   [2]vram // 1 on GB, 2 on GBC
	vramID uint8
}
