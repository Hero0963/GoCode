package p46_test

import (
	"GoCode/Practices57/p45"
	"testing"
)

type testData struct {
	input  TestIn
	output TestOut
}

type TestIn = p45.Input
type TestOut = p45.Output

func generateTestData() []testData {

	data := []testData{
		{
			input:  TestIn{},
			output: TestOut{},
		},
	}

	return data
}

func TestPractice45ormal(t *testing.T) {

}
