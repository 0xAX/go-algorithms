package numerical

//using recursion
func fibo(num int) int {
	if num <= 1 {
		return num
	}
	return fibo(num-1) + fibo(num-2)
}
