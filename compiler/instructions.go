package main

type InstructionType byte

const (
	Plus          InstructionType = '+'
	Minus         InstructionType = '-'
	Right         InstructionType = '>'
	Left          InstructionType = '<'
	WriteChar     InstructionType = '.'
	ReadChar      InstructionType = ','
	JumpIfZero    InstructionType = '['
	JumpIfNonZero InstructionType = ']'
)

type Instruction struct {
	Type     InstructionType
	Arugment int
}
