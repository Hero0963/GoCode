package p23

import "fmt"

func isLegalInput(s string) bool {

	if s == "y" || s == "n" {
		return true
	}

	printWrongInputMessage()
	return false

}

func isInputYes(s string) bool {
	if s == "y" {
		return true
	}

	return false
}

func printWrongInputMessage() {
	fmt.Println("Wrong input !")
}

func carTroubleShooting(s string) {

	if !isLegalInput(s) {
		return
	}

	if isInputYes(s) {
		carSilentWhenTurnTheKey()
	} else {
		carNotSilentWhenTurnTheKey()
	}

}

func carSilentWhenTurnTheKey() {

	fmt.Println("Are the battery terminals corroded?")
	printInputRemind()
	var str string
	fmt.Scanln(&str)

	if !isLegalInput(str) {
		return
	}

	if isInputYes(str) {
		batteryTerminalsCorroded()
	} else {
		batteryTerminalsNotCorroded()
	}
}

func carNotSilentWhenTurnTheKey() {
	fmt.Println("Does the car make a clicking noise?")
	printInputRemind()
	var str string
	fmt.Scanln(&str)

	if !isLegalInput(str) {
		return
	}

	if isInputYes(str) {
		carMakeAClickingNoise()
	} else {
		carNotMakeAClickingNoise()
	}

}

func batteryTerminalsCorroded() {
	fmt.Println("Clean terminals and try starting again.")
}

func batteryTerminalsNotCorroded() {
	fmt.Println("Replace cables and try again.")
}

func carMakeAClickingNoise() {
	fmt.Println("Replace the battery")
}

func carNotMakeAClickingNoise() {
	fmt.Println("Does the car crank up but fail to start?")
	printInputRemind()
	var str string
	fmt.Scanln(&str)

	if !isLegalInput(str) {
		return
	}

	if isInputYes(str) {
		carCrankUpButFailToStart()
	} else {
		carCrankUpNotFailToStart()
	}
}

func carCrankUpButFailToStart() {
	fmt.Println("Check spark plug connections.")
}

func carCrankUpNotFailToStart() {
	fmt.Println("Does the engine start and then die?")
	printInputRemind()
	var str string
	fmt.Scanln(&str)

	if !isLegalInput(str) {
		return
	}

	if isInputYes(str) {
		engineStartAndThenDie()
	} else {
		engineStartAndThenNotDie()
	}
}

func engineStartAndThenDie() {
	fmt.Println("Does your car have fuel injection?")
	printInputRemind()
	var str string
	fmt.Scanln(&str)

	if !isLegalInput(str) {
		return
	}

	if isInputYes(str) {
		carHasFuelInjection()
	} else {
		carHasNoFuelInjection()
	}

}

func engineStartAndThenNotDie() {
	fmt.Println("Sorry, we have no idea.")
}

func carHasFuelInjection() {
	fmt.Println("Get it in for service.")
}

func carHasNoFuelInjection() {
	fmt.Println("Check to ensure the choke is opening and closing.")
}

func printInputRemind() {
	fmt.Println("Please enter y/n.")
}
