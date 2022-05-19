package p26

import "math"

func getMonthUnit(NeedMonth int) string {

	if NeedMonth == 1 {
		return "month"
	}

	return "months"

}

func getPowerHelper(x float64) float64 {

	return math.Pow(x, 30)

}
