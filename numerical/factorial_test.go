package numerical

import "testing"

// Testfactorial tests factorial function
func Testfactorial(t *testing.T) {

	if factorial(2) != 2 {
		t.Error("[Error] factorial(2) is wrong")
	}

	if factorial(3) != 6 {
		t.Error("[Error] factorial(3) is wrong")
	}

	if factorial(0) != 1 {
		t.Error("[Error] factorial(0) is wrong")
	}

	if factorial(5) != 120 {
		t.Error("[Error] factorial(5) is wrong")
	}
}
