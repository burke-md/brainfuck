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
	for m.ip < len(m.code) { // Iterate through code string
		instruction := m.code[m.ip]

		switch instruction {
		case '<':
		case '>':
		case '+':
		case '-':
		case '.':
		case ',':
		case '[':
		case ']':
		}

		m.ip++ // Move instruction pointer forward
	}
}

func (m *Machine) readByte() {

}

func (m *Machine) writeByte() {

}
