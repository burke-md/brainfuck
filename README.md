# Brain Fuck

An implementation of the Brain Fuck programming language. See 
[wikipedia](https://en.wikipedia.org/wiki/Brainfuck#P%E2%80%B2%E2%80%B2:_Brainfuck's_formal_%22parent_language%22) here.

## The interpreter:

Below is a check list I used while building the basic interpreter - in the 
future things may evolve beyond its current state.

To build and run: 

- Navigate inside the interpreter dir
- Run command to build ` go build -o machine`
- Run command to see output in terminal ` ./machine ../sample_bf/hello_world.b`

### TODOS:

- [x] Create Machine:
    - [x] Handle input
    - [x] Handle output
    - [x] Accepts code as string 
    - [x] Has  instruction pointer for array (len 30000)
    - [x] Has context pointer (slot in memory to be updated)

- [x] Create Instantiate func:
    - [x] Pass in code, reader, writer, and set buf

- [ ] Create Run func:
    - [x] Moves instruction pointer
    - [x] Switch statement to asses instruction at pointer
    - [x] Special attention/helper funcs for:
        - [x] ReadByte
        - [X] WriteByte
        - [x] Loops(not implemented as a seperate method at this time)

- [x] Create Main func that:
    - [x] Reads Brain Fuck code in from file
    - [x] Calls new instance of Machine

- [x] Handle bug => `go build -o machine && ./machine ./hello_world.b` causes 
terminal to hang.

Solution:

The switch statement case associated with the ']' char has a mistake.
This has been resolved.

### Beyond the basics

- [ ] Tests
    - [x] Define tests for each basic command
    - [ ] Implement tests as defined
    
- [ ] Create lexer with:
    - [ ] Struct for instruction w/ data for optimizing repeated operations
    - [ ] Similar struct w/ data for more efficient loops

## The compiler:

Below is the working checklist to build a functioning brainfuck compiler.
The end goal is to benchmark the interpreter and compiler (perhaps even an
improved compiler) against the classic `mandelbrot.b` (a mandelbrot fractal 
viewer) written by Erik Bosman.

- [ ] main.go
    - [ ] Accept CLI arg (brainfuck file location)
    - [ ] Basic error handling
    - [ ] Instantiate new instance of compiler
    - [ ] Run/return output from compiler
    - [ ] Instantiate new instance of bf machine
    - [ ] Run/execute


- [ ] machine.go
    - [ ] Define Machine as struct
    - [ ] Define `InstantiateMachine` func
    - [ ] Extend `Machine` w/ `Run` method
        - [ ] Iterate through new instruction list
        - [ ] Switch/case w/ appropriate case for each instruction
        - [ ] Read/Write helper functions

- [ ] compiler.go
    - [ ] Define `Compiler` struct 
        - [ ] code, code length, position, instructions array
    - [ ] Define `InstantiateCompiler` func
    - [ ] Extend `Compiler` w/ `Run` method
        - [ ] Iterate through code
        - [ ] Switch/case
            - [ ] Append instruction w/ arg for `[` / `]`
            - [ ] Call helper to handle repetitive commands

- [ ] instructions.go
    - [ ] Define `InstructionType` as `byte`
    - [ ] Define `Instruction` struct w/ `InstructionType` & `Arg`
    - [ ] Map instruction names to `InstructionType` (all 8 chars)

