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
