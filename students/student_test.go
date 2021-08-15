package students

import (
	"testing"
)

func TestStudent_TakeCourse(t *testing.T) {
	tables := []struct{
		input 	Student
		output 	bool
		word   	string
	}{
		{
			input: Students[2],
			output: true,
			word: "True Value",
		},
		{
			input: Students[1],
			output: true,
			word: "False Value",
		},
	}

	for _, table := range tables {
		t.Run(table.word, func(t *testing.T) {
			result := table.input.TakeCourse()

			if result != table.output {
				t.Errorf("Expected %v; Got %v", table.output, result)
			}
		})
	}
}

func TestStudent_CheckGrade(t *testing.T) {
	tables := []struct{
		input 	Student
		output 	int
		word   	string
	}{
		{
			input: Students[1],
			output: 0,
			word: "Check Grade 1",
		},
		{
			input: Students[2],
			output: 0,
			word: "Check Grade 2",
		},
	}

	for _, table := range tables {
		t.Run(table.word, func(t *testing.T) {
			result := table.input.CheckGrade()

			if len(result) != table.output {
				t.Errorf("Expected %v; Got %v", table.output, len(result))
			}
		})
	}
}
