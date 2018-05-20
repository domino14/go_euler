package sixtyone

import (
	"fmt"
)

type intSet map[int]bool

type polygonalNumbers struct {
	order   int
	numbers map[int]intSet
}

type candidateSet struct {
	set   intSet
	order int
}

type candidate struct {
	order int
	num   int
}

func addToMap(gen int, m map[int]intSet) {
	// modify m in place.

	prefix := gen / 100

	if _, ok := m[prefix]; !ok {
		m[prefix] = make(map[int]bool)
	}
	m[prefix][gen] = true
}

func find4dig(order int) polygonalNumbers {
	// Return a map of 2-digit prefixes to lists of 4-digit integers
	// with those prefixes.
	out := make(map[int]intSet)
	i := 1
	gen := 0
	for gen < 10000 {
		switch order {
		case 3:
			gen = i * (i + 1) / 2
		case 4:
			gen = i * i
		case 5:
			gen = i * (3*i - 1) / 2
		case 6:
			gen = i * (2*i - 1)
		case 7:
			gen = i * (5*i - 3) / 2
		case 8:
			gen = i * (3*i - 2)
		}

		if gen >= 1010 && gen < 10000 {
			addToMap(gen, out)
		}

		i++
	}

	return polygonalNumbers{
		order:   order,
		numbers: out,
	}
}

func Solve() {

	// Find all 4-digit polygonal numbers
	figurates := map[int]polygonalNumbers{}

	for i := 3; i <= 8; i++ {
		figurates[i] = find4dig(i)
	}

	order := 3
	for _, val := range figurates[order].numbers {
		visited := map[int]bool{}
		visited[order] = true
		solveWithRecurse(val, order, visited, []candidate{}, figurates, 1)
	}

}

func sum(candidates ...candidate) int {
	sum := 0
	for _, c := range candidates {
		sum += c.num
	}
	return sum
}

func solveWithRecurse(values intSet, order int, visited map[int]bool,
	candidates []candidate, figurates map[int]polygonalNumbers, level int) {

	for val := range values {

		setToVisit := []polygonalNumbers{}
		for order, f := range figurates {
			if _, ok := visited[order]; !ok {
				setToVisit = append(setToVisit, f)
			}
		}

		c := append(candidates, candidate{order: order, num: val})

		if len(setToVisit) == 0 {
			// We have already visited all sets.
			if c[0].num/100 == c[len(c)-1].num%100 {
				fmt.Println("Candidates", c, sum(c...))
			}
			continue
		}

		loopableSet := findInSet(val%100, setToVisit...)
		if len(loopableSet) == 0 {
			// Fail, not found. Try the next value.
			continue
		}
		for _, possible := range loopableSet {
			// Create a new visited map.
			newVisited := map[int]bool{}
			for k, v := range visited {
				newVisited[k] = v
			}
			newVisited[possible.order] = true

			solveWithRecurse(possible.set, possible.order, newVisited, c,
				figurates, level+1)
		}
	}

}

func findInSet(num int, polys ...polygonalNumbers) []candidateSet {
	cs := []candidateSet{}
	for _, p := range polys {
		if set, ok := p.numbers[num]; ok {
			cs = append(cs, candidateSet{
				set: set, order: p.order,
			})
		}
	}
	return cs
}
