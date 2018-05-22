package sixtytwo

// from https://gist.github.com/schwern/93f7d93c1c80f8720e6de85bdc92096f
// Based on https://stackoverflow.com/a/30230552/14660 by Paul Hankin

// perm := NewPermutation( slice )
// for next := perm.Next(); next != nil; next = perm.Next() {
//    ...
// }
//
// or
//
// permCh := day09.NewPermutationChan( test.arg )
// for next := range ch {
//     ...
// }

type Permutation struct {
	orig []int
	perm []int
}

func NewPermutation(orig []int) *Permutation {
	self := Permutation{orig: orig, perm: make([]int, len(orig))}
	return &self
}

func NewPermutationChan(orig []int) chan []int {
	perm := NewPermutation(orig)
	ch := make(chan []int)

	go func() {
		defer close(ch)
		for next := perm.Next(); next != nil; next = perm.Next() {
			ch <- next
		}
	}()

	return ch
}

func (self *Permutation) nextPerm() {
	for i := len(self.perm) - 1; i >= 0; i-- {
		if i == 0 || self.perm[i] < len(self.perm)-i-1 {
			self.perm[i]++
			return
		}
		self.perm[i] = 0
	}
}

func (self *Permutation) Next() []int {
	defer self.nextPerm()

	if len(self.perm) == 0 || self.perm[0] >= len(self.perm) {
		return nil
	}

	result := append([]int{}, self.orig...)
	for i, v := range self.perm {
		result[i], result[i+v] = result[i+v], result[i]
	}
	return result
}
