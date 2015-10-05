package fft

import (
	"reflect"
	"testing"
)

func TestBitReversal(t *testing.T) {
	tab := [][]int{
		[]int{0},
		[]int{0, 1},
		[]int{0, 2, 1, 3},
		[]int{0, 4, 2, 6, 1, 5, 3, 7},
		[]int{0, 8, 4, 12, 2, 10, 6, 14, 1, 9, 5, 13, 3, 11, 7, 15},
	}
	for i := 0; i < len(tab); i++ {
		got := permutationIndex(i)
		//		got := getRevIndex(i)
		expect := tab[i]
		if !reflect.DeepEqual(got, expect) {
			t.Errorf("%d expected: %v, got: %v\n", i, expect, got)
		}
	}
}
