package p42_test

import (
	"GoCode/Practices57/p42"
	"reflect"
	"testing"
)

type testData struct {
	input  TestIn
	output TestOut
}

type TestIn = p42.Input
type TestOut = p42.Output

func generateTestData() []testData {

	data := []testData{
		{
			input: TestIn{

				RawData: []string{"Ling,Mai,55900",
					"Johnson,Jim,56500",
					"Jones,Aaron,46000",
					"Jones,Chris,34500",
					"Swift,Geoffrey,14200",
					"Xiong,Fong,65000",
					"Zarnecki,Sabrina,51500"},
			},
			output: TestOut{

				ProcessedData: []string{"Ling Mai 55900",
					"Johnson Jim 56500",
					"Jones Aaron 46000",
					"Jones Chris 34500",
					"Swift Geoffrey 14200",
					"Xiong Fong 65000",
					"Zarnecki Sabrina 51500"},
			},
		},
	}

	return data
}

func TestPractice42ormal(t *testing.T) {

	data := generateTestData()

	for _, v := range data {

		var actual TestOut

		actual = p42.Practice42Normal(v.input)

		expected := v.output

		if reflect.DeepEqual(expected, actual) {
			t.Log(true)
		} else {
			t.Error("not equal!!")
			t.Error("input= ", v.input)
			t.Error("actual= ", actual)
			t.Error("expected= ", expected)
		}
	}
}
