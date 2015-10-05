package fft

import "testing"

func TestLastPow2(t *testing.T) {
	tab := []int{2, 4, 8, 16, 32, 64, 128, 256, 512, 1024}
	for i := 2; i <= tab[len(tab)-1]; i++ {
		expectN := 0
		expectP := 0
		for k, N := range tab {
			if i >= N {
				expectN = N
				expectP = k + 1
			}
		}
		n, p, err := lastPow2(i)
		if err != nil {
			t.Errorf("lastPow2: %d: %s\n", i, err)
		}
		if n != expectN || p != expectP {
			t.Errorf("lastPow2(%d), got: %d %d, expected: %d %d\n", i, n, p, expectN, expectP)
		}
	}
}
