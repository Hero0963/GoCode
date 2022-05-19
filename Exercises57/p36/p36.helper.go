package p36

import "math"

func finishEnteringNumbers(s string) bool {
	if s == "done" {
		return true
	}

	return false
}

func getMean(numbers []float64) (mean float64) {

	l := len(numbers)

	if l == 0 {
		return
	}

	var sum float64

	for _, v := range numbers {
		sum += v
	}

	mean = sum / float64(l)
	return
}

func getMaxNumber(numbers []float64) (max float64) {

	l := len(numbers)

	if l == 0 {
		return
	}

	max = numbers[0]

	for _, v := range numbers {
		if v > max {
			max = v
		}
	}

	return
}

func getMinNumber(numbers []float64) (min float64) {

	l := len(numbers)

	if l == 0 {
		return
	}

	min = numbers[0]

	for _, v := range numbers {
		if v < min {
			min = v
		}
	}

	return
}

func getStandardDeviation(numbers []float64) (stdD float64) {

	mean := getMean(numbers)
	mss := getMeanSquareSum(numbers)

	squareOfStdD := mss - mean*mean

	if squareOfStdD < 0 {
		return
	}

	stdD = math.Sqrt(squareOfStdD)

	tmp := int(stdD * 100)
	stdD = float64(tmp) / 100.0

	return
}

func getMeanSquareSum(numbers []float64) (mss float64) {

	l := len(numbers)

	if l == 0 {
		return
	}

	for _, v := range numbers {
		mss += v * v
	}

	mss /= float64(l)

	return

}
