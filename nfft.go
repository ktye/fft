package fft

import (
	"math"
	"math/bits"
)

type NFFT struct {
	n, h, l int
	p       []int
	e       [][]complex128
	i       [][]int
	z       []complex128
}

func Prepare(n int) NFFT { //n: power of two
	l := bits.TrailingZeros(uint(n))
	e := make([][]complex128, l)
	i := make([][]int, l)
	r := rots(n)
	p := perm(n)
	s := n
	h := n >> 1
	t := 1
	for k := range e {
		E := make([]complex128, n)
		I := make([]int, n)
		s >>= 1
		c := 0
		for b := 0; b < s; b++ {
			o := 2 * b * t
			for j := 0; j < t; j++ {
				I[c] = j + o
				I[c+h] = j + o + t
				E[c] = r[s*j]
				E[c+h] = r[s*(j+t)]
				c++
			}
		}
		e[k] = E
		i[k] = I
		t <<= 1
	}
	return NFFT{n: n, h: h, l: l, p: p, e: e, i: i}
}
func (f NFFT) Complex(x []complex128) {
	h := f.h
	for i, el := range f.e {
		l := f.i[i]
		for k := 0; k < f.h; k++ {
			ii := l[k]
			jj := l[h+k]
			//xi := x[ii]
			//xj := x[jj]
			//x[ii] += xj * el[k]
			//x[jj] = xi + xj*el[k+h]
			x[ii], x[jj] = x[ii] + x[jj] * el[k], x[ii]+x[jj]*el[k+h]
			//x[jj] = xi + xj*el[k+h]
		}
	}
}
func perm(n int) []int {
	r := make([]int, n)
	k := 1
	for n > 1 {
		n >>= 1
		for i := 0; i < k; i++ {
			r[i] <<= 1
			r[i+k] = 1 + r[i]
		}
		k <<= 1
	}
	return r
}
func rots(N int) []complex128 {
	E := make([]complex128, N)
	for n := 0; n < N; n++ {
		phi := -2.0 * math.Pi * float64(n) / float64(N)
		s, c := math.Sincos(phi)
		E[n] = complex(c, s)
	}
	return E
}
