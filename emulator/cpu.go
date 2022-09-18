package emulator

import (
	"fmt"
	"time"
)

type CPU struct {
	programCounter uint16       // 16-bit program counter
	stackPointer   StackPointer // 8-bit stack pointer

	registers struct {
		A uint8 // Accumulator
		X uint8 // Index register X
		Y uint8 // Index register Y
	}

	cycleCount         uint64 // The number of cycles that have passed
	previousCycleCount uint64 // The number of cycles that have passed in the previous step

	flags Flags

	clock     <-chan time.Time
	bus       *Bus
	interrupt <-chan struct{}

	instructionSet iInstructionSet
}

func (cpu *CPU) Reset() {
	cpu.programCounter = uint16(cpu.bus.Read(0xFFFC)) | uint16(cpu.bus.Read(0xFFFD))<<8
	cpu.stackPointer = 0
	cpu.registers.A = 0
	cpu.registers.X = 0
	cpu.registers.Y = 0

	cpu.flags.FromByte(0)
}

func (cpu *CPU) Start() {
	cpu.Reset()

	for {
		select {
		case <-cpu.interrupt:
			if cpu.flags.InterruptDisable {
				continue
			}

			cpu.Interrupt()
		case <-cpu.clock:
			cpu.cycleCount++
			cpu.Step()
		}
	}
}

func (cpu *CPU) Step() {
	code := cpu.bus.Read(cpu.programCounter)
	cpu.programCounter++

	cpu.ProcessInstruction(code)
}

func (cpu *CPU) ConnectBus(bus *Bus) {
	cpu.bus = bus
}

func (cpu *CPU) ConnectClock(clock <-chan time.Time) {
	cpu.clock = clock
}

func (cpu *CPU) ConnectInterrupt(interrupt <-chan struct{}) {
	cpu.interrupt = interrupt
}

func (cpu *CPU) Interrupt() {
	// Write the program counter to the stack
	cpu.PushStack16(cpu.programCounter)
	// Write the flags to the stack
	cpu.PushToStack(cpu.flags.ToByte())

	// Set the interrupt disable flag so that we don't get interrupted while handling the interrupt
	cpu.flags.InterruptDisable = true

	// Little-endian
	// Read the program counter from the interrupt vector (0xFFFE) and (0xFFFF)
	cpu.programCounter = uint16(cpu.bus.Read(0xFFFE)) | uint16(cpu.bus.Read(0xFFFF))<<8

	fmt.Println("Number of cycles passed: ", cpu.cycleCount, " (", cpu.cycleCount-cpu.previousCycleCount, ")")
	cpu.previousCycleCount = cpu.cycleCount
}

func (cpu *CPU) Pulse() {
	<-cpu.clock
	cpu.cycleCount++
}

func (cpu *CPU) PushToStack(b uint8) {
	cpu.bus.Write(cpu.stackPointer.Address(), b)
	cpu.stackPointer.Push()
}

func (cpu *CPU) PopFromStack() uint8 {
	cpu.stackPointer.Pop()
	return cpu.bus.Read(cpu.stackPointer.Address())
}

func (cpu *CPU) PushStack16(b uint16) {
	// Little endian
	cpu.PushToStack(uint8(b))
	cpu.PushToStack(uint8(b >> 8))
}

func (cpu *CPU) PopStack16() uint16 {
	// Little endian
	high := cpu.PopFromStack()
	low := cpu.PopFromStack()
	return uint16(high)<<8 | uint16(low)
}

func NewCPU() *CPU {
	cpu := &CPU{}

	return cpu
}
