package fft

import (
	"math"
	"math/bits"
)

type NFFT struct {
	n, h, l uint16
	p       []uint16
	e       [][]complex128
	i       [][]uint16
}

func Prepare(n uint16) NFFT { //n: power of two
	l := uint16(bits.TrailingZeros16(n))
	e := make([][]complex128, l)
	i := make([][]uint16, l)
	r := rots(n)
	p := perm(n)
	s := n
	h := n >> 1
	t := uint16(1)
	for k := range e {
		E := make([]complex128, h)
		I := make([]uint16, h)
		s >>= 1
		c := 0
		for b := uint16(0); b < s; b++ {
			o := 2 * b * t
			for j := uint16(0); j < t; j++ {
				I[c] = j + o
				E[c] = r[s*j]
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
	for i, el := range f.e {
		l := f.i[i]
		for k := uint16(0); k < f.h; k++ {
			ii := l[k]
			jj := uint16(i) + ii
			xi := x[ii]
			xj := x[jj]
			ek := el[k]
			x[ii] += xj * ek
			x[jj] = xi - xj*ek
		}
	}
}
func (f NFFT) Real2(x, y []float64, out []complex128) {
	for i := uint16(0); i < f.n; i++ {
		out[i] = complex(x[i], y[i])
	}
	f.Complex(out)
	// X, Y (complex) from real2:
	//for i := uint16(0); i<f.n; i++ {
	//	k := f.n - i
	//	X[i] = 0.5*(out[i] + out[k])
	//	Y[i] = 0.5*(out[i] - out[k]) // factor -i omitted
	//}
}
func perm(n uint16) []uint16 {
	r := make([]uint16, n)
	k := uint16(1)
	for n > 1 {
		n >>= 1
		for i := uint16(0); i < k; i++ {
			r[i] <<= 1
			r[i+k] = 1 + r[i]
		}
		k <<= 1
	}
	return r
}
func rots(N uint16) []complex128 {
	E := make([]complex128, N)
	for n := uint16(0); n < N; n++ {
		phi := -2.0 * math.Pi * float64(n) / float64(N)
		s, c := math.Sincos(phi)
		E[n] = complex(c, s)
	}
	return E
}
