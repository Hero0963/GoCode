package p31

import "fmt"

func showIntensityRateGraph(age int, rHR int) {

	fmt.Println("Intensity,Rate:")

	var tHR int

	for intensity := IntensityLB; intensity <= IntensityUB; intensity += IntensityGap {

		tHR = calTargetHeartRate(age, rHR, intensity)

		fmt.Println(intensity, "%    ", tHR)

	}

}

func calTargetHeartRate(age int, rHR int, intensity int) (tHR int) {

	var tmp float64

	tmp = float64(((220-age)-rHR)*intensity)*0.01 + float64(rHR)

	tHR = int(tmp)

	return tHR
}
