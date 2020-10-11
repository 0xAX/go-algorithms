package numerical

import "testing"

func TestFibonacci(t *testing.T) {

	if fibo(10) != 55 {
		t.Error("[Error] Fibonacci(10) is wrong")
	}

	if fibo(0) != 0 {
		t.Error("[Error] Fibonacci(0) is wrong")
	}

	if fibo(3) != 2 {
		t.Error("[Error] Fibonacci(3) is wrong")
	}

}
