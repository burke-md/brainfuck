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
		inst := m.code[m.ip]

		switch inst.Type {
		case Plus:
			m.memory[m.cp] += inst.Argument
			if m.memory[m.cp] == 256 {
				m.memory[m.cp] = 0 // underflow
			}
		case Minus:
			m.memory[m.cp] -= inst.Argument
			if m.memory[m.cp] == -1 {
				m.memory[m.cp] = 255 // overflow
			}
		case Right:
			m.cp += inst.Argument
		case Left:
			m.cp -= inst.Argument
		case WriteChar:
			for i := 0; i < inst.Argument; i++ {
				m.writeChar()
			}
		case ReadChar:
			for i := 0; i < inst.Argument; i++ {
				m.readChar()
			}
		case JumpIfZero:
		case JumpIfNonZero:
		}

		m.ip++
	}
}

func (m *Machine) readChar() {
	n, err := m.input.Read(m.readBuf)
	if err != nil {
		panic(err)
	}
	if n != 1 {
		panic("Wrong number of bytes read")
	}

	m.memory[m.cp] = int(m.readBuf[0])
}

func (m *Machine) writeChar() {
	m.readBuf[0] = byte(m.memory[m.cp])

	n, err := m.output.Write(m.readBuf)
	if err != nil {
		panic(err)
	}
	if n != 1 {
		panic("Wrong number of bytes written")
	}
}
