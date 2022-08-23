package main

import "io"

type Machine struct {
	code string
	ip   int // Instruction pointer

	memory [30000]int
	cp     int // Context pointer (value to be effected by instruction)

	input  io.Reader
	output io.Writer

	buf []byte
}

func Instantiate(code string, in io.Reader, out io.Writer) *Machine {
	return &Machine{
		code:   code,
		input:  in,
		output: out,
		buf:    make([]byte, 1),
	}
}

func (m *Machine) Run() {
}
