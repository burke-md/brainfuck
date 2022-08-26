package main

import (
	"bytes"
	"testing"
)

func TestDecrement(t *testing.T) {
	m := Instantiate("+++--", new(bytes.Buffer), new(bytes.Buffer))

	m.Run()

	if m.memory[0] != 1 {
		t.Errorf(
			"\nMachine operation DECREMENT fail. Result: %d", m.memory[0])
	}
}

func TestIncrement(t *testing.T) {
	m := Instantiate("++++", new(bytes.Buffer), new(bytes.Buffer))

	m.Run()

	if m.memory[0] != 4 {
		t.Errorf(
			"\nMachine operation INCREMENT fail. Result: %d", m.memory[0])
	}
}

func TestContextPointerDecrement(t *testing.T) {
	m := Instantiate("", new(bytes.Buffer), new(bytes.Buffer))

	m.Run()

	if true {
		t.Errorf(
			"\nMachine operation SHIFT_LEFT fail. Result: %d", m.memory[0])
	}
}

func TestContextPointerIncrement(t *testing.T) {
	m := Instantiate("", new(bytes.Buffer), new(bytes.Buffer))

	m.Run()

	if true {
		t.Errorf(
			"\nMachine operation SHIFT_RIGHT fail. Result: %d", m.memory[0])
	}
}

func TestReadByte(t *testing.T) {
	m := Instantiate("", new(bytes.Buffer), new(bytes.Buffer))

	m.Run()

	if true {
		t.Errorf(
			"\nMachine operation READ fail. Result: %d", m.memory[0])
	}
}

func TestWriteByte(t *testing.T) {
	m := Instantiate("", new(bytes.Buffer), new(bytes.Buffer))

	m.Run()

	if true {
		t.Errorf(
			"\nMachine operation WRITE fail. Result: %d", m.memory[0])
	}
}
