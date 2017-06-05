package hegb

import (
	"encoding/binary"
)

// Instruction ID type
type instruction uint16

// All instructions
const (
	// Z80 instructions
	OpNop                      instruction = iota // 00 NOP
	OpLoadImmediateBC                             // 01 LD  BC,d16
	OpLoadIndirectBCA                             // 02 LD  (BC),A
	OpIncrementBC                                 // 03 INC BC
	OpIncrementB                                  // 04 INC B
	OpDecrementB                                  // 05 DEC B
	OpLoadImmediateB                              // 06 LD  B,d8
	OpRotateAccLeftDrop                           // 07 RLCA
	OpStoreMemSP                                  // 08 LD  (a16),SP
	OpAddDirectHLBC                               // 09 ADD HL,BC
	OpLoadIndirectABC                             // 0a LD  A,(BC)
	OpDecrementBC                                 // 0b DEC BC
	OpIncrementC                                  // 0c INC C
	OpDecrementC                                  // 0d DEC C
	OpLoadImmediateC                              // 0e LD  C,d8
	OpRotateAccRightDrop                          // 0f RRCA
	OpStop                                        // 10 STOP
	OpLoadImmediateDE                             // 11 LD  DE,d16
	OpLoadIndirectDEA                             // 12 LD  (DE,A
	OpIncrementDE                                 // 13 INC DE
	OpIncrementD                                  // 14 INC D
	OpDecrementD                                  // 15 DEC D
	OpLoadImmediateD                              // 16 LD  D,d8
	OpRotateAccLeftThrough                        // 17 RLA
	OpJumpRelativeNO                              // 18 JR  r8
	OpAddDirectHLDE                               // 19 ADD HL,DE
	OpLoadIndirectADE                             // 1a LD  A,(DE
	OpDecrementDE                                 // 1b DEC DE
	OpIncrementE                                  // 1c INC E
	OpDecrementE                                  // 1d DEC E
	OpLoadImmediateE                              // 1e LD  E,d8
	OpRotateAccRightThrough                       // 1f RRA
	OpJumpRelativeNZ                              // 20 JR  NZ,r8
	OpLoadImmediateHL                             // 21 LD  HL,d16
	OpLoadIndirectHLAIncrement                    // 22 LDI (HL),A
	OpIncrementHL                                 // 23 INC HL
	OpIncrementH                                  // 24 INC H
	OpDecrementH                                  // 25 DEC H
	OpLoadImmediateH                              // 26 LD  H,d8
	OpDecimalToBCD                                // 27 DAA
	OpJumpRelativeZE                              // 28 JR  Z,r8
	OpAddDirectHLHL                               // 29 ADD HL,HL
	OpLoadIndirectAHLIncrement                    // 2a LDI A,(HL)
	OpDecrementHL                                 // 2b DEC HL
	OpIncrementL                                  // 2c INC L
	OpDecrementL                                  // 2d DEC L
	OpLoadImmediateL                              // 2e LD  L,d8
	OpInvertA                                     // 2f CPL
	OpJumpRelativeNC                              // 30 JR  NC,r8
	OpLoadImmediateSP                             // 31 LD  SP,d16
	OpLoadIndirectHLADecrement                    // 32 LDD (HL),A
	OpIncrementSP                                 // 33 INC SP
	OpIncrementIndirectHL                         // 34 INC (HL)
	OpDecrementIndirectHL                         // 35 DEC (HL)
	OpLoadImmediateIndirectHL                     // 36 LD  (HL),d8
	OpResetCarry                                  // 37 SCF
	OpJumpRelativeCA                              // 38 JR  C,r8
	OpAddDirectHLSP                               // 39 ADD HL,SP
	OpLoadIndirectAHLDecrement                    // 3a LDD A,(HL)
	OpDecrementSP                                 // 3b DEC SP
	OpIncrementA                                  // 3c INC A
	OpDecrementA                                  // 3d DEC A
	OpLoadImmediateA                              // 3e LD  A,d8
	OpSetCarry                                    // 3f CCF
	OpLoadDirectBB                                // 40 LD B,B
	OpLoadDirectBC                                // 41 LD B,C
	OpLoadDirectBD                                // 42 LD B,D
	OpLoadDirectBE                                // 43 LD B,E
	OpLoadDirectBH                                // 44 LD B,H
	OpLoadDirectBL                                // 45 LD B,L
	OpLoadIndirectBHL                             // 46 LD B,(HL)
	OpLoadDirectBA                                // 47 LD B,A
	OpLoadDirectCB                                // 48 LD C,B
	OpLoadDirectCC                                // 49 LD C,C
	OpLoadDirectCD                                // 4a LD C,D
	OpLoadDirectCE                                // 4b LD C,E
	OpLoadDirectCH                                // 4c LD C,H
	OpLoadDirectCL                                // 4d LD C,L
	OpLoadIndirectCHL                             // 4e LD C,(HL)
	OpLoadDirectCA                                // 4f LD C,A
	OpLoadDirectDB                                // 50 LD D,B
	OpLoadDirectDC                                // 51 LD D,C
	OpLoadDirectDD                                // 52 LD D,D
	OpLoadDirectDE                                // 53 LD D,E
	OpLoadDirectDH                                // 54 LD D,H
	OpLoadDirectDL                                // 55 LD D,L
	OpLoadIndirectDHL                             // 56 LD D,(HL)
	OpLoadDirectDA                                // 57 LD D,A
	OpLoadDirectEB                                // 58 LD E,B
	OpLoadDirectEC                                // 59 LD E,C
	OpLoadDirectED                                // 5a LD E,D
	OpLoadDirectEE                                // 5b LD E,E
	OpLoadDirectEH                                // 5c LD E,H
	OpLoadDirectEL                                // 5d LD E,L
	OpLoadIndirectEHL                             // 5e LD E,(HL)
	OpLoadDirectEA                                // 5f LD E,A
	OpLoadDirectHB                                // 60 LD H,B
	OpLoadDirectHC                                // 61 LD H,C
	OpLoadDirectHD                                // 62 LD H,D
	OpLoadDirectHE                                // 63 LD H,E
	OpLoadDirectHH                                // 64 LD H,H
	OpLoadDirectHL                                // 65 LD H,L
	OpLoadIndirectHHL                             // 66 LD H,(HL
	OpLoadDirectHA                                // 67 LD H,A
	OpLoadDirectLB                                // 68 LD L,B
	OpLoadDirectLC                                // 69 LD L,C
	OpLoadDirectLD                                // 6a LD L,D
	OpLoadDirectLE                                // 6b LD L,E
	OpLoadDirectLH                                // 6c LD L,H
	OpLoadDirectLL                                // 6d LD L,L
	OpLoadIndirectLHL                             // 6e LD L,(HL)
	OpLoadDirectLA                                // 6f LD L,A
	OpLoadIndirectHLB                             // 70 LD (HL),B
	OpLoadIndirectHLC                             // 71 LD (HL),C
	OpLoadIndirectHLD                             // 72 LD (HL),D
	OpLoadIndirectHLE                             // 73 LD (HL),E
	OpLoadIndirectHLH                             // 74 LD (HL),H
	OpLoadIndirectHLL                             // 75 LD (HL),L
	OpHalt                                        // 76 HALT
	OpLoadIndirectHLA                             // 77 LD (HL),A
	OpLoadDirectAB                                // 78 LD A,B
	OpLoadDirectAC                                // 79 LD A,C
	OpLoadDirectAD                                // 7a LD A,D
	OpLoadDirectAE                                // 7b LD A,E
	OpLoadDirectAH                                // 7c LD A,H
	OpLoadDirectAL                                // 7d LD A,L
	OpLoadIndirectAHL                             // 7e LD A,(HL)
	OpLoadDirectAA                                // 7f LD A,A
	OpAddDirectABNoCarry                          // 80 ADD A,B
	OpAddDirectACNoCarry                          // 81 ADD A,C
	OpAddDirectADNoCarry                          // 82 ADD A,D
	OpAddDirectAENoCarry                          // 83 ADD A,E
	OpAddDirectAHNoCarry                          // 84 ADD A,H
	OpAddDirectALNoCarry                          // 85 ADD A,L
	OpAddIndirectAHLNoCarry                       // 86 ADD A,(HL)
	OpAddDirectAANoCarry                          // 87 ADD A,A
	OpAddDirectABCarry                            // 88 ADC A,B
	OpAddDirectACCarry                            // 89 ADC A,C
	OpAddDirectADCarry                            // 8a ADC A,D
	OpAddDirectAECarry                            // 8b ADC A,E
	OpAddDirectAHCarry                            // 8c ADC A,H
	OpAddDirectALCarry                            // 8d ADC A,L
	OpAddIndirectAHLCarry                         // 8e ADC A,(HL)
	OpAddDirectAAtrue                             // 8f ADC A,A
	OpSubDirectABNoCarry                          // 90 SUB A,B
	OpSubDirectACNoCarry                          // 91 SUB A,C
	OpSubDirectADNoCarry                          // 92 SUB A,D
	OpSubDirectAENoCarry                          // 93 SUB A,E
	OpSubDirectAHNoCarry                          // 94 SUB A,H
	OpSubDirectALNoCarry                          // 95 SUB A,L
	OpSubIndirectAHLNoCarry                       // 96 SUB A,(HL)
	OpSubDirectAANoCarry                          // 97 SUB A,A
	OpSubDirectABCarry                            // 98 SBC A,B
	OpSubDirectACCarry                            // 99 SBC A,C
	OpSubDirectADCarry                            // 9a SBC A,D
	OpSubDirectAECarry                            // 9b SBC A,E
	OpSubDirectAHCarry                            // 9c SBC A,H
	OpSubDirectALCarry                            // 9d SBC A,L
	OpSubIndirectAHLCarry                         // 9e SBC A,(HL)
	OpSubDirectAACarry                            // 9f SBC A,A
	OpAndDirectAB                                 // a0 AND A,B
	OpAndDirectAC                                 // a1 AND A,C
	OpAndDirectAD                                 // a2 AND A,D
	OpAndDirectAE                                 // a3 AND A,E
	OpAndDirectAH                                 // a4 AND A,H
	OpAndDirectAL                                 // a5 AND A,L
	OpAndIndirectAHL                              // a6 AND A,(HL)
	OpAndDirectAA                                 // a7 AND A,A
	OpXorDirectAB                                 // a8 XOR A,B
	OpXorDirectAC                                 // a9 XOR A,C
	OpXorDirectAD                                 // aa XOR A,D
	OpXorDirectAE                                 // ab XOR A,E
	OpXorDirectAH                                 // ac XOR A,H
	OpXorDirectAL                                 // ad XOR A,L
	OpXorIndirectAHL                              // ae XOR A,(HL)
	OpXorDirectAA                                 // af XOR A,A
	OpOrDirectAB                                  // b0 OR  A,B
	OpOrDirectAC                                  // b1 OR  A,C
	OpOrDirectAD                                  // b2 OR  A,D
	OpOrDirectAE                                  // b3 OR  A,E
	OpOrDirectAH                                  // b4 OR  A,H
	OpOrDirectAL                                  // b5 OR  A,L
	OpOrIndirectAHL                               // b6 OR  A,(HL)
	OpOrDirectAA                                  // b7 OR  A,A
	OpCmpDirectAB                                 // b8 CP  A,B
	OpCmpDirectAC                                 // b9 CP  A,C
	OpCmpDirectAD                                 // ba CP  A,D
	OpCmpDirectAE                                 // bb CP  A,E
	OpCmpDirectAH                                 // bc CP  A,H
	OpCmpDirectAL                                 // bd CP  A,L
	OpCmpIndirectAHL                              // be CP  A,(HL)
	OpCmpDirectAA                                 // bf CP  A,A
	OpReturnNZ                                    // c0 RET NZ
	OpPopBC                                       // c1 POP BC
	OpJumpAbsoluteNZ                              // c2 JP  NZ,a16
	OpJumpAbsoluteNO                              // c3 JP  a16
	OpCallNZ                                      // c4 CALL NZ,a16
	OpPushBC                                      // c5 PUSH BC
	OpAddImmediateANoCarry                        // c6 ADD A,d8
	OpRestart00                                   // c7 RST 00h
	OpReturnZE                                    // c8 RET Z
	OpReturnNO                                    // c9 RET
	OpJumpAbsoluteZE                              // ca JP  Z,a16
	OpCBPrefix                                    // cb PREFIX: See cbhandlers below (01xx)
	OpCallZE                                      // cc CALL Z,a16
	OpCallNO                                      // cd CALL a16
	OpAddImmediateACarry                          // ce ADC A,d8
	OpRestart08                                   // cf RST 08h
	OpReturnNC                                    // d0 RET NC
	OpPopRegDE                                    // d1 POP DE
	OpJumpAbsoluteNC                              // d2 JP  NC,a16
	_                                             // d3 --
	OpCallNC                                      // d4 CALL NC,a16
	OpPushDE                                      // d5 PUSH DE
	OpSubImmediateANoCarry                        // d6 SUB A,d8
	OpRestart10                                   // d7 RST 10h
	OpReturnCA                                    // d8 RET C
	OpRETI                                        // d9 RETI
	OpJumpAbsoluteCA                              // da JP  C,a16
	_                                             // db --
	OpCallCA                                      // dc CALL C,a16
	_                                             // dd --
	OpSubImmediateACarry                          // de SBC A,d8
	OpRestart18                                   // df RST 18h
	OpLoadHighAbsA                                // e0 LDH (a8),A
	OpPopHL                                       // e1 POP HL
	OpLoadHighMemCA                               // e2 LD  (C),A
	_                                             // e3 --
	_                                             // e4 --
	OpPushHL                                      // e5 PUSH HL
	OpAndImmediateA                               // e6 AND A,d8
	OpRestart20                                   // e7 RST 20h
	OpAddImmediateSignedSP                        // e8 ADD SP,r8
	OpJumpAbsoluteHL                              // e9 JP  (HL)
	OpStoreMemA                                   // ea LD  (a16),A
	_                                             // eb --
	_                                             // ec --
	_                                             // ed --
	OpXorImmediateA                               // ee XOR A,d8
	OpRestart28                                   // ef RST 28h
	OpLoadHighRegA                                // f0 LDH A,(a8)
	OpPopAF                                       // f1 POP AF
	OpLoadHighRegAC                               // f2 LD  A,(C)
	OpResetInt                                    // f3 DI
	_                                             // f4 --
	OpPushAF                                      // f5 PUSH AF
	OpOrImmediateA                                // f6 OR  A,d8
	OpRestart30                                   // f7 RES 30h
	OpLoadOffsetHLSP                              // f8 LD  HL,SP+r8
	OpLoadDirectSPHL                              // f9 LD  SP,HL
	OpLoadMemA                                    // fa LD  A,(a16)
	OpSetInt                                      // fb EI
	_                                             // fc --
	_                                             // fd --
	OpCmpImmediateA                               // fe CP  A,d8
	OpRestart38                                   // ff RST 38h

	// CB prefix operations
	OpCbRotateRegBLeftRot   instruction = 0x100 + iota // 00 RLC B
	OpCbRotateRegCLeftRot                              // 01 RLC C
	OpCbRotateRegDLeftRot                              // 02 RLC D
	OpCbRotateRegELeftRot                              // 03 RLC E
	OpCbRotateRegHLeftRot                              // 04 RLC H
	OpCbRotateRegLLeftRot                              // 05 RLC L
	OpCbRotateIndHLLeftRot                             // 06 RLC (HL)
	OpCbRotateRegALeftRot                              // 07 RLC A
	OpCbRotateRegBRightRot                             // 08 RRC B
	OpCbRotateRegCRightRot                             // 09 RRC C
	OpCbRotateRegDRightRot                             // 0a RRC D
	OpCbRotateRegERightRot                             // 0b RRC E
	OpCbRotateRegHRightRot                             // 0c RRC H
	OpCbRotateRegLRightRot                             // 0d RRC L
	OpCbRotateIndHLRightRot                            // 0e RRC (HL)
	OpCbRotateRegARightRot                             // 0f RRC A
	OpCbRotateRegBLeftThC                              // 10 RL  B
	OpCbRotateRegCLeftThC                              // 11 RL  C
	OpCbRotateRegDLeftThC                              // 12 RL  D
	OpCbRotateRegELeftThC                              // 13 RL  E
	OpCbRotateRegHLeftThC                              // 14 RL  H
	OpCbRotateRegLLeftThC                              // 15 RL  L
	OpCbRotateIndHLLeftThC                             // 16 RL  (HL)
	OpCbRotateRegALeftThC                              // 17 RL  A
	OpCbRotateRegBRightThC                             // 18 RR  B
	OpCbRotateRegCRightThC                             // 19 RR  C
	OpCbRotateRegDRightThC                             // 1a RR  D
	OpCbRotateRegERightThC                             // 1b RR  E
	OpCbRotateRegHRightThC                             // 1c RR  H
	OpCbRotateRegLRightThC                             // 1d RR  L
	OpCbRotateIndHLRightThC                            // 1e RR  (HL)
	OpCbRotateRegARightThC                             // 1f RR  A
	OpCbRotateRegBLeftShf                              // 20 SLA B
	OpCbRotateRegCLeftShf                              // 21 SLA C
	OpCbRotateRegDLeftShf                              // 22 SLA D
	OpCbRotateRegELeftShf                              // 23 SLA E
	OpCbRotateRegHLeftShf                              // 24 SLA H
	OpCbRotateRegLLeftShf                              // 25 SLA L
	OpCbRotateIndHLLeftShf                             // 26 SLA (HL)
	OpCbRotateRegALeftShf                              // 27 SLA A
	OpCbRotateRegBRightRep                             // 28 SRA B
	OpCbRotateRegCRightRep                             // 29 SRA C
	OpCbRotateRegDRightRep                             // 2a SRA D
	OpCbRotateRegERightRep                             // 2b SRA E
	OpCbRotateRegHRightRep                             // 2c SRA H
	OpCbRotateRegLRightRep                             // 2d SRA L
	OpCbRotateIndHLRightRep                            // 2e SRA (HL)
	OpCbRotateRegARightRep                             // 2f SRA A
	OpCbSwapDirectB                                    // 30 SWAP B
	OpCbSwapDirectC                                    // 31 SWAP C
	OpCbSwapDirectD                                    // 32 SWAP D
	OpCbSwapDirectE                                    // 33 SWAP E
	OpCbSwapDirectH                                    // 34 SWAP H
	OpCbSwapDirectL                                    // 35 SWAP L
	OpCbSwapIndirectHL                                 // 36 SWAP (HL)
	OpCbSwapDirectA                                    // 37 SWAP A
	OpCbRotateRegBRightShf                             // 38 SRL B
	OpCbRotateRegCRightShf                             // 39 SRL C
	OpCbRotateRegDRightShf                             // 3a SRL D
	OpCbRotateRegERightShf                             // 3b SRL E
	OpCbRotateRegHRightShf                             // 3c SRL H
	OpCbRotateRegLRightShf                             // 3d SRL L
	OpCbRotateIndHLRightShf                            // 3e SRL (HL)
	OpCbRotateRegARightShf                             // 3f SRL A
	OpCbBitDirectB0                                    // 40 BIT 0,B
	OpCbBitDirectC0                                    // 41 BIT 0,C
	OpCbBitDirectD0                                    // 42 BIT 0,D
	OpCbBitDirectE0                                    // 43 BIT 0,E
	OpCbBitDirectH0                                    // 44 BIT 0,H
	OpCbBitDirectL0                                    // 45 BIT 0,L
	OpCbBitIndirectHL0                                 // 46 BIT 0,(HL)
	OpCbBitDirectA0                                    // 47 BIT 0,A
	OpCbBitDirectB1                                    // 48 BIT 1,B
	OpCbBitDirectC1                                    // 49 BIT 1,C
	OpCbBitDirectD1                                    // 4a BIT 1,D
	OpCbBitDirectE1                                    // 4b BIT 1,E
	OpCbBitDirectH1                                    // 4c BIT 1,H
	OpCbBitDirectL1                                    // 4d BIT 1,L
	OpCbBitIndirectHL1                                 // 4e BIT 1,(HL)
	OpCbBitDirectA1                                    // 4f BIT 1,A
	OpCbBitDirectB2                                    // 50 BIT 2,B
	OpCbBitDirectC2                                    // 51 BIT 2,C
	OpCbBitDirectD2                                    // 52 BIT 2,D
	OpCbBitDirectE2                                    // 53 BIT 2,E
	OpCbBitDirectH2                                    // 54 BIT 2,H
	OpCbBitDirectL2                                    // 55 BIT 2,L
	OpCbBitIndirectHL2                                 // 56 BIT 2,(HL)
	OpCbBitDirectA2                                    // 57 BIT 2,A
	OpCbBitDirectB3                                    // 58 BIT 3,B
	OpCbBitDirectC3                                    // 59 BIT 3,C
	OpCbBitDirectD3                                    // 5a BIT 3,D
	OpCbBitDirectE3                                    // 5b BIT 3,E
	OpCbBitDirectH3                                    // 5c BIT 3,H
	OpCbBitDirectL3                                    // 5d BIT 3,L
	OpCbBitIndirectHL3                                 // 5e BIT 3,(HL)
	OpCbBitDirectA3                                    // 5f BIT 3,A
	OpCbBitDirectB4                                    // 60 BIT 4,B
	OpCbBitDirectC4                                    // 61 BIT 4,C
	OpCbBitDirectD4                                    // 62 BIT 4,D
	OpCbBitDirectE4                                    // 63 BIT 4,E
	OpCbBitDirectH4                                    // 64 BIT 4,H
	OpCbBitDirectL4                                    // 65 BIT 4,L
	OpCbBitIndirectHL4                                 // 66 BIT 4,(HL)
	OpCbBitDirectA4                                    // 67 BIT 4,A
	OpCbBitDirectB5                                    // 68 BIT 5,B
	OpCbBitDirectC5                                    // 69 BIT 5,C
	OpCbBitDirectD5                                    // 6a BIT 5,D
	OpCbBitDirectE5                                    // 6b BIT 5,E
	OpCbBitDirectH5                                    // 6c BIT 5,H
	OpCbBitDirectL5                                    // 6d BIT 5,L
	OpCbBitIndirectHL5                                 // 6e BIT 5,(HL)
	OpCbBitDirectA5                                    // 6f BIT 5,A
	OpCbBitDirectB6                                    // 70 BIT 6,B
	OpCbBitDirectC6                                    // 71 BIT 6,C
	OpCbBitDirectD6                                    // 72 BIT 6,D
	OpCbBitDirectE6                                    // 73 BIT 6,E
	OpCbBitDirectH6                                    // 74 BIT 6,H
	OpCbBitDirectL6                                    // 75 BIT 6,L
	OpCbBitIndirectHL6                                 // 76 BIT 6,(HL)
	OpCbBitDirectA6                                    // 77 BIT 6,A
	OpCbBitDirectB7                                    // 78 BIT 7,B
	OpCbBitDirectC7                                    // 79 BIT 7,C
	OpCbBitDirectD7                                    // 7a BIT 7,D
	OpCbBitDirectE7                                    // 7b BIT 7,E
	OpCbBitDirectH7                                    // 7c BIT 7,H
	OpCbBitDirectL7                                    // 7d BIT 7,L
	OpCbBitIndirectHL7                                 // 7e BIT 7,(HL)
	OpCbBitDirectA7                                    // 7f BIT 7,A
	OpCbResetDirectB0                                  // 80 RES 0,B
	OpCbResetDirectC0                                  // 81 RES 0,C
	OpCbResetDirectD0                                  // 82 RES 0,D
	OpCbResetDirectE0                                  // 83 RES 0,E
	OpCbResetDirectH0                                  // 84 RES 0,H
	OpCbResetDirectL0                                  // 85 RES 0,L
	OpCbResetIndirectHL0                               // 86 RES 0,(HL)
	OpCbResetDirectA0                                  // 87 RES 0,A
	OpCbResetDirectB1                                  // 88 RES 1,B
	OpCbResetDirectC1                                  // 89 RES 1,C
	OpCbResetDirectD1                                  // 8a RES 1,D
	OpCbResetDirectE1                                  // 8b RES 1,E
	OpCbResetDirectH1                                  // 8c RES 1,H
	OpCbResetDirectL1                                  // 8d RES 1,L
	OpCbResetIndirectHL1                               // 8e RES 1,(HL)
	OpCbResetDirectA1                                  // 8f RES 1,A
	OpCbResetDirectB2                                  // 90 RES 2,B
	OpCbResetDirectC2                                  // 91 RES 2,C
	OpCbResetDirectD2                                  // 92 RES 2,D
	OpCbResetDirectE2                                  // 93 RES 2,E
	OpCbResetDirectH2                                  // 94 RES 2,H
	OpCbResetDirectL2                                  // 95 RES 2,L
	OpCbResetIndirectHL2                               // 96 RES 2,(HL)
	OpCbResetDirectA2                                  // 97 RES 2,A
	OpCbResetDirectB3                                  // 98 RES 3,B
	OpCbResetDirectC3                                  // 99 RES 3,C
	OpCbResetDirectD3                                  // 9a RES 3,D
	OpCbResetDirectE3                                  // 9b RES 3,E
	OpCbResetDirectH3                                  // 9c RES 3,H
	OpCbResetDirectL3                                  // 9d RES 3,L
	OpCbResetIndirectHL3                               // 9e RES 3,(HL)
	OpCbResetDirectA3                                  // 9f RES 3,A
	OpCbResetDirectB4                                  // a0 RES 4,B
	OpCbResetDirectC4                                  // a1 RES 4,C
	OpCbResetDirectD4                                  // a2 RES 4,D
	OpCbResetDirectE4                                  // a3 RES 4,E
	OpCbResetDirectH4                                  // a4 RES 4,H
	OpCbResetDirectL4                                  // a5 RES 4,L
	OpCbResetIndirectHL4                               // a6 RES 4,(HL)
	OpCbResetDirectA4                                  // a7 RES 4,A
	OpCbResetDirectB5                                  // a8 RES 5,B
	OpCbResetDirectC5                                  // a9 RES 5,C
	OpCbResetDirectD5                                  // aa RES 5,D
	OpCbResetDirectE5                                  // ab RES 5,E
	OpCbResetDirectH5                                  // ac RES 5,H
	OpCbResetDirectL5                                  // ad RES 5,L
	OpCbResetIndirectHL5                               // ae RES 5,(HL)
	OpCbResetDirectA5                                  // af RES 5,A
	OpCbResetDirectB6                                  // b0 RES 6,B
	OpCbResetDirectC6                                  // b1 RES 6,C
	OpCbResetDirectD6                                  // b2 RES 6,D
	OpCbResetDirectE6                                  // b3 RES 6,E
	OpCbResetDirectH6                                  // b4 RES 6,H
	OpCbResetDirectL6                                  // b5 RES 6,L
	OpCbResetIndirectHL6                               // b6 RES 6,(HL)
	OpCbResetDirectA6                                  // b7 RES 6,A
	OpCbResetDirectB7                                  // b8 RES 7,B
	OpCbResetDirectC7                                  // b9 RES 7,C
	OpCbResetDirectD7                                  // ba RES 7,D
	OpCbResetDirectE7                                  // bb RES 7,E
	OpCbResetDirectH7                                  // bc RES 7,H
	OpCbResetDirectL7                                  // bd RES 7,L
	OpCbResetIndirectHL7                               // be RES 7,(HL)
	OpCbResetDirectA7                                  // bf RES 7,A
	OpCbSetDirectB0                                    // c0 SET 0,B
	OpCbSetDirectC0                                    // c1 SET 0,C
	OpCbSetDirectD0                                    // c2 SET 0,D
	OpCbSetDirectE0                                    // c3 SET 0,E
	OpCbSetDirectH0                                    // c4 SET 0,H
	OpCbSetDirectL0                                    // c5 SET 0,L
	OpCbSetIndirectHL0                                 // c6 SET 0,(HL)
	OpCbSetDirectA0                                    // c7 SET 0,A
	OpCbSetDirectB1                                    // c8 SET 1,B
	OpCbSetDirectC1                                    // c9 SET 1,C
	OpCbSetDirectD1                                    // ca SET 1,D
	OpCbSetDirectE1                                    // cb SET 1,E
	OpCbSetDirectH1                                    // cc SET 1,H
	OpCbSetDirectL1                                    // cd SET 1,L
	OpCbSetIndirectHL1                                 // ce SET 1,(HL)
	OpCbSetDirectA1                                    // cf SET 1,A
	OpCbSetDirectB2                                    // d0 SET 2,B
	OpCbSetDirectC2                                    // d1 SET 2,C
	OpCbSetDirectD2                                    // d2 SET 2,D
	OpCbSetDirectE2                                    // d3 SET 2,E
	OpCbSetDirectH2                                    // d4 SET 2,H
	OpCbSetDirectL2                                    // d5 SET 2,L
	OpCbSetIndirectHL2                                 // d6 SET 2,(HL)
	OpCbSetDirectA2                                    // d7 SET 2,A
	OpCbSetDirectB3                                    // d8 SET 3,B
	OpCbSetDirectC3                                    // d9 SET 3,C
	OpCbSetDirectD3                                    // da SET 3,D
	OpCbSetDirectE3                                    // db SET 3,E
	OpCbSetDirectH3                                    // dc SET 3,H
	OpCbSetDirectL3                                    // dd SET 3,L
	OpCbSetIndirectHL3                                 // de SET 3,(HL)
	OpCbSetDirectA3                                    // df SET 3,A
	OpCbSetDirectB4                                    // e0 SET 4,B
	OpCbSetDirectC4                                    // e1 SET 4,C
	OpCbSetDirectD4                                    // e2 SET 4,D
	OpCbSetDirectE4                                    // e3 SET 4,E
	OpCbSetDirectH4                                    // e4 SET 4,H
	OpCbSetDirectL4                                    // e5 SET 4,L
	OpCbSetIndirectHL4                                 // e6 SET 4,(HL)
	OpCbSetDirectA4                                    // e7 SET 4,A
	OpCbSetDirectB5                                    // e8 SET 5,B
	OpCbSetDirectC5                                    // e9 SET 5,C
	OpCbSetDirectD5                                    // ea SET 5,D
	OpCbSetDirectE5                                    // eb SET 5,E
	OpCbSetDirectH5                                    // ec SET 5,H
	OpCbSetDirectL5                                    // ed SET 5,L
	OpCbSetIndirectHL5                                 // ee SET 5,(HL)
	OpCbSetDirectA5                                    // ef SET 5,A
	OpCbSetDirectB6                                    // f0 SET 6,B
	OpCbSetDirectC6                                    // f1 SET 6,C
	OpCbSetDirectD6                                    // f2 SET 6,D
	OpCbSetDirectE6                                    // f3 SET 6,E
	OpCbSetDirectH6                                    // f4 SET 6,H
	OpCbSetDirectL6                                    // f5 SET 6,L
	OpCbSetIndirectHL6                                 // f6 SET 6,(HL)
	OpCbSetDirectA6                                    // f7 SET 6,A
	OpCbSetDirectB7                                    // f8 SET 7,B
	OpCbSetDirectC7                                    // f9 SET 7,C
	OpCbSetDirectD7                                    // fa SET 7,D
	OpCbSetDirectE7                                    // fb SET 7,E
	OpCbSetDirectH7                                    // fc SET 7,H
	OpCbSetDirectL7                                    // fd SET 7,L
	OpCbSetIndirectHL7                                 // fe SET 7,(HL)
	OpCbSetDirectA7                                    // ff SET 7,A
)

// Handler handles exactly one instruction
type Handler func(c *CPU)

func noop(c *CPU) {
	c.Cycles.Add(1, 4)
}

var cpuhandlers = map[instruction]Handler{
	OpNop:             noop,
	OpLoadImmediateBC: loadImmediate16(RegBC),
	OpLoadImmediateDE: loadImmediate16(RegDE),
	OpLoadImmediateHL: loadImmediate16(RegHL),
	OpLoadImmediateSP: loadImmediate16(RegSP),
	OpLoadImmediateA:  loadImmediate8(RegA),
	OpLoadImmediateB:  loadImmediate8(RegB),
	OpLoadImmediateC:  loadImmediate8(RegC),
	OpLoadImmediateD:  loadImmediate8(RegD),
	OpLoadImmediateE:  loadImmediate8(RegE),
	OpLoadImmediateH:  loadImmediate8(RegH),
	OpLoadImmediateL:  loadImmediate8(RegL),
	OpIncrementBC:     increment16(RegBC),
	OpIncrementDE:     increment16(RegDE),
	OpIncrementHL:     increment16(RegHL),
	OpIncrementSP:     increment16(RegSP),
	OpDecrementBC:     decrement16(RegBC),
	OpDecrementDE:     decrement16(RegDE),
	OpDecrementHL:     decrement16(RegHL),
	OpDecrementSP:     decrement16(RegSP),
	OpStop:            halt,
}

func loadImmediate8(regid RegID) Handler {
	return func(c *CPU) {
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

func increment16(regid RegID) Handler {
	return func(c *CPU) {
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

func halt(c *CPU) {
	//TODO Handle properly
	if c.Test {
		c.Running = false
	}
}

type RegID uint8

const (
	RegAF RegID = iota
	RegBC
	RegDE
	RegHL
	RegSP
	RegA
	RegF
	RegB
	RegC
	RegD
	RegE
	RegH
	RegL
)

func (r RegID) String() string {
	switch r {
	case RegAF:
		return "AF"
	case RegBC:
		return "BC"
	case RegDE:
		return "DE"
	case RegHL:
		return "HL"
	case RegSP:
		return "SP"
	case RegA:
		return "A"
	case RegF:
		return "F"
	case RegB:
		return "B"
	case RegC:
		return "C"
	case RegD:
		return "D"
	case RegE:
		return "E"
	case RegH:
		return "H"
	case RegL:
		return "L"
	}
	panic("unknown RegID")
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

instructionStr:=map[instruction]string{
		{OpNop, "NOP"},
		{OpLoadImmediateBC, "LD  BC,d16"},
		{OpLoadIndirectBCA, "LD  (BC),A"},
		{OpIncrementBC, "INC BC"},
		{OpIncrementB, "INC B"},
		{OpDecrementB, "DEC B"},
		{OpLoadImmediateB, "LD  B,d8"},
		{OpRotateAccLeftDrop, "RLCA"},
		{OpStoreMemSP, "LD  (a16),SP"},
		{OpAddDirectHLBC, "ADD HL,BC"},
		{OpLoadIndirectABC, "LD  A,(BC)"},
		{OpDecrementBC, "DEC BC"},
		{OpIncrementC, "INC C"},
		{OpDecrementC, "DEC C"},
		{OpLoadImmediateC, "LD  C,d8"},
		{OpRotateAccRightDrop, "RRCA"},
		{OpStop, "STOP"},
		{OpLoadImmediateDE, "LD  DE,d16"},
		{OpLoadIndirectDEA, "LD  (DE,A"},
		{OpIncrementDE, "INC DE"},
		{OpIncrementD, "INC D"},
		{OpDecrementD, "DEC D"},
		{OpLoadImmediateD, "LD  D,d8"},
		{OpRotateAccLeftThrough, "RLA"},
		{OpJumpRelativeNO, "JR  r8"},
		{OpAddDirectHLDE, "ADD HL,DE"},
		{OpLoadIndirectADE, "LD  A,(DE"},
		{OpDecrementDE, "DEC DE"},
		{OpIncrementE, "INC E"},
		{OpDecrementE, "DEC E"},
		{OpLoadImmediateE, "LD  E,d8"},
		{OpRotateAccRightThrough, "RRA"},
		{OpJumpRelativeNZ, "JR  NZ,r8"},
		{OpLoadImmediateHL, "LD  HL,d16"},
		{OpLoadIndirectHLAIncrement, "LDI (HL),A"},
		{OpIncrementHL, "INC HL"},
		{OpIncrementH, "INC H"},
		{OpDecrementH, "DEC H"},
		{OpLoadImmediateH, "LD  H,d8"},
		{OpDecimalToBCD, "DAA"},
		{OpJumpRelativeZE, "JR  Z,r8"},
		{OpAddDirectHLHL, "ADD HL,HL"},
		{OpLoadIndirectAHLIncrement, "LDI A,(HL)"},
		{OpDecrementHL, "DEC HL"},
		{OpIncrementL, "INC L"},
		{OpDecrementL, "DEC L"},
		{OpLoadImmediateL, "LD  L,d8"},
		{OpInvertA, "CPL"},
		{OpJumpRelativeNC, "JR  NC,r8"},
		{OpLoadImmediateSP, "LD  SP,d16"},
		{OpLoadIndirectHLADecrement, "LDD (HL),A"},
		{OpIncrementSP, "INC SP"},
		{OpIncrementIndirectHL, "INC (HL)"},
		{OpDecrementIndirectHL, "DEC (HL)"},
		{OpLoadImmediateIndirectHL, "LD  (HL),d8"},
		{OpResetCarry, "SCF"},
		{OpJumpRelativeCA, "JR  C,r8"},
		{OpAddDirectHLSP, "ADD HL,SP"},
		{OpLoadIndirectAHLDecrement, "LDD A,(HL)"},
		{OpDecrementSP, "DEC SP"},
		{OpIncrementA, "INC A"},
		{OpDecrementA, "DEC A"},
		{OpLoadImmediateA, "LD  A,d8"},
		{OpSetCarry, "CCF"},
		{OpLoadDirectBB, "LD B,B"},
		{OpLoadDirectBC, "LD B,C"},
		{OpLoadDirectBD, "LD B,D"},
		{OpLoadDirectBE, "LD B,E"},
		{OpLoadDirectBH, "LD B,H"},
		{OpLoadDirectBL, "LD B,L"},
		{OpLoadIndirectBHL, "LD B,(HL)"},
		{OpLoadDirectBA, "LD B,A"},
		{OpLoadDirectCB, "LD C,B"},
		{OpLoadDirectCC, "LD C,C"},
		{OpLoadDirectCD, "LD C,D"},
		{OpLoadDirectCE, "LD C,E"},
		{OpLoadDirectCH, "LD C,H"},
		{OpLoadDirectCL, "LD C,L"},
		{OpLoadIndirectCHL, "LD C,(HL)"},
		{OpLoadDirectCA, "LD C,A"},
		{OpLoadDirectDB, "LD D,B"},
		{OpLoadDirectDC, "LD D,C"},
		{OpLoadDirectDD, "LD D,D"},
		{OpLoadDirectDE, "LD D,E"},
		{OpLoadDirectDH, "LD D,H"},
		{OpLoadDirectDL, "LD D,L"},
		{OpLoadIndirectDHL, "LD D,(HL)"},
		{OpLoadDirectDA, "LD D,A"},
		{OpLoadDirectEB, "LD E,B"},
		{OpLoadDirectEC, "LD E,C"},
		{OpLoadDirectED, "LD E,D"},
		{OpLoadDirectEE, "LD E,E"},
		{OpLoadDirectEH, "LD E,H"},
		{OpLoadDirectEL, "LD E,L"},
		{OpLoadIndirectEHL, "LD E,(HL)"},
		{OpLoadDirectEA, "LD E,A"},
		{OpLoadDirectHB, "LD H,B"},
		{OpLoadDirectHC, "LD H,C"},
		{OpLoadDirectHD, "LD H,D"},
		{OpLoadDirectHE, "LD H,E"},
		{OpLoadDirectHH, "LD H,H"},
		{OpLoadDirectHL, "LD H,L"},
		{OpLoadIndirectHHL, "LD H,(HL"},
		{OpLoadDirectHA, "LD H,A"},
		{OpLoadDirectLB, "LD L,B"},
		{OpLoadDirectLC, "LD L,C"},
		{OpLoadDirectLD, "LD L,D"},
		{OpLoadDirectLE, "LD L,E"},
		{OpLoadDirectLH, "LD L,H"},
		{OpLoadDirectLL, "LD L,L"},
		{OpLoadIndirectLHL, "LD L,(HL)"},
		{OpLoadDirectLA, "LD L,A"},
		{OpLoadIndirectHLB, "LD (HL),B"},
		{OpLoadIndirectHLC, "LD (HL),C"},
		{OpLoadIndirectHLD, "LD (HL),D"},
		{OpLoadIndirectHLE, "LD (HL),E"},
		{OpLoadIndirectHLH, "LD (HL),H"},
		{OpLoadIndirectHLL, "LD (HL),L"},
		{OpHalt, "HALT"},
		{OpLoadIndirectHLA, "LD (HL),A"},
		{OpLoadDirectAB, "LD A,B"},
		{OpLoadDirectAC, "LD A,C"},
		{OpLoadDirectAD, "LD A,D"},
		{OpLoadDirectAE, "LD A,E"},
		{OpLoadDirectAH, "LD A,H"},
		{OpLoadDirectAL, "LD A,L"},
		{OpLoadIndirectAHL, "LD A,(HL)"},
		{OpLoadDirectAA, "LD A,A"},
		{OpAddDirectABNoCarry, "ADD A,B"},
		{OpAddDirectACNoCarry, "ADD A,C"},
		{OpAddDirectADNoCarry, "ADD A,D"},
		{OpAddDirectAENoCarry, "ADD A,E"},
		{OpAddDirectAHNoCarry, "ADD A,H"},
		{OpAddDirectALNoCarry, "ADD A,L"},
		{OpAddIndirectAHLNoCarry, "ADD A,(HL)"},
		{OpAddDirectAANoCarry, "ADD A,A"},
		{OpAddDirectABCarry, "ADC A,B"},
		{OpAddDirectACCarry, "ADC A,C"},
		{OpAddDirectADCarry, "ADC A,D"},
		{OpAddDirectAECarry, "ADC A,E"},
		{OpAddDirectAHCarry, "ADC A,H"},
		{OpAddDirectALCarry, "ADC A,L"},
		{OpAddIndirectAHLCarry, "ADC A,(HL)"},
		{OpAddDirectAAtrue, "ADC A,A"},
		{OpSubDirectABNoCarry, "SUB A,B"},
		{OpSubDirectACNoCarry, "SUB A,C"},
		{OpSubDirectADNoCarry, "SUB A,D"},
		{OpSubDirectAENoCarry, "SUB A,E"},
		{OpSubDirectAHNoCarry, "SUB A,H"},
		{OpSubDirectALNoCarry, "SUB A,L"},
		{OpSubIndirectAHLNoCarry, "SUB A,(HL)"},
		{OpSubDirectAANoCarry, "SUB A,A"},
		{OpSubDirectABCarry, "SBC A,B"},
		{OpSubDirectACCarry, "SBC A,C"},
		{OpSubDirectADCarry, "SBC A,D"},
		{OpSubDirectAECarry, "SBC A,E"},
		{OpSubDirectAHCarry, "SBC A,H"},
		{OpSubDirectALCarry, "SBC A,L"},
		{OpSubIndirectAHLCarry, "SBC A,(HL)"},
		{OpSubDirectAACarry, "SBC A,A"},
		{OpAndDirectAB, "AND A,B"},
		{OpAndDirectAC, "AND A,C"},
		{OpAndDirectAD, "AND A,D"},
		{OpAndDirectAE, "AND A,E"},
		{OpAndDirectAH, "AND A,H"},
		{OpAndDirectAL, "AND A,L"},
		{OpAndIndirectAHL, "AND A,(HL)"},
		{OpAndDirectAA, "AND A,A"},
		{OpXorDirectAB, "XOR A,B"},
		{OpXorDirectAC, "XOR A,C"},
		{OpXorDirectAD, "XOR A,D"},
		{OpXorDirectAE, "XOR A,E"},
		{OpXorDirectAH, "XOR A,H"},
		{OpXorDirectAL, "XOR A,L"},
		{OpXorIndirectAHL, "XOR A,(HL)"},
		{OpXorDirectAA, "XOR A,A"},
		{OpOrDirectAB, "OR  A,B"},
		{OpOrDirectAC, "OR  A,C"},
		{OpOrDirectAD, "OR  A,D"},
		{OpOrDirectAE, "OR  A,E"},
		{OpOrDirectAH, "OR  A,H"},
		{OpOrDirectAL, "OR  A,L"},
		{OpOrIndirectAHL, "OR  A,(HL)"},
		{OpOrDirectAA, "OR  A,A"},
		{OpCmpDirectAB, "CP  A,B"},
		{OpCmpDirectAC, "CP  A,C"},
		{OpCmpDirectAD, "CP  A,D"},
		{OpCmpDirectAE, "CP  A,E"},
		{OpCmpDirectAH, "CP  A,H"},
		{OpCmpDirectAL, "CP  A,L"},
		{OpCmpIndirectAHL, "CP  A,(HL)"},
		{OpCmpDirectAA, "CP  A,A"},
		{OpReturnNZ, "RET NZ"},
		{OpPopBC, "POP BC"},
		{OpJumpAbsoluteNZ, "JP  NZ,a16"},
		{OpJumpAbsoluteNO, "JP  a16"},
		{OpCallNZ, "CALL NZ,a16"},
		{OpPushBC, "PUSH BC"},
		{OpAddImmediateANoCarry, "ADD A,d8"},
		{OpRestart00, "RST 00h"},
		{OpReturnZE, "RET Z"},
		{OpReturnNO, "RET"},
		{OpJumpAbsoluteZE, "JP  Z,a16"},
		{OpCBPrefix, "PREFIX: See cbhandlers below (01xx)"},
		{OpCallZE, "CALL Z,a16"},
		{OpCallNO, "CALL a16"},
		{OpAddImmediateACarry, "ADC A,d8"},
		{OpRestart08, "RST 08h"},
		{OpReturnNC, "RET NC"},
		{OpPopRegDE, "POP DE"},
		{OpJumpAbsoluteNC, "JP  NC,a16"},
		{OpCallNC, "CALL NC,a16"},
		{OpPushDE, "PUSH DE"},
		{OpSubImmediateANoCarry, "SUB A,d8"},
		{OpRestart10, "RST 10h"},
		{OpReturnCA, "RET C"},
		{OpRETI, "RETI"},
		{OpJumpAbsoluteCA, "JP  C,a16"},
		{OpCallCA, "CALL C,a16"},
		{OpSubImmediateACarry, "SBC A,d8"},
		{OpRestart18, "RST 18h"},
		{OpLoadHighAbsA, "LDH (a8),A"},
		{OpPopHL, "POP HL"},
		{OpLoadHighMemCA, "LD  (C),A"},
		{OpPushHL, "PUSH HL"},
		{OpAndImmediateA, "AND A,d8"},
		{OpRestart20, "RST 20h"},
		{OpAddImmediateSignedSP, "ADD SP,r8"},
		{OpJumpAbsoluteHL, "JP  (HL)"},
		{OpStoreMemA, "LD  (a16),A"},
		{OpXorImmediateA, "XOR A,d8"},
		{OpRestart28, "RST 28h"},
		{OpLoadHighRegA, "LDH A,(a8)"},
		{OpPopAF, "POP AF"},
		{OpLoadHighRegAC, "LD  A,(C)"},
		{OpResetInt, "DI"},
		{OpPushAF, "PUSH AF"},
		{OpOrImmediateA, "OR  A,d8"},
		{OpRestart30, "RES 30h"},
		{OpLoadOffsetHLSP, "LD  HL,SP+r8"},
		{OpLoadDirectSPHL, "LD  SP,HL"},
		{OpLoadMemA, "LD  A,(a16)"},
		{OpSetInt, "EI"},
		{OpCmpImmediateA, "CP  A,d8"},
		{OpRestart38, "RST 38h"},
		{OpCbRotateRegBLeftRot, "RLC B"},
		{OpCbRotateRegCLeftRot, "RLC C"},
		{OpCbRotateRegDLeftRot, "RLC D"},
		{OpCbRotateRegELeftRot, "RLC E"},
		{OpCbRotateRegHLeftRot, "RLC H"},
		{OpCbRotateRegLLeftRot, "RLC L"},
		{OpCbRotateIndHLLeftRot, "RLC (HL)"},
		{OpCbRotateRegALeftRot, "RLC A"},
		{OpCbRotateRegBRightRot, "RRC B"},
		{OpCbRotateRegCRightRot, "RRC C"},
		{OpCbRotateRegDRightRot, "RRC D"},
		{OpCbRotateRegERightRot, "RRC E"},
		{OpCbRotateRegHRightRot, "RRC H"},
		{OpCbRotateRegLRightRot, "RRC L"},
		{OpCbRotateIndHLRightRot, "RRC (HL)"},
		{OpCbRotateRegARightRot, "RRC A"},
		{OpCbRotateRegBLeftThC, "RL  B"},
		{OpCbRotateRegCLeftThC, "RL  C"},
		{OpCbRotateRegDLeftThC, "RL  D"},
		{OpCbRotateRegELeftThC, "RL  E"},
		{OpCbRotateRegHLeftThC, "RL  H"},
		{OpCbRotateRegLLeftThC, "RL  L"},
		{OpCbRotateIndHLLeftThC, "RL  (HL)"},
		{OpCbRotateRegALeftThC, "RL  A"},
		{OpCbRotateRegBRightThC, "RR  B"},
		{OpCbRotateRegCRightThC, "RR  C"},
		{OpCbRotateRegDRightThC, "RR  D"},
		{OpCbRotateRegERightThC, "RR  E"},
		{OpCbRotateRegHRightThC, "RR  H"},
		{OpCbRotateRegLRightThC, "RR  L"},
		{OpCbRotateIndHLRightThC, "RR  (HL)"},
		{OpCbRotateRegARightThC, "RR  A"},
		{OpCbRotateRegBLeftShf, "SLA B"},
		{OpCbRotateRegCLeftShf, "SLA C"},
		{OpCbRotateRegDLeftShf, "SLA D"},
		{OpCbRotateRegELeftShf, "SLA E"},
		{OpCbRotateRegHLeftShf, "SLA H"},
		{OpCbRotateRegLLeftShf, "SLA L"},
		{OpCbRotateIndHLLeftShf, "SLA (HL)"},
		{OpCbRotateRegALeftShf, "SLA A"},
		{OpCbRotateRegBRightRep, "SRA B"},
		{OpCbRotateRegCRightRep, "SRA C"},
		{OpCbRotateRegDRightRep, "SRA D"},
		{OpCbRotateRegERightRep, "SRA E"},
		{OpCbRotateRegHRightRep, "SRA H"},
		{OpCbRotateRegLRightRep, "SRA L"},
		{OpCbRotateIndHLRightRep, "SRA (HL)"},
		{OpCbRotateRegARightRep, "SRA A"},
		{OpCbSwapDirectB, "SWAP B"},
		{OpCbSwapDirectC, "SWAP C"},
		{OpCbSwapDirectD, "SWAP D"},
		{OpCbSwapDirectE, "SWAP E"},
		{OpCbSwapDirectH, "SWAP H"},
		{OpCbSwapDirectL, "SWAP L"},
		{OpCbSwapIndirectHL, "SWAP (HL)"},
		{OpCbSwapDirectA, "SWAP A"},
		{OpCbRotateRegBRightShf, "SRL B"},
		{OpCbRotateRegCRightShf, "SRL C"},
		{OpCbRotateRegDRightShf, "SRL D"},
		{OpCbRotateRegERightShf, "SRL E"},
		{OpCbRotateRegHRightShf, "SRL H"},
		{OpCbRotateRegLRightShf, "SRL L"},
		{OpCbRotateIndHLRightShf, "SRL (HL)"},
		{OpCbRotateRegARightShf, "SRL A"},
		{OpCbBitDirectB0, "BIT 0,B"},
		{OpCbBitDirectC0, "BIT 0,C"},
		{OpCbBitDirectD0, "BIT 0,D"},
		{OpCbBitDirectE0, "BIT 0,E"},
		{OpCbBitDirectH0, "BIT 0,H"},
		{OpCbBitDirectL0, "BIT 0,L"},
		{OpCbBitIndirectHL0, "BIT 0,(HL)"},
		{OpCbBitDirectA0, "BIT 0,A"},
		{OpCbBitDirectB1, "BIT 1,B"},
		{OpCbBitDirectC1, "BIT 1,C"},
		{OpCbBitDirectD1, "BIT 1,D"},
		{OpCbBitDirectE1, "BIT 1,E"},
		{OpCbBitDirectH1, "BIT 1,H"},
		{OpCbBitDirectL1, "BIT 1,L"},
		{OpCbBitIndirectHL1, "BIT 1,(HL)"},
		{OpCbBitDirectA1, "BIT 1,A"},
		{OpCbBitDirectB2, "BIT 2,B"},
		{OpCbBitDirectC2, "BIT 2,C"},
		{OpCbBitDirectD2, "BIT 2,D"},
		{OpCbBitDirectE2, "BIT 2,E"},
		{OpCbBitDirectH2, "BIT 2,H"},
		{OpCbBitDirectL2, "BIT 2,L"},
		{OpCbBitIndirectHL2, "BIT 2,(HL)"},
		{OpCbBitDirectA2, "BIT 2,A"},
		{OpCbBitDirectB3, "BIT 3,B"},
		{OpCbBitDirectC3, "BIT 3,C"},
		{OpCbBitDirectD3, "BIT 3,D"},
		{OpCbBitDirectE3, "BIT 3,E"},
		{OpCbBitDirectH3, "BIT 3,H"},
		{OpCbBitDirectL3, "BIT 3,L"},
		{OpCbBitIndirectHL3, "BIT 3,(HL)"},
		{OpCbBitDirectA3, "BIT 3,A"},
		{OpCbBitDirectB4, "BIT 4,B"},
		{OpCbBitDirectC4, "BIT 4,C"},
		{OpCbBitDirectD4, "BIT 4,D"},
		{OpCbBitDirectE4, "BIT 4,E"},
		{OpCbBitDirectH4, "BIT 4,H"},
		{OpCbBitDirectL4, "BIT 4,L"},
		{OpCbBitIndirectHL4, "BIT 4,(HL)"},
		{OpCbBitDirectA4, "BIT 4,A"},
		{OpCbBitDirectB5, "BIT 5,B"},
		{OpCbBitDirectC5, "BIT 5,C"},
		{OpCbBitDirectD5, "BIT 5,D"},
		{OpCbBitDirectE5, "BIT 5,E"},
		{OpCbBitDirectH5, "BIT 5,H"},
		{OpCbBitDirectL5, "BIT 5,L"},
		{OpCbBitIndirectHL5, "BIT 5,(HL)"},
		{OpCbBitDirectA5, "BIT 5,A"},
		{OpCbBitDirectB6, "BIT 6,B"},
		{OpCbBitDirectC6, "BIT 6,C"},
		{OpCbBitDirectD6, "BIT 6,D"},
		{OpCbBitDirectE6, "BIT 6,E"},
		{OpCbBitDirectH6, "BIT 6,H"},
		{OpCbBitDirectL6, "BIT 6,L"},
		{OpCbBitIndirectHL6, "BIT 6,(HL)"},
		{OpCbBitDirectA6, "BIT 6,A"},
		{OpCbBitDirectB7, "BIT 7,B"},
		{OpCbBitDirectC7, "BIT 7,C"},
		{OpCbBitDirectD7, "BIT 7,D"},
		{OpCbBitDirectE7, "BIT 7,E"},
		{OpCbBitDirectH7, "BIT 7,H"},
		{OpCbBitDirectL7, "BIT 7,L"},
		{OpCbBitIndirectHL7, "BIT 7,(HL)"},
		{OpCbBitDirectA7, "BIT 7,A"},
		{OpCbResetDirectB0, "RES 0,B"},
		{OpCbResetDirectC0, "RES 0,C"},
		{OpCbResetDirectD0, "RES 0,D"},
		{OpCbResetDirectE0, "RES 0,E"},
		{OpCbResetDirectH0, "RES 0,H"},
		{OpCbResetDirectL0, "RES 0,L"},
		{OpCbResetIndirectHL0, "RES 0,(HL)"},
		{OpCbResetDirectA0, "RES 0,A"},
		{OpCbResetDirectB1, "RES 1,B"},
		{OpCbResetDirectC1, "RES 1,C"},
		{OpCbResetDirectD1, "RES 1,D"},
		{OpCbResetDirectE1, "RES 1,E"},
		{OpCbResetDirectH1, "RES 1,H"},
		{OpCbResetDirectL1, "RES 1,L"},
		{OpCbResetIndirectHL1, "RES 1,(HL)"},
		{OpCbResetDirectA1, "RES 1,A"},
		{OpCbResetDirectB2, "RES 2,B"},
		{OpCbResetDirectC2, "RES 2,C"},
		{OpCbResetDirectD2, "RES 2,D"},
		{OpCbResetDirectE2, "RES 2,E"},
		{OpCbResetDirectH2, "RES 2,H"},
		{OpCbResetDirectL2, "RES 2,L"},
		{OpCbResetIndirectHL2, "RES 2,(HL)"},
		{OpCbResetDirectA2, "RES 2,A"},
		{OpCbResetDirectB3, "RES 3,B"},
		{OpCbResetDirectC3, "RES 3,C"},
		{OpCbResetDirectD3, "RES 3,D"},
		{OpCbResetDirectE3, "RES 3,E"},
		{OpCbResetDirectH3, "RES 3,H"},
		{OpCbResetDirectL3, "RES 3,L"},
		{OpCbResetIndirectHL3, "RES 3,(HL)"},
		{OpCbResetDirectA3, "RES 3,A"},
		{OpCbResetDirectB4, "RES 4,B"},
		{OpCbResetDirectC4, "RES 4,C"},
		{OpCbResetDirectD4, "RES 4,D"},
		{OpCbResetDirectE4, "RES 4,E"},
		{OpCbResetDirectH4, "RES 4,H"},
		{OpCbResetDirectL4, "RES 4,L"},
		{OpCbResetIndirectHL4, "RES 4,(HL)"},
		{OpCbResetDirectA4, "RES 4,A"},
		{OpCbResetDirectB5, "RES 5,B"},
		{OpCbResetDirectC5, "RES 5,C"},
		{OpCbResetDirectD5, "RES 5,D"},
		{OpCbResetDirectE5, "RES 5,E"},
		{OpCbResetDirectH5, "RES 5,H"},
		{OpCbResetDirectL5, "RES 5,L"},
		{OpCbResetIndirectHL5, "RES 5,(HL)"},
		{OpCbResetDirectA5, "RES 5,A"},
		{OpCbResetDirectB6, "RES 6,B"},
		{OpCbResetDirectC6, "RES 6,C"},
		{OpCbResetDirectD6, "RES 6,D"},
		{OpCbResetDirectE6, "RES 6,E"},
		{OpCbResetDirectH6, "RES 6,H"},
		{OpCbResetDirectL6, "RES 6,L"},
		{OpCbResetIndirectHL6, "RES 6,(HL)"},
		{OpCbResetDirectA6, "RES 6,A"},
		{OpCbResetDirectB7, "RES 7,B"},
		{OpCbResetDirectC7, "RES 7,C"},
		{OpCbResetDirectD7, "RES 7,D"},
		{OpCbResetDirectE7, "RES 7,E"},
		{OpCbResetDirectH7, "RES 7,H"},
		{OpCbResetDirectL7, "RES 7,L"},
		{OpCbResetIndirectHL7, "RES 7,(HL)"},
		{OpCbResetDirectA7, "RES 7,A"},
		{OpCbSetDirectB0, "SET 0,B"},
		{OpCbSetDirectC0, "SET 0,C"},
		{OpCbSetDirectD0, "SET 0,D"},
		{OpCbSetDirectE0, "SET 0,E"},
		{OpCbSetDirectH0, "SET 0,H"},
		{OpCbSetDirectL0, "SET 0,L"},
		{OpCbSetIndirectHL0, "SET 0,(HL)"},
		{OpCbSetDirectA0, "SET 0,A"},
		{OpCbSetDirectB1, "SET 1,B"},
		{OpCbSetDirectC1, "SET 1,C"},
		{OpCbSetDirectD1, "SET 1,D"},
		{OpCbSetDirectE1, "SET 1,E"},
		{OpCbSetDirectH1, "SET 1,H"},
		{OpCbSetDirectL1, "SET 1,L"},
		{OpCbSetIndirectHL1, "SET 1,(HL)"},
		{OpCbSetDirectA1, "SET 1,A"},
		{OpCbSetDirectB2, "SET 2,B"},
		{OpCbSetDirectC2, "SET 2,C"},
		{OpCbSetDirectD2, "SET 2,D"},
		{OpCbSetDirectE2, "SET 2,E"},
		{OpCbSetDirectH2, "SET 2,H"},
		{OpCbSetDirectL2, "SET 2,L"},
		{OpCbSetIndirectHL2, "SET 2,(HL)"},
		{OpCbSetDirectA2, "SET 2,A"},
		{OpCbSetDirectB3, "SET 3,B"},
		{OpCbSetDirectC3, "SET 3,C"},
		{OpCbSetDirectD3, "SET 3,D"},
		{OpCbSetDirectE3, "SET 3,E"},
		{OpCbSetDirectH3, "SET 3,H"},
		{OpCbSetDirectL3, "SET 3,L"},
		{OpCbSetIndirectHL3, "SET 3,(HL)"},
		{OpCbSetDirectA3, "SET 3,A"},
		{OpCbSetDirectB4, "SET 4,B"},
		{OpCbSetDirectC4, "SET 4,C"},
		{OpCbSetDirectD4, "SET 4,D"},
		{OpCbSetDirectE4, "SET 4,E"},
		{OpCbSetDirectH4, "SET 4,H"},
		{OpCbSetDirectL4, "SET 4,L"},
		{OpCbSetIndirectHL4, "SET 4,(HL)"},
		{OpCbSetDirectA4, "SET 4,A"},
		{OpCbSetDirectB5, "SET 5,B"},
		{OpCbSetDirectC5, "SET 5,C"},
		{OpCbSetDirectD5, "SET 5,D"},
		{OpCbSetDirectE5, "SET 5,E"},
		{OpCbSetDirectH5, "SET 5,H"},
		{OpCbSetDirectL5, "SET 5,L"},
		{OpCbSetIndirectHL5, "SET 5,(HL)"},
		{OpCbSetDirectA5, "SET 5,A"},
		{OpCbSetDirectB6, "SET 6,B"},
		{OpCbSetDirectC6, "SET 6,C"},
		{OpCbSetDirectD6, "SET 6,D"},
		{OpCbSetDirectE6, "SET 6,E"},
		{OpCbSetDirectH6, "SET 6,H"},
		{OpCbSetDirectL6, "SET 6,L"},
		{OpCbSetIndirectHL6, "SET 6,(HL)"},
		{OpCbSetDirectA6, "SET 6,A"},
		{OpCbSetDirectB7, "SET 7,B"},
		{OpCbSetDirectC7, "SET 7,C"},
		{OpCbSetDirectD7, "SET 7,D"},
		{OpCbSetDirectE7, "SET 7,E"},
		{OpCbSetDirectH7, "SET 7,H"},
		{OpCbSetDirectL7, "SET 7,L"},
		{OpCbSetIndirectHL7, "SET 7,(HL)"},
		{OpCbSetDirectA7, "SET 7,A"},
	}