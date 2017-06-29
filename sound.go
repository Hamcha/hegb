package hegb

type Sound struct {
	SoundEnable bool

	PlayLeft    bool
	PlayRight   bool
	VolumeLeft  uint8 // 0-7
	VolumeRight uint8 // 0-7

	ChToneSweep soundChannel
	ChTone      soundChannel
	ChWave      soundChannel
	ChNoise     soundChannel
}

type soundChannel struct {
	// Global
	Enable        bool
	Frequency     uint16 // Only 11 bits are used
	CounterConsec bool
	Restart       bool
	OutputLeft    bool
	OutputRight   bool
	Envelope      envelope // Wave channel does not have this!

	// Tone+Sweep specific
	SweepTime      sweepTime
	SweepDirection sweepDirection
	SweepShift     uint8 // 0-7

	// Tone channels specific
	ToneLength uint8 // 0-63
	ToneDuty   toneDuty

	// Wave channel specific
	WaveEnable      bool
	WaveLength      uint8
	WaveOutputLevel waveOutputLevel
	WavePattern     [0x10]byte

	// Noise channel specific
	NoiseLength        uint8 // 0-63
	NoisePolyCounter   uint8
	NoiseCounterConsec uint8
}

type sweepTime uint8

const (
	stOff sweepTime = 0 // Off
	st078 sweepTime = 1 // 7.8ms (1/128Hz)
	st156 sweepTime = 2 // 15.6ms (2/128Hz)
	st234 sweepTime = 3 // 23.4ms (3/128Hz)
	st313 sweepTime = 4 // 31.3ms (4/128Hz)
	st391 sweepTime = 5 // 39.1ms (5/128Hz)
	st469 sweepTime = 6 // 46.9ms (6/128Hz)
	st547 sweepTime = 7 // 54.7ms (7/128Hz)
)

type toneDuty uint8

const (
	tone12 toneDuty = 0 // 12.5%
	tone25 toneDuty = 1 // 25%
	tone50 toneDuty = 2 // 50%
	tone75 toneDuty = 3 // 75%
)

type envelope struct {
	InitialVolume uint8 // 0-16 (0 = No sound)
	Direction     sweepDirection
	Sweep         sweepTime
}

type sweepDirection uint8

const (
	swpDecrease sweepDirection = 0
	swpIncrease sweepDirection = 1
)

type waveOutputLevel uint8

const (
	wolMute waveOutputLevel = 0 // Mute
	wol100  waveOutputLevel = 1 // 100% Volume
	wol50   waveOutputLevel = 2 // 50% Volume
	wol25   waveOutputLevel = 3 // 25% Volume
)

type channelType uint8

const (
	sndchToneSweep channelType = iota
	sndchTone
	sndchWave
	sndchNoise
)

// MMU IO functions

func soundEnableRead(c *CPU) (out uint8) {
	if c.ChToneSweep.Enable {
		out |= 0x01
	}
	if c.ChTone.Enable {
		out |= 0x02
	}
	if c.ChWave.Enable {
		out |= 0x04
	}
	if c.ChNoise.Enable {
		out |= 0x08
	}
	if c.SoundEnable {
		out |= 0x80
	}
	return
}

func soundEnableWrite(c *CPU, val uint8) {
	c.ChToneSweep.Enable = val&0x01 == 0x01
	c.ChTone.Enable = val&0x02 == 0x02
	c.ChWave.Enable = val&0x04 == 0x04
	c.ChNoise.Enable = val&0x08 == 0x08
	c.SoundEnable = val&0x80 == 0x80
}

func soundLengthRead(ch channelType) IOReadHandler {
	return func(c *CPU) (out uint8) {
		switch ch {
		case sndchToneSweep:
			out |= uint8(c.ChToneSweep.ToneDuty) << 6
		case sndchTone:
			out |= uint8(c.ChTone.ToneDuty) << 6
		case sndchWave:
			out |= c.ChWave.WaveLength
		case sndchNoise:
			out |= c.ChNoise.NoiseLength
		}
		return
	}
}

func soundLengthWrite(ch channelType) IOWriteHandler {
	return func(c *CPU, val uint8) {
		switch ch {
		case sndchToneSweep:
			c.ChToneSweep.ToneDuty = toneDuty((val >> 6) & 0x03)
			c.ChToneSweep.ToneLength = val & 0x3f
		case sndchTone:
			c.ChTone.ToneDuty = toneDuty((val >> 6) & 0x03)
			c.ChTone.ToneLength = val & 0x3f
		case sndchWave:
			c.ChWave.WaveLength = val
		case sndchNoise:
			c.ChNoise.NoiseLength = val & 0x3f
		}
		return
	}
}

func soundSweepRead(c *CPU) (out uint8) {
	out |= c.ChToneSweep.SweepShift
	out |= uint8(c.ChToneSweep.SweepDirection) << 3
	out |= uint8(c.ChToneSweep.SweepShift) << 4
	return
}

func soundSweepWrite(c *CPU, val uint8) {
	c.ChToneSweep.SweepShift = val & 0x07
	c.ChToneSweep.SweepDirection = sweepDirection((val >> 3) & 0x1)
	c.ChToneSweep.SweepTime = sweepTime((val >> 4) & 0x07)
}

func soundEnvelopeRead(ch channelType) IOReadHandler {
	return func(c *CPU) (out uint8) {
		sndch := getchannel(c, ch)
		out |= uint8(sndch.Envelope.Sweep)
		out |= uint8(sndch.Envelope.Direction) << 3
		out |= sndch.Envelope.InitialVolume << 4
		return
	}
}

func soundEnvelopeWrite(ch channelType) IOWriteHandler {
	return func(c *CPU, val uint8) {
		sndch := getchannel(c, ch)
		sndch.Envelope.Sweep = sweepTime(val & 0x07)
		sndch.Envelope.Direction = sweepDirection((val >> 3) & 0x1)
		sndch.Envelope.InitialVolume = (val >> 4) & 0xf
	}
}

func soundFreqLowWrite(ch channelType) IOWriteHandler {
	return func(c *CPU, val uint8) {
		sndch := getchannel(c, ch)
		sndch.Frequency = (sndch.Frequency & 0xf0) | uint16(val)
	}
}

func soundFreqHighRead(ch channelType) IOReadHandler {
	return func(c *CPU) (out uint8) {
		sndch := getchannel(c, ch)
		if sndch.CounterConsec {
			out |= 0x20
		}
		return
	}
}

func soundFreqHighWrite(ch channelType) IOWriteHandler {
	return func(c *CPU, val uint8) {
		sndch := getchannel(c, ch)
		sndch.Restart = val&0x40 == 0x40
		sndch.CounterConsec = val&0x20 == 0x20
		sndch.Frequency = (uint16(val&0x7) << 8) | (sndch.Frequency & 0xf)
	}
}

func soundWaveReadByte(byteNum uint8) IOReadHandler {
	return func(c *CPU) uint8 {
		return c.ChWave.WavePattern[byteNum]
	}
}

func soundWaveWriteByte(byteNum uint8) IOWriteHandler {
	return func(c *CPU, val uint8) {
		c.ChWave.WavePattern[byteNum] = val
	}
}

// Util functions

func getchannel(c *CPU, ch channelType) *soundChannel {
	switch ch {
	case sndchToneSweep:
		return &c.ChToneSweep
	case sndchTone:
		return &c.ChTone
	case sndchWave:
		return &c.ChWave
	case sndchNoise:
		return &c.ChNoise
	}
	panic("unreacheable")
}
