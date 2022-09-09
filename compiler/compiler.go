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
		case ']':
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
