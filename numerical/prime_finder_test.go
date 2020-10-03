package numerical

import "testing"

func arrayEquals(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// TestPrimeFinder tests prime finding
func TestPrimeFinder(t *testing.T) {

	if !arrayEquals(PrimesUpTo(10), []int{2, 3, 5, 7}) {
		t.Error("[Error] PrimesUpTo(10) is wrong")
	}
	if len(PrimesUpTo(100)) != 25 {
		t.Error("[Error] PrimesUpTo(100) is wrong")
	}
}
