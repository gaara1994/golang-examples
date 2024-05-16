package cpu

func CalculatePi(n int) float64 {
	pi := 0.0
	sign := 1.0
	for i := 0; i < n; i++ {
		pi += sign / float64(2*i + 1)
		sign = -sign
	}
	return pi * 4
}
