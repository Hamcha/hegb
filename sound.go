package hegb

type Sound struct {
	Enable bool

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

	// Tone+Sweep specific
	SweepTime sweepTime

	// Tone channels specific
	ToneLength   uint8 // 0-63
	ToneDuty     toneDuty
	ToneEnvelope envelope

	// Wave channel specific
	WaveEnable      bool
	WaveLength      uint8
	WaveOutputLevel waveOutputLevel
	WavePattern     [0x10]byte

	// Noise channel specific
	NoiseLength        uint8 // 0-63
	NoiseEnvelope      envelope
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
	Direction     envelopeDirection
	Sweep         sweepTime
}

type envelopeDirection uint8

const (
	envDecrease envelopeDirection = 0
	envIncrease envelopeDirection = 1
)

type waveOutputLevel uint8

const (
	wolMute waveOutputLevel = 0 // Mute
	wol100  waveOutputLevel = 1 // 100% Volume
	wol50   waveOutputLevel = 2 // 50% Volume
	wol25   waveOutputLevel = 3 // 25% Volume
)
