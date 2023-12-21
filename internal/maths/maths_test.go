package maths

import "testing"

func TestGCD(t *testing.T) {
	testCases := []struct {
		A, B, Expected int
	}{
		{8, 12, 4},
		{54, 24, 6},
		{48, 18, 6},
	}

	for _, c := range testCases {
		if actual := GCD(c.A, c.B); actual != c.Expected {
			t.Errorf("GCD of %d and %d was expected to be %d, but was %d", c.A, c.B, c.Expected, actual)
		}
	}
}

func TestLCM(t *testing.T) {
	testCases := []struct {
		A, B, Expected int
	}{
		{12, 15, 60},
	}

	for _, c := range testCases {
		if actual := LCM(c.A, c.B); actual != c.Expected {
			t.Errorf("LCM of %d and %d was expected to be %d, but was %d", c.A, c.B, c.Expected, actual)
		}
	}

	if actual := LCM(12, 15, 20, 30, 40, 50, 60, 70, 80, 90, 100); actual != 25200 {
		t.Errorf("LCM of 12, 15, 20, 30, 40, 50, 60, 70, 80, 90, 100 was expected to be 25200, but was %d", actual)
	}
}
