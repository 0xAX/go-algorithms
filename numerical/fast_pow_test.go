package numerical

import "testing"

// TestFastPow tests fastpow function
func TestFastPow(t *testing.T) {

	if FastPow(2, 10) != 1024 {
		t.Error("[Error] FastPow(2, 10) is wrong")
	}

	if FastPow(1, 10) != 1 {
		t.Error("[Error] FastPow(1, 10) is wrong")
	}

	if FastPow(0, 15) != 0 {
		t.Error("[Error] FastPow(0, 15) is wrong")
	}

	if FastPow(10, 2) != 100 {
		t.Error("[Error] FastPow(0, 15) is wrong")
	}
}
