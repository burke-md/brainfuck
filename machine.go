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

func Instantiate() *Machine {
	return &Machine{}
}

func (m *Machine) Run() {
}
