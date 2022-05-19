package p40_test

import (
	"GoCode/Practices57/p40"
	"reflect"
	"testing"
)

type testData struct {
	input  []TestIn
	output []TestOut
}

type TestIn = p40.Input
type TestOut = p40.Output

func generateTestData() []testData {

	data := []testData{
		{
			input: []TestIn{
				{
					FirstName:      "John",
					LastName:       "Johnson",
					Position:       "Manager",
					SeparationDate: "2016-12-31",
				},

				{
					FirstName:      "Tou",
					LastName:       "Xiong",
					Position:       "Software Engineer",
					SeparationDate: "2016-10-05",
				},

				{
					FirstName:      "Michaela",
					LastName:       "Michaelson",
					Position:       "District Manager",
					SeparationDate: "2015-12-19",
				},

				{
					FirstName:      "Jake",
					LastName:       "Jacobson",
					Position:       "Programmer",
					SeparationDate: "2021-02-01",
				},

				{
					FirstName:      "Jacquelyn",
					LastName:       "Jackson",
					Position:       "DBA",
					SeparationDate: "2021-02-01",
				},

				{
					FirstName:      "Sally",
					LastName:       "Webber",
					Position:       "Web Developer",
					SeparationDate: "2015-12-18",
				},
			},
			output: []TestOut{
				{
					FullName:       "Jacquelyn Jackson",
					Position:       "DBA",
					SeparationDate: "2021-02-01",
				},

				{
					FullName:       "Jake Jacobson",
					Position:       "Programmer",
					SeparationDate: "2021-02-01",
				},
			},
		},
	}

	return data
}

func TestPractice40Normal(t *testing.T) {

	data := generateTestData()

	for _, v := range data {

		var actual []TestOut

		actual = p40.Practice40Normal(v.input, p40.TestSearchStr)

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
