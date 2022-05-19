package p32

import (
	"math/rand"
	"time"
)

func isLegalLevel(level int) bool {
	if level < 1 || level > LevelMax {
		return false
	}

	return true
}

func getGaolNumber(level int) (n int) {

	var min int = 1
	var max int = 1

	for i := 0; i < level; i++ {
		max *= 10
	}

	//fmt.Println("max=", max)

	rand.Seed(time.Now().UnixNano())
	n = rand.Intn(max-min+1) + min

	//fmt.Println("n= ", n)

	return

}

func getCountUnitSuffix(count int) string {

	if count == 1 {
		return " time."
	}

	return " times."

}

func willPlayAgain(s string) bool {

	if s == "Y" || s == "y" {
		return true
	}

	return false
}
