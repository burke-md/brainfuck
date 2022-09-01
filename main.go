package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	//grab file
	fileName := os.Args[1]
	code, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
		os.Exit(-1)
	}
	// Create new machine instance
	m := InstantiateMachine(string(code), os.Stdin, os.Stdout)
	// Run/Start/Execute
	m.Run()
}
