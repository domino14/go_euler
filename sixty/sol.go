package sixty

import (
	"fmt"
	"log"
	"strconv"
	"sync"

	"github.com/domino14/go_euler/primes"
)

func concat(p int, q int) uint64 {
	s := fmt.Sprintf("%d%d", p, q)
	i, _ := strconv.ParseInt(s, 10, 64)
	return uint64(i)
}

func catenatedPrimes(sieve map[int]bool, maxSieveInt uint64,
	ps ...int) bool {

	checks := []uint64{}
	for _, a := range ps {
		for _, b := range ps {
			if a == b {
				continue
			}
			checks = append(checks, concat(a, b))
		}
	}

	for _, c := range checks {
		// log.Println("checking", c)
		if c <= maxSieveInt {
			if !sieve[int(c)] {
				return false
			}
		} else {
			if !primes.CheckPrime(c) {
				return false
			}
		}

	}

	return true
}

func anyEqual(ps ...int) bool {
	eqMap := map[int]bool{}
	for _, a := range ps {
		eqMap[a] = true
	}
	return len(eqMap) != len(ps)
}

func sum(ps ...int) int {
	sum := 0
	for _, c := range ps {
		sum += c
	}
	return sum
}

func tryCandidates(ps []int, pd map[int]bool, sieveLim uint64) {
	/**
	3 37 67 2377 ; sum = 2484   P=11010547
	3 17 449 2069 ; sum = 2538  (None found)
	11 23 743 1871 ; sum = 2648  (None found)
	3 11 2069 2297 ; sum = 4380  (None found)
	3 17 2297 2069 ; sum = 4386  P=2244623   - This is an upper bound i've found :/

	23 311 677 827 ; sum = 1838

	11 239 1049 1847 ; sum = 3146
	11 239 1091 1847 ; sum = 3188

	*/

	candidates := []int{3, 17, 2297, 2069}
	s := sum(candidates...)

	for p := 0; p < 1000000; p++ {
		newArgs := append(candidates, ps[p])

		if anyEqual(newArgs...) {
			continue
		}
		if catenatedPrimes(pd, uint64(sieveLim), newArgs...) {
			log.Println("New prime:", ps[p], "sum:", ps[p]+s)
			break
		}
	}
}

func searchCandidates(ps []int, pd map[int]bool, sieveLim uint64) {
	countLim := 400

	var wg sync.WaitGroup
	count := func(pbegin, pend int, ch chan []int) {
		for p := pbegin; p < pend; p++ {
			for q := p + 1; q < countLim; q++ {
				for r := q + 1; r < countLim; r++ {
					for s := r + 1; s < countLim; s++ {
						if catenatedPrimes(pd, sieveLim,
							ps[p], ps[q], ps[r], ps[s]) {

							ch <- []int{ps[p], ps[q], ps[r], ps[s]}
						}
					}
				}
			}
			log.Println("p now", p)
		}
		wg.Done()
	}

	ch := make(chan []int)
	wg.Add(4)

	go count(0, 25, ch)
	go count(25, 50, ch)
	go count(50, 75, ch)
	go count(75, 100, ch)

	go func() {
		for x := range ch {
			log.Println(x, "; sum =", sum(x...))
		}
	}()
	wg.Wait()
}

func Solve() {

	sieveLim := 1e8
	ps := primes.Sieve(int(sieveLim))
	pd := primes.SieveDict(int(sieveLim))
	log.Println("Generated sieve")

	searchCandidates(ps, pd, uint64(sieveLim))

	// tryCandidates(ps, pd, uint64(sieveLim))

}
