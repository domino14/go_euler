package sixtytwo

import (
	"reflect"
	"testing"
)

func TestDecomposeInt(t *testing.T) {
	eq := reflect.DeepEqual(
		decomposeInt(12334309553),
		[]int{1, 2, 3, 3, 4, 3, 0, 9, 5, 5, 3})
	if !eq {
		t.Errorf("Not equal: %v", decomposeInt(123353))
	}
}

func TestJoinIntoInt(t *testing.T) {
	joined := joinIntoInt([]int{1, 2, 3, 3, 4, 3, 0, 9, 5, 5, 3})
	if joined != 12334309553 {
		t.Errorf("Failed: %v", joined)
	}
}

func TestPermsAreCubes(t *testing.T) {
	populateCubes()
	np := permutationsAreCubes(41063625)
	if np != 3 {
		t.Errorf("Failed: %v", np)
	}

	np = permutationsAreCubes(1000000)
	if np != 1 {
		t.Errorf("Failed: %v", np)
	}
}
