package p39_test

import (
	"GoCode/Practices57/p39"
	"reflect"
	"testing"
)

type testData struct {
	input  []TestIn
	output []TestOut
}

type TestIn = p39.Input
type TestOut = p39.Output

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

				{
					FullName:       "John Johnson",
					Position:       "Manager",
					SeparationDate: "2016-12-31",
				},

				{
					FullName:       "Michaela Michaelson",
					Position:       "District Manager",
					SeparationDate: "2015-12-19",
				},

				{
					FullName:       "Sally Webber",
					Position:       "Web Developer",
					SeparationDate: "2015-12-18",
				},

				{
					FullName:       "Tou Xiong",
					Position:       "Software Engineer",
					SeparationDate: "2016-10-05",
				},
			},
		},
	}

	return data
}

func TestPractice39Normal(t *testing.T) {

	data := generateTestData()

	for _, v := range data {

		var actual []TestOut

		actual = p39.Practice39Normal(v.input)

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
