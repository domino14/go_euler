package sixty

import (
	"fmt"
	"log"
	"sort"
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

	for _, a := range ps {
		for _, b := range ps {
			if a == b {
				continue
			}
			c := concat(a, b)
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

func tryCandidates(c candidate, ps []int, pd map[int]bool,
	sieveLim uint64) []int {
	/**
	3 37 67 2377 ; sum = 2484   P=11010547
	3 17 449 2069 ; sum = 2538  (None found)
	11 23 743 1871 ; sum = 2648  (None found)
	3 11 2069 2297 ; sum = 4380  (None found)
	3 17 2069 2297 ; sum = 4386  P=2244623   - This is an upper bound i've found :/

	23 311 677 827 ; sum = 1838 ; none
	11 239 1049 1847 ; sum = 3146 ; none
	11 239 1091 1847 ; sum = 3188 ; none

	269 617 887 2741 ; sum = 4514 ; none
	23 677 827 1871 ; sum = 3398 ; none
	31 1123 2029 2281 ; sum = 5464  P=11803459 ; s=11808923
	7 61 1693 3181 ; sum = 4942  P = 2359207
	7 829 2671 3361 ; sum = 6868 ; none

	7 1237 1549 3019 ; sum = 5812 ; none
	7 2089 2953 3181 ; sum = 8230 ; none
	37 1549 2707 3463 ; sum = 7756 ; none
	79 967 1117 3511 ; sum = 5674  ; s=8942491

	79 1801 3253 3547 ; sum = 8680 ; p=15396127

	----
	3 37 67 5923  ; 194119 ; s=200149  - new upper bound! (idx = 17510)

	3 17 449 6353 - none
	3 17 449 6599 - none

	11 23 743 115259 - too big
	11 23 743 164663 - too big

	3 11 2069 8219 - none
	3 11 2069 8747 - none

	23 311 677 - too big
	11 239 1049 11057 - none
	11 239 1049 93503 - none
	11 239 1091 56249 - none

	269 617 887 - none
	23 677 827 - none

	31 1123 2029 5281 - 495199, too big (s 503663)
	7 61 1693 - too big

	3 7 109 29059 - none
	3 7 109 79693 - none
	3 7 109 91159 - none
	3 7 109 93187 - none
	7 829 2671 - too big
	7 829 3361 - too big
	7 1237 1549 - too big

	7 2089 2953 - too big
	37 1549 2707 - nah

	*/

	// todo, should try all 4c3 combinations of found groups of 4 above
	// find bigger numbers (so new groups of 4)
	// try groups of 4 until idx roughly 18K. can't be any bigger

	newCandidates := []int{}

	highestPrimeIndex := 2140
	upperBoundSum := 35000

	for p := 0; p < highestPrimeIndex; p++ {
		newArgs := append(c.primes, ps[p])

		if anyEqual(newArgs...) {
			continue
		}
		if ps[p] > upperBoundSum-c.sum {
			break
		}
		if catenatedPrimes(pd, uint64(sieveLim), newArgs...) {
			// log.Println("New prime:", ps[p], "sum:", ps[p]+s, "idx", p)
			//break
			newCandidates = append(newCandidates, ps[p])
		}
	}
	return newCandidates
}

type candidate struct {
	primes []int
	sum    int
}

type bySum []candidate

func (c bySum) Len() int           { return len(c) }
func (c bySum) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c bySum) Less(i, j int) bool { return c[i].sum < c[j].sum }

// Search for candidate triplets
func searchTriplets(ps []int, pd map[int]bool, sieveLim uint64) []candidate {
	var wg sync.WaitGroup
	// Cannot be more than 1400. This is prime 11657, which times 3
	// is higher than our upper bound sum of 34427
	countLim := 1400

	count := func(pbegin, pend int, ch chan []int) {
		for q := pbegin; q < pend; q++ {
			for r := q + 1; r < countLim; r++ {
				for s := r + 1; s < countLim; s++ {
					if catenatedPrimes(pd, sieveLim,
						ps[q], ps[r], ps[s]) {

						ch <- []int{ps[q], ps[r], ps[s]}
					}
				}
			}
			if q%10 == 0 {
				log.Println(q)
			}
		}

		wg.Done()
	}

	ch := make(chan []int)
	wg.Add(4) // cores on my computer :P

	go count(0, 175, ch)
	go count(175, 350, ch)
	go count(350, 700, ch)
	go count(700, 1400, ch)

	candidates := []candidate{}

	go func() {
		for x := range ch {
			candidates = append(candidates, candidate{
				primes: x,
				sum:    sum(x...),
			})
		}
	}()
	wg.Wait()
	sort.Sort(bySum(candidates))
	return candidates
}

// func searchCandidates(ps []int, pd map[int]bool, sieveLim uint64) {
// 	countLim := 500

// 	var wg sync.WaitGroup
// 	count := func(pbegin, pend int, ch chan []int) {
// 		for p := pbegin; p < pend; p++ {
// 			for q := p + 1; q < countLim; q++ {
// 				for r := q + 1; r < countLim; r++ {
// 					for s := r + 1; s < countLim; s++ {
// 						if catenatedPrimes(pd, sieveLim,
// 							ps[p], ps[q], ps[r], ps[s]) {

// 							ch <- []int{ps[p], ps[q], ps[r], ps[s]}
// 						}
// 					}
// 				}
// 			}
// 			log.Println("p now", p)
// 		}
// 		wg.Done()
// 	}

// 	ch := make(chan []int)
// 	wg.Add(4)

// 	go count(0, 62, ch)
// 	go count(62, 125, ch)
// 	go count(125, 250, ch)
// 	go count(250, 500, ch)

// 	go func() {
// 		for x := range ch {
// 			log.Println(x, "; sum =", sum(x...))
// 		}
// 	}()
// 	wg.Wait()
// }

func finalSolver() {

	sieveLim := uint64(1e8)
	ps := primes.Sieve(int(sieveLim))
	pd := primes.SieveDict(int(sieveLim))
	log.Println("Generated sieve")

	// searchCandidates(ps, pd, uint64(sieveLim))

	// tryCandidates(ps, pd, uint64(sieveLim))
	triplets := searchTriplets(ps, pd, sieveLim)

	// We have an upper bound of sum = 200149. What this means is that if we have
	// a triplet candidate, the largest number can be no more than about 66K.
	// 66K + 66k + 66k ~= 200K (or rather the sum of the three can't be > 66K)

	log.Println("Found", len(triplets), "candidate triplets")

	newCandidates := []candidate{}

	for i, t := range triplets {
		// For each triplet, find a set of fourths, as long as the fourth
		// doesn't exceed 200K - sum.
		// i.e. 3 7 11 are triplets (eg), 120K won't work, 200K - 120K
		// is 80K (we are searching in increasing order).
		possibilities := tryCandidates(t, ps, pd, sieveLim)
		if i%10 == 0 {
			log.Println("Trying 3-candidate", i)
		}
		if len(possibilities) == 0 {
			continue // Discard this basically.
		}

		for _, p := range possibilities {
			newCandidates = append(newCandidates, candidate{
				primes: append(t.primes, p),
				sum:    t.sum + p,
			})
		}

	}
	log.Println("Found quad-candidates: count =", len(newCandidates))
	sort.Sort(bySum(newCandidates))
	// Finally, for each set of 4-candidates, find the 5-candidates.
	finalCandidates := []candidate{}

	for i, t := range newCandidates {
		possibilities := tryCandidates(t, ps, pd, sieveLim)
		if i%10 == 0 {
			log.Println("Trying 4-candidate", i)
		}
		if len(possibilities) == 0 {
			continue
		}
		for _, p := range possibilities {
			finalCandidates = append(finalCandidates, candidate{
				primes: append(t.primes, p),
				sum:    t.sum + p,
			})
		}
	}
	sort.Sort(bySum(finalCandidates))

	log.Println("Final candidates:", finalCandidates)

	// This found a new upper bound:
	// [7 1237 2341 12409 18433] 34427
	// 18433 is the 2111st prime
	// 3571 is the 500th prime
	/**
	Final candidates: [{[7 1237 2341 12409 18433] 34427}
	{[7 1237 2341 18433 12409] 34427} {[467 941 2099 19793 25253] 48553}
	{[467 941 2099 25253 19793] 48553} {[733 883 991 55621 18493] 76721}
	{[733 883 991 18493 55621] 76721} {[1109 1889 2999 57089 101681] 164767}
	{[1109 1889 2999 101681 57089] 164767} {[3 37 67 5923 194119] 200149}
	{[3 37 67 194119 5923] 200149}]
	*/

	// final final:
	// [{[5197 5701 6733 8389 13] 26033}
}

func Solve() {
	finalSolver()
	// sieveLim := uint64(1e8)
	// ps := primes.Sieve(int(sieveLim))
	// //pd := primes.SieveDict(int(sieveLim))
	// log.Println("Generated sieve")

	// for i := 0; i < 5000; i++ {
	// 	fmt.Println(i+1, ps[i])
	// }
}
