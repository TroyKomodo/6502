package emulator

type iRAM struct {
	Offset uint16
	Data   []uint8
}

func NewRAM(size uint16, offset uint16) Memory {
	return &iRAM{
		Offset: offset,
		Data:   make([]uint8, size),
	}
}

func (ram *iRAM) Read(address uint16) uint8 {
	return ram.Data[address-ram.Offset]
}

func (ram *iRAM) Write(address uint16, data uint8) {
	ram.Data[address-ram.Offset] = data
}

func (ram *iRAM) Contains(address uint16) bool {
	return address >= ram.Offset && address < ram.Offset+uint16(len(ram.Data))
}
