package main

import "testing"

func TestAnything(t *testing.T) {

	foo := "bar"

	if foo != "bar" {
		t.Errorf("this is an error")
	}
}
