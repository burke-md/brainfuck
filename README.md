# Brain Fuck

An implementation of the Brain Fuck programming language. See 
[wikipedia](https://en.wikipedia.org/wiki/Brainfuck#P%E2%80%B2%E2%80%B2:_Brainfuck's_formal_%22parent_language%22) here.

### TODOS:

- [ ] Create Machine:
    - [x] Handle input
    - [x] Handle output
    - [x] Accepts code as string 
    - [x] Has  instruction pointer for array (len 30000)
    - [x] Has context pointer (slot in memory to be updated)

- [x] Create Instantiate func:
    - [x] Pass in code, reader, writer, and set buf

- [ ] Create Run func:
    - [x] Moves instruction pointer
    - [ ] Switch statement to asses instruction at pointer
    - [ ] Special attention/helper funcs for:
        - [ ] Read char
        - [ ] Write char
        - [ ] Loops

- [ ] Create Main func that:
    - [ ] Reads Brain Fuck code in from file
    - [ ] Calls new instance of Machine

### Beyond the basics

- [ ] Create lexer with:
    - [ ] Struct for instruction w/ data for optimizing repeated operations
    - [ ] Similar struct w/ data for more efficient loops
