package main

import (
	"testing"
)

func TestMainUn(t *testing.T) {
	if aaa() != "asd" {
		t.Log("pass1")
	}
}

func TestMainEq(t *testing.T) {
	if aaa() == "asd" {
		t.Log("pass2")
	}
}
