package emulator

import (
	"bufio"
	"fmt"
	"os"
)

type iIO struct {
	Address uint16

	In  <-chan uint8
	Out chan<- uint8
}

func NewIO(address uint16, in <-chan uint8, out chan<- uint8) Memory {
	return &iIO{
		Address: address,
		In:      in,
		Out:     out,
	}
}

func (io *iIO) Contains(address uint16) bool {
	return address == io.Address
}

func (io *iIO) Read(address uint16) uint8 {
	if io.In == nil {
		return 0
	}

	select {
	case data := <-io.In:
		return data
	default:
		fmt.Println("IO read failed")
		return 0
	}
}

func (io *iIO) Write(address uint16, data uint8) {
	if io.Out == nil {
		return
	}

	io.Out <- data
	// select {
	// case io.Out <- data:
	// default:
	// fmt.Println("IO write failed")
	// }
}

type ReadWrite byte

const (
	Read  ReadWrite = 1
	Write ReadWrite = 2
)

func InOutFromFile(file *os.File, mode ReadWrite) (<-chan uint8, chan<- uint8) {
	var (
		in  chan uint8
		out chan uint8
	)

	if mode&Read != 0 {
		reader := bufio.NewReader(file)
		in = make(chan uint8)
		go func() {
			for {
				data, err := reader.ReadByte()
				if err != nil {
					return
				}

				in <- data
			}
		}()
	}
	if mode&Write != 0 {
		out = make(chan uint8)
		go func() {
			for data := range out {
				_, err := file.Write([]byte{data})
				if err != nil {
					return
				}
			}
		}()
	}

	return in, out
}
