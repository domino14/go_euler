package sixtytwo

import "fmt"

var cubes map[uint64]uint64 // key is the cube, value is the root
const (
	Limit = 100000
)

func populateCubes() {
	cubes = make(map[uint64]uint64)
	var i uint64
	for i = 1; i < Limit; i++ {
		cubes[i*i*i] = i
	}
}

// Decompose into an array of ints
func decomposeInt(n uint64) []int {
	ret := []int{}
	rem := uint64(0)
	for n != 0 {
		n, rem = n/10, n%10
		ret = append([]int{int(rem)}, ret...)
	}
	return ret
}

func joinIntoInt(arr []int) uint64 {
	x := uint64(0)
	o := 1
	for i := len(arr) - 1; i >= 0; i-- {
		x += uint64(o) * uint64(arr[i])
		o *= 10
	}
	return x
}

func permutationsAreCubes(cube uint64) int {
	// Return how many permutations are cubes.
	orig := decomposeInt(cube)

	p := NewPermutation(orig)

	uniqueCubes := map[uint64]bool{}

	for next := p.Next(); next != nil; next = p.Next() {
		perm := joinIntoInt(next)
		if _, found := cubes[perm]; found {
			// Let's not count leading zeros, as the PE problem doesn't
			// seem to count them properly. (otherwise 1000000 should be
			// the first answer: 0000001, 0010000, and 1000000 are cubes)
			if next[0] != 0 {
				uniqueCubes[perm] = true
			}
		}
	}
	return len(uniqueCubes)
}

func Solve() {
	populateCubes()
	for i := uint64(1); i < Limit; i++ {
		cube := i * i * i
		if permutationsAreCubes(cube) == 5 {
			fmt.Println(cube)
			break
		}
		if i%10 == 0 {
			fmt.Println("Checking", i, cube)
		}
	}

}
