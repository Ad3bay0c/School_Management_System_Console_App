package staff

import (
	"github.com/Ad3bay0c/SchoolManagement/school"
	"github.com/Ad3bay0c/SchoolManagement/students"
	"testing"
	"time"
)

//var stud = new(students.Student)
var s = []students.Student{
	{
		ID: 1,
		FirstName: "Clin",
		LastName: "Ade",
		Age: 15,
		Class: school.SSS1,
	},
	{
		ID: 2,
		FirstName: "Clin",
		LastName: "Ade",
		Age: 15,
		Class: school.SSS2,
	},
	{
		ID: 3,
		FirstName: "Clin",
		LastName: "Ade",
		Age: 15,
		Class: school.SSS3,
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
				Age: 10,
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
			output: 5,
			word: "Admit Student 2",
		},
	}

	for _, applicant := range applicants {
		t.Run(applicant.word, func(t *testing.T) {
			result, err := principal.Admit(applicant.input, s)
			if result != applicant.output {
				t.Errorf("%v\n", err)
			}
		})
	}
}