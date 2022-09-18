package emulator

type StackPointer uint8

func (sp *StackPointer) Push() {
	*sp--
}

func (sp *StackPointer) Pop() {
	*sp++
}

func (sp *StackPointer) Address() uint16 {
	return uint16(*sp) + 0x0100
}

type Flags struct {
	Carry            bool // Carry flag
	Zero             bool // Zero flag
	InterruptDisable bool // Interrupt disable flag
	Decimal          bool // Decimal mode flag
	BreakCommand     bool // Break command flag
	Overflow         bool // Overflow flag
	Negative         bool // Negative flag
}

func (f Flags) ToByte() uint8 {
	var b uint8
	if f.Carry {
		b |= 0x01
	}
	if f.Zero {
		b |= 0x02
	}
	if f.InterruptDisable {
		b |= 0x04
	}
	if f.Decimal {
		b |= 0x08
	}
	if f.BreakCommand {
		b |= 0x10
	}
	if f.Overflow {
		b |= 0x20
	}
	if f.Negative {
		b |= 0x40
	}

	return b
}

func (f *Flags) FromByte(b uint8) {
	f.Carry = b&0x01 != 0
	f.Zero = b&0x02 != 0
	f.InterruptDisable = b&0x04 != 0
	f.Decimal = b&0x08 != 0
	f.BreakCommand = b&0x10 != 0
	f.Overflow = b&0x20 != 0
	f.Negative = b&0x40 != 0
}
