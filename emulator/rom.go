package emulator

type ROM struct {
	Offset uint16
	Data   []uint8
}

func NewROM(data []uint8, offset uint16) Memory {
	return &ROM{
		Offset: offset,
		Data:   data,
	}
}

func (rom *ROM) Read(address uint16) uint8 {
	return rom.Data[address-rom.Offset]
}

func (rom *ROM) Write(address uint16, data uint8) {
	// do nothing
}

func (rom *ROM) Contains(address uint16) bool {
	return address >= rom.Offset && (address-rom.Offset) < uint16(len(rom.Data))
}
