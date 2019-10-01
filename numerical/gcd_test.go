package numerical

import "testing"

// TestGcd tests gcd
func TestGcd(t *testing.T) {

	if GCD(100, 200) != 50 {
		t.Error("[Error] GCD(100, 200) is wrong")
	}

	if GCD(4, 2) != 1 {
		t.Error("[Error] GCD(4,2) is wrong")
	}

	if GCD(6, 3) != 3 {
		t.Error("[Error] GCD(6,3) is wrong")
	}
}
