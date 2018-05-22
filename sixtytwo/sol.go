package sixtytwo

import (
	"fmt"
	"sort"
	"strconv"
)

const (
	Limit = 20000
)

// A map of "alphagrams" to slices of cube numbers
var alphas map[string][]string

func alphagrammize(n uint64) string {
	// Sort alphabetically.
	x := decomposeInt(n)
	sort.Ints(x)
	return joinIntoString(x)
}

// // Decompose into an array of ints
func decomposeInt(n uint64) []int {
	ret := []int{}
	rem := uint64(0)
	for n != 0 {
		n, rem = n/10, n%10
		ret = append([]int{int(rem)}, ret...)
	}
	return ret
}

func joinIntoString(arr []int) string {
	str := ""
	for _, d := range arr {
		str += strconv.Itoa(d)
	}
	return str
}

func Solve() {
	alphas = make(map[string][]string)
	for i := uint64(1); i < Limit; i++ {
		c := i * i * i
		alph := alphagrammize(c)
		if _, ok := alphas[alph]; !ok {
			alphas[alph] = []string{}
		}
		alphas[alph] = append(alphas[alph], joinIntoString(decomposeInt(c)))
		sort.Strings(alphas[alph])

		if i%100000 == 0 {
			fmt.Println("Computing...", i)
		}
	}
	final := []string{}
	// Then go through the map
	for _, sl := range alphas {
		if len(sl) == 5 {
			final = append(final, sl[0])
		}

	}
	sort.Strings(final)
	//fmt.Println(final)
	fmt.Println(final[0])
}
