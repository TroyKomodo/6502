package main

import (
	"6502emulator/emulator"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	// The range of the RAM is 0x0000 - 0x7FFF

	// RAM_SIZE is 32KB
	RAM_SIZE = 1 << 15
	// RAM_OFFSET is at 0
	RAM_OFFSET = 0x0000

	// StdIn is at 0x8000
	StdInOut = 0x8000

	// The max range of the ROM is 0x80FF - 0xFFFF
	// the Lower 256 bytes are reserved for the I/O
	// The ROM offset is calculated by subtracting the size of the ROM from the max range
	MAX_ROM_SIZE = (1 << 15) - (1 << 8)
)

func main() {
	cpu := emulator.NewCPU()
	bus := &emulator.Bus{}

	bus.AddMemory(emulator.NewRAM(RAM_SIZE, RAM_OFFSET))

	stdInChan, _ := emulator.InOutFromFile(os.Stdin, emulator.Read)
	_, stdOutChan := emulator.InOutFromFile(os.Stderr, emulator.Write)

	bus.AddMemory(emulator.NewIO(StdInOut, stdInChan, stdOutChan))

	cpu.Reset()

	cpu.ConnectBus(bus)

	interupt := make(chan struct{})

	cpu.ConnectInterrupt(interupt)

	// Load the ROM
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	if len(data) > MAX_ROM_SIZE {
		panic("ROM too large")
	}

	offset := 0xFFFF - len(data) + 1

	bus.AddMemory(emulator.NewROM(data, uint16(offset)))

	clock := make(chan time.Time)
	close(clock)
	cpu.ConnectClock(clock)

	go func() {
		for {
			<-stdInChan
			interupt <- struct{}{}
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	go func() {
		<-sig
		os.Exit(0)
	}()

	cpu.Start()
}
