package main

import "testing"

func TestCmdOk(t *testing.T) {
	err := Cmd("echo", "test")
	if err != nil {
		t.Fatalf("Expected err == nil, got %s", err)
	}
}

func TestCmdErrorNonFound(t *testing.T) {
	err := Cmd("anonexistentcommand")
	if err == nil {
		t.Fatalf("Expected err =! nil, got %s", err)
	}
}

func TestCmdBadReturnCode(t *testing.T) {
	err := Cmd("cd", "/anonexistenpath")
	if err == nil {
		t.Fatalf("Expected err =! nil, got %s", err)
	}
}
