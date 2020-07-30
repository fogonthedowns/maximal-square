package main

import (
	"testing"
)

func TestFun(t *testing.T) {
	a := [][]byte{{'1'}}

	out1 := maximalSquare(a)

	if out1 != 1 {
		t.Errorf("got %d, want %d", out1, 1)
	}

	b := [][]byte{{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'}}
	out2 := maximalSquare(b)
	if out2 != 4 {
		t.Errorf("got %d, want %d", out2, 4)

	}
}
