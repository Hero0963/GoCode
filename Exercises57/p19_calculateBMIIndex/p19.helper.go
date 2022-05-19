package p19

func calBMI(height float64, weight float64) (bmi float64) {

	bmi = (weight / (height * height)) * BMIConvertConst1

	tmp := int(bmi * 100)
	bmi = float64(tmp) / 100.0

	return
}

func checkBMIRange(bmi float64) (s string) {

	s = "Ideal Weight"

	if bmi < BMILowerBound {
		s = "UnderWeight"
	}

	if bmi > BMIUpperBound {
		s = "OverWeight"
	}

	return
}
