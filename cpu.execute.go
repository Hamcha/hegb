package hegb

import (
	"encoding/binary"
)

// Handler handles exactly one instruction
type Handler func(c *CPU)

func noop(c *CPU) {
	c.Cycles.Add(1, 4)
}

var cpuhandlers = map[instruction]Handler{
	OpNop:                     noop,
	OpLoadImmediateBC:         loadImmediate16(RegBC),
	OpLoadImmediateDE:         loadImmediate16(RegDE),
	OpLoadImmediateHL:         loadImmediate16(RegHL),
	OpLoadImmediateSP:         loadImmediate16(RegSP),
	OpLoadImmediateA:          loadImmediate8(RegA),
	OpLoadImmediateB:          loadImmediate8(RegB),
	OpLoadImmediateC:          loadImmediate8(RegC),
	OpLoadImmediateD:          loadImmediate8(RegD),
	OpLoadImmediateE:          loadImmediate8(RegE),
	OpLoadImmediateH:          loadImmediate8(RegH),
	OpLoadImmediateL:          loadImmediate8(RegL),
	OpLoadImmediateIndirectHL: loadImmediate8(RegHLInd),
	OpLoadDirectAA:            loadRegister(RegA, RegA),
	OpLoadDirectAB:            loadRegister(RegA, RegB),
	OpLoadDirectAC:            loadRegister(RegA, RegC),
	OpLoadDirectAD:            loadRegister(RegA, RegD),
	OpLoadDirectAE:            loadRegister(RegA, RegE),
	OpLoadDirectAH:            loadRegister(RegA, RegH),
	OpLoadDirectAL:            loadRegister(RegA, RegL),
	OpLoadIndirectAHL:         loadRegister(RegA, RegHLInd),
	OpLoadDirectBA:            loadRegister(RegB, RegA),
	OpLoadDirectBB:            loadRegister(RegB, RegB),
	OpLoadDirectBC:            loadRegister(RegB, RegC),
	OpLoadDirectBD:            loadRegister(RegB, RegD),
	OpLoadDirectBE:            loadRegister(RegB, RegE),
	OpLoadDirectBH:            loadRegister(RegB, RegH),
	OpLoadDirectBL:            loadRegister(RegB, RegL),
	OpLoadIndirectBHL:         loadRegister(RegB, RegHLInd),
	OpLoadDirectCA:            loadRegister(RegC, RegA),
	OpLoadDirectCB:            loadRegister(RegC, RegB),
	OpLoadDirectCC:            loadRegister(RegC, RegC),
	OpLoadDirectCD:            loadRegister(RegC, RegD),
	OpLoadDirectCE:            loadRegister(RegC, RegE),
	OpLoadDirectCH:            loadRegister(RegC, RegH),
	OpLoadDirectCL:            loadRegister(RegC, RegL),
	OpLoadIndirectCHL:         loadRegister(RegC, RegHLInd),
	OpLoadDirectDA:            loadRegister(RegD, RegA),
	OpLoadDirectDB:            loadRegister(RegD, RegB),
	OpLoadDirectDC:            loadRegister(RegD, RegC),
	OpLoadDirectDD:            loadRegister(RegD, RegD),
	OpLoadDirectDE:            loadRegister(RegD, RegE),
	OpLoadDirectDH:            loadRegister(RegD, RegH),
	OpLoadDirectDL:            loadRegister(RegD, RegL),
	OpLoadIndirectDHL:         loadRegister(RegD, RegHLInd),
	OpLoadDirectEA:            loadRegister(RegE, RegA),
	OpLoadDirectEB:            loadRegister(RegE, RegB),
	OpLoadDirectEC:            loadRegister(RegE, RegC),
	OpLoadDirectED:            loadRegister(RegE, RegD),
	OpLoadDirectEE:            loadRegister(RegE, RegE),
	OpLoadDirectEH:            loadRegister(RegE, RegH),
	OpLoadDirectEL:            loadRegister(RegE, RegL),
	OpLoadIndirectEHL:         loadRegister(RegE, RegHLInd),
	OpLoadDirectHA:            loadRegister(RegH, RegA),
	OpLoadDirectHB:            loadRegister(RegH, RegB),
	OpLoadDirectHC:            loadRegister(RegH, RegC),
	OpLoadDirectHD:            loadRegister(RegH, RegD),
	OpLoadDirectHE:            loadRegister(RegH, RegE),
	OpLoadDirectHH:            loadRegister(RegH, RegH),
	OpLoadDirectHL:            loadRegister(RegH, RegL),
	OpLoadIndirectHHL:         loadRegister(RegH, RegHLInd),
	OpLoadDirectLA:            loadRegister(RegL, RegA),
	OpLoadDirectLB:            loadRegister(RegL, RegB),
	OpLoadDirectLC:            loadRegister(RegL, RegC),
	OpLoadDirectLD:            loadRegister(RegL, RegD),
	OpLoadDirectLE:            loadRegister(RegL, RegE),
	OpLoadDirectLH:            loadRegister(RegL, RegH),
	OpLoadDirectLL:            loadRegister(RegL, RegL),
	OpLoadIndirectLHL:         loadRegister(RegL, RegHLInd),
	OpLoadIndirectHLA:         loadRegister(RegHLInd, RegA),
	OpLoadIndirectHLB:         loadRegister(RegHLInd, RegB),
	OpLoadIndirectHLC:         loadRegister(RegHLInd, RegC),
	OpLoadIndirectHLD:         loadRegister(RegHLInd, RegD),
	OpLoadIndirectHLE:         loadRegister(RegHLInd, RegE),
	OpLoadIndirectHLH:         loadRegister(RegHLInd, RegH),
	OpLoadIndirectHLL:         loadRegister(RegHLInd, RegL),
	OpLoadIndirectBCA:         loadRegister(RegBCInd, RegA),
	OpLoadIndirectDEA:         loadRegister(RegDEInd, RegA),
	OpLoadIndirectABC:         loadRegister(RegA, RegBCInd),
	OpLoadIndirectADE:         loadRegister(RegA, RegDEInd),
	OpLoadHighMemCA:           loadRegister(RegCInd, RegA),
	OpLoadHighRegAC:           loadRegister(RegA, RegCInd),
	OpLoadHighAbsA:            storeHighA,
	OpLoadHighRegA:            loadHighA,
	OpLoadDirectSPHL:          loadRegister16(RegSP, RegHL),
	OpStoreMemSP:              storeSP,
	OpStoreMemA:               storeA,
	OpLoadMemA:                loadA,
	OpIncrementBC:             increment16(RegBC),
	OpIncrementDE:             increment16(RegDE),
	OpIncrementHL:             increment16(RegHL),
	OpIncrementSP:             increment16(RegSP),
	OpDecrementBC:             decrement16(RegBC),
	OpDecrementDE:             decrement16(RegDE),
	OpDecrementHL:             decrement16(RegHL),
	OpDecrementSP:             decrement16(RegSP),
	OpIncrementA:              increment8(RegA),
	OpIncrementB:              increment8(RegB),
	OpIncrementC:              increment8(RegC),
	OpIncrementD:              increment8(RegD),
	OpIncrementE:              increment8(RegE),
	OpIncrementH:              increment8(RegH),
	OpIncrementL:              increment8(RegL),
	OpIncrementIndirectHL:     increment8(RegHLInd),
	OpDecrementA:              decrement8(RegA),
	OpDecrementB:              decrement8(RegB),
	OpDecrementC:              decrement8(RegC),
	OpDecrementD:              decrement8(RegD),
	OpDecrementE:              decrement8(RegE),
	OpDecrementH:              decrement8(RegH),
	OpDecrementL:              decrement8(RegL),
	OpDecrementIndirectHL:     decrement8(RegHLInd),
	OpStop:                    halt,
	OpHalt:                    halt,
	OpInvertA:                 invertA,
	OpSetCarry:                setCarry(false),
	OpFlipCarry:               setCarry(true),
	OpPushAF:                  push16(RegAF),
	OpPushBC:                  push16(RegBC),
	OpPushDE:                  push16(RegDE),
	OpPushHL:                  push16(RegHL),
	OpPopAF:                   pop16(RegAF),
	OpPopBC:                   pop16(RegBC),
	OpPopDE:                   pop16(RegDE),
	OpPopHL:                   pop16(RegHL),
	OpAndDirectAA:             andReg(RegA),
	OpAndDirectAB:             andReg(RegB),
	OpAndDirectAC:             andReg(RegC),
	OpAndDirectAD:             andReg(RegD),
	OpAndDirectAE:             andReg(RegE),
	OpAndDirectAH:             andReg(RegH),
	OpAndDirectAL:             andReg(RegL),
	OpAndIndirectAHL:          andReg(RegHLInd),
	OpAndImmediateA:           andImmediate,
	OpOrDirectAA:              orReg(RegA),
	OpOrDirectAB:              orReg(RegB),
	OpOrDirectAC:              orReg(RegC),
	OpOrDirectAD:              orReg(RegD),
	OpOrDirectAE:              orReg(RegE),
	OpOrDirectAH:              orReg(RegH),
	OpOrDirectAL:              orReg(RegL),
	OpOrIndirectAHL:           orReg(RegHLInd),
	OpOrImmediateA:            orImmediate,
	OpXorDirectAA:             xorReg(RegA),
	OpXorDirectAB:             xorReg(RegB),
	OpXorDirectAC:             xorReg(RegC),
	OpXorDirectAD:             xorReg(RegD),
	OpXorDirectAE:             xorReg(RegE),
	OpXorDirectAH:             xorReg(RegH),
	OpXorDirectAL:             xorReg(RegL),
	OpXorIndirectAHL:          xorReg(RegHLInd),
	OpXorImmediateA:           xorImmediate,
	OpRestart00:               restart(0x00),
	OpRestart08:               restart(0x08),
	OpRestart10:               restart(0x10),
	OpRestart18:               restart(0x18),
	OpRestart20:               restart(0x20),
	OpRestart28:               restart(0x28),
	OpRestart30:               restart(0x30),
	OpRestart38:               restart(0x38),
	OpJumpAbsoluteNO:          jumpa16(fNone),
	OpJumpAbsoluteCA:          jumpa16(fCarry),
	OpJumpAbsoluteNC:          jumpa16(fNotCarry),
	OpJumpAbsoluteZE:          jumpa16(fZero),
	OpJumpAbsoluteNZ:          jumpa16(fNotZero),
	OpJumpAbsoluteHL:          jumpHL,
	OpJumpRelativeNO:          jumpr8(fNone),
	OpJumpRelativeCA:          jumpr8(fCarry),
	OpJumpRelativeNC:          jumpr8(fNotCarry),
	OpJumpRelativeZE:          jumpr8(fZero),
	OpJumpRelativeNZ:          jumpr8(fNotZero),
	OpAddDirectAANoCarry:      add8(RegA, RegA, false),
	OpAddDirectABNoCarry:      add8(RegA, RegB, false),
	OpAddDirectACNoCarry:      add8(RegA, RegC, false),
	OpAddDirectADNoCarry:      add8(RegA, RegD, false),
	OpAddDirectAENoCarry:      add8(RegA, RegE, false),
	OpAddDirectAHNoCarry:      add8(RegA, RegH, false),
	OpAddDirectALNoCarry:      add8(RegA, RegL, false),
	OpAddIndirectAHLNoCarry:   add8(RegA, RegHLInd, false),
	OpAddDirectAACarry:        add8(RegA, RegA, true),
	OpAddDirectABCarry:        add8(RegA, RegB, true),
	OpAddDirectACCarry:        add8(RegA, RegC, true),
	OpAddDirectADCarry:        add8(RegA, RegD, true),
	OpAddDirectAECarry:        add8(RegA, RegE, true),
	OpAddDirectAHCarry:        add8(RegA, RegH, true),
	OpAddDirectALCarry:        add8(RegA, RegL, true),
	OpAddIndirectAHLCarry:     add8(RegA, RegHLInd, true),
	OpSubDirectAANoCarry:      sub8(RegA, RegA, false),
	OpSubDirectABNoCarry:      sub8(RegA, RegB, false),
	OpSubDirectACNoCarry:      sub8(RegA, RegC, false),
	OpSubDirectADNoCarry:      sub8(RegA, RegD, false),
	OpSubDirectAENoCarry:      sub8(RegA, RegE, false),
	OpSubDirectAHNoCarry:      sub8(RegA, RegH, false),
	OpSubDirectALNoCarry:      sub8(RegA, RegL, false),
	OpSubIndirectAHLNoCarry:   sub8(RegA, RegHLInd, false),
	OpSubDirectAACarry:        sub8(RegA, RegA, true),
	OpSubDirectABCarry:        sub8(RegA, RegB, true),
	OpSubDirectACCarry:        sub8(RegA, RegC, true),
	OpSubDirectADCarry:        sub8(RegA, RegD, true),
	OpSubDirectAECarry:        sub8(RegA, RegE, true),
	OpSubDirectAHCarry:        sub8(RegA, RegH, true),
	OpSubDirectALCarry:        sub8(RegA, RegL, true),
	OpSubIndirectAHLCarry:     sub8(RegA, RegHLInd, true),
	OpAddDirectHLBC:           add16HL(RegBC),
	OpAddDirectHLDE:           add16HL(RegDE),
	OpAddDirectHLHL:           add16HL(RegHL),
	OpAddDirectHLSP:           add16HL(RegSP),
	OpCbBitDirectA0:           bit(RegA, 0),
	OpCbBitDirectA1:           bit(RegA, 1),
	OpCbBitDirectA2:           bit(RegA, 2),
	OpCbBitDirectA3:           bit(RegA, 3),
	OpCbBitDirectA4:           bit(RegA, 4),
	OpCbBitDirectA5:           bit(RegA, 5),
	OpCbBitDirectA6:           bit(RegA, 6),
	OpCbBitDirectA7:           bit(RegA, 7),
	OpCbBitDirectB0:           bit(RegB, 0),
	OpCbBitDirectB1:           bit(RegB, 1),
	OpCbBitDirectB2:           bit(RegB, 2),
	OpCbBitDirectB3:           bit(RegB, 3),
	OpCbBitDirectB4:           bit(RegB, 4),
	OpCbBitDirectB5:           bit(RegB, 5),
	OpCbBitDirectB6:           bit(RegB, 6),
	OpCbBitDirectB7:           bit(RegB, 7),
	OpCbBitDirectC0:           bit(RegC, 0),
	OpCbBitDirectC1:           bit(RegC, 1),
	OpCbBitDirectC2:           bit(RegC, 2),
	OpCbBitDirectC3:           bit(RegC, 3),
	OpCbBitDirectC4:           bit(RegC, 4),
	OpCbBitDirectC5:           bit(RegC, 5),
	OpCbBitDirectC6:           bit(RegC, 6),
	OpCbBitDirectC7:           bit(RegC, 7),
	OpCbBitDirectD0:           bit(RegD, 0),
	OpCbBitDirectD1:           bit(RegD, 1),
	OpCbBitDirectD2:           bit(RegD, 2),
	OpCbBitDirectD3:           bit(RegD, 3),
	OpCbBitDirectD4:           bit(RegD, 4),
	OpCbBitDirectD5:           bit(RegD, 5),
	OpCbBitDirectD6:           bit(RegD, 6),
	OpCbBitDirectD7:           bit(RegD, 7),
	OpCbBitDirectE0:           bit(RegE, 0),
	OpCbBitDirectE1:           bit(RegE, 1),
	OpCbBitDirectE2:           bit(RegE, 2),
	OpCbBitDirectE3:           bit(RegE, 3),
	OpCbBitDirectE4:           bit(RegE, 4),
	OpCbBitDirectE5:           bit(RegE, 5),
	OpCbBitDirectE6:           bit(RegE, 6),
	OpCbBitDirectE7:           bit(RegE, 7),
	OpCbBitDirectH0:           bit(RegH, 0),
	OpCbBitDirectH1:           bit(RegH, 1),
	OpCbBitDirectH2:           bit(RegH, 2),
	OpCbBitDirectH3:           bit(RegH, 3),
	OpCbBitDirectH4:           bit(RegH, 4),
	OpCbBitDirectH5:           bit(RegH, 5),
	OpCbBitDirectH6:           bit(RegH, 6),
	OpCbBitDirectH7:           bit(RegH, 7),
	OpCbBitDirectL0:           bit(RegL, 0),
	OpCbBitDirectL1:           bit(RegL, 1),
	OpCbBitDirectL2:           bit(RegL, 2),
	OpCbBitDirectL3:           bit(RegL, 3),
	OpCbBitDirectL4:           bit(RegL, 4),
	OpCbBitDirectL5:           bit(RegL, 5),
	OpCbBitDirectL6:           bit(RegL, 6),
	OpCbBitDirectL7:           bit(RegL, 7),
	OpCbBitIndirectHL0:        bit(RegHLInd, 0),
	OpCbBitIndirectHL1:        bit(RegHLInd, 1),
	OpCbBitIndirectHL2:        bit(RegHLInd, 2),
	OpCbBitIndirectHL3:        bit(RegHLInd, 3),
	OpCbBitIndirectHL4:        bit(RegHLInd, 4),
	OpCbBitIndirectHL5:        bit(RegHLInd, 5),
	OpCbBitIndirectHL6:        bit(RegHLInd, 6),
	OpCbBitIndirectHL7:        bit(RegHLInd, 7),
	OpCbSetDirectA0:           bset(RegA, 0),
	OpCbSetDirectA1:           bset(RegA, 1),
	OpCbSetDirectA2:           bset(RegA, 2),
	OpCbSetDirectA3:           bset(RegA, 3),
	OpCbSetDirectA4:           bset(RegA, 4),
	OpCbSetDirectA5:           bset(RegA, 5),
	OpCbSetDirectA6:           bset(RegA, 6),
	OpCbSetDirectA7:           bset(RegA, 7),
	OpCbSetDirectB0:           bset(RegB, 0),
	OpCbSetDirectB1:           bset(RegB, 1),
	OpCbSetDirectB2:           bset(RegB, 2),
	OpCbSetDirectB3:           bset(RegB, 3),
	OpCbSetDirectB4:           bset(RegB, 4),
	OpCbSetDirectB5:           bset(RegB, 5),
	OpCbSetDirectB6:           bset(RegB, 6),
	OpCbSetDirectB7:           bset(RegB, 7),
	OpCbSetDirectC0:           bset(RegC, 0),
	OpCbSetDirectC1:           bset(RegC, 1),
	OpCbSetDirectC2:           bset(RegC, 2),
	OpCbSetDirectC3:           bset(RegC, 3),
	OpCbSetDirectC4:           bset(RegC, 4),
	OpCbSetDirectC5:           bset(RegC, 5),
	OpCbSetDirectC6:           bset(RegC, 6),
	OpCbSetDirectC7:           bset(RegC, 7),
	OpCbSetDirectD0:           bset(RegD, 0),
	OpCbSetDirectD1:           bset(RegD, 1),
	OpCbSetDirectD2:           bset(RegD, 2),
	OpCbSetDirectD3:           bset(RegD, 3),
	OpCbSetDirectD4:           bset(RegD, 4),
	OpCbSetDirectD5:           bset(RegD, 5),
	OpCbSetDirectD6:           bset(RegD, 6),
	OpCbSetDirectD7:           bset(RegD, 7),
	OpCbSetDirectE0:           bset(RegE, 0),
	OpCbSetDirectE1:           bset(RegE, 1),
	OpCbSetDirectE2:           bset(RegE, 2),
	OpCbSetDirectE3:           bset(RegE, 3),
	OpCbSetDirectE4:           bset(RegE, 4),
	OpCbSetDirectE5:           bset(RegE, 5),
	OpCbSetDirectE6:           bset(RegE, 6),
	OpCbSetDirectE7:           bset(RegE, 7),
	OpCbSetDirectH0:           bset(RegH, 0),
	OpCbSetDirectH1:           bset(RegH, 1),
	OpCbSetDirectH2:           bset(RegH, 2),
	OpCbSetDirectH3:           bset(RegH, 3),
	OpCbSetDirectH4:           bset(RegH, 4),
	OpCbSetDirectH5:           bset(RegH, 5),
	OpCbSetDirectH6:           bset(RegH, 6),
	OpCbSetDirectH7:           bset(RegH, 7),
	OpCbSetDirectL0:           bset(RegL, 0),
	OpCbSetDirectL1:           bset(RegL, 1),
	OpCbSetDirectL2:           bset(RegL, 2),
	OpCbSetDirectL3:           bset(RegL, 3),
	OpCbSetDirectL4:           bset(RegL, 4),
	OpCbSetDirectL5:           bset(RegL, 5),
	OpCbSetDirectL6:           bset(RegL, 6),
	OpCbSetDirectL7:           bset(RegL, 7),
	OpCbSetIndirectHL0:        bset(RegHLInd, 0),
	OpCbSetIndirectHL1:        bset(RegHLInd, 1),
	OpCbSetIndirectHL2:        bset(RegHLInd, 2),
	OpCbSetIndirectHL3:        bset(RegHLInd, 3),
	OpCbSetIndirectHL4:        bset(RegHLInd, 4),
	OpCbSetIndirectHL5:        bset(RegHLInd, 5),
	OpCbSetIndirectHL6:        bset(RegHLInd, 6),
	OpCbSetIndirectHL7:        bset(RegHLInd, 7),
	OpCbResetDirectA0:         breset(RegA, 0),
	OpCbResetDirectA1:         breset(RegA, 1),
	OpCbResetDirectA2:         breset(RegA, 2),
	OpCbResetDirectA3:         breset(RegA, 3),
	OpCbResetDirectA4:         breset(RegA, 4),
	OpCbResetDirectA5:         breset(RegA, 5),
	OpCbResetDirectA6:         breset(RegA, 6),
	OpCbResetDirectA7:         breset(RegA, 7),
	OpCbResetDirectB0:         breset(RegB, 0),
	OpCbResetDirectB1:         breset(RegB, 1),
	OpCbResetDirectB2:         breset(RegB, 2),
	OpCbResetDirectB3:         breset(RegB, 3),
	OpCbResetDirectB4:         breset(RegB, 4),
	OpCbResetDirectB5:         breset(RegB, 5),
	OpCbResetDirectB6:         breset(RegB, 6),
	OpCbResetDirectB7:         breset(RegB, 7),
	OpCbResetDirectC0:         breset(RegC, 0),
	OpCbResetDirectC1:         breset(RegC, 1),
	OpCbResetDirectC2:         breset(RegC, 2),
	OpCbResetDirectC3:         breset(RegC, 3),
	OpCbResetDirectC4:         breset(RegC, 4),
	OpCbResetDirectC5:         breset(RegC, 5),
	OpCbResetDirectC6:         breset(RegC, 6),
	OpCbResetDirectC7:         breset(RegC, 7),
	OpCbResetDirectD0:         breset(RegD, 0),
	OpCbResetDirectD1:         breset(RegD, 1),
	OpCbResetDirectD2:         breset(RegD, 2),
	OpCbResetDirectD3:         breset(RegD, 3),
	OpCbResetDirectD4:         breset(RegD, 4),
	OpCbResetDirectD5:         breset(RegD, 5),
	OpCbResetDirectD6:         breset(RegD, 6),
	OpCbResetDirectD7:         breset(RegD, 7),
	OpCbResetDirectE0:         breset(RegE, 0),
	OpCbResetDirectE1:         breset(RegE, 1),
	OpCbResetDirectE2:         breset(RegE, 2),
	OpCbResetDirectE3:         breset(RegE, 3),
	OpCbResetDirectE4:         breset(RegE, 4),
	OpCbResetDirectE5:         breset(RegE, 5),
	OpCbResetDirectE6:         breset(RegE, 6),
	OpCbResetDirectE7:         breset(RegE, 7),
	OpCbResetDirectH0:         breset(RegH, 0),
	OpCbResetDirectH1:         breset(RegH, 1),
	OpCbResetDirectH2:         breset(RegH, 2),
	OpCbResetDirectH3:         breset(RegH, 3),
	OpCbResetDirectH4:         breset(RegH, 4),
	OpCbResetDirectH5:         breset(RegH, 5),
	OpCbResetDirectH6:         breset(RegH, 6),
	OpCbResetDirectH7:         breset(RegH, 7),
	OpCbResetDirectL0:         breset(RegL, 0),
	OpCbResetDirectL1:         breset(RegL, 1),
	OpCbResetDirectL2:         breset(RegL, 2),
	OpCbResetDirectL3:         breset(RegL, 3),
	OpCbResetDirectL4:         breset(RegL, 4),
	OpCbResetDirectL5:         breset(RegL, 5),
	OpCbResetDirectL6:         breset(RegL, 6),
	OpCbResetDirectL7:         breset(RegL, 7),
	OpCbResetIndirectHL0:      breset(RegHLInd, 0),
	OpCbResetIndirectHL1:      breset(RegHLInd, 1),
	OpCbResetIndirectHL2:      breset(RegHLInd, 2),
	OpCbResetIndirectHL3:      breset(RegHLInd, 3),
	OpCbResetIndirectHL4:      breset(RegHLInd, 4),
	OpCbResetIndirectHL5:      breset(RegHLInd, 5),
	OpCbResetIndirectHL6:      breset(RegHLInd, 6),
	OpCbResetIndirectHL7:      breset(RegHLInd, 7),
	OpCbSwapDirectA:           bswap(RegA),
	OpCbSwapDirectB:           bswap(RegB),
	OpCbSwapDirectC:           bswap(RegC),
	OpCbSwapDirectD:           bswap(RegD),
	OpCbSwapDirectE:           bswap(RegE),
	OpCbSwapDirectH:           bswap(RegH),
	OpCbSwapDirectL:           bswap(RegL),
	OpCbSwapIndirectHL:        bswap(RegHLInd),
	OpSetInt:                  setInterrupt(true),
	OpResetInt:                setInterrupt(false),
	OpRotateAccLeftRotate:     rotateReg(RegA, rLeft, rRotate, rAcc),
	OpRotateAccRightRotate:    rotateReg(RegA, rRight, rRotate, rAcc),
	OpRotateAccLeftThrough:    rotateReg(RegA, rLeft, rWithCarry, rAcc),
	OpRotateAccRightThrough:   rotateReg(RegA, rRight, rWithCarry, rAcc),
	OpCbRotateRegALeftRot:     rotateReg(RegA, rLeft, rRotate, rCB),
	OpCbRotateRegBLeftRot:     rotateReg(RegB, rLeft, rRotate, rCB),
	OpCbRotateRegCLeftRot:     rotateReg(RegC, rLeft, rRotate, rCB),
	OpCbRotateRegDLeftRot:     rotateReg(RegD, rLeft, rRotate, rCB),
	OpCbRotateRegELeftRot:     rotateReg(RegE, rLeft, rRotate, rCB),
	OpCbRotateRegHLeftRot:     rotateReg(RegH, rLeft, rRotate, rCB),
	OpCbRotateRegLLeftRot:     rotateReg(RegL, rLeft, rRotate, rCB),
	OpCbRotateIndHLLeftRot:    rotateReg(RegHLInd, rLeft, rRotate, rCB),
	OpCbRotateRegARightRot:    rotateReg(RegA, rRight, rRotate, rCB),
	OpCbRotateRegBRightRot:    rotateReg(RegB, rRight, rRotate, rCB),
	OpCbRotateRegCRightRot:    rotateReg(RegC, rRight, rRotate, rCB),
	OpCbRotateRegDRightRot:    rotateReg(RegD, rRight, rRotate, rCB),
	OpCbRotateRegERightRot:    rotateReg(RegE, rRight, rRotate, rCB),
	OpCbRotateRegHRightRot:    rotateReg(RegH, rRight, rRotate, rCB),
	OpCbRotateRegLRightRot:    rotateReg(RegL, rRight, rRotate, rCB),
	OpCbRotateIndHLRightRot:   rotateReg(RegHLInd, rRight, rRotate, rCB),
	OpCbRotateRegALeftThC:     rotateReg(RegA, rLeft, rWithCarry, rCB),
	OpCbRotateRegBLeftThC:     rotateReg(RegB, rLeft, rWithCarry, rCB),
	OpCbRotateRegCLeftThC:     rotateReg(RegC, rLeft, rWithCarry, rCB),
	OpCbRotateRegDLeftThC:     rotateReg(RegD, rLeft, rWithCarry, rCB),
	OpCbRotateRegELeftThC:     rotateReg(RegE, rLeft, rWithCarry, rCB),
	OpCbRotateRegHLeftThC:     rotateReg(RegH, rLeft, rWithCarry, rCB),
	OpCbRotateRegLLeftThC:     rotateReg(RegL, rLeft, rWithCarry, rCB),
	OpCbRotateIndHLLeftThC:    rotateReg(RegHLInd, rLeft, rWithCarry, rCB),
	OpCbRotateRegARightThC:    rotateReg(RegA, rRight, rWithCarry, rCB),
	OpCbRotateRegBRightThC:    rotateReg(RegB, rRight, rWithCarry, rCB),
	OpCbRotateRegCRightThC:    rotateReg(RegC, rRight, rWithCarry, rCB),
	OpCbRotateRegDRightThC:    rotateReg(RegD, rRight, rWithCarry, rCB),
	OpCbRotateRegERightThC:    rotateReg(RegE, rRight, rWithCarry, rCB),
	OpCbRotateRegHRightThC:    rotateReg(RegH, rRight, rWithCarry, rCB),
	OpCbRotateRegLRightThC:    rotateReg(RegL, rRight, rWithCarry, rCB),
	OpCbRotateIndHLRightThC:   rotateReg(RegHLInd, rRight, rWithCarry, rCB),
	OpCbRotateRegALeftShf:     rotateReg(RegA, rLeft, rShift, rCB),
	OpCbRotateRegBLeftShf:     rotateReg(RegB, rLeft, rShift, rCB),
	OpCbRotateRegCLeftShf:     rotateReg(RegC, rLeft, rShift, rCB),
	OpCbRotateRegDLeftShf:     rotateReg(RegD, rLeft, rShift, rCB),
	OpCbRotateRegELeftShf:     rotateReg(RegE, rLeft, rShift, rCB),
	OpCbRotateRegHLeftShf:     rotateReg(RegH, rLeft, rShift, rCB),
	OpCbRotateRegLLeftShf:     rotateReg(RegL, rLeft, rShift, rCB),
	OpCbRotateIndHLLeftShf:    rotateReg(RegHLInd, rLeft, rShift, rCB),
	OpCbRotateRegARightShf:    rotateReg(RegA, rRight, rShift, rCB),
	OpCbRotateRegBRightShf:    rotateReg(RegB, rRight, rShift, rCB),
	OpCbRotateRegCRightShf:    rotateReg(RegC, rRight, rShift, rCB),
	OpCbRotateRegDRightShf:    rotateReg(RegD, rRight, rShift, rCB),
	OpCbRotateRegERightShf:    rotateReg(RegE, rRight, rShift, rCB),
	OpCbRotateRegHRightShf:    rotateReg(RegH, rRight, rShift, rCB),
	OpCbRotateRegLRightShf:    rotateReg(RegL, rRight, rShift, rCB),
	OpCbRotateIndHLRightShf:   rotateReg(RegHLInd, rRight, rShift, rCB),
	OpCbRotateRegARightRep:    rotateReg(RegA, rRight, rRepeat, rCB),
	OpCbRotateRegBRightRep:    rotateReg(RegB, rRight, rRepeat, rCB),
	OpCbRotateRegCRightRep:    rotateReg(RegC, rRight, rRepeat, rCB),
	OpCbRotateRegDRightRep:    rotateReg(RegD, rRight, rRepeat, rCB),
	OpCbRotateRegERightRep:    rotateReg(RegE, rRight, rRepeat, rCB),
	OpCbRotateRegHRightRep:    rotateReg(RegH, rRight, rRepeat, rCB),
	OpCbRotateRegLRightRep:    rotateReg(RegL, rRight, rRepeat, rCB),
	OpCbRotateIndHLRightRep:   rotateReg(RegHLInd, rRight, rRepeat, rCB),
}

func loadImmediate8(regid RegID) Handler {
	return func(c *CPU) {
		checkind(c, regid)
		setreg8(c, regid, nextu8(c))
		c.Cycles.Add(2, 8)
	}
}

func loadImmediate16(regid RegID) Handler {
	return func(c *CPU) {
		*reg16(c, regid) = Register(nextu16(c))
		c.Cycles.Add(3, 12)
	}
}

func loadRegister(regtgt, regsrc RegID) Handler {
	return func(c *CPU) {
		checkind(c, regtgt)
		checkind(c, regsrc)
		setreg8(c, regtgt, getreg8(c, regsrc))
		c.Cycles.Add(1, 4)
	}
}

func loadRegister16(regtgt, regsrc RegID) Handler {
	return func(c *CPU) {
		checkind(c, regtgt)
		checkind(c, regsrc)
		*reg16(c, regtgt) = *reg16(c, regsrc)
		c.Cycles.Add(1, 8)
	}
}

func storeHighA(c *CPU) {
	c.MMU.Write(0xff00+uint16(nextu8(c)), c.AF.Left())
	c.Cycles.Add(2, 12)
}

func loadHighA(c *CPU) {
	c.AF.SetLeft(c.MMU.Read(0xff00 + uint16(nextu8(c))))
	c.Cycles.Add(2, 12)
}

func storeA(c *CPU) {
	addr := nextu16(c)
	c.MMU.Write(addr, c.AF.Left())
	c.Cycles.Add(3, 16)
}

func loadA(c *CPU) {
	addr := nextu16(c)
	c.AF.SetLeft(c.MMU.Read(addr))
	c.Cycles.Add(3, 16)
}

func storeSP(c *CPU) {
	addr := nextu16(c)
	c.MMU.Write(addr, c.SP.Right())
	c.MMU.Write(addr+1, c.SP.Left())
	c.Cycles.Add(3, 20)
}

func increment16(regid RegID) Handler {
	return func(c *CPU) {
		checkind(c, regid)
		*reg16(c, regid)++
		c.Cycles.Add(1, 8)
	}
}

func decrement16(regid RegID) Handler {
	return func(c *CPU) {
		*reg16(c, regid)--
		c.Cycles.Add(1, 8)
	}
}

func increment8(regid RegID) Handler {
	return func(c *CPU) {
		checkind(c, regid)
		val := getreg8(c, regid)
		hcbit := val & 0xf
		val++
		c.SetFlags(Flags{
			Zero:      val == 0,
			AddSub:    false,
			HalfCarry: hcbit == 0xf,
			Carry:     c.Flags().Carry,
		})
		setreg8(c, regid, val)
		c.Cycles.Add(1, 4)
	}
}

func decrement8(regid RegID) Handler {
	return func(c *CPU) {
		checkind(c, regid)
		val := getreg8(c, regid)
		hcbit := val & 0xf
		val--
		c.SetFlags(Flags{
			Zero:      val == 0,
			AddSub:    true,
			HalfCarry: hcbit < (val & 0xf),
			Carry:     c.Flags().Carry,
		})
		setreg8(c, regid, val)
		c.Cycles.Add(1, 4)
	}
}

func halt(c *CPU) {
	//TODO Handle properly
	if c.Test {
		c.Running = false
	}
}

func restart(offset uint8) Handler {
	return func(c *CPU) {
		push16(RegPC)(c)
		c.PC = Register(offset)
	}
}

func setInterrupt(val bool) Handler {
	return func(c *CPU) {
		c.InterruptEnable = val
		c.Cycles.Add(1, 4)
	}
}

func invertA(c *CPU) {
	c.AF.SetLeft(^c.AF.Left())
	flags := c.Flags()
	flags.AddSub = true
	flags.HalfCarry = true
	c.SetFlags(flags)
	c.Cycles.Add(1, 4)
}

func setCarry(invert bool) Handler {
	return func(c *CPU) {
		flags := c.Flags()
		flags.AddSub = false
		flags.HalfCarry = false
		if invert {
			flags.Carry = !flags.Carry
		} else {
			flags.Carry = true
		}
		c.SetFlags(flags)
		c.Cycles.Add(1, 4)
	}
}

func add8(regop1 RegID, regop2 RegID, useCarry bool) Handler {
	return func(c *CPU) {
		checkind(c, regop2)
		oldval := getreg8(c, regop1)
		added := getreg8(c, regop2)
		carry := uint8(0)
		if useCarry && c.Flags().Carry {
			carry = 1
		}
		newval := oldval + added + carry

		setreg8(c, regop1, newval)
		c.SetFlags(Flags{
			Carry:     oldval > newval,
			Zero:      newval == 0,
			AddSub:    false,
			HalfCarry: (oldval&0xf)+((added+carry)&0xf) >= 0x10,
		})
		c.Cycles.Add(1, 4)
	}
}

func add16HL(regop2 RegID) Handler {
	return func(c *CPU) {
		current := c.HL
		added := *reg16(c, regop2)
		newval := current + added
		c.SetFlags(Flags{
			Zero:      c.Flags().Zero,
			Carry:     newval < current,
			AddSub:    false,
			HalfCarry: ((current>>8)&0xf)+((added>>8)&0xf) >= 0x10,
		})
		c.HL = newval
		c.Cycles.Add(1, 8)
	}
}

func sub8(regop1 RegID, regop2 RegID, useCarry bool) Handler {
	return func(c *CPU) {
		checkind(c, regop2)
		oldval := getreg8(c, regop1)
		subbed := getreg8(c, regop2)
		carry := uint8(0)
		if useCarry && c.Flags().Carry {
			carry = 1
		}
		newval := oldval - subbed - carry

		setreg8(c, regop1, newval)
		c.SetFlags(Flags{
			Carry:     oldval < newval,
			Zero:      newval == 0,
			AddSub:    true,
			HalfCarry: (newval & 0xf) > (oldval & 0xf),
		})
		c.Cycles.Add(1, 4)
	}
}

func xorReg(regid RegID) Handler {
	return func(c *CPU) {
		val := c.AF.Left() ^ getreg8(c, regid)
		c.AF.SetLeft(val)

		c.SetFlags(Flags{
			Zero:      val == 0,
			AddSub:    false,
			HalfCarry: false,
			Carry:     false,
		})
		c.Cycles.Add(1, 4)
	}
}

func xorImmediate(c *CPU) {
	val := c.AF.Left() ^ nextu8(c)
	c.AF.SetLeft(val)

	c.SetFlags(Flags{
		Zero:      val == 0,
		AddSub:    false,
		HalfCarry: false,
		Carry:     false,
	})
	c.Cycles.Add(2, 8)
}

func andReg(regid RegID) Handler {
	return func(c *CPU) {
		val := c.AF.Left() & getreg8(c, regid)
		c.AF.SetLeft(val)

		c.SetFlags(Flags{
			Zero:      val == 0,
			AddSub:    false,
			HalfCarry: true,
			Carry:     false,
		})
		c.Cycles.Add(1, 4)
	}
}

func andImmediate(c *CPU) {
	val := c.AF.Left() & nextu8(c)
	c.AF.SetLeft(val)

	c.SetFlags(Flags{
		Zero:      val == 0,
		AddSub:    false,
		HalfCarry: true,
		Carry:     false,
	})
	c.Cycles.Add(2, 8)
}

func orReg(regid RegID) Handler {
	return func(c *CPU) {
		val := c.AF.Left() | getreg8(c, regid)
		c.AF.SetLeft(val)

		c.SetFlags(Flags{
			Zero:      val == 0,
			AddSub:    false,
			HalfCarry: false,
			Carry:     false,
		})
		c.Cycles.Add(1, 4)
	}
}

func orImmediate(c *CPU) {
	val := c.AF.Left() | nextu8(c)
	c.AF.SetLeft(val)

	c.SetFlags(Flags{
		Zero:      val == 0,
		AddSub:    false,
		HalfCarry: false,
		Carry:     false,
	})
	c.Cycles.Add(2, 8)
}

func bit(regid RegID, bitNum uint8) Handler {
	return func(c *CPU) {
		checkind(c, regid)
		val := (getreg8(c, regid) >> bitNum) & 0x1
		c.SetFlags(Flags{
			Zero:      val == 0,
			AddSub:    false,
			HalfCarry: true,
			Carry:     c.Flags().Carry,
		})
		c.Cycles.Add(2, 8)
	}
}

func bset(regid RegID, bitNum uint8) Handler {
	return func(c *CPU) {
		checkind(c, regid)
		val := getreg8(c, regid)
		val |= 1 << bitNum
		setreg8(c, regid, val)
		c.Cycles.Add(2, 8)
	}
}

func breset(regid RegID, bitNum uint8) Handler {
	return func(c *CPU) {
		checkind(c, regid)
		val := getreg8(c, regid)
		val &= ^(1 << bitNum)
		setreg8(c, regid, val)
		c.Cycles.Add(2, 8)
	}
}

func bswap(regid RegID) Handler {
	return func(c *CPU) {
		checkind(c, regid)
		val := getreg8(c, regid)
		low := val & 0xf
		high := (val >> 4) & 0xf
		newval := high | (low << 4)
		setreg8(c, regid, newval)
		c.SetFlags(Flags{
			Carry:     false,
			AddSub:    false,
			HalfCarry: false,
			Zero:      newval == 0,
		})
		c.Cycles.Add(2, 8)
	}
}

func push16(regid RegID) Handler {
	return func(c *CPU) {
		val := *reg16(c, regid)
		c.MMU.Write(uint16(c.SP)-1, val.Left())
		c.MMU.Write(uint16(c.SP)-2, val.Right())
		c.SP -= 2
		c.Cycles.Add(1, 16)
	}
}

func pop16(regid RegID) Handler {
	return func(c *CPU) {
		val := reg16(c, regid)
		val.SetRight(c.MMU.Read(uint16(c.SP)))
		val.SetLeft(c.MMU.Read(uint16(c.SP) + 1))
		c.SP += 2
		c.Cycles.Add(1, 12)
	}
}

type rDir int

const (
	rLeft rDir = iota
	rRight
)

type rOp int

const (
	rShift rOp = iota
	rRotate
	rWithCarry
	rRepeat
)

type rOpcodeType int

const (
	rAcc rOpcodeType = iota
	rCB
)

func rotateReg(regid RegID, dir rDir, rop rOp, typ rOpcodeType) Handler {
	return func(c *CPU) {
		checkind(c, regid)
		val := getreg8(c, regid)

		// Save lost bit
		lost := uint8(0)

		// Save last bit (for repeat)
		last := uint8(0)

		// Shift in either direction
		switch dir {
		case rLeft:
			lost = (val >> 7) & 0x1
			last = val & 0x1
			val = val << 1
		case rRight:
			last = val & 0x80
			lost = val & 0x1
			val = val >> 1
		}

		flags := c.Flags()

		// Save old carry flag for WithCarry op
		oldcarry := uint8(0)
		if flags.Carry {
			// If shifting to the right, put as MSB
			if dir == rRight {
				oldcarry = 0x80
			} else {
				oldcarry = 1
			}
		}

		// Set carry to lost bit (true in all forms)
		flags.Carry = lost == 1

		// Operation-specific steps
		switch rop {
		case rRepeat:
			// Repeat last bit being shifted
			val |= last
		case rWithCarry:
			// Apply carry to last bit
			val |= oldcarry
		case rRotate:
			if dir == rRight {
				lost = lost << 7
			}
			val |= lost
		}

		// Set zero flag if using a CB instruction
		if typ == rCB {
			flags.Zero = val == 0
		}

		// Set flags
		c.SetFlags(flags)

		// Set result
		setreg8(c, regid, val)

		if typ == rAcc {
			c.Cycles.Add(1, 4)
		} else {
			c.Cycles.Add(2, 8)
		}
	}
}

type flagID uint8

const (
	fNone flagID = iota
	fCarry
	fZero
	fNotCarry
	fNotZero
)

func jumpa16(flag flagID) Handler {
	return func(c *CPU) {
		flags := c.Flags()
		addr := nextu16(c)
		taken := flag == fNone ||
			(flag == fCarry && flags.Carry) ||
			(flag == fZero && flags.Zero) ||
			(flag == fNotCarry && !flags.Carry) ||
			(flag == fNotZero && !flags.Zero)
		if taken {
			c.PC = Register(addr)
			c.Cycles.Add(0, 4)
		}
		c.Cycles.Add(3, 12)
	}
}

func jumpHL(c *CPU) {
	c.PC = c.HL
	c.Cycles.Add(1, 4)
}

func jumpr8(flag flagID) Handler {
	return func(c *CPU) {
		flags := c.Flags()
		addr := nextu8(c)
		taken := flag == fNone ||
			(flag == fCarry && flags.Carry) ||
			(flag == fZero && flags.Zero) ||
			(flag == fNotCarry && !flags.Carry) ||
			(flag == fNotZero && !flags.Zero)
		if taken {
			c.PC = Register(int32(c.PC) + int32(addr) - 1)
			c.Cycles.Add(0, 4)
		}
		c.Cycles.Add(2, 8)
	}
}

func getreg8(c *CPU, id RegID) uint8 {
	switch id {
	case RegA:
		return c.AF.Left()
	case RegF:
		return c.AF.Right()
	case RegB:
		return c.BC.Left()
	case RegC:
		return c.BC.Right()
	case RegD:
		return c.DE.Left()
	case RegE:
		return c.DE.Right()
	case RegH:
		return c.HL.Left()
	case RegL:
		return c.HL.Right()
	case RegBCInd:
		return c.MMU.Read(uint16(c.BC))
	case RegDEInd:
		return c.MMU.Read(uint16(c.DE))
	case RegHLInd:
		return c.MMU.Read(uint16(c.HL))
	case RegCInd:
		return c.MMU.Read(0xff00 + uint16(c.BC.Right()))
	}
	panic("invalid RegID provided to getreg8")
}

func setreg8(c *CPU, id RegID, val uint8) {
	switch id {
	case RegA:
		c.AF.SetLeft(val)
	case RegF:
		c.AF.SetRight(val)
	case RegB:
		c.BC.SetLeft(val)
	case RegC:
		c.BC.SetRight(val)
	case RegD:
		c.DE.SetLeft(val)
	case RegE:
		c.DE.SetRight(val)
	case RegH:
		c.HL.SetLeft(val)
	case RegL:
		c.HL.SetRight(val)
	case RegBCInd:
		c.MMU.Write(uint16(c.BC), val)
	case RegDEInd:
		c.MMU.Write(uint16(c.DE), val)
	case RegHLInd:
		c.MMU.Write(uint16(c.HL), val)
	case RegCInd:
		c.MMU.Write(0xff00+uint16(c.BC.Right()), val)
	default:
		panic("invalid RegID provided to setreg8")
	}
}

func reg16(c *CPU, id RegID) *Register {
	switch id {
	case RegAF:
		return &c.AF
	case RegBC:
		return &c.BC
	case RegDE:
		return &c.DE
	case RegHL:
		return &c.HL
	case RegSP:
		return &c.SP
	case RegPC:
		return &c.PC
	}
	panic("invalid RegID provided to reg16")
}

// Read uint16 from memory
func nextu16(c *CPU) uint16 {
	val := binary.LittleEndian.Uint16([]byte{c.MMU.Read(uint16(c.PC)), c.MMU.Read(uint16(c.PC) + 1)})
	c.PC += 2
	return val
}

// Read uint8 from memory
func nextu8(c *CPU) uint8 {
	val := c.MMU.Read(uint16(c.PC))
	c.PC++
	return val
}

// Add extra cycles if using the indirect register (for memory fetching)
func checkind(c *CPU, regid RegID) {
	switch regid {
	case RegHLInd, RegBCInd, RegDEInd:
		c.Cycles.Add(0, 4)
	case RegCInd:
		c.Cycles.Add(1, 4)
	}
}
