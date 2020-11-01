package bowling

import (
    "testing"
)

func TestValueGet(t *testing.T) {
	if ValueGet("x") != 10 {
		t.Fail()
	}
	
	if ValueGet("/") != 10 {
		t.Fail()
	}

	if ValueGet("-") != 0 {
		t.Fail()
	}

	if ValueGet("5") != 5 {
		t.Fail()
	}

	if ValueGet("1337") != 0 {
		t.Fail()
	}
}
