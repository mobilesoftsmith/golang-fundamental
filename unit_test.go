package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	if 3 != add(1, 2) {
		t.Error("an error happened...")
	}
}