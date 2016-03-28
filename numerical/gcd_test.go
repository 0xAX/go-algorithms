package gcd

import "testing"

func Test_gcd(t *testing.T) {

	if gcd(100, 200) != 50 {
		t.Error("[Error] gcd(100, 200) is wrong")
	}

	if gcd(4, 2) != 1 {
		t.Error("[Error] gcd(4,2) is wrong")
	}

	if gcd(6, 3) != 3 {
		t.Error("[Error] gcd(6,3) is wrong")
	}
}
