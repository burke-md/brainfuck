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
			m.cp--
		case '>':
			m.cp++
		case '+':
			current := m.memory[cp]
			if current == 255 {
				m.memory[cp] = 0
				break
			}
			m.memory[m.cp]++
		case '-':
			current := m.memory[m.cp]
			if current == 0 {
				m.memory[cp] = 255
				break
			}
			m.memory[cp]--
		case '.':
			m.writeByte()
		case ',':
			m.readByte()
		case '[':
			if m.memory[m.cp] == 0 {
				// See lang spec.
				// Continue execution if current cell(m.cp) has non zero value
				depth := 1
				for depth != 0 {
					m.ip++
					switch m.code[m.ip] {
					case '[':
						depth++ // Handle nested loops
					case ']':
						// Re-execute loop if current cell (m.cp)
						// has non zero value
						depth--
					}
				}
			}
		case ']':
			if m.memory[m.cp] != 0 {
				depth := 1
				for depth != 0 {
					m.ip--
					switch m.code[m.cp] {
					case ']':
						depth++
					case '[':
						depth--
					}
				}
			}
		}

		m.ip++ // Move instruction pointer forward
	}
}

func (m *Machine) readByte() {
	n, err := m.input.Read(m.buf)
	if err != nil {
		panic(err)
	}
	if n != 1 {
		panic("Wrong number of bytes read.")
	}
	m.memory[m.cp] = int(m.buf[0])
}

func (m *Machine) writeByte() {
	m.buf[0] = byte(m.memory[m.cp])

	n, err := m.output.Write(m.buf)
	if err != nil {
		panic(err)
	}
	if n != 1 {
		panic("Wrong number of bytes written.")
	}
}
