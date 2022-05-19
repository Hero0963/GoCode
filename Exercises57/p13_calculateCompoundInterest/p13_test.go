package p13_test

import (
	"reflect"
	"selfTesting/chall57/p13"
	"testing"
)

type testData struct {
	input  TestIn
	output TestOut
}

type TestIn = p13.Input
type TestOut = p13.Output

func generateTestData() []testData {

	data := []testData{
		testData{
			input: TestIn{
				Principal:     1500,
				Rate:          4.3,
				NumberOfYears: 6,
				InterestTimes: 4,
			},
			output: TestOut{
				Worth: 1938.83,
			},
		},
	}

	return data
}

func TestPractice13Normal(t *testing.T) {

	data := generateTestData()

	for _, v := range data {

		var actual TestOut

		actual = p13.Practice13Normal(v.input)

		expected := v.output

		if reflect.DeepEqual(expected, actual) {
			t.Log(true)
		} else {
			t.Error("not equal!!")
			t.Error("actual= ", actual)
			t.Error("expected= ", expected)
		}
	}
}
