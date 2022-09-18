package emulator

type MemoryMode uint8

const (
	Implicit MemoryMode = iota
	Accumulator
	Immediate
	ZeroPage
	ZeroPageX
	ZeroPageY
	Relative
	Absolute
	AbsoluteX
	AbsoluteY
	Indirect
	IndirectX
	IndirectY
)

func (m MemoryMode) Size() uint16 {
	switch m {
	case Accumulator, Implicit:
		return 0
	case Immediate, ZeroPage, ZeroPageX, ZeroPageY, IndirectX, IndirectY:
		return 1
	case Absolute, AbsoluteX, AbsoluteY, Indirect:
		return 2
	case Relative:
		return 1
	}

	panic("unreachable")
}

func (m MemoryMode) String() string {
	switch m {
	case Implicit:
		return "implicit"
	case Accumulator:
		return "accumulator"
	case Immediate:
		return "immediate"
	case ZeroPage:
		return "zero page"
	case ZeroPageX:
		return "zero page, X"
	case ZeroPageY:
		return "zero page, Y"
	case Relative:
		return "relative"
	case Absolute:
		return "absolute"
	case AbsoluteX:
		return "absolute, X"
	case AbsoluteY:
		return "absolute, Y"
	case Indirect:
		return "indirect"
	case IndirectX:
		return "indirect, X"
	case IndirectY:
		return "indirect, Y"
	}

	panic("unreachable")
}

type Instruction uint8

const (
	_ Instruction = iota
	ADC
	AND
	ASL
	BCC
	BCS
	BEQ
	BIT
	BMI
	BNE
	BPL
	BRK
	BVC
	BVS
	CLC
	CLD
	CLI
	CLV
	CMP
	CPX
	CPY
	DEC
	DEX
	DEY
	EOR
	INC
	INX
	INY
	JMP
	JSR
	LDA
	LDX
	LDY
	LSR
	NOP
	ORA
	PHA
	PHP
	PLA
	PLP
	ROL
	ROR
	RTI
	RTS
	SBC
	SEC
	SED
	SEI
	STA
	STX
	STY
	TAX
	TAY
	TSX
	TXA
	TXS
	TYA
)

func (i Instruction) String() string {
	switch i {
	case ADC:
		return "ADC"
	case AND:
		return "AND"
	case ASL:
		return "ASL"
	case BCC:
		return "BCC"
	case BCS:
		return "BCS"
	case BEQ:
		return "BEQ"
	case BIT:
		return "BIT"
	case BMI:
		return "BMI"
	case BNE:
		return "BNE"
	case BPL:
		return "BPL"
	case BRK:
		return "BRK"
	case BVC:
		return "BVC"
	case BVS:
		return "BVS"
	case CLC:
		return "CLC"
	case CLD:
		return "CLD"
	case CLI:
		return "CLI"
	case CLV:
		return "CLV"
	case CMP:
		return "CMP"
	case CPX:
		return "CPX"
	case CPY:
		return "CPY"
	case DEC:
		return "DEC"
	case DEX:
		return "DEX"
	case DEY:
		return "DEY"
	case EOR:
		return "EOR"
	case INC:
		return "INC"
	case INX:
		return "INX"
	case INY:
		return "INY"
	case JMP:
		return "JMP"
	case JSR:
		return "JSR"
	case LDA:
		return "LDA"
	case LDX:
		return "LDX"
	case LDY:
		return "LDY"
	case LSR:
		return "LSR"
	case NOP:
		return "NOP"
	case ORA:
		return "ORA"
	case PHA:
		return "PHA"
	case PHP:
		return "PHP"
	case PLA:
		return "PLA"
	case PLP:
		return "PLP"
	case ROL:
		return "ROL"
	case ROR:
		return "ROR"
	case RTI:
		return "RTI"
	case RTS:
		return "RTS"
	case SBC:
		return "SBC"
	case SEC:
		return "SEC"
	case SED:
		return "SED"
	case SEI:
		return "SEI"
	case STA:
		return "STA"
	case STX:
		return "STX"
	case STY:
		return "STY"
	case TAX:
		return "TAX"
	case TAY:
		return "TAY"
	case TSX:
		return "TSX"
	case TXA:
		return "TXA"
	case TXS:
		return "TXS"
	case TYA:
		return "TYA"
	}

	panic("unreachable")
}

type OpCode struct {
	Instruction Instruction
	MemoryMode  MemoryMode
	Cycles      int
}

var OpCodeMap = map[uint8]OpCode{
	0x69: {ADC, Immediate, 2}, // ADC
	0x65: {ADC, ZeroPage, 3},  // ADC
	0x75: {ADC, ZeroPageX, 4}, // ADC
	0x6D: {ADC, Absolute, 4},  // ADC
	0x7D: {ADC, AbsoluteX, 4}, // ADC
	0x79: {ADC, AbsoluteY, 4}, // ADC
	0x61: {ADC, IndirectX, 6}, // ADC
	0x71: {ADC, IndirectY, 5}, // ADC

	0x29: {AND, Immediate, 2}, // AND
	0x25: {AND, ZeroPage, 3},  // AND
	0x35: {AND, ZeroPageX, 4}, // AND
	0x2D: {AND, Absolute, 4},  // AND
	0x3D: {AND, AbsoluteX, 4}, // AND
	0x39: {AND, AbsoluteY, 4}, // AND
	0x21: {AND, IndirectX, 6}, // AND
	0x31: {AND, IndirectY, 5}, // AND

	0x0A: {ASL, Accumulator, 2}, // ASL
	0x06: {ASL, ZeroPage, 5},    // ASL
	0x16: {ASL, ZeroPageX, 6},   // ASL
	0x0E: {ASL, Absolute, 6},    // ASL
	0x1E: {ASL, AbsoluteX, 7},   // ASL

	0x90: {BCC, Relative, 2}, // BCC

	0xB0: {BCS, Relative, 2}, // BCS

	0xF0: {BEQ, Relative, 2}, // BEQ

	0x24: {BIT, ZeroPage, 3}, // BIT
	0x2C: {BIT, Absolute, 4}, // BIT

	0x30: {BMI, Relative, 2}, // BMI

	0xD0: {BNE, Relative, 2}, // BNE

	0x10: {BPL, Relative, 2}, // BPL

	0x00: {BRK, Implicit, 7}, // BRK

	0x50: {BVC, Relative, 2}, // BVC

	0x70: {BVS, Relative, 2}, // BVS

	0x18: {CLC, Implicit, 2}, // CLC

	0xD8: {CLD, Implicit, 2}, // CLD

	0x58: {CLI, Implicit, 2}, // CLI

	0xB8: {CLV, Implicit, 2}, // CLV

	0xC9: {CMP, Immediate, 2}, // CMP
	0xC5: {CMP, ZeroPage, 3},  // CMP
	0xD5: {CMP, ZeroPageX, 4}, // CMP
	0xCD: {CMP, Absolute, 4},  // CMP
	0xDD: {CMP, AbsoluteX, 4}, // CMP
	0xD9: {CMP, AbsoluteY, 4}, // CMP
	0xC1: {CMP, IndirectX, 6}, // CMP
	0xD1: {CMP, IndirectY, 5}, // CMP

	0xE0: {CPX, Immediate, 2}, // CPX
	0xE4: {CPX, ZeroPage, 3},  // CPX
	0xEC: {CPX, Absolute, 4},  // CPX

	0xC0: {CPY, Immediate, 2}, // CPY
	0xC4: {CPY, ZeroPage, 3},  // CPY
	0xCC: {CPY, Absolute, 4},  // CPY

	0xC6: {DEC, ZeroPage, 5},  // DEC
	0xD6: {DEC, ZeroPageX, 6}, // DEC
	0xCE: {DEC, Absolute, 6},  // DEC
	0xDE: {DEC, AbsoluteX, 7}, // DEC

	0xCA: {DEX, Implicit, 2}, // DEX

	0x88: {DEY, Implicit, 2}, // DEY

	0x49: {EOR, Immediate, 2}, // EOR
	0x45: {EOR, ZeroPage, 3},  // EOR
	0x55: {EOR, ZeroPageX, 4}, // EOR
	0x4D: {EOR, Absolute, 4},  // EOR
	0x5D: {EOR, AbsoluteX, 4}, // EOR
	0x59: {EOR, AbsoluteY, 4}, // EOR
	0x41: {EOR, IndirectX, 6}, // EOR
	0x51: {EOR, IndirectY, 5}, // EOR

	0xE6: {INC, ZeroPage, 5},  // INC
	0xF6: {INC, ZeroPageX, 6}, // INC
	0xEE: {INC, Absolute, 6},  // INC
	0xFE: {INC, AbsoluteX, 7}, // INC

	0xE8: {INX, Implicit, 2}, // INX

	0xC8: {INY, Implicit, 2}, // INY

	0x4C: {JMP, Absolute, 3}, // JMP

	0x6C: {JMP, Indirect, 5}, // JMP

	0x20: {JSR, Absolute, 6}, // JSR

	0xA9: {LDA, Immediate, 2}, // LDA
	0xA5: {LDA, ZeroPage, 3},  // LDA
	0xB5: {LDA, ZeroPageX, 4}, // LDA
	0xAD: {LDA, Absolute, 4},  // LDA
	0xBD: {LDA, AbsoluteX, 4}, // LDA
	0xB9: {LDA, AbsoluteY, 4}, // LDA
	0xA1: {LDA, IndirectX, 6}, // LDA
	0xB1: {LDA, IndirectY, 5}, // LDA

	0xA2: {LDX, Immediate, 2}, // LDX
	0xA6: {LDX, ZeroPage, 3},  // LDX
	0xB6: {LDX, ZeroPageY, 4}, // LDX
	0xAE: {LDX, Absolute, 4},  // LDX
	0xBE: {LDX, AbsoluteY, 4}, // LDX

	0xA0: {LDY, Immediate, 2}, // LDY
	0xA4: {LDY, ZeroPage, 3},  // LDY
	0xB4: {LDY, ZeroPageX, 4}, // LDY
	0xAC: {LDY, Absolute, 4},  // LDY
	0xBC: {LDY, AbsoluteX, 4}, // LDY

	0x4A: {LSR, Accumulator, 2}, // LSR
	0x46: {LSR, ZeroPage, 5},    // LSR
	0x56: {LSR, ZeroPageX, 6},   // LSR
	0x4E: {LSR, Absolute, 6},    // LSR
	0x5E: {LSR, AbsoluteX, 7},   // LSR

	0xEA: {NOP, Implicit, 2}, // NOP

	0x09: {ORA, Immediate, 2}, // ORA
	0x05: {ORA, ZeroPage, 3},  // ORA
	0x15: {ORA, ZeroPageX, 4}, // ORA
	0x0D: {ORA, Absolute, 4},  // ORA
	0x1D: {ORA, AbsoluteX, 4}, // ORA
	0x19: {ORA, AbsoluteY, 4}, // ORA
	0x01: {ORA, IndirectX, 6}, // ORA
	0x11: {ORA, IndirectY, 5}, // ORA

	0x48: {PHA, Implicit, 3}, // PHA

	0x08: {PHP, Implicit, 3}, // PHP

	0x68: {PLA, Implicit, 4}, // PLA

	0x28: {PLP, Implicit, 4}, // PLP

	0x2A: {ROL, Accumulator, 2}, // ROL
	0x26: {ROL, ZeroPage, 5},    // ROL
	0x36: {ROL, ZeroPageX, 6},   // ROL
	0x2E: {ROL, Absolute, 6},    // ROL
	0x3E: {ROL, AbsoluteX, 7},   // ROL

	0x6A: {ROR, Accumulator, 2}, // ROR
	0x66: {ROR, ZeroPage, 5},    // ROR
	0x76: {ROR, ZeroPageX, 6},   // ROR
	0x6E: {ROR, Absolute, 6},    // ROR
	0x7E: {ROR, AbsoluteX, 7},   // ROR

	0x40: {RTI, Implicit, 6}, // RTI

	0x60: {RTS, Implicit, 6}, // RTS

	0xE9: {SBC, Immediate, 2}, // SBC

	0xE5: {SBC, ZeroPage, 3},  // SBC
	0xF5: {SBC, ZeroPageX, 4}, // SBC
	0xED: {SBC, Absolute, 4},  // SBC
	0xFD: {SBC, AbsoluteX, 4}, // SBC
	0xF9: {SBC, AbsoluteY, 4}, // SBC
	0xE1: {SBC, IndirectX, 6}, // SBC
	0xF1: {SBC, IndirectY, 5}, // SBC

	0x38: {SEC, Implicit, 2}, // SEC

	0xF8: {SED, Implicit, 2}, // SED

	0x78: {SEI, Implicit, 2}, // SEI

	0x85: {STA, ZeroPage, 3},  // STA
	0x95: {STA, ZeroPageX, 4}, // STA
	0x8D: {STA, Absolute, 4},  // STA
	0x9D: {STA, AbsoluteX, 5}, // STA
	0x99: {STA, AbsoluteY, 5}, // STA
	0x81: {STA, IndirectX, 6}, // STA
	0x91: {STA, IndirectY, 6}, // STA

	0x86: {STX, ZeroPage, 3},  // STX
	0x96: {STX, ZeroPageY, 4}, // STX
	0x8E: {STX, Absolute, 4},  // STX

	0x84: {STY, ZeroPage, 3},  // STY
	0x94: {STY, ZeroPageX, 4}, // STY
	0x8C: {STY, Absolute, 4},  // STY

	0xAA: {TAX, Implicit, 2}, // TAX

	0xA8: {TAY, Implicit, 2}, // TAY

	0xBA: {TSX, Implicit, 2}, // TSX

	0x8A: {TXA, Implicit, 2}, // TXA

	0x9A: {TXS, Implicit, 2}, // TXS

	0x98: {TYA, Implicit, 2}, // TYA
}
