package maths

func Pow(base int, times int) int {
	if times == 0 {
		return 1
	} else {
		return base * Pow(base, times-1)
	}

}

func Average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

// Yeah, yeah, I know
func HalfTheNumberAndCheckIfTheHalfIsEven(number int) (x int, isEven bool) {
	x = number / 2
	if x%2 == 0 {
		return x, false
	}
	return x, true
}
