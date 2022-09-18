package emulator

// in the 6502, the address bus is 16 bits wide
// and the data bus is 8 bits wide
// the address bus is used to specify the location of data
// the data bus is used to transfer data between the CPU and memory

type Memory interface {
	Read(address uint16) uint8
	Write(address uint16, data uint8)
	Contains(address uint16) bool
}

type Bus struct {
	Memory []Memory
}

func (bus *Bus) Read(address uint16) uint8 {
	if bus == nil {
		return 0
	}

	var result uint8
	for _, memory := range bus.Memory {
		if memory.Contains(address) {
			result |= memory.Read(address)
		}
	}

	// fmt.Printf("R %04x %02x\n", address, result)

	return result
}

func (bus *Bus) Write(address uint16, data uint8) {
	if bus == nil {
		return
	}

	// fmt.Printf("W %04x %02x\n", address, data)

	for _, memory := range bus.Memory {
		if memory.Contains(address) {
			memory.Write(address, data)
		}
	}
}

func (bus *Bus) AddMemory(memory Memory) {
	bus.Memory = append(bus.Memory, memory)
}
