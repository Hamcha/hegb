package hegb

import "errors"

// Test framework

func runCode(code []byte) *Gameboy {
	rom := makeTestROM(code)
	gb := MakeGB(rom, EmulatorOptions{
		SkipBootstrap: true,
	})
	gb.Run()
	return gb
}

func makeTestROM(code []byte) *ROM {
	return &ROM{
		Header: ROMHeader{
			Entrypoint: 0x200,
			Title:      "TEST",
			Type:       ROMTypeONLY,
			ROMSize:    ROMSize32K,
			RAMSize:    RAMSizeNONE,
		},
		Controller: &testController{data: code},
	}
}

type testController struct {
	data []byte
}

func (t *testController) Read(addr uint16) (uint8, error) {
	if int(addr) >= len(t.data) {
		return 0, errors.New("out of bound read")
	}
	return uint8(t.data[addr]), nil
}

func (t *testController) Write(addr uint16, data uint8) error {
	if int(addr) >= len(t.data) {
		return errors.New("out of bound write")
	}
	t.data[addr] = data
	return nil
}
