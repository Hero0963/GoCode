package p06

import (
	"fmt"
	"time"
)

func RunP06() {

	var curAge, retireAge int

	fmt.Println("What is your current age?")
	fmt.Scanln(&curAge)

	fmt.Println("At what age would you like to retire?")
	fmt.Scanln(&retireAge)

	diff := Practice06Normal(curAge, retireAge)

	if diff > 0 {
		fmt.Println("You have ", diff, "years left until you can retire.")

		t := time.Now()
		yearNow := t.Year()
		yearRetire := yearNow + diff

		fmt.Println("It's ", yearNow, " so you can retire in ", yearRetire)

	} else {
		fmt.Println("You can retire.")
	}

}

func Practice06Normal(curAge int, retireAge int) (diff int) {

	if curAge < 0 {
		fmt.Println(" curAge <0 ", curAge)
		return
	}

	if retireAge < 0 {
		fmt.Println(" retireAge <0 ", retireAge)
		return
	}

	diff = retireAge - curAge
	return diff
}
