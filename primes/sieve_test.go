package primes

import (
	"reflect"
	"testing"
)

func TestSieve(t *testing.T) {
	x := Sieve(25)
	if !reflect.DeepEqual(x, []int{2, 3, 5, 7, 11, 13, 17, 19, 23}) {
		t.Errorf("Not equal: %v", x)
	}
}

func TestSieveDict(t *testing.T) {
	x := SieveDict(25)
	if !reflect.DeepEqual(x, map[int]bool{
		2: true, 3: true, 5: true, 7: true, 11: true, 13: true,
		17: true, 19: true, 23: true}) {
		t.Errorf("Not equal: %v", x)
	}
}

type primetest struct {
	num uint64
	p   bool
}

var primeTests = []primetest{
	{0, false},
	{1, false},
	{2, true},
	{3, true},
	{4, false},
	{123234324, false},
	{17, true},
	{49, false},
	{47, true},
	{193, true},
	{195, false},
	{197, true},
	{199, true},
	{179426549, true},
	{32416190071, true},
}

func TestCheckPrime(t *testing.T) {
	for _, pair := range primeTests {
		p := CheckPrime(pair.num)
		if p != pair.p {
			t.Error("For", pair.num, "expected", pair.p, "got", p)
		}
	}
}
