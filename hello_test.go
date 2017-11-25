package hello_go

import "testing"

func TestHello(t *testing.T) {
	expected := "Hello Go!"
	actual := hello("Go")
	if actual != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}
}
