package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

var M int

func main() {
	n, e := strconv.Atoi(os.Args[1])
	if e != nil {
		panic("unroll n")
	}
	O("#include<complex.h>\n")
	O("#define z double complex\n")
	M = 1<<n
	O("int fft%d(z *x){\n", M)
	O("z p,q;\n")
	for i := 0; i<M; i++ {
		O("z y%d;\n", i)
	}
	f(M, 1, 0)
	O("}\n")
}

// func O(f string, a ...interface{}) { fmt.Printf(f, a...) }
var O = fmt.Printf

func f(N, s, o int) {
	if N == 1 {
		O("z x1_%d=x[%d];\n", o, o)
		return
	}
	f(N/2, 2*s, o)
	for i := 0; i < N/2; i++ {
		O("y%d=x%d_%d;\n", i, N/2, i+o)
	}
	f(N/2, 2*s, o+s)
	for i := 0; i < N/2; i++ {
		O("y%d=x%d_%d;\n", i+N/2, N/2, i+o+s)
	}
	for i := 0; i < N/2; i++ {
		O("p=y%d;\n", i)
		re := math.Cos(2.0 * math.Pi * float64(i) / float64(N))
		im := math.Sin(2.0 * math.Pi * float64(i) / float64(N))
		re = small(re)
		im = small(im)
		O("q=((%v)*I*(%v))*y%d;\n", re, -im, i+N/2)
		if N == M {
			O("x[%d]=p+q;\n", i+o)
			O("x[%d]=p-q;\n", i+o+N/2)
		} else {
			O("z x%d_%d=p+q;\n", N, i+o)
			O("z x%d_%d=p-q;\n", N, i+o+N/2)
		}
	}
}
func small(f float64) float64 {
	if math.Abs(f) < 1e-16 {
		return 0
	}
	return f
}
