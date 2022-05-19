package p18

func isLegalTemperatureUnit(s string) bool {

	if s == "F" {
		return true
	}

	if s == "C" {
		return true
	}

	return false
}

func getDualTemperatureUnit(s string) (str string) {
	if s == "C" {
		str = TempF
	}

	if s == "F" {
		str = TempC
	}

	return
}

func getTemperatureUnit(s string) (str string) {
	if s == "C" {
		str = TempC
	}

	if s == "F" {
		str = TempF
	}

	return
}

func convertFToC(v float64) (val float64) {

	val = (v - convertConst1) * convertConst2

	return

}

func convertCToF(v float64) (val float64) {

	val = v/(convertConst2) + convertConst1

	return

}
