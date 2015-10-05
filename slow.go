package fft

import (
	"math"
)

// Implementations of the DFT in this file are for testing only.
// These are direct but slow implementations.

// Slow is the simplest and slowest FFT transform.
type slow struct {
}

func (s slow) Transform(x []complex128) []complex128 {
	N := len(x)
	y := make([]complex128, N)
	for k := 0; k < N; k++ {
		for n := 0; n < N; n++ {
			phi := -2.0 * math.Pi * float64(k*n) / float64(N)
			s, c := math.Sincos(phi)
			y[k] += x[n] * complex(c, s)
		}
	}
	return y
}

// SlowPre uses a precomputed roots table.
type slowPre struct {
	E []complex128
	N int
}

func newSlowPre(N int) slowPre {
	var s slowPre
	s.E = roots(N)
	s.N = N
	return s
}

func (s slowPre) Transform(x []complex128) []complex128 {
	if len(x) != len(s.E) {
		panic("SlowPre has been initialized with a long length, or not at all.")
	}
	y := make([]complex128, s.N)
	for k := 0; k < s.N; k++ {
		for n := 0; n < s.N; n++ {
			y[k] += x[n] * s.E[k*n%s.N]
		}
	}
	return y
}
