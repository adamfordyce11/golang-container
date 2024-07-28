package main

import (
	"testing"
)

func TestGetOutput(t *testing.T) {
	expected := "Value: Start\n"
	actual := GetOutput("Start\n")
	if actual != expected {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}
