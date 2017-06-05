package hegb

import (
	"encoding/binary"
)

// Instruction ID type
type instruction uint16

// Z80 instructions
const (
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
	OpLoadIndirectDEA                             // 12 LD  (DE),A
	OpIncrementDE                                 // 13 INC DE
	OpIncrementD                                  // 14 INC D
	OpDecrementD                                  // 15 DEC D
	OpLoadImmediateD                              // 16 LD  D,d8
	OpRotateAccLeftThrough                        // 17 RLA
	OpJumpRelativeNO                              // 18 JR  r8
	OpAddDirectHLDE                               // 19 ADD HL,DE
	OpLoadIndirectADE                             // 1a LD  A,(DE)
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
)

// CB prefix operations
const (
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
	OpLoadDirectAA:    loadRegister(RegA, RegA),
	OpLoadDirectAB:    loadRegister(RegA, RegB),
	OpLoadDirectAC:    loadRegister(RegA, RegC),
	OpLoadDirectAD:    loadRegister(RegA, RegD),
	OpLoadDirectAE:    loadRegister(RegA, RegE),
	OpLoadDirectAH:    loadRegister(RegA, RegH),
	OpLoadDirectAL:    loadRegister(RegA, RegL),
	OpLoadDirectBA:    loadRegister(RegB, RegA),
	OpLoadDirectBB:    loadRegister(RegB, RegB),
	OpLoadDirectBC:    loadRegister(RegB, RegC),
	OpLoadDirectBD:    loadRegister(RegB, RegD),
	OpLoadDirectBE:    loadRegister(RegB, RegE),
	OpLoadDirectBH:    loadRegister(RegB, RegH),
	OpLoadDirectBL:    loadRegister(RegB, RegL),
	OpLoadDirectCA:    loadRegister(RegC, RegA),
	OpLoadDirectCB:    loadRegister(RegC, RegB),
	OpLoadDirectCC:    loadRegister(RegC, RegC),
	OpLoadDirectCD:    loadRegister(RegC, RegD),
	OpLoadDirectCE:    loadRegister(RegC, RegE),
	OpLoadDirectCH:    loadRegister(RegC, RegH),
	OpLoadDirectCL:    loadRegister(RegC, RegL),
	OpLoadDirectDA:    loadRegister(RegD, RegA),
	OpLoadDirectDB:    loadRegister(RegD, RegB),
	OpLoadDirectDC:    loadRegister(RegD, RegC),
	OpLoadDirectDD:    loadRegister(RegD, RegD),
	OpLoadDirectDE:    loadRegister(RegD, RegE),
	OpLoadDirectDH:    loadRegister(RegD, RegH),
	OpLoadDirectDL:    loadRegister(RegD, RegL),
	OpLoadDirectEA:    loadRegister(RegE, RegA),
	OpLoadDirectEB:    loadRegister(RegE, RegB),
	OpLoadDirectEC:    loadRegister(RegE, RegC),
	OpLoadDirectED:    loadRegister(RegE, RegD),
	OpLoadDirectEE:    loadRegister(RegE, RegE),
	OpLoadDirectEH:    loadRegister(RegE, RegH),
	OpLoadDirectEL:    loadRegister(RegE, RegL),
	OpLoadDirectHA:    loadRegister(RegH, RegA),
	OpLoadDirectHB:    loadRegister(RegH, RegB),
	OpLoadDirectHC:    loadRegister(RegH, RegC),
	OpLoadDirectHD:    loadRegister(RegH, RegD),
	OpLoadDirectHE:    loadRegister(RegH, RegE),
	OpLoadDirectHH:    loadRegister(RegH, RegH),
	OpLoadDirectHL:    loadRegister(RegH, RegL),
	OpLoadDirectLA:    loadRegister(RegL, RegA),
	OpLoadDirectLB:    loadRegister(RegL, RegB),
	OpLoadDirectLC:    loadRegister(RegL, RegC),
	OpLoadDirectLD:    loadRegister(RegL, RegD),
	OpLoadDirectLE:    loadRegister(RegL, RegE),
	OpLoadDirectLH:    loadRegister(RegL, RegH),
	OpLoadDirectLL:    loadRegister(RegL, RegL),
	OpIncrementBC:     increment16(RegBC),
	OpIncrementDE:     increment16(RegDE),
	OpIncrementHL:     increment16(RegHL),
	OpIncrementSP:     increment16(RegSP),
	OpDecrementBC:     decrement16(RegBC),
	OpDecrementDE:     decrement16(RegDE),
	OpDecrementHL:     decrement16(RegHL),
	OpDecrementSP:     decrement16(RegSP),
	OpIncrementA:      increment8(RegA),
	OpIncrementB:      increment8(RegB),
	OpIncrementC:      increment8(RegC),
	OpIncrementD:      increment8(RegD),
	OpIncrementE:      increment8(RegE),
	OpIncrementH:      increment8(RegH),
	OpIncrementL:      increment8(RegL),
	OpDecrementA:      decrement8(RegA),
	OpDecrementB:      decrement8(RegB),
	OpDecrementC:      decrement8(RegC),
	OpDecrementD:      decrement8(RegD),
	OpDecrementE:      decrement8(RegE),
	OpDecrementH:      decrement8(RegH),
	OpDecrementL:      decrement8(RegL),
	OpStop:            halt,
	OpCbBitDirectA0:   bit(RegA, 0),
	OpCbBitDirectA1:   bit(RegA, 1),
	OpCbBitDirectA2:   bit(RegA, 2),
	OpCbBitDirectA3:   bit(RegA, 3),
	OpCbBitDirectA4:   bit(RegA, 4),
	OpCbBitDirectA5:   bit(RegA, 5),
	OpCbBitDirectA6:   bit(RegA, 6),
	OpCbBitDirectA7:   bit(RegA, 7),
	OpCbBitDirectB0:   bit(RegB, 0),
	OpCbBitDirectB1:   bit(RegB, 1),
	OpCbBitDirectB2:   bit(RegB, 2),
	OpCbBitDirectB3:   bit(RegB, 3),
	OpCbBitDirectB4:   bit(RegB, 4),
	OpCbBitDirectB5:   bit(RegB, 5),
	OpCbBitDirectB6:   bit(RegB, 6),
	OpCbBitDirectB7:   bit(RegB, 7),
	OpCbBitDirectC0:   bit(RegC, 0),
	OpCbBitDirectC1:   bit(RegC, 1),
	OpCbBitDirectC2:   bit(RegC, 2),
	OpCbBitDirectC3:   bit(RegC, 3),
	OpCbBitDirectC4:   bit(RegC, 4),
	OpCbBitDirectC5:   bit(RegC, 5),
	OpCbBitDirectC6:   bit(RegC, 6),
	OpCbBitDirectC7:   bit(RegC, 7),
	OpCbBitDirectD0:   bit(RegD, 0),
	OpCbBitDirectD1:   bit(RegD, 1),
	OpCbBitDirectD2:   bit(RegD, 2),
	OpCbBitDirectD3:   bit(RegD, 3),
	OpCbBitDirectD4:   bit(RegD, 4),
	OpCbBitDirectD5:   bit(RegD, 5),
	OpCbBitDirectD6:   bit(RegD, 6),
	OpCbBitDirectD7:   bit(RegD, 7),
	OpCbBitDirectE0:   bit(RegE, 0),
	OpCbBitDirectE1:   bit(RegE, 1),
	OpCbBitDirectE2:   bit(RegE, 2),
	OpCbBitDirectE3:   bit(RegE, 3),
	OpCbBitDirectE4:   bit(RegE, 4),
	OpCbBitDirectE5:   bit(RegE, 5),
	OpCbBitDirectE6:   bit(RegE, 6),
	OpCbBitDirectE7:   bit(RegE, 7),
	OpCbBitDirectH0:   bit(RegH, 0),
	OpCbBitDirectH1:   bit(RegH, 1),
	OpCbBitDirectH2:   bit(RegH, 2),
	OpCbBitDirectH3:   bit(RegH, 3),
	OpCbBitDirectH4:   bit(RegH, 4),
	OpCbBitDirectH5:   bit(RegH, 5),
	OpCbBitDirectH6:   bit(RegH, 6),
	OpCbBitDirectH7:   bit(RegH, 7),
	OpCbBitDirectL0:   bit(RegL, 0),
	OpCbBitDirectL1:   bit(RegL, 1),
	OpCbBitDirectL2:   bit(RegL, 2),
	OpCbBitDirectL3:   bit(RegL, 3),
	OpCbBitDirectL4:   bit(RegL, 4),
	OpCbBitDirectL5:   bit(RegL, 5),
	OpCbBitDirectL6:   bit(RegL, 6),
	OpCbBitDirectL7:   bit(RegL, 7),
	OpCbSetDirectA0:   bset(RegA, 0),
	OpCbSetDirectA1:   bset(RegA, 1),
	OpCbSetDirectA2:   bset(RegA, 2),
	OpCbSetDirectA3:   bset(RegA, 3),
	OpCbSetDirectA4:   bset(RegA, 4),
	OpCbSetDirectA5:   bset(RegA, 5),
	OpCbSetDirectA6:   bset(RegA, 6),
	OpCbSetDirectA7:   bset(RegA, 7),
	OpCbSetDirectB0:   bset(RegB, 0),
	OpCbSetDirectB1:   bset(RegB, 1),
	OpCbSetDirectB2:   bset(RegB, 2),
	OpCbSetDirectB3:   bset(RegB, 3),
	OpCbSetDirectB4:   bset(RegB, 4),
	OpCbSetDirectB5:   bset(RegB, 5),
	OpCbSetDirectB6:   bset(RegB, 6),
	OpCbSetDirectB7:   bset(RegB, 7),
	OpCbSetDirectC0:   bset(RegC, 0),
	OpCbSetDirectC1:   bset(RegC, 1),
	OpCbSetDirectC2:   bset(RegC, 2),
	OpCbSetDirectC3:   bset(RegC, 3),
	OpCbSetDirectC4:   bset(RegC, 4),
	OpCbSetDirectC5:   bset(RegC, 5),
	OpCbSetDirectC6:   bset(RegC, 6),
	OpCbSetDirectC7:   bset(RegC, 7),
	OpCbSetDirectD0:   bset(RegD, 0),
	OpCbSetDirectD1:   bset(RegD, 1),
	OpCbSetDirectD2:   bset(RegD, 2),
	OpCbSetDirectD3:   bset(RegD, 3),
	OpCbSetDirectD4:   bset(RegD, 4),
	OpCbSetDirectD5:   bset(RegD, 5),
	OpCbSetDirectD6:   bset(RegD, 6),
	OpCbSetDirectD7:   bset(RegD, 7),
	OpCbSetDirectE0:   bset(RegE, 0),
	OpCbSetDirectE1:   bset(RegE, 1),
	OpCbSetDirectE2:   bset(RegE, 2),
	OpCbSetDirectE3:   bset(RegE, 3),
	OpCbSetDirectE4:   bset(RegE, 4),
	OpCbSetDirectE5:   bset(RegE, 5),
	OpCbSetDirectE6:   bset(RegE, 6),
	OpCbSetDirectE7:   bset(RegE, 7),
	OpCbSetDirectH0:   bset(RegH, 0),
	OpCbSetDirectH1:   bset(RegH, 1),
	OpCbSetDirectH2:   bset(RegH, 2),
	OpCbSetDirectH3:   bset(RegH, 3),
	OpCbSetDirectH4:   bset(RegH, 4),
	OpCbSetDirectH5:   bset(RegH, 5),
	OpCbSetDirectH6:   bset(RegH, 6),
	OpCbSetDirectH7:   bset(RegH, 7),
	OpCbSetDirectL0:   bset(RegL, 0),
	OpCbSetDirectL1:   bset(RegL, 1),
	OpCbSetDirectL2:   bset(RegL, 2),
	OpCbSetDirectL3:   bset(RegL, 3),
	OpCbSetDirectL4:   bset(RegL, 4),
	OpCbSetDirectL5:   bset(RegL, 5),
	OpCbSetDirectL6:   bset(RegL, 6),
	OpCbSetDirectL7:   bset(RegL, 7),
	OpCbResetDirectA0: breset(RegA, 0),
	OpCbResetDirectA1: breset(RegA, 1),
	OpCbResetDirectA2: breset(RegA, 2),
	OpCbResetDirectA3: breset(RegA, 3),
	OpCbResetDirectA4: breset(RegA, 4),
	OpCbResetDirectA5: breset(RegA, 5),
	OpCbResetDirectA6: breset(RegA, 6),
	OpCbResetDirectA7: breset(RegA, 7),
	OpCbResetDirectB0: breset(RegB, 0),
	OpCbResetDirectB1: breset(RegB, 1),
	OpCbResetDirectB2: breset(RegB, 2),
	OpCbResetDirectB3: breset(RegB, 3),
	OpCbResetDirectB4: breset(RegB, 4),
	OpCbResetDirectB5: breset(RegB, 5),
	OpCbResetDirectB6: breset(RegB, 6),
	OpCbResetDirectB7: breset(RegB, 7),
	OpCbResetDirectC0: breset(RegC, 0),
	OpCbResetDirectC1: breset(RegC, 1),
	OpCbResetDirectC2: breset(RegC, 2),
	OpCbResetDirectC3: breset(RegC, 3),
	OpCbResetDirectC4: breset(RegC, 4),
	OpCbResetDirectC5: breset(RegC, 5),
	OpCbResetDirectC6: breset(RegC, 6),
	OpCbResetDirectC7: breset(RegC, 7),
	OpCbResetDirectD0: breset(RegD, 0),
	OpCbResetDirectD1: breset(RegD, 1),
	OpCbResetDirectD2: breset(RegD, 2),
	OpCbResetDirectD3: breset(RegD, 3),
	OpCbResetDirectD4: breset(RegD, 4),
	OpCbResetDirectD5: breset(RegD, 5),
	OpCbResetDirectD6: breset(RegD, 6),
	OpCbResetDirectD7: breset(RegD, 7),
	OpCbResetDirectE0: breset(RegE, 0),
	OpCbResetDirectE1: breset(RegE, 1),
	OpCbResetDirectE2: breset(RegE, 2),
	OpCbResetDirectE3: breset(RegE, 3),
	OpCbResetDirectE4: breset(RegE, 4),
	OpCbResetDirectE5: breset(RegE, 5),
	OpCbResetDirectE6: breset(RegE, 6),
	OpCbResetDirectE7: breset(RegE, 7),
	OpCbResetDirectH0: breset(RegH, 0),
	OpCbResetDirectH1: breset(RegH, 1),
	OpCbResetDirectH2: breset(RegH, 2),
	OpCbResetDirectH3: breset(RegH, 3),
	OpCbResetDirectH4: breset(RegH, 4),
	OpCbResetDirectH5: breset(RegH, 5),
	OpCbResetDirectH6: breset(RegH, 6),
	OpCbResetDirectH7: breset(RegH, 7),
	OpCbResetDirectL0: breset(RegL, 0),
	OpCbResetDirectL1: breset(RegL, 1),
	OpCbResetDirectL2: breset(RegL, 2),
	OpCbResetDirectL3: breset(RegL, 3),
	OpCbResetDirectL4: breset(RegL, 4),
	OpCbResetDirectL5: breset(RegL, 5),
	OpCbResetDirectL6: breset(RegL, 6),
	OpCbResetDirectL7: breset(RegL, 7),
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

func loadRegister(regtgt, regsrc RegID) Handler {
	return func(c *CPU) {
		setreg8(c, regtgt, getreg8(c, regsrc))
		c.Cycles.Add(1, 4)
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

func increment8(regid RegID) Handler {
	return func(c *CPU) {
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

func bit(regid RegID, bitNum uint8) Handler {
	return func(c *CPU) {
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
		val := getreg8(c, regid)
		val |= 1 << bitNum
		setreg8(c, regid, val)
		c.Cycles.Add(2, 8)
	}
}

func breset(regid RegID, bitNum uint8) Handler {
	return func(c *CPU) {
		val := getreg8(c, regid)
		val &= ^(1 << bitNum)
		setreg8(c, regid, val)
		c.Cycles.Add(2, 8)
	}
}

// RegID identifies a register
type RegID uint8

// Registers
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

func (i instruction) String() string {
	switch i {
	case OpNop:
		return "NOP"
	case OpLoadImmediateBC:
		return "LD  BC,d16"
	case OpLoadIndirectBCA:
		return "LD  (BC),A"
	case OpIncrementBC:
		return "INC BC"
	case OpIncrementB:
		return "INC B"
	case OpDecrementB:
		return "DEC B"
	case OpLoadImmediateB:
		return "LD  B,d8"
	case OpRotateAccLeftDrop:
		return "RLCA"
	case OpStoreMemSP:
		return "LD  (a16),SP"
	case OpAddDirectHLBC:
		return "ADD HL,BC"
	case OpLoadIndirectABC:
		return "LD  A,(BC)"
	case OpDecrementBC:
		return "DEC BC"
	case OpIncrementC:
		return "INC C"
	case OpDecrementC:
		return "DEC C"
	case OpLoadImmediateC:
		return "LD  C,d8"
	case OpRotateAccRightDrop:
		return "RRCA"
	case OpStop:
		return "STOP"
	case OpLoadImmediateDE:
		return "LD  DE,d16"
	case OpLoadIndirectDEA:
		return "LD  (DE),A"
	case OpIncrementDE:
		return "INC DE"
	case OpIncrementD:
		return "INC D"
	case OpDecrementD:
		return "DEC D"
	case OpLoadImmediateD:
		return "LD  D,d8"
	case OpRotateAccLeftThrough:
		return "RLA"
	case OpJumpRelativeNO:
		return "JR  r8"
	case OpAddDirectHLDE:
		return "ADD HL,DE"
	case OpLoadIndirectADE:
		return "LD  A,(DE)"
	case OpDecrementDE:
		return "DEC DE"
	case OpIncrementE:
		return "INC E"
	case OpDecrementE:
		return "DEC E"
	case OpLoadImmediateE:
		return "LD  E,d8"
	case OpRotateAccRightThrough:
		return "RRA"
	case OpJumpRelativeNZ:
		return "JR  NZ,r8"
	case OpLoadImmediateHL:
		return "LD  HL,d16"
	case OpLoadIndirectHLAIncrement:
		return "LDI (HL),A"
	case OpIncrementHL:
		return "INC HL"
	case OpIncrementH:
		return "INC H"
	case OpDecrementH:
		return "DEC H"
	case OpLoadImmediateH:
		return "LD  H,d8"
	case OpDecimalToBCD:
		return "DAA"
	case OpJumpRelativeZE:
		return "JR  Z,r8"
	case OpAddDirectHLHL:
		return "ADD HL,HL"
	case OpLoadIndirectAHLIncrement:
		return "LDI A,(HL)"
	case OpDecrementHL:
		return "DEC HL"
	case OpIncrementL:
		return "INC L"
	case OpDecrementL:
		return "DEC L"
	case OpLoadImmediateL:
		return "LD  L,d8"
	case OpInvertA:
		return "CPL"
	case OpJumpRelativeNC:
		return "JR  NC,r8"
	case OpLoadImmediateSP:
		return "LD  SP,d16"
	case OpLoadIndirectHLADecrement:
		return "LDD (HL),A"
	case OpIncrementSP:
		return "INC SP"
	case OpIncrementIndirectHL:
		return "INC (HL)"
	case OpDecrementIndirectHL:
		return "DEC (HL)"
	case OpLoadImmediateIndirectHL:
		return "LD  (HL),d8"
	case OpResetCarry:
		return "SCF"
	case OpJumpRelativeCA:
		return "JR  C,r8"
	case OpAddDirectHLSP:
		return "ADD HL,SP"
	case OpLoadIndirectAHLDecrement:
		return "LDD A,(HL)"
	case OpDecrementSP:
		return "DEC SP"
	case OpIncrementA:
		return "INC A"
	case OpDecrementA:
		return "DEC A"
	case OpLoadImmediateA:
		return "LD  A,d8"
	case OpSetCarry:
		return "CCF"
	case OpLoadDirectBB:
		return "LD B,B"
	case OpLoadDirectBC:
		return "LD B,C"
	case OpLoadDirectBD:
		return "LD B,D"
	case OpLoadDirectBE:
		return "LD B,E"
	case OpLoadDirectBH:
		return "LD B,H"
	case OpLoadDirectBL:
		return "LD B,L"
	case OpLoadIndirectBHL:
		return "LD B,(HL)"
	case OpLoadDirectBA:
		return "LD B,A"
	case OpLoadDirectCB:
		return "LD C,B"
	case OpLoadDirectCC:
		return "LD C,C"
	case OpLoadDirectCD:
		return "LD C,D"
	case OpLoadDirectCE:
		return "LD C,E"
	case OpLoadDirectCH:
		return "LD C,H"
	case OpLoadDirectCL:
		return "LD C,L"
	case OpLoadIndirectCHL:
		return "LD C,(HL)"
	case OpLoadDirectCA:
		return "LD C,A"
	case OpLoadDirectDB:
		return "LD D,B"
	case OpLoadDirectDC:
		return "LD D,C"
	case OpLoadDirectDD:
		return "LD D,D"
	case OpLoadDirectDE:
		return "LD D,E"
	case OpLoadDirectDH:
		return "LD D,H"
	case OpLoadDirectDL:
		return "LD D,L"
	case OpLoadIndirectDHL:
		return "LD D,(HL)"
	case OpLoadDirectDA:
		return "LD D,A"
	case OpLoadDirectEB:
		return "LD E,B"
	case OpLoadDirectEC:
		return "LD E,C"
	case OpLoadDirectED:
		return "LD E,D"
	case OpLoadDirectEE:
		return "LD E,E"
	case OpLoadDirectEH:
		return "LD E,H"
	case OpLoadDirectEL:
		return "LD E,L"
	case OpLoadIndirectEHL:
		return "LD E,(HL)"
	case OpLoadDirectEA:
		return "LD E,A"
	case OpLoadDirectHB:
		return "LD H,B"
	case OpLoadDirectHC:
		return "LD H,C"
	case OpLoadDirectHD:
		return "LD H,D"
	case OpLoadDirectHE:
		return "LD H,E"
	case OpLoadDirectHH:
		return "LD H,H"
	case OpLoadDirectHL:
		return "LD H,L"
	case OpLoadIndirectHHL:
		return "LD H,(HL"
	case OpLoadDirectHA:
		return "LD H,A"
	case OpLoadDirectLB:
		return "LD L,B"
	case OpLoadDirectLC:
		return "LD L,C"
	case OpLoadDirectLD:
		return "LD L,D"
	case OpLoadDirectLE:
		return "LD L,E"
	case OpLoadDirectLH:
		return "LD L,H"
	case OpLoadDirectLL:
		return "LD L,L"
	case OpLoadIndirectLHL:
		return "LD L,(HL)"
	case OpLoadDirectLA:
		return "LD L,A"
	case OpLoadIndirectHLB:
		return "LD (HL),B"
	case OpLoadIndirectHLC:
		return "LD (HL),C"
	case OpLoadIndirectHLD:
		return "LD (HL),D"
	case OpLoadIndirectHLE:
		return "LD (HL),E"
	case OpLoadIndirectHLH:
		return "LD (HL),H"
	case OpLoadIndirectHLL:
		return "LD (HL),L"
	case OpHalt:
		return "HALT"
	case OpLoadIndirectHLA:
		return "LD (HL),A"
	case OpLoadDirectAB:
		return "LD A,B"
	case OpLoadDirectAC:
		return "LD A,C"
	case OpLoadDirectAD:
		return "LD A,D"
	case OpLoadDirectAE:
		return "LD A,E"
	case OpLoadDirectAH:
		return "LD A,H"
	case OpLoadDirectAL:
		return "LD A,L"
	case OpLoadIndirectAHL:
		return "LD A,(HL)"
	case OpLoadDirectAA:
		return "LD A,A"
	case OpAddDirectABNoCarry:
		return "ADD A,B"
	case OpAddDirectACNoCarry:
		return "ADD A,C"
	case OpAddDirectADNoCarry:
		return "ADD A,D"
	case OpAddDirectAENoCarry:
		return "ADD A,E"
	case OpAddDirectAHNoCarry:
		return "ADD A,H"
	case OpAddDirectALNoCarry:
		return "ADD A,L"
	case OpAddIndirectAHLNoCarry:
		return "ADD A,(HL)"
	case OpAddDirectAANoCarry:
		return "ADD A,A"
	case OpAddDirectABCarry:
		return "ADC A,B"
	case OpAddDirectACCarry:
		return "ADC A,C"
	case OpAddDirectADCarry:
		return "ADC A,D"
	case OpAddDirectAECarry:
		return "ADC A,E"
	case OpAddDirectAHCarry:
		return "ADC A,H"
	case OpAddDirectALCarry:
		return "ADC A,L"
	case OpAddIndirectAHLCarry:
		return "ADC A,(HL)"
	case OpAddDirectAAtrue:
		return "ADC A,A"
	case OpSubDirectABNoCarry:
		return "SUB A,B"
	case OpSubDirectACNoCarry:
		return "SUB A,C"
	case OpSubDirectADNoCarry:
		return "SUB A,D"
	case OpSubDirectAENoCarry:
		return "SUB A,E"
	case OpSubDirectAHNoCarry:
		return "SUB A,H"
	case OpSubDirectALNoCarry:
		return "SUB A,L"
	case OpSubIndirectAHLNoCarry:
		return "SUB A,(HL)"
	case OpSubDirectAANoCarry:
		return "SUB A,A"
	case OpSubDirectABCarry:
		return "SBC A,B"
	case OpSubDirectACCarry:
		return "SBC A,C"
	case OpSubDirectADCarry:
		return "SBC A,D"
	case OpSubDirectAECarry:
		return "SBC A,E"
	case OpSubDirectAHCarry:
		return "SBC A,H"
	case OpSubDirectALCarry:
		return "SBC A,L"
	case OpSubIndirectAHLCarry:
		return "SBC A,(HL)"
	case OpSubDirectAACarry:
		return "SBC A,A"
	case OpAndDirectAB:
		return "AND A,B"
	case OpAndDirectAC:
		return "AND A,C"
	case OpAndDirectAD:
		return "AND A,D"
	case OpAndDirectAE:
		return "AND A,E"
	case OpAndDirectAH:
		return "AND A,H"
	case OpAndDirectAL:
		return "AND A,L"
	case OpAndIndirectAHL:
		return "AND A,(HL)"
	case OpAndDirectAA:
		return "AND A,A"
	case OpXorDirectAB:
		return "XOR A,B"
	case OpXorDirectAC:
		return "XOR A,C"
	case OpXorDirectAD:
		return "XOR A,D"
	case OpXorDirectAE:
		return "XOR A,E"
	case OpXorDirectAH:
		return "XOR A,H"
	case OpXorDirectAL:
		return "XOR A,L"
	case OpXorIndirectAHL:
		return "XOR A,(HL)"
	case OpXorDirectAA:
		return "XOR A,A"
	case OpOrDirectAB:
		return "OR  A,B"
	case OpOrDirectAC:
		return "OR  A,C"
	case OpOrDirectAD:
		return "OR  A,D"
	case OpOrDirectAE:
		return "OR  A,E"
	case OpOrDirectAH:
		return "OR  A,H"
	case OpOrDirectAL:
		return "OR  A,L"
	case OpOrIndirectAHL:
		return "OR  A,(HL)"
	case OpOrDirectAA:
		return "OR  A,A"
	case OpCmpDirectAB:
		return "CP  A,B"
	case OpCmpDirectAC:
		return "CP  A,C"
	case OpCmpDirectAD:
		return "CP  A,D"
	case OpCmpDirectAE:
		return "CP  A,E"
	case OpCmpDirectAH:
		return "CP  A,H"
	case OpCmpDirectAL:
		return "CP  A,L"
	case OpCmpIndirectAHL:
		return "CP  A,(HL)"
	case OpCmpDirectAA:
		return "CP  A,A"
	case OpReturnNZ:
		return "RET NZ"
	case OpPopBC:
		return "POP BC"
	case OpJumpAbsoluteNZ:
		return "JP  NZ,a16"
	case OpJumpAbsoluteNO:
		return "JP  a16"
	case OpCallNZ:
		return "CALL NZ,a16"
	case OpPushBC:
		return "PUSH BC"
	case OpAddImmediateANoCarry:
		return "ADD A,d8"
	case OpRestart00:
		return "RST 00h"
	case OpReturnZE:
		return "RET Z"
	case OpReturnNO:
		return "RET"
	case OpJumpAbsoluteZE:
		return "JP  Z,a16"
	case OpCBPrefix:
		return "PREFIX: See cbhandlers below (01xx)"
	case OpCallZE:
		return "CALL Z,a16"
	case OpCallNO:
		return "CALL a16"
	case OpAddImmediateACarry:
		return "ADC A,d8"
	case OpRestart08:
		return "RST 08h"
	case OpReturnNC:
		return "RET NC"
	case OpPopRegDE:
		return "POP DE"
	case OpJumpAbsoluteNC:
		return "JP  NC,a16"
	case OpCallNC:
		return "CALL NC,a16"
	case OpPushDE:
		return "PUSH DE"
	case OpSubImmediateANoCarry:
		return "SUB A,d8"
	case OpRestart10:
		return "RST 10h"
	case OpReturnCA:
		return "RET C"
	case OpRETI:
		return "RETI"
	case OpJumpAbsoluteCA:
		return "JP  C,a16"
	case OpCallCA:
		return "CALL C,a16"
	case OpSubImmediateACarry:
		return "SBC A,d8"
	case OpRestart18:
		return "RST 18h"
	case OpLoadHighAbsA:
		return "LDH (a8),A"
	case OpPopHL:
		return "POP HL"
	case OpLoadHighMemCA:
		return "LD  (C),A"
	case OpPushHL:
		return "PUSH HL"
	case OpAndImmediateA:
		return "AND A,d8"
	case OpRestart20:
		return "RST 20h"
	case OpAddImmediateSignedSP:
		return "ADD SP,r8"
	case OpJumpAbsoluteHL:
		return "JP  (HL)"
	case OpStoreMemA:
		return "LD  (a16),A"
	case OpXorImmediateA:
		return "XOR A,d8"
	case OpRestart28:
		return "RST 28h"
	case OpLoadHighRegA:
		return "LDH A,(a8)"
	case OpPopAF:
		return "POP AF"
	case OpLoadHighRegAC:
		return "LD  A,(C)"
	case OpResetInt:
		return "DI"
	case OpPushAF:
		return "PUSH AF"
	case OpOrImmediateA:
		return "OR  A,d8"
	case OpRestart30:
		return "RES 30h"
	case OpLoadOffsetHLSP:
		return "LD  HL,SP+r8"
	case OpLoadDirectSPHL:
		return "LD  SP,HL"
	case OpLoadMemA:
		return "LD  A,(a16)"
	case OpSetInt:
		return "EI"
	case OpCmpImmediateA:
		return "CP  A,d8"
	case OpRestart38:
		return "RST 38h"
	case OpCbRotateRegBLeftRot:
		return "RLC B"
	case OpCbRotateRegCLeftRot:
		return "RLC C"
	case OpCbRotateRegDLeftRot:
		return "RLC D"
	case OpCbRotateRegELeftRot:
		return "RLC E"
	case OpCbRotateRegHLeftRot:
		return "RLC H"
	case OpCbRotateRegLLeftRot:
		return "RLC L"
	case OpCbRotateIndHLLeftRot:
		return "RLC (HL)"
	case OpCbRotateRegALeftRot:
		return "RLC A"
	case OpCbRotateRegBRightRot:
		return "RRC B"
	case OpCbRotateRegCRightRot:
		return "RRC C"
	case OpCbRotateRegDRightRot:
		return "RRC D"
	case OpCbRotateRegERightRot:
		return "RRC E"
	case OpCbRotateRegHRightRot:
		return "RRC H"
	case OpCbRotateRegLRightRot:
		return "RRC L"
	case OpCbRotateIndHLRightRot:
		return "RRC (HL)"
	case OpCbRotateRegARightRot:
		return "RRC A"
	case OpCbRotateRegBLeftThC:
		return "RL  B"
	case OpCbRotateRegCLeftThC:
		return "RL  C"
	case OpCbRotateRegDLeftThC:
		return "RL  D"
	case OpCbRotateRegELeftThC:
		return "RL  E"
	case OpCbRotateRegHLeftThC:
		return "RL  H"
	case OpCbRotateRegLLeftThC:
		return "RL  L"
	case OpCbRotateIndHLLeftThC:
		return "RL  (HL)"
	case OpCbRotateRegALeftThC:
		return "RL  A"
	case OpCbRotateRegBRightThC:
		return "RR  B"
	case OpCbRotateRegCRightThC:
		return "RR  C"
	case OpCbRotateRegDRightThC:
		return "RR  D"
	case OpCbRotateRegERightThC:
		return "RR  E"
	case OpCbRotateRegHRightThC:
		return "RR  H"
	case OpCbRotateRegLRightThC:
		return "RR  L"
	case OpCbRotateIndHLRightThC:
		return "RR  (HL)"
	case OpCbRotateRegARightThC:
		return "RR  A"
	case OpCbRotateRegBLeftShf:
		return "SLA B"
	case OpCbRotateRegCLeftShf:
		return "SLA C"
	case OpCbRotateRegDLeftShf:
		return "SLA D"
	case OpCbRotateRegELeftShf:
		return "SLA E"
	case OpCbRotateRegHLeftShf:
		return "SLA H"
	case OpCbRotateRegLLeftShf:
		return "SLA L"
	case OpCbRotateIndHLLeftShf:
		return "SLA (HL)"
	case OpCbRotateRegALeftShf:
		return "SLA A"
	case OpCbRotateRegBRightRep:
		return "SRA B"
	case OpCbRotateRegCRightRep:
		return "SRA C"
	case OpCbRotateRegDRightRep:
		return "SRA D"
	case OpCbRotateRegERightRep:
		return "SRA E"
	case OpCbRotateRegHRightRep:
		return "SRA H"
	case OpCbRotateRegLRightRep:
		return "SRA L"
	case OpCbRotateIndHLRightRep:
		return "SRA (HL)"
	case OpCbRotateRegARightRep:
		return "SRA A"
	case OpCbSwapDirectB:
		return "SWAP B"
	case OpCbSwapDirectC:
		return "SWAP C"
	case OpCbSwapDirectD:
		return "SWAP D"
	case OpCbSwapDirectE:
		return "SWAP E"
	case OpCbSwapDirectH:
		return "SWAP H"
	case OpCbSwapDirectL:
		return "SWAP L"
	case OpCbSwapIndirectHL:
		return "SWAP (HL)"
	case OpCbSwapDirectA:
		return "SWAP A"
	case OpCbRotateRegBRightShf:
		return "SRL B"
	case OpCbRotateRegCRightShf:
		return "SRL C"
	case OpCbRotateRegDRightShf:
		return "SRL D"
	case OpCbRotateRegERightShf:
		return "SRL E"
	case OpCbRotateRegHRightShf:
		return "SRL H"
	case OpCbRotateRegLRightShf:
		return "SRL L"
	case OpCbRotateIndHLRightShf:
		return "SRL (HL)"
	case OpCbRotateRegARightShf:
		return "SRL A"
	case OpCbBitDirectB0:
		return "BIT 0,B"
	case OpCbBitDirectC0:
		return "BIT 0,C"
	case OpCbBitDirectD0:
		return "BIT 0,D"
	case OpCbBitDirectE0:
		return "BIT 0,E"
	case OpCbBitDirectH0:
		return "BIT 0,H"
	case OpCbBitDirectL0:
		return "BIT 0,L"
	case OpCbBitIndirectHL0:
		return "BIT 0,(HL)"
	case OpCbBitDirectA0:
		return "BIT 0,A"
	case OpCbBitDirectB1:
		return "BIT 1,B"
	case OpCbBitDirectC1:
		return "BIT 1,C"
	case OpCbBitDirectD1:
		return "BIT 1,D"
	case OpCbBitDirectE1:
		return "BIT 1,E"
	case OpCbBitDirectH1:
		return "BIT 1,H"
	case OpCbBitDirectL1:
		return "BIT 1,L"
	case OpCbBitIndirectHL1:
		return "BIT 1,(HL)"
	case OpCbBitDirectA1:
		return "BIT 1,A"
	case OpCbBitDirectB2:
		return "BIT 2,B"
	case OpCbBitDirectC2:
		return "BIT 2,C"
	case OpCbBitDirectD2:
		return "BIT 2,D"
	case OpCbBitDirectE2:
		return "BIT 2,E"
	case OpCbBitDirectH2:
		return "BIT 2,H"
	case OpCbBitDirectL2:
		return "BIT 2,L"
	case OpCbBitIndirectHL2:
		return "BIT 2,(HL)"
	case OpCbBitDirectA2:
		return "BIT 2,A"
	case OpCbBitDirectB3:
		return "BIT 3,B"
	case OpCbBitDirectC3:
		return "BIT 3,C"
	case OpCbBitDirectD3:
		return "BIT 3,D"
	case OpCbBitDirectE3:
		return "BIT 3,E"
	case OpCbBitDirectH3:
		return "BIT 3,H"
	case OpCbBitDirectL3:
		return "BIT 3,L"
	case OpCbBitIndirectHL3:
		return "BIT 3,(HL)"
	case OpCbBitDirectA3:
		return "BIT 3,A"
	case OpCbBitDirectB4:
		return "BIT 4,B"
	case OpCbBitDirectC4:
		return "BIT 4,C"
	case OpCbBitDirectD4:
		return "BIT 4,D"
	case OpCbBitDirectE4:
		return "BIT 4,E"
	case OpCbBitDirectH4:
		return "BIT 4,H"
	case OpCbBitDirectL4:
		return "BIT 4,L"
	case OpCbBitIndirectHL4:
		return "BIT 4,(HL)"
	case OpCbBitDirectA4:
		return "BIT 4,A"
	case OpCbBitDirectB5:
		return "BIT 5,B"
	case OpCbBitDirectC5:
		return "BIT 5,C"
	case OpCbBitDirectD5:
		return "BIT 5,D"
	case OpCbBitDirectE5:
		return "BIT 5,E"
	case OpCbBitDirectH5:
		return "BIT 5,H"
	case OpCbBitDirectL5:
		return "BIT 5,L"
	case OpCbBitIndirectHL5:
		return "BIT 5,(HL)"
	case OpCbBitDirectA5:
		return "BIT 5,A"
	case OpCbBitDirectB6:
		return "BIT 6,B"
	case OpCbBitDirectC6:
		return "BIT 6,C"
	case OpCbBitDirectD6:
		return "BIT 6,D"
	case OpCbBitDirectE6:
		return "BIT 6,E"
	case OpCbBitDirectH6:
		return "BIT 6,H"
	case OpCbBitDirectL6:
		return "BIT 6,L"
	case OpCbBitIndirectHL6:
		return "BIT 6,(HL)"
	case OpCbBitDirectA6:
		return "BIT 6,A"
	case OpCbBitDirectB7:
		return "BIT 7,B"
	case OpCbBitDirectC7:
		return "BIT 7,C"
	case OpCbBitDirectD7:
		return "BIT 7,D"
	case OpCbBitDirectE7:
		return "BIT 7,E"
	case OpCbBitDirectH7:
		return "BIT 7,H"
	case OpCbBitDirectL7:
		return "BIT 7,L"
	case OpCbBitIndirectHL7:
		return "BIT 7,(HL)"
	case OpCbBitDirectA7:
		return "BIT 7,A"
	case OpCbResetDirectB0:
		return "RES 0,B"
	case OpCbResetDirectC0:
		return "RES 0,C"
	case OpCbResetDirectD0:
		return "RES 0,D"
	case OpCbResetDirectE0:
		return "RES 0,E"
	case OpCbResetDirectH0:
		return "RES 0,H"
	case OpCbResetDirectL0:
		return "RES 0,L"
	case OpCbResetIndirectHL0:
		return "RES 0,(HL)"
	case OpCbResetDirectA0:
		return "RES 0,A"
	case OpCbResetDirectB1:
		return "RES 1,B"
	case OpCbResetDirectC1:
		return "RES 1,C"
	case OpCbResetDirectD1:
		return "RES 1,D"
	case OpCbResetDirectE1:
		return "RES 1,E"
	case OpCbResetDirectH1:
		return "RES 1,H"
	case OpCbResetDirectL1:
		return "RES 1,L"
	case OpCbResetIndirectHL1:
		return "RES 1,(HL)"
	case OpCbResetDirectA1:
		return "RES 1,A"
	case OpCbResetDirectB2:
		return "RES 2,B"
	case OpCbResetDirectC2:
		return "RES 2,C"
	case OpCbResetDirectD2:
		return "RES 2,D"
	case OpCbResetDirectE2:
		return "RES 2,E"
	case OpCbResetDirectH2:
		return "RES 2,H"
	case OpCbResetDirectL2:
		return "RES 2,L"
	case OpCbResetIndirectHL2:
		return "RES 2,(HL)"
	case OpCbResetDirectA2:
		return "RES 2,A"
	case OpCbResetDirectB3:
		return "RES 3,B"
	case OpCbResetDirectC3:
		return "RES 3,C"
	case OpCbResetDirectD3:
		return "RES 3,D"
	case OpCbResetDirectE3:
		return "RES 3,E"
	case OpCbResetDirectH3:
		return "RES 3,H"
	case OpCbResetDirectL3:
		return "RES 3,L"
	case OpCbResetIndirectHL3:
		return "RES 3,(HL)"
	case OpCbResetDirectA3:
		return "RES 3,A"
	case OpCbResetDirectB4:
		return "RES 4,B"
	case OpCbResetDirectC4:
		return "RES 4,C"
	case OpCbResetDirectD4:
		return "RES 4,D"
	case OpCbResetDirectE4:
		return "RES 4,E"
	case OpCbResetDirectH4:
		return "RES 4,H"
	case OpCbResetDirectL4:
		return "RES 4,L"
	case OpCbResetIndirectHL4:
		return "RES 4,(HL)"
	case OpCbResetDirectA4:
		return "RES 4,A"
	case OpCbResetDirectB5:
		return "RES 5,B"
	case OpCbResetDirectC5:
		return "RES 5,C"
	case OpCbResetDirectD5:
		return "RES 5,D"
	case OpCbResetDirectE5:
		return "RES 5,E"
	case OpCbResetDirectH5:
		return "RES 5,H"
	case OpCbResetDirectL5:
		return "RES 5,L"
	case OpCbResetIndirectHL5:
		return "RES 5,(HL)"
	case OpCbResetDirectA5:
		return "RES 5,A"
	case OpCbResetDirectB6:
		return "RES 6,B"
	case OpCbResetDirectC6:
		return "RES 6,C"
	case OpCbResetDirectD6:
		return "RES 6,D"
	case OpCbResetDirectE6:
		return "RES 6,E"
	case OpCbResetDirectH6:
		return "RES 6,H"
	case OpCbResetDirectL6:
		return "RES 6,L"
	case OpCbResetIndirectHL6:
		return "RES 6,(HL)"
	case OpCbResetDirectA6:
		return "RES 6,A"
	case OpCbResetDirectB7:
		return "RES 7,B"
	case OpCbResetDirectC7:
		return "RES 7,C"
	case OpCbResetDirectD7:
		return "RES 7,D"
	case OpCbResetDirectE7:
		return "RES 7,E"
	case OpCbResetDirectH7:
		return "RES 7,H"
	case OpCbResetDirectL7:
		return "RES 7,L"
	case OpCbResetIndirectHL7:
		return "RES 7,(HL)"
	case OpCbResetDirectA7:
		return "RES 7,A"
	case OpCbSetDirectB0:
		return "SET 0,B"
	case OpCbSetDirectC0:
		return "SET 0,C"
	case OpCbSetDirectD0:
		return "SET 0,D"
	case OpCbSetDirectE0:
		return "SET 0,E"
	case OpCbSetDirectH0:
		return "SET 0,H"
	case OpCbSetDirectL0:
		return "SET 0,L"
	case OpCbSetIndirectHL0:
		return "SET 0,(HL)"
	case OpCbSetDirectA0:
		return "SET 0,A"
	case OpCbSetDirectB1:
		return "SET 1,B"
	case OpCbSetDirectC1:
		return "SET 1,C"
	case OpCbSetDirectD1:
		return "SET 1,D"
	case OpCbSetDirectE1:
		return "SET 1,E"
	case OpCbSetDirectH1:
		return "SET 1,H"
	case OpCbSetDirectL1:
		return "SET 1,L"
	case OpCbSetIndirectHL1:
		return "SET 1,(HL)"
	case OpCbSetDirectA1:
		return "SET 1,A"
	case OpCbSetDirectB2:
		return "SET 2,B"
	case OpCbSetDirectC2:
		return "SET 2,C"
	case OpCbSetDirectD2:
		return "SET 2,D"
	case OpCbSetDirectE2:
		return "SET 2,E"
	case OpCbSetDirectH2:
		return "SET 2,H"
	case OpCbSetDirectL2:
		return "SET 2,L"
	case OpCbSetIndirectHL2:
		return "SET 2,(HL)"
	case OpCbSetDirectA2:
		return "SET 2,A"
	case OpCbSetDirectB3:
		return "SET 3,B"
	case OpCbSetDirectC3:
		return "SET 3,C"
	case OpCbSetDirectD3:
		return "SET 3,D"
	case OpCbSetDirectE3:
		return "SET 3,E"
	case OpCbSetDirectH3:
		return "SET 3,H"
	case OpCbSetDirectL3:
		return "SET 3,L"
	case OpCbSetIndirectHL3:
		return "SET 3,(HL)"
	case OpCbSetDirectA3:
		return "SET 3,A"
	case OpCbSetDirectB4:
		return "SET 4,B"
	case OpCbSetDirectC4:
		return "SET 4,C"
	case OpCbSetDirectD4:
		return "SET 4,D"
	case OpCbSetDirectE4:
		return "SET 4,E"
	case OpCbSetDirectH4:
		return "SET 4,H"
	case OpCbSetDirectL4:
		return "SET 4,L"
	case OpCbSetIndirectHL4:
		return "SET 4,(HL)"
	case OpCbSetDirectA4:
		return "SET 4,A"
	case OpCbSetDirectB5:
		return "SET 5,B"
	case OpCbSetDirectC5:
		return "SET 5,C"
	case OpCbSetDirectD5:
		return "SET 5,D"
	case OpCbSetDirectE5:
		return "SET 5,E"
	case OpCbSetDirectH5:
		return "SET 5,H"
	case OpCbSetDirectL5:
		return "SET 5,L"
	case OpCbSetIndirectHL5:
		return "SET 5,(HL)"
	case OpCbSetDirectA5:
		return "SET 5,A"
	case OpCbSetDirectB6:
		return "SET 6,B"
	case OpCbSetDirectC6:
		return "SET 6,C"
	case OpCbSetDirectD6:
		return "SET 6,D"
	case OpCbSetDirectE6:
		return "SET 6,E"
	case OpCbSetDirectH6:
		return "SET 6,H"
	case OpCbSetDirectL6:
		return "SET 6,L"
	case OpCbSetIndirectHL6:
		return "SET 6,(HL)"
	case OpCbSetDirectA6:
		return "SET 6,A"
	case OpCbSetDirectB7:
		return "SET 7,B"
	case OpCbSetDirectC7:
		return "SET 7,C"
	case OpCbSetDirectD7:
		return "SET 7,D"
	case OpCbSetDirectE7:
		return "SET 7,E"
	case OpCbSetDirectH7:
		return "SET 7,H"
	case OpCbSetDirectL7:
		return "SET 7,L"
	case OpCbSetIndirectHL7:
		return "SET 7,(HL)"
	case OpCbSetDirectA7:
		return "SET 7,A"
	}
	panic("invalid opcode")
}
