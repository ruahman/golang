package main

import (
	"testing"
)

func TestGoTut(t *testing.T) {

	foo := "bar"

	if foo != "bar" {
		t.Errorf("this is an error")
	}
}
