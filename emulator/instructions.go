package emulator

type iInstructionSet struct{}

func (cpu *CPU) ProcessInstruction(opcode uint8) {
	opCode := OpCodeMap[opcode]

	// fmt.Printf("%04x - %v\n", cpu.programCounter-1, opCode)

	cpu.programCounter += opCode.MemoryMode.Size()

	switch opCode.Instruction {
	case ADC:
		cpu.instructionSet.ADC(cpu, opCode.MemoryMode)
	case AND:
		cpu.instructionSet.AND(cpu, opCode.MemoryMode)
	case ASL:
		cpu.instructionSet.ASL(cpu, opCode.MemoryMode)
	case BCC:
		cpu.instructionSet.BCC(cpu)
	case BCS:
		cpu.instructionSet.BCS(cpu)
	case BEQ:
		cpu.instructionSet.BEQ(cpu)
	case BIT:
		cpu.instructionSet.BIT(cpu, opCode.MemoryMode)
	case BMI:
		cpu.instructionSet.BMI(cpu)
	case BNE:
		cpu.instructionSet.BNE(cpu)
	case BPL:
		cpu.instructionSet.BPL(cpu)
	case BRK:
		cpu.instructionSet.BRK(cpu)
	case BVC:
		cpu.instructionSet.BVC(cpu)
	case BVS:
		cpu.instructionSet.BVS(cpu)
	case CLC:
		cpu.instructionSet.CLC(cpu)
	case CLD:
		cpu.instructionSet.CLD(cpu)
	case CLI:
		cpu.instructionSet.CLI(cpu)
	case CLV:
		cpu.instructionSet.CLV(cpu)
	case CMP:
		cpu.instructionSet.CMP(cpu, opCode.MemoryMode)
	case CPX:
		cpu.instructionSet.CPX(cpu, opCode.MemoryMode)
	case CPY:
		cpu.instructionSet.CPY(cpu, opCode.MemoryMode)
	case DEC:
		cpu.instructionSet.DEC(cpu, opCode.MemoryMode)
	case DEX:
		cpu.instructionSet.DEX(cpu)
	case DEY:
		cpu.instructionSet.DEY(cpu)
	case EOR:
		cpu.instructionSet.EOR(cpu, opCode.MemoryMode)
	case INC:
		cpu.instructionSet.INC(cpu, opCode.MemoryMode)
	case INX:
		cpu.instructionSet.INX(cpu)
	case INY:
		cpu.instructionSet.INY(cpu)
	case JMP:
		cpu.instructionSet.JMP(cpu, opCode.MemoryMode)
	case JSR:
		cpu.instructionSet.JSR(cpu)
	case LDA:
		cpu.instructionSet.LDA(cpu, opCode.MemoryMode)
	case LDX:
		cpu.instructionSet.LDX(cpu, opCode.MemoryMode)
	case LDY:
		cpu.instructionSet.LDY(cpu, opCode.MemoryMode)
	case LSR:
		cpu.instructionSet.LSR(cpu, opCode.MemoryMode)
	case NOP:
		cpu.instructionSet.NOP(cpu)
	case ORA:
		cpu.instructionSet.ORA(cpu, opCode.MemoryMode)
	case PHA:
		cpu.instructionSet.PHA(cpu)
	case PHP:
		cpu.instructionSet.PHP(cpu)
	case PLA:
		cpu.instructionSet.PLA(cpu)
	case PLP:
		cpu.instructionSet.PLP(cpu)
	case ROL:
		cpu.instructionSet.ROL(cpu, opCode.MemoryMode)
	case ROR:
		cpu.instructionSet.ROR(cpu, opCode.MemoryMode)
	case RTI:
		cpu.instructionSet.RTI(cpu)
	case RTS:
		cpu.instructionSet.RTS(cpu)
	case SBC:
		cpu.instructionSet.SBC(cpu, opCode.MemoryMode)
	case SEC:
		cpu.instructionSet.SEC(cpu)
	case SED:
		cpu.instructionSet.SED(cpu)
	case SEI:
		cpu.instructionSet.SEI(cpu)
	case STA:
		cpu.instructionSet.STA(cpu, opCode.MemoryMode)
	case STX:
		cpu.instructionSet.STX(cpu, opCode.MemoryMode)
	case STY:
		cpu.instructionSet.STY(cpu, opCode.MemoryMode)
	case TAX:
		cpu.instructionSet.TAX(cpu)
	case TAY:
		cpu.instructionSet.TAY(cpu)
	case TSX:
		cpu.instructionSet.TSX(cpu)
	case TXA:
		cpu.instructionSet.TXA(cpu)
	case TXS:
		cpu.instructionSet.TXS(cpu)
	case TYA:
		cpu.instructionSet.TYA(cpu)
	}

	// For emulation purposes, we need to delay the return by the number of cycles the instruction takes
	for i := 0; i < opCode.Cycles; i++ {
		cpu.Pulse()
	}
}

func (i iInstructionSet) MemoryMode(cpu *CPU, mode MemoryMode, read bool) (uint8, uint16, bool) {
	var (
		val         uint8
		address     uint16
		accumulator bool
	)

	switch mode {
	case Implicit: // No memory access
	case Accumulator: // A Register (Accumulator)
		val = cpu.registers.A
		accumulator = true
	case Immediate: // Immediate
		//	the next byte is the value to add
		address = cpu.programCounter - mode.Size()
		if read {
			val = cpu.bus.Read(address)
		}
	case ZeroPage: // Zero page
		// the next byte is the lower bits of the address of the value to add
		address = uint16(cpu.bus.Read(cpu.programCounter - mode.Size()))
		if read {
			val = cpu.bus.Read(address)
		}
	case ZeroPageX: // Zero page, X
		// the next byte is the lower bits of the address of the value to add, offset by X
		address = uint16(cpu.bus.Read(cpu.programCounter-mode.Size()) + cpu.registers.X)
		if read {
			val = cpu.bus.Read(address)
		}
	case ZeroPageY: // Zero page, Y
		// the next byte is the lower bits of the address of the value to add, offset by X
		address = uint16(cpu.bus.Read(cpu.programCounter-mode.Size()) + cpu.registers.Y)
		if read {
			val = cpu.bus.Read(address)
		}
	case Relative: // Relative
		address = cpu.programCounter - mode.Size()
		if read {
			val = cpu.bus.Read(address) // this value is signed and will be added to the program counter
		}
	case Absolute: // Absolute
		// the next two bytes are the address of the value to add
		address = uint16(cpu.bus.Read(cpu.programCounter-mode.Size())) | uint16(cpu.bus.Read(cpu.programCounter-mode.Size()+1))<<8
		if read {
			val = cpu.bus.Read(address)
		}
	case AbsoluteX: // Absolute, X
		// the next two bytes are the address of the value to add, offset by X
		address = (uint16(cpu.bus.Read(cpu.programCounter-mode.Size())) | uint16(cpu.bus.Read(cpu.programCounter-mode.Size()+1))<<8) + uint16(cpu.registers.X)
		if read {
			val = cpu.bus.Read(address)
		}
	case AbsoluteY: // Absolute, Y
		// the next two bytes are the address of the value to add, offset by Y
		address = (uint16(cpu.bus.Read(cpu.programCounter-mode.Size())) | uint16(cpu.bus.Read(cpu.programCounter-mode.Size()+1))<<8) + uint16(cpu.registers.Y)
		if read {
			val = cpu.bus.Read(address)
		}
	case IndirectX: // Indirect, X
		// the next byte is the lower bits of the address of the value to add, offset by X
		tmp := uint16(cpu.bus.Read(cpu.programCounter-mode.Size())) + uint16(cpu.registers.X)
		address = uint16(cpu.bus.Read(tmp)) | uint16(cpu.bus.Read(tmp+1))<<8
		if read {
			val = cpu.bus.Read(address)
		}
	case IndirectY: // Indirect, Y
		tmp := uint16(cpu.bus.Read(cpu.programCounter - mode.Size()))
		address = (uint16(cpu.bus.Read(tmp)) | uint16(cpu.bus.Read(tmp+1))<<8) + uint16(cpu.registers.Y)
		if read {
			val = cpu.bus.Read(address)
		}
	}

	return val, address, accumulator
}

func (i iInstructionSet) ADC(cpu *CPU, mode MemoryMode) {
	val, _, _ := i.MemoryMode(cpu, mode, true)

	cpu.registers.A += val
	if cpu.flags.Carry {
		cpu.registers.A++
	}

	cpu.flags.Carry = cpu.registers.A < val
	cpu.flags.Zero = cpu.registers.A == 0
	cpu.flags.Negative = cpu.registers.A&0x80 != 0
	cpu.flags.Overflow = cpu.registers.A&0x80 != val&0x80
}

func (i iInstructionSet) AND(cpu *CPU, mode MemoryMode) {
	val, _, _ := i.MemoryMode(cpu, mode, true)

	cpu.registers.A = cpu.registers.A & val
	if cpu.registers.A == 0 {
		cpu.flags.Zero = true
	}
	if cpu.registers.A > 0x7F {
		cpu.flags.Negative = true
	}
}

func (i iInstructionSet) ASL(cpu *CPU, mode MemoryMode) {
	val, addr, accumulator := i.MemoryMode(cpu, mode, true)

	cpu.flags.Carry = val > 0x7F
	val <<= 1
	if accumulator {
		cpu.registers.A = val
	} else {
		cpu.bus.Write(addr, val)
	}

	cpu.flags.Zero = val == 0
	cpu.flags.Negative = val&0x80 != 0
}

func (i iInstructionSet) BCC(cpu *CPU) {
	if !cpu.flags.Carry {
		val, _, _ := i.MemoryMode(cpu, Relative, true)
		if val > 0x7F {
			cpu.programCounter -= uint16(0xff-val) + 1
		} else {
			cpu.programCounter += uint16(val)
		}
	}
}

func (i iInstructionSet) BCS(cpu *CPU) {
	if cpu.flags.Carry {
		val, _, _ := i.MemoryMode(cpu, Relative, true)
		if val > 0x7F {
			cpu.programCounter -= uint16(0xff-val) + 1
		} else {
			cpu.programCounter += uint16(val)
		}
	}
}

func (i iInstructionSet) BEQ(cpu *CPU) {
	if cpu.flags.Zero {
		val, _, _ := i.MemoryMode(cpu, Relative, true)
		if val > 0x7F {
			cpu.programCounter -= uint16(0xff-val) + 1
		} else {
			cpu.programCounter += uint16(val)
		}
	}
}

func (i iInstructionSet) BIT(cpu *CPU, mode MemoryMode) {
	val, _, _ := i.MemoryMode(cpu, mode, true)

	cpu.flags.Zero = val&cpu.registers.A == 0
	cpu.flags.Negative = val&0x80 == 0x80
	cpu.flags.Overflow = val&0x40 == 0x40
}

func (i iInstructionSet) BMI(cpu *CPU) {
	if cpu.flags.Negative {
		val, _, _ := i.MemoryMode(cpu, Relative, true)
		if val > 0x7F {
			cpu.programCounter -= uint16(0xff-val) + 1
		} else {
			cpu.programCounter += uint16(val)
		}
	}
}

func (i iInstructionSet) BNE(cpu *CPU) {
	if !cpu.flags.Zero {
		val, _, _ := i.MemoryMode(cpu, Relative, true)
		if val > 0x7F {
			cpu.programCounter -= uint16(0xff-val) + 1
		} else {
			cpu.programCounter += uint16(val)
		}
	}
}

func (i iInstructionSet) BPL(cpu *CPU) {
	if !cpu.flags.Negative {
		val, _, _ := i.MemoryMode(cpu, Relative, true)
		if val > 0x7F {
			cpu.programCounter -= uint16(0xff-val) + 1
		} else {
			cpu.programCounter += uint16(val)
		}
	}
}

func (i iInstructionSet) BRK(cpu *CPU) {
	cpu.Interrupt()
	cpu.flags.BreakCommand = true
}

func (i iInstructionSet) BVC(cpu *CPU) {
	if !cpu.flags.Overflow {
		val, _, _ := i.MemoryMode(cpu, Relative, true)
		if val > 0x7F {
			cpu.programCounter -= uint16(0xff-val) + 1
		} else {
			cpu.programCounter += uint16(val)
		}
	}
}

func (i iInstructionSet) BVS(cpu *CPU) {
	if cpu.flags.Overflow {
		val, _, _ := i.MemoryMode(cpu, Relative, true)
		if val > 0x7F {
			cpu.programCounter -= uint16(0xff-val) + 1
		} else {
			cpu.programCounter += uint16(val)
		}
	}
}

func (i iInstructionSet) CLC(cpu *CPU) {
	cpu.flags.Carry = false
}

func (i iInstructionSet) CLD(cpu *CPU) {
	cpu.flags.Decimal = false
}

func (i iInstructionSet) CLI(cpu *CPU) {
	cpu.flags.InterruptDisable = false
}

func (i iInstructionSet) CLV(cpu *CPU) {
	cpu.flags.Overflow = false
}

func (i iInstructionSet) CMP(cpu *CPU, mode MemoryMode) {
	val, _, _ := i.MemoryMode(cpu, mode, true)

	result := cpu.registers.A - val

	cpu.flags.Carry = cpu.registers.A >= val
	cpu.flags.Zero = cpu.registers.A == val
	cpu.flags.Negative = result&0x80 != 0
}

func (i iInstructionSet) CPX(cpu *CPU, mode MemoryMode) {
	val, _, _ := i.MemoryMode(cpu, mode, true)

	result := cpu.registers.X - val

	cpu.flags.Carry = cpu.registers.X >= val
	cpu.flags.Zero = cpu.registers.X == val
	cpu.flags.Negative = result&0x80 != 0
}

func (i iInstructionSet) CPY(cpu *CPU, mode MemoryMode) {
	val, _, _ := i.MemoryMode(cpu, mode, true)

	result := cpu.registers.Y - val

	cpu.flags.Carry = cpu.registers.Y >= val
	cpu.flags.Zero = cpu.registers.Y == val
	cpu.flags.Negative = result&0x80 != 0
}

func (i iInstructionSet) DEC(cpu *CPU, mode MemoryMode) {
	val, addr, _ := i.MemoryMode(cpu, mode, true)

	val--
	cpu.bus.Write(addr, val)

	cpu.flags.Zero = val == 0
	cpu.flags.Negative = val&0x80 != 0
}

func (i iInstructionSet) DEX(cpu *CPU) {
	cpu.registers.X--

	cpu.flags.Zero = cpu.registers.X == 0
	cpu.flags.Negative = cpu.registers.X&0x80 != 0
}

func (i iInstructionSet) DEY(cpu *CPU) {
	cpu.registers.Y--

	cpu.flags.Zero = cpu.registers.Y == 0
	cpu.flags.Negative = cpu.registers.Y&0x80 != 0
}

func (i iInstructionSet) EOR(cpu *CPU, mode MemoryMode) {
	val, _, _ := i.MemoryMode(cpu, mode, true)

	cpu.registers.A ^= val

	cpu.flags.Zero = cpu.registers.A == 0
	cpu.flags.Negative = cpu.registers.A&0x80 != 0
}

func (i iInstructionSet) INC(cpu *CPU, mode MemoryMode) {
	val, addr, _ := i.MemoryMode(cpu, mode, true)

	val++
	cpu.bus.Write(addr, val)

	cpu.flags.Zero = val == 0
	cpu.flags.Negative = val&0x80 != 0
}

func (i iInstructionSet) INX(cpu *CPU) {
	cpu.registers.X++

	cpu.flags.Zero = cpu.registers.X == 0
	cpu.flags.Negative = cpu.registers.X&0x80 != 0
}

func (i iInstructionSet) INY(cpu *CPU) {
	cpu.registers.Y++

	cpu.flags.Zero = cpu.registers.Y == 0
	cpu.flags.Negative = cpu.registers.Y&0x80 != 0
}

func (i iInstructionSet) JMP(cpu *CPU, mode MemoryMode) {
	_, address, _ := i.MemoryMode(cpu, mode, false)
	cpu.programCounter = address
}

func (i iInstructionSet) JSR(cpu *CPU) {
	cpu.PushStack16(cpu.programCounter - 1)
	_, address, _ := i.MemoryMode(cpu, Absolute, true)
	cpu.programCounter = address
}

func (i iInstructionSet) LDA(cpu *CPU, mode MemoryMode) {
	cpu.registers.A, _, _ = i.MemoryMode(cpu, mode, true)

	cpu.flags.Zero = cpu.registers.A == 0
	cpu.flags.Negative = cpu.registers.A&0x80 != 0
}

func (i iInstructionSet) LDX(cpu *CPU, mode MemoryMode) {
	cpu.registers.X, _, _ = i.MemoryMode(cpu, mode, true)

	cpu.flags.Zero = cpu.registers.X == 0
	cpu.flags.Negative = cpu.registers.X&0x80 != 0
}

func (i iInstructionSet) LDY(cpu *CPU, mode MemoryMode) {
	cpu.registers.Y, _, _ = i.MemoryMode(cpu, mode, true)

	cpu.flags.Zero = cpu.registers.Y == 0
	cpu.flags.Negative = cpu.registers.Y&0x80 != 0
}

func (i iInstructionSet) LSR(cpu *CPU, mode MemoryMode) {
	val, addr, accumulator := i.MemoryMode(cpu, mode, true)

	cpu.flags.Carry = val&0x01 == 0x01
	val >>= 1
	if accumulator {
		cpu.registers.A = val
	} else {
		cpu.bus.Write(addr, val)
	}

	cpu.flags.Zero = val == 0
	cpu.flags.Negative = val&0x80 != 0
}

func (i iInstructionSet) NOP(cpu *CPU) {
	// Do nothing
}

func (i iInstructionSet) ORA(cpu *CPU, mode MemoryMode) {
	val, _, _ := i.MemoryMode(cpu, mode, true)

	cpu.registers.A |= val

	cpu.flags.Zero = cpu.registers.A == 0
	cpu.flags.Negative = cpu.registers.A&0x80 != 0
}

func (i iInstructionSet) PHA(cpu *CPU) {
	cpu.PushToStack(cpu.registers.A)
}

func (i iInstructionSet) PHP(cpu *CPU) {
	cpu.PushToStack(cpu.flags.ToByte())
}

func (i iInstructionSet) PLA(cpu *CPU) {
	cpu.registers.A = cpu.PopFromStack()

	cpu.flags.Zero = cpu.registers.A == 0
	cpu.flags.Negative = cpu.registers.A&0x80 != 0
}

func (i iInstructionSet) PLP(cpu *CPU) {
	cpu.flags.FromByte(cpu.PopFromStack())
}

func (i iInstructionSet) ROL(cpu *CPU, mode MemoryMode) {
	val, addr, accumulator := i.MemoryMode(cpu, mode, true)

	carry := cpu.flags.Carry
	cpu.flags.Carry = val&0x80 == 0x80
	val <<= 1
	if carry {
		val |= 0x01
	}
	if accumulator {
		cpu.registers.A = val
	} else {
		cpu.bus.Write(addr, val)
	}

	cpu.flags.Zero = val == 0
	cpu.flags.Negative = val&0x80 != 0
}

func (i iInstructionSet) ROR(cpu *CPU, mode MemoryMode) {
	val, addr, accumulator := i.MemoryMode(cpu, mode, true)

	carry := cpu.flags.Carry
	cpu.flags.Carry = val&0x01 == 0x01
	val >>= 1
	if carry {
		val |= 0x80
	}
	if accumulator {
		cpu.registers.A = val
	} else {
		cpu.bus.Write(addr, val)
	}

	cpu.flags.Zero = val == 0
	cpu.flags.Negative = val&0x80 != 0
}

func (i iInstructionSet) RTI(cpu *CPU) {
	cpu.flags.FromByte(cpu.PopFromStack())
	cpu.programCounter = cpu.PopStack16()
}

func (i iInstructionSet) RTS(cpu *CPU) {
	cpu.programCounter = cpu.PopStack16() + 1
}

func (i iInstructionSet) SBC(cpu *CPU, mode MemoryMode) {
	val, _, _ := i.MemoryMode(cpu, mode, true)

	cpu.registers.A -= val
	if !cpu.flags.Carry {
		cpu.registers.A--
	}

	cpu.flags.Carry = cpu.registers.A <= val
	cpu.flags.Zero = cpu.registers.A == 0
	cpu.flags.Negative = cpu.registers.A&0x80 != 0
	cpu.flags.Overflow = cpu.registers.A&0x80 != val&0x80
}

func (i iInstructionSet) SEC(cpu *CPU) {
	cpu.flags.Carry = true
}

func (i iInstructionSet) SED(cpu *CPU) {
	cpu.flags.Decimal = true
}

func (i iInstructionSet) SEI(cpu *CPU) {
	cpu.flags.InterruptDisable = true
}

func (i iInstructionSet) STA(cpu *CPU, mode MemoryMode) {
	_, addr, _ := i.MemoryMode(cpu, mode, false)
	cpu.bus.Write(addr, cpu.registers.A)
}

func (i iInstructionSet) STX(cpu *CPU, mode MemoryMode) {
	_, addr, _ := i.MemoryMode(cpu, mode, false)
	cpu.bus.Write(addr, cpu.registers.X)
}

func (i iInstructionSet) STY(cpu *CPU, mode MemoryMode) {
	_, addr, _ := i.MemoryMode(cpu, mode, false)
	cpu.bus.Write(addr, cpu.registers.Y)
}

func (i iInstructionSet) TAX(cpu *CPU) {
	cpu.registers.X = cpu.registers.A

	cpu.flags.Zero = cpu.registers.X == 0
	cpu.flags.Negative = cpu.registers.X&0x80 != 0
}

func (i iInstructionSet) TAY(cpu *CPU) {
	cpu.registers.Y = cpu.registers.A

	cpu.flags.Zero = cpu.registers.Y == 0
	cpu.flags.Negative = cpu.registers.Y&0x80 != 0
}

func (i iInstructionSet) TSX(cpu *CPU) {
	cpu.registers.X = uint8(cpu.stackPointer)

	cpu.flags.Zero = cpu.registers.X == 0
	cpu.flags.Negative = cpu.registers.X&0x80 != 0
}

func (i iInstructionSet) TXA(cpu *CPU) {
	cpu.registers.A = cpu.registers.X

	cpu.flags.Zero = cpu.registers.A == 0
	cpu.flags.Negative = cpu.registers.A&0x80 != 0
}

func (i iInstructionSet) TXS(cpu *CPU) {
	cpu.stackPointer = StackPointer(cpu.registers.X)
}

func (i iInstructionSet) TYA(cpu *CPU) {
	cpu.registers.A = cpu.registers.Y

	cpu.flags.Zero = cpu.registers.A == 0
	cpu.flags.Negative = cpu.registers.A&0x80 != 0
}

func (i *iInstructionSet) ReturnFromInterrupt(cpu *CPU) {
	// Pop the flags from the stack
	cpu.stackPointer.Pop()
	cpu.flags.FromByte(cpu.bus.Read(cpu.stackPointer.Address()))
	// Pop the program counter from the stack
	cpu.stackPointer.Pop()
	cpu.programCounter = uint16(cpu.bus.Read(cpu.stackPointer.Address()))
	cpu.stackPointer.Pop()
	cpu.programCounter |= uint16(cpu.bus.Read(cpu.stackPointer.Address())) << 8
}
