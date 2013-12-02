package main

import (
	"testing"
)

func MainUnTest(t *testing.T) {
	if aaa() != "asd" {
		t.Log("pass1")
	}
}

func MainEqTest(t *testing.T) {
	if aaa() != "asd" {
		t.Log("pass2")
	}
}
