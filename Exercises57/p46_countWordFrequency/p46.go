package p46

import "fmt"

type Input struct {
}

type Output struct {
}

type StatData = map[string]int

func RunP46() {

	fmt.Println("This program will show how many times each string appearence : ")

	Practice46Normal()

	fmt.Println("Finished.")

}

func Practice46Normal() {

	statData := GetData()
	SaveStatData(statData)

	return
}
