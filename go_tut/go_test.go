package main

import (
	"fmt"
	"testing"
)

func TestAnything(t *testing.T) {

	foo := "bar"

	if foo != "bar" {
		t.Errorf("this is an error")
	}
	fmt.Println("...It passed")
}
