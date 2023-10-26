package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"math"
	"os"
)

var (
	// 8 registers
	reg = make([]uint32, 8)
	// execution finger
	ef = uint32(0)
	// collection of array of platters, init with 100 arrays
	col = make([][]uint32, 10, 100)
	// Create stack structure
	stack = memStack{}
)

type memStack []uint32

func (s *memStack) isEmpty() bool {
	return len(*s) == 0
}

func (s *memStack) push(x uint32) {
	*s = append(*s, x)
}

func (s *memStack) pop() (uint32, error) {
	if s.isEmpty() {
		return 0, fmt.Errorf("stack is empty")
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, nil
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("You need to specify path to input program")
		os.Exit(1)
	}
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Something went wrong while reading file!")
		os.Exit(1)
	}
	// let's see what our data holds, and cast it to uint32 and change endian
	col[0] = make([]uint32, len(data)/4)
	for i := 0; i < len(col[0]); i++ {
		col[0][i] = binary.BigEndian.Uint32(data[i*4 : (i+1)*4])
	}
	for {
		instruction := col[0][ef]
		// move instruction finger
		ef++
		// decode
		opCode := instruction >> (32 - 4)
		if opCode == 13 {
			A := (instruction >> 25) & 0x7
			// 0x 01 ff ff ff
			// (1 << 25) - 1
			reg[A] = instruction & 0x1ffffff
		} else {
			// Fetch register values from platter
			A := (instruction >> 6) & 0x7
			B := (instruction >> 3) & 0x7
			C := (instruction >> 0) & 0x7

			switch opCode {
			case 0:
				if reg[C] != 0 {
					reg[A] = reg[B]
				}
			case 1:
				reg[A] = col[reg[B]][reg[C]]
			case 2:
				col[reg[A]][reg[B]] = reg[C]
			case 3:
				// this won't overflow in golang
				reg[A] = reg[B] + reg[C]
			case 4:
				reg[A] = reg[B] * reg[C]
			case 5:
				if reg[C] != 0 {
					reg[A] = reg[B] / reg[C]
				}
			case 6:
				reg[A] = ^(reg[B] & reg[C])
			case 7:
				os.Exit(0)
			case 8:
				// create empty array
				emptyArr := make([]uint32, reg[C])
				// pop from stack
				sId, err := stack.pop()
				if err != nil {
					col = append(col, emptyArr)
					reg[B] = uint32(len(col)) - 1
				} else {
					col[sId] = emptyArr
					reg[B] = sId
				}
			case 9:
				// Remove array from memory
				col[reg[C]] = make([]uint32, 0)
				// push array identifier on stack
				stack.push(reg[C])
			case 10:
				// print ascii char
				if reg[C] <= 255 {
					// This was my first real "idk what is actually happening" problem
					// Figuring out that fmt.Printf and terminal somehow additionaly encode data
					// So I couldn't unpack codex.umz properly, but sandmark was passing and everything else
					// was working as expected
					// Thanks @ALPHA-60
					os.Stdout.Write([]byte{byte(reg[C])})
				} else {
					fmt.Println("Printing char > 255. Exiting")
					os.Exit(1)
				}
			case 11:
				// Input from console
				bRead := make([]byte, 1)
				_, err := os.Stdin.Read(bRead)
				if err == io.EOF {
					// if we hit EOF, MaxUint32
					reg[C] = math.MaxUint32
				}
				// else
				reg[C] = uint32(bRead[0])
			case 12:
				// free array at 0 and create it to be of size col[reg[B]]
				if reg[B] != 0 {
					col[0] = make([]uint32, len(col[reg[B]]))

					// copy array to col[0]
					copy(col[0], col[reg[B]])
				}
				ef = reg[C]
			default:
				fmt.Println("Not in spec")
				os.Exit(1)
			}

		}
	}
}
