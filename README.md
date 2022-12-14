# Brain Fuck

### The results:

Using the mandlebrot visualization (Implemented by [Erik](https://github.com/erikdubbelboer) )
as the bench mark:

The interpreter with no optimization: `71.31s`

The compiler with instructions to reduce duplicate commands and searching for
corrosponding brackets: `13.17s`

![Screen Shot 2022-09-21 at 4 33 16 PM](https://user-images.githubusercontent.com/22263098/191605032-354a602e-1403-4be7-9af2-e70cbe70fd2a.png)


### Instructions to run tests:

This can be run locally by cloning this repo and running the following command
to build, run and time the execution in either the interpreter and or the compiler
directory :
`go build -o machine && time ./machine ../sample_bf/mandelbrot.b`

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
    
- [x] Create lexer with:
    - [x] Struct for instruction w/ data for optimizing repeated operations
    - [x] Similar struct w/ data for more efficient loops

## The compiler:

Below is the working checklist to build a functioning brainfuck compiler.
The end goal is to benchmark the interpreter and compiler (perhaps even an
improved compiler) against the classic `mandelbrot.b` (a mandelbrot fractal 
viewer) written by Erik Bosman.

- [ ] main.go
    - [x] Accept arg (brainfuck file location)
    - [x] Basic error handling
    - [x] Instantiate new instance of compiler
    - [x] Run/return output from compiler
    - [x] Instantiate new instance of bf machine
    - [x] Run/execute


- [ ] machine.go
    - [x] Define Machine as struct
    - [x] Define `InstantiateMachine` func
    - [x] Extend `Machine` w/ `Run` method
        - [x] Iterate through new instruction list
        - [x] Switch/case w/ appropriate case for each instruction
        - [x] Read/Write helper functions

- [x] compiler.go
    - [x] Define `Compiler` struct 
        - [x] code, code length, position, instructions array
    - [x] Define `InstantiateCompiler` func
    - [x] Extend `Compiler` w/ `Run` method
        - [x] Iterate through code
        - [x] Switch/case
            - [x] Append instruction w/ arg for `[` / `]`
            - [x] Call helper to handle repetitive commands

- [x] instructions.go
    - [x] Define `InstructionType` as `byte`
    - [x] Define `Instruction` struct w/ `InstructionType` & `Arg`
    - [x] Map instruction names to `InstructionType` (all 8 chars)

