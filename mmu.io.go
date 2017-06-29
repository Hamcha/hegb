package hegb

// MMU IO Register
type ioregister uint16

// IOReadHandler handles read access to a single IO register
type IOReadHandler func(c *CPU) uint8

// IOWriteHandler handles write access to a single IO register
type IOWriteHandler func(c *CPU, val uint8)

// All MMU IO registers
const (
	MIOJoypad             ioregister = 0xff00 + iota // ff00 Joypad port
	MIOSerialData                                    // ff01 Serial IO data
	MIOSerialControl                                 // ff02 Serial IO control
	_                                                // ff03 <empty>
	MIODivider                                       // ff04 Divider
	MIOTimerCounter                                  // ff05 Timer counter
	MIOTimerModulo                                   // ff06 Timer modulo
	MIOTimerControl                                  // ff07 Timer control
	_                                                // ff08 <empty>
	_                                                // ff09 <empty>
	_                                                // ff0a <empty>
	_                                                // ff0b <empty>
	_                                                // ff0c <empty>
	_                                                // ff0d <empty>
	_                                                // ff0e <empty>
	MIOInterruptFlags                                // ff0f Interrupt flags
	MIOSound1Sweep                                   // ff10 Sweep (Sound mode #1)
	MIOSound1Length                                  // ff11 Sound length / Pattern duty (Sound mode #1)
	MIOSound1Control                                 // ff12 Control (Sound mode #1)
	MIOSound1FreqLow                                 // ff13 Frequency low (Sound mode #1)
	MIOSound1FreqHigh                                // ff14 Frequency high (Sound mode #1)
	_                                                // ff15 <empty>
	MIOSound2Length                                  // ff16 Sound length / Pattern duty (Sound mode #2)
	MIOSound2Control                                 // ff17 Control (Sound mode #2)
	MIOSound2FreqLow                                 // ff18 Frequency low (Sound mode #2)
	MIOSound2FreqHigh                                // ff19 Frequency high (Sound mode #2)
	MIOSound3Control                                 // ff1a Control (Sound mode #3)
	MIOSound3Length                                  // ff1b Sound length (Sound mode #3)
	MIOSound3Level                                   // ff1c Output level (Sound mode #3)
	MIOSound3FreqLow                                 // ff1d Frequency low (Sound mode #3)
	MIOSound3FreqHigh                                // ff1e Frequency high (Sound mode #3)
	_                                                // ff1f <empty>
	MIOSound4Length                                  // ff20 Sound length / Pattern duty (Sound mode #4)
	MIOSound4Control                                 // ff21 Control (Sound mode #4)
	MIOSound4Counter                                 // ff22 Polynomial counter (Sound mode #4)
	MIOSound4FreqHigh                                // ff23 Frequency high (Sound mode #4)
	MIOSoundChanVol                                  // ff24 Channel / Volume control
	MIOSoundTermSelect                               // ff25 Sound output terminal selector
	MIOSoundEnable                                   // ff26 Sound ON/OFF
	_                                                // ff27 <empty>
	_                                                // ff28 <empty>
	_                                                // ff29 <empty>
	_                                                // ff2a <empty>
	_                                                // ff2b <empty>
	_                                                // ff2c <empty>
	_                                                // ff2d <empty>
	_                                                // ff2e <empty>
	_                                                // ff2f <empty>
	MIOSoundWave0                                    // ff30 Wave channel data # 1
	MIOSoundWave1                                    // ff31 Wave channel data # 2
	MIOSoundWave2                                    // ff32 Wave channel data # 3
	MIOSoundWave3                                    // ff33 Wave channel data # 4
	MIOSoundWave4                                    // ff34 Wave channel data # 5
	MIOSoundWave5                                    // ff35 Wave channel data # 6
	MIOSoundWave6                                    // ff36 Wave channel data # 7
	MIOSoundWave7                                    // ff37 Wave channel data # 8
	MIOSoundWave8                                    // ff38 Wave channel data # 9
	MIOSoundWave9                                    // ff39 Wave channel data # 10
	MIOSoundWaveA                                    // ff3a Wave channel data # 11
	MIOSoundWaveB                                    // ff3b Wave channel data # 12
	MIOSoundWaveC                                    // ff3c Wave channel data # 13
	MIOSoundWaveD                                    // ff3d Wave channel data # 14
	MIOSoundWaveE                                    // ff3e Wave channel data # 15
	MIOSoundWaveF                                    // ff3f Wave channel data # 16
	MIOLCDControl                                    // ff40 LCD Control
	MIOLCDStatus                                     // ff41 LCD Status
	MIOBGVerticalScroll                              // ff42 Background vertical scrolling
	MIOBGHorizontalScroll                            // ff43 Background horizontal scrolling
	MIOLCDCurrentScanline                            // ff44 Current scanline
	MIOLCDScanlineCompare                            // ff45 Scanline comparison
	MIODMAControl                                    // ff46 DMA transfer control
	MIOBGPalette                                     // ff47 Background palette
	MIOSpritePalette0                                // ff48 Sprite palette #0
	MIOSpritePalette1                                // ff49 Sprite palette #1
	MIOWindowYPosition                               // ff4a Window Y position
	MIOWindowXPosition                               // ff4b Window X position
	_                                                // ff4c <empty>
	_                                                // ff4d <empty>
	_                                                // ff4e <empty>
	_                                                // ff4f <empty>
	_                                                // ff50 <empty>
	_                                                // ff51 <empty>
	_                                                // ff52 <empty>
	_                                                // ff53 <empty>
	_                                                // ff54 <empty>
	_                                                // ff55 <empty>
	_                                                // ff56 <empty>
	_                                                // ff57 <empty>
	_                                                // ff58 <empty>
	_                                                // ff59 <empty>
	_                                                // ff5a <empty>
	_                                                // ff5b <empty>
	_                                                // ff5c <empty>
	_                                                // ff5d <empty>
	_                                                // ff5e <empty>
	_                                                // ff5f <empty>
	_                                                // ff60 <empty>
	_                                                // ff61 <empty>
	_                                                // ff62 <empty>
	_                                                // ff63 <empty>
	_                                                // ff64 <empty>
	_                                                // ff65 <empty>
	_                                                // ff66 <empty>
	_                                                // ff67 <empty>
	_                                                // ff68 <empty>
	_                                                // ff69 <empty>
	_                                                // ff6a <empty>
	_                                                // ff6b <empty>
	_                                                // ff6c <empty>
	_                                                // ff6d <empty>
	_                                                // ff6e <empty>
	_                                                // ff6f <empty>
	_                                                // ff70 <empty>
	_                                                // ff71 <empty>
	_                                                // ff72 <empty>
	_                                                // ff73 <empty>
	_                                                // ff74 <empty>
	_                                                // ff75 <empty>
	_                                                // ff76 <empty>
	_                                                // ff77 <empty>
	_                                                // ff78 <empty>
	_                                                // ff79 <empty>
	_                                                // ff7a <empty>
	_                                                // ff7b <empty>
	_                                                // ff7c <empty>
	_                                                // ff7d <empty>
	_                                                // ff7e <empty>
	_                                                // ff7f <empty>
)

func (r ioregister) String() string {
	switch r {
	case MIOJoypad:
		return "Joypad port"
	case MIOSerialData:
		return "Serial IO data"
	case MIOSerialControl:
		return "Serial IO control"
	case MIODivider:
		return "Divider"
	case MIOTimerCounter:
		return "Timer counter"
	case MIOTimerModulo:
		return "Timer modulo"
	case MIOTimerControl:
		return "Timer control"
	case MIOInterruptFlags:
		return "Interrupt flags"
	case MIOSound1Sweep:
		return "Sweep (Sound mode #1)"
	case MIOSound1Length:
		return "Sound length / Pattern duty (Sound mode #1)"
	case MIOSound1Control:
		return "Control (Sound mode #1)"
	case MIOSound1FreqLow:
		return "Frequency low (Sound mode #1)"
	case MIOSound1FreqHigh:
		return "Frequency high (Sound mode #1)"
	case MIOSound2Length:
		return "Sound length / Pattern duty (Sound mode #2)"
	case MIOSound2Control:
		return "Control (Sound mode #2)"
	case MIOSound2FreqLow:
		return "Frequency low (Sound mode #2)"
	case MIOSound2FreqHigh:
		return "Frequency high (Sound mode #2)"
	case MIOSound3Control:
		return "Control (Sound mode #3)"
	case MIOSound3Length:
		return "Sound length (Sound mode #3)"
	case MIOSound3Level:
		return "Output level (Sound mode #3)"
	case MIOSound3FreqLow:
		return "Frequency low (Sound mode #3)"
	case MIOSound3FreqHigh:
		return "Frequency high (Sound mode #3)"
	case MIOSound4Length:
		return "Sound length / Pattern duty (Sound mode #4)"
	case MIOSound4Control:
		return "Control (Sound mode #4)"
	case MIOSound4Counter:
		return "Polynomial counter (Sound mode #4)"
	case MIOSound4FreqHigh:
		return "Frequency high (Sound mode #4)"
	case MIOSoundChanVol:
		return "Channel / Volume control"
	case MIOSoundTermSelect:
		return "Sound output terminal selector"
	case MIOSoundEnable:
		return "Sound ON/OFF"
	case MIOSoundWave0:
		return "Wave channel data # 1"
	case MIOSoundWave1:
		return "Wave channel data # 2"
	case MIOSoundWave2:
		return "Wave channel data # 3"
	case MIOSoundWave3:
		return "Wave channel data # 4"
	case MIOSoundWave4:
		return "Wave channel data # 5"
	case MIOSoundWave5:
		return "Wave channel data # 6"
	case MIOSoundWave6:
		return "Wave channel data # 7"
	case MIOSoundWave7:
		return "Wave channel data # 8"
	case MIOSoundWave8:
		return "Wave channel data # 9"
	case MIOSoundWave9:
		return "Wave channel data # 10"
	case MIOSoundWaveA:
		return "Wave channel data # 11"
	case MIOSoundWaveB:
		return "Wave channel data # 12"
	case MIOSoundWaveC:
		return "Wave channel data # 13"
	case MIOSoundWaveD:
		return "Wave channel data # 14"
	case MIOSoundWaveE:
		return "Wave channel data # 15"
	case MIOSoundWaveF:
		return "Wave channel data # 16"
	case MIOLCDControl:
		return "LCD Control"
	case MIOLCDStatus:
		return "LCD Status"
	case MIOBGVerticalScroll:
		return "Background vertical scrolling"
	case MIOBGHorizontalScroll:
		return "Background horizontal scrolling"
	case MIOLCDCurrentScanline:
		return "Current scanline"
	case MIOLCDScanlineCompare:
		return "Scanline comparison"
	case MIODMAControl:
		return "DMA transfer control"
	case MIOBGPalette:
		return "Background palette"
	case MIOSpritePalette0:
		return "Sprite palette #0"
	case MIOSpritePalette1:
		return "Sprite palette #1"
	case MIOWindowYPosition:
		return "Window Y position"
	case MIOWindowXPosition:
		return "Window X position"
	}
	if r < 0xff80 {
		return "<unused IO register>"
	}
	return "<invalid IO register>"
}

var ioreadhandlers = map[ioregister]IOReadHandler{
	MIOInterruptFlags: func(c *CPU) uint8 { return c.interruptFlags() },
	MIOSoundEnable:    soundEnableRead,
	MIOSound1Sweep:    soundSweepRead,
	MIOSound1Length:   soundLengthRead(sndchToneSweep),
	MIOSound1Control:  soundEnvelopeRead(sndchToneSweep),
	MIOSound1FreqLow:  nil,
	MIOSound1FreqHigh: soundFreqHighRead(sndchToneSweep),
	MIOSound2Length:   soundLengthRead(sndchTone),
	MIOSound2Control:  soundEnvelopeRead(sndchTone),
	MIOSound2FreqLow:  nil,
	MIOSound2FreqHigh: soundFreqHighRead(sndchTone),
	MIOSound3Length:   soundLengthRead(sndchWave),
	MIOSound3FreqLow:  nil,
	MIOSound3FreqHigh: soundFreqHighRead(sndchWave),
	MIOSound4Length:   soundLengthRead(sndchNoise),
	MIOSound4Control:  soundEnvelopeRead(sndchNoise),
	MIOSound4FreqHigh: soundFreqHighRead(sndchNoise),
	MIOSoundWave0:     soundWaveReadByte(0),
	MIOSoundWave1:     soundWaveReadByte(0x1),
	MIOSoundWave2:     soundWaveReadByte(0x2),
	MIOSoundWave3:     soundWaveReadByte(0x3),
	MIOSoundWave4:     soundWaveReadByte(0x4),
	MIOSoundWave5:     soundWaveReadByte(0x5),
	MIOSoundWave6:     soundWaveReadByte(0x6),
	MIOSoundWave7:     soundWaveReadByte(0x7),
	MIOSoundWave8:     soundWaveReadByte(0x8),
	MIOSoundWave9:     soundWaveReadByte(0x9),
	MIOSoundWaveA:     soundWaveReadByte(0xa),
	MIOSoundWaveB:     soundWaveReadByte(0xb),
	MIOSoundWaveC:     soundWaveReadByte(0xc),
	MIOSoundWaveD:     soundWaveReadByte(0xd),
	MIOSoundWaveE:     soundWaveReadByte(0xe),
	MIOSoundWaveF:     soundWaveReadByte(0xf),
}

var iowritehandlers = map[ioregister]IOWriteHandler{
	MIOInterruptFlags: func(c *CPU, val uint8) { c.setInterruptFlags(val) },
	MIOSoundEnable:    soundEnableWrite,
	MIOSound1Sweep:    soundSweepWrite,
	MIOSound1Length:   soundLengthWrite(sndchToneSweep),
	MIOSound1Control:  soundEnvelopeWrite(sndchToneSweep),
	MIOSound1FreqHigh: soundFreqHighWrite(sndchToneSweep),
	MIOSound1FreqLow:  soundFreqLowWrite(sndchToneSweep),
	MIOSound2Length:   soundLengthWrite(sndchTone),
	MIOSound2Control:  soundEnvelopeWrite(sndchTone),
	MIOSound2FreqHigh: soundFreqHighWrite(sndchTone),
	MIOSound2FreqLow:  soundFreqLowWrite(sndchTone),
	MIOSound3Length:   soundLengthWrite(sndchWave),
	MIOSound3FreqHigh: soundFreqHighWrite(sndchWave),
	MIOSound3FreqLow:  soundFreqLowWrite(sndchWave),
	MIOSound4Length:   soundLengthWrite(sndchNoise),
	MIOSound4Control:  soundEnvelopeWrite(sndchNoise),
	MIOSound4FreqHigh: soundFreqHighWrite(sndchNoise),
	MIOSoundWave0:     soundWaveWriteByte(0),
	MIOSoundWave1:     soundWaveWriteByte(0x1),
	MIOSoundWave2:     soundWaveWriteByte(0x2),
	MIOSoundWave3:     soundWaveWriteByte(0x3),
	MIOSoundWave4:     soundWaveWriteByte(0x4),
	MIOSoundWave5:     soundWaveWriteByte(0x5),
	MIOSoundWave6:     soundWaveWriteByte(0x6),
	MIOSoundWave7:     soundWaveWriteByte(0x7),
	MIOSoundWave8:     soundWaveWriteByte(0x8),
	MIOSoundWave9:     soundWaveWriteByte(0x9),
	MIOSoundWaveA:     soundWaveWriteByte(0xa),
	MIOSoundWaveB:     soundWaveWriteByte(0xb),
	MIOSoundWaveC:     soundWaveWriteByte(0xc),
	MIOSoundWaveD:     soundWaveWriteByte(0xd),
	MIOSoundWaveE:     soundWaveWriteByte(0xe),
	MIOSoundWaveF:     soundWaveWriteByte(0xf),
}
