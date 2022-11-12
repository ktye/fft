package fft

import (
	"math/cmplx"
	"math/rand"
	"testing"
)

// Run the benchmark against direct implementations with:
// go test -bench=.
//
// A benchmark against other more sophisticated implementations would be nice.

func complexRand(N int) []complex128 {
	x := make([]complex128, N)
	for i := 0; i < N; i++ {
		x[i] = complex(2.0*rand.Float64()-1.0, 2.0*rand.Float64()-1.0)
	}
	return x
}

func copyVector(v []complex128) []complex128 {
	y := make([]complex128, len(v))
	copy(y, v)
	return y
}

func TestFFT(t *testing.T) {
	N := 1024
	x := complexRand(N)
	slow := slow{}
	slowPre := newSlowPre(N)
	fast, err := New(N)
	if err != nil {
		t.Error(err)
	}
	faster := Prepare(N)

	y1 := slow.Transform(copyVector(x))
	y2 := slowPre.Transform(copyVector(x))
	y3 := fast.Transform(copyVector(x))
	y4 := copyVector(x); faster.Complex(y4)
	for i := 0; i < N; i++ {
		if e := cmplx.Abs(y1[i] - y2[i]); e > 1E-9 {
			t.Errorf("slow and slowPre differ: i=%d diff=%v\n", i, e)
		}
		if e := cmplx.Abs(y1[i] - y3[i]); e > 1E-9 {
			t.Errorf("slow and fast differ: i=%d diff=%v\n", i, e)
		}
	}
}

func BenchmarkSlow000(t *testing.B) {
	N := 8192

	slow := slow{}
	x := complexRand(N)

	for i := 0; i < t.N; i++ {
		_ = slow.Transform(x)
	}
}

func BenchmarkSlowPre(t *testing.B) {
	N := 8192

	slowPre := newSlowPre(N)
	x := complexRand(N)

	for i := 0; i < t.N; i++ {
		_ = slowPre.Transform(x)
	}
}

func BenchmarkFast000(t *testing.B) {
	N := 8192

	fast, err := New(N)
	if err != nil {
		t.Error(err)
	}
	x := complexRand(N)

	for i := 0; i < t.N; i++ {
		_ = fast.Transform(copyVector(x))
	}
}

func BenchmarkFaster0(t *testing.B) {
	N := 8192

	faster := Prepare(N)
	x := complexRand(N)

	for i := 0; i < t.N; i++ {
		faster.Complex(copyVector(x))
	}
}
