package p12_test

import (
	"GoCode/Practices57/p12"
	"reflect"
	"testing"
)

type testData struct {
	input  TestIn
	output TestOut
}

type TestIn = p12.Input
type TestOut = p12.Output

func generateTestData() []testData {

	data := []testData{
		testData{
			input: TestIn{
				Principal:      1500,
				RateOfInterest: 4.3,
				NumberOfYears:  4,
			},
			output: TestOut{
				Worth: 1758,
			},
		},
	}

	return data
}

func TestPractice10Normal(t *testing.T) {

	data := generateTestData()

	for _, v := range data {

		var actual TestOut

		actual = p12.Practice12Normal(v.input)

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
