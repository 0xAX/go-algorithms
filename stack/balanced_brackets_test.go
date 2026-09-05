package stack

import "testing"

func Test_isExpressionBalanced(t *testing.T) {

	if !isExpressionBalanced("(())") {
		t.Error("[Error] Expression Balanced is wrong")
	}

	if isExpressionBalanced("(()") {
		t.Error("[Error] Expression Balanced is wrong")
	}

}
