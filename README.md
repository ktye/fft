# fft
radix-2 fast Fourier transform

Package fft provides a fast discrete Fourier transformation algorithm.

Implemented is the 1-dimensional DFT of complex input data
for with input lengths which are powers of 2.

The algorithm is non-recursive and works in-place overwriting
the input array.

Before doing the transform on acutal data, allocate
an FFT object with t := fft.New(N) where N is the length of the
input array.
Then multiple calls to t.Transform(x) can be done with
different input vectors having the same length.
