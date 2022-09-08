package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	filename := os.Args[1]
	code, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(-1)
	}

	compiler := InstantiateCompiler(string(code))
	instructionSet := compiler.Run()

	machine := InstantiateMachine(instructionSet, os.Stdin, os.Stdout)
	machine.Run()
}
