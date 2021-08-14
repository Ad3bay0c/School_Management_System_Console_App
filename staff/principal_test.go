package staff

import (
	"github.com/Ad3bay0c/SchoolManagement/school"
	"github.com/Ad3bay0c/SchoolManagement/students"
	"testing"
	"time"
)

//var stud = new(students.Student)
var stud = map[int]students.Student{
		1: {
			FirstName: "Clin",
			LastName:  "Ade",
			Age:       15,
			Class:     school.SSS1,
		},
		2: {
			ID:        2,
			FirstName: "Clint",
			LastName:  "Ade",
			Age:       15,
			Class:     school.SSS2,
		},
		3: {
			ID:        3,
			FirstName: "Clinto",
			LastName:  "Ade",
			Age:       15,
			Class:     school.SSS3,
		},
}

var principal = new(Principal)

func TestPrincipal_Admit(t *testing.T) {
	principal.ID = int(time.Now().Unix())
	principal.FirstName = "Mr Dada"
	principal.LastName = "O"
	principal.Position = school.PRINCIPAL

	applicants := []struct{
		input students.Student
		output int
		word string
	} {
		{
			input: students.Student{
				FirstName: "Bamidele",
				LastName: "Johnson",
				Age: 12,
			},
			output: 4,
			word: "Admit Student 1",
		},
		{
			input: students.Student{
				FirstName: "Bami",
				LastName: "John",
				Age: 0,
			},
			output: 0,
			word: "Admit Student 2",
		},
		{
			input: students.Student{
				FirstName: "Bami",
				LastName: "John",
				Age: 8,
			},
			output: 5,
			word: "Admit Student 3",
		},
		{
			input: students.Student{
				FirstName: "Bami",
				LastName: "John",
				Age: 9,
			},
			output: 6,
			word: "Admit Student 4",
		},
		{
			input: students.Student{
				FirstName: "Bami",
				LastName: "John",
				Age: 10,
			},
			output: 7,
			word: "Admit Student 5",
		},
		{
			input: students.Student{
				FirstName: "Bami",
				LastName: "John",
				Age: 11,
			},
			output: 8,
			word: "Admit Student 6",
		},
		{
			input: students.Student{
				FirstName: "Bami",
				LastName: "John",
				Age: 13,
			},
			output: 10,
			word: "Admit Student 7",
		},
	}

	for _, applicant := range applicants {
		t.Run(applicant.word, func(t *testing.T) {
			result, err, pid := principal.Admit(applicant.input, &stud)
			if result != applicant.output && principal.ID != pid{
				t.Errorf("%v %v\n", err)
			}
		})
	}
}