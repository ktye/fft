package fft

import (
	"math/cmplx"
	"testing"
)

func TestInverse(t *testing.T) {
	N := 256
	x := complexRand(N)
	f, err := New(N)
	if err != nil {
		t.Error(err)
	}
	y := f.Transform(copyVector(x))
	y = f.Inverse(y)
	for i := range x {
		if e := cmplx.Abs(x[i] - y[i]); e > 1E-9 {
			t.Errorf("inverse differs %d: %v %v\n", i, x[i], y[i])
		}
	}
}
