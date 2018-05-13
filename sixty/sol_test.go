package sixty

import (
	"testing"

	"github.com/domino14/go_euler/primes"
)

func TestCatenatedPrimes(t *testing.T) {
	sieveLim := uint64(1e6)
	sieve := primes.SieveDict(int(sieveLim))

	sol := catenatedPrimes(sieve, sieveLim, 3, 7, 109, 673)
	if !sol {
		t.Error("Sol is false")
	}
}
