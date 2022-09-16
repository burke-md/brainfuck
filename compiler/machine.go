package main

import "io"

type Machine struct {
	code []*Instruction
	ip   int // Instruction pointer

	memory [30000]int
	cp     int // Context pointer

	input  io.Reader
	output io.Writer

	readBuf []byte
}

func InstantiateMachine(
	instructions []*Instruction,
	in io.Reader,
	out io.Writer) *Machine {
	return &Machine{
		code:    instructions,
		input:   in,
		output:  out,
		readBuf: make([]byte, 1),
	}
}

func (m *Machine) Run() {
	for m.ip < len(m.code) {
		instruction := m.code[m.ip]

		switch instruction.Type {
		case Plus:
		case Minus:
		case Right:
		case Left:
		case WriteChar:
		case ReadChar:
		case JumpIfZero:
		case JumpIfNonZero:
		}

		m.ip++
	}
}
