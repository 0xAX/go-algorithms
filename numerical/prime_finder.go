package numerical

// PrimesUpTo finds all prime numbers from 1 to upperBound
// It's implemented using eratosthenes sieve
// Works in O(upperBound) time and space
func PrimesUpTo(upperBound int) []int {
	// lp array stores minimal prime divisor for every number from 2 to upperBound
	var lp []int
	lp = make([]int, upperBound+1)

	// primes array stores primes
	var primes []int

	for i := 2; i <= upperBound; i++ {
		if lp[i] == 0 {
			lp[i] = i
			primes = append(primes, i)
		}
		for j := 0; j < len(primes) && primes[j] <= lp[i] && i*primes[j] <= upperBound; j++ {
			lp[i*primes[j]] = primes[j]
		}
	}

	return primes
}
