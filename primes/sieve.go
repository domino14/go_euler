package primes

import "math"

// Sieve returns a list of prime numbers from 2 to a prime < n
func Sieve(n int) []int {
	var sieve []bool
	sieve = make([]bool, n/2)
	for i := 0; i < len(sieve); i++ {
		sieve[i] = true
	}
	for i := 3; i < int(math.Pow(float64(n), 0.5)); i += 2 {
		if sieve[i/2] {
			for j := i * i / 2; j < len(sieve); j += i {
				sieve[j] = false
			}
		}
	}
	vals := []int{2}
	for i := 1; i < n/2; i++ {
		if sieve[i] {
			vals = append(vals, 2*i+1)
		}
	}
	return vals
}

// SieveDict is a dictionary version of the sieve above
func SieveDict(n int) map[int]bool {
	// basically just a set
	var sieve map[int]bool
	sieve = make(map[int]bool)
	intsieve := Sieve(n)
	for _, p := range intsieve {
		sieve[p] = true
	}
	return sieve
}

// CheckPrime checks if n is prime
func CheckPrime(n uint64) bool {
	if n == 0 || n == 1 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	for i := uint64(3); i <= uint64(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
