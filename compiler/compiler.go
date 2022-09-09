package main

type Compiler struct {
	code       string
	codeLength int
	position   int

	instructions []*Instruction
}

func InstantiateCompiler(code tring) *Compiler {
	return &Compiler{
		code:         code,
		codeLength:   len(code),
		instructions: []*Instruction{},
	}
}

func (c *Compiler) Run() []*Instruction {
	loopStack := []int{}

	for c.position < c.codeLength {
		current := c.code[c.position]

		switch current {
		case '[':
			insPos := c.EmitWithArg(JumpIfZero, 0)
			loopStack = append(loopStack, insPos)
		case ']':
			// Pop last open bracket off stack
			openInstruction := loopStack[len(loopStack)-1]
			loopStack = loopStack[:len(loopStack)-1]

			closeInstructionPos := c.EmitWithArg(JumpIfNonZero, openInstruction)
			c.instructions[openInstruction].Argument = closeInstructionPos

		case '+':
		case '-':
		case '<':
		case '>':
		case '.':
		case ',':
		}

		c.position++
	}

	return c.instructions
}

func (c *Compiler) CompileFoldableInstruction(char byte, insType InstructionType) {
	count := 1

	for c.position < c.codeLength-1 && c.code[c.position+1] == char {
		count++
		c.position++
	}

	c.EmitWithArg(insType, count)
}

func (c *Compiler) EmitWithArg(insType InstructionType, arg int) int {
	ins := &Instruction{Type: insType, Argument: arg}
	c.instructions = append(c.instructions, ins)
	return len(c.instructions) - 1
}
