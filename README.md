# Brain Fuck

An implementation of the Brain Fuck programming language. See 
[wikipedia](https://en.wikipedia.org/wiki/Brainfuck#P%E2%80%B2%E2%80%B2:_Brainfuck's_formal_%22parent_language%22) here.

# TODOS:

[ ] Create Machine:
    [ ] Handle input
    [ ] Handle output
    [ ] Accepts code as string 
    [ ] Has  instruction pointer for array (len 30000)

[ ] Creat Run func:
    [ ] Moves pointer
    [ ] Switch statement to asses instruction at pointer
    [ ] Special attention/helper funcs for:
        [ ] Read char
        [ ] Write char
        [ ] Loops

[ ] Create Main func that:
    [ ] Reads Brain Fuck code in from file
    [ ] Calls new instance of Machine

### Beyond the basics

[ ] Create lexer with:
    [ ] Struct for instruction w/ data for optimizing repeated operations
    [ ] Similar struct w/ data for more efficient loops
