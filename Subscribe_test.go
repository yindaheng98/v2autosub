package main

import (
	"testing"
)

func TestArithmeticMean(t *testing.T) {
	s, err := Subscribe(addr)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(s)
}
