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
