package staff

import (
	"github.com/Ad3bay0c/SchoolManagement/school"
	"github.com/Ad3bay0c/SchoolManagement/students"
	"testing"
	"time"
)

//var stud = new(students.Student)
var principal = new(Principal)

//var principal = Principal{
//	ID:  int(time.Now().Unix()),
//	FirstName: "Mr Ojo",
//	LastName: "O",
//	Role: school.PRINCIPAL,
//}

func TestPrincipal_Admit(t *testing.T) {
	//iPrincipal := IPrincipal(principal)
	principal.ID =  int(time.Now().Unix())
	principal.FirstName = "Mr Dada"
	principal.LastName = "O"
	principal.Role = school.PRINCIPAL
	applicants := []struct{
		input students.Applicant
		output int
		word string
	} {
		{
			input: students.Applicant{
				FirstName: "Bamidele",
				LastName: "Johnson",
				Age: 12,
			},
			output: 4,
			word: "Admit Student 1",
		},
		{
			input: students.Applicant{
				FirstName: "Bami",
				LastName: "John",
				Age: 0,
			},
			output: 0,
			word: "Admit Student 2",
		},
		{
			input: students.Applicant{
				FirstName: "Bami",
				LastName: "John",
				Age: 8,
			},
			output: 5,
			word: "Admit Student 3",
		},
		{
			input: students.Applicant{
				FirstName: "Bami",
				LastName: "John",
				Age: 9,
			},
			output: 6,
			word: "Admit Student 4",
		},
		{
			input: students.Applicant{
				FirstName: "Bami",
				LastName: "John",
				Age: 10,
			},
			output: 7,
			word: "Admit Student 5",
		},
		{
			input: students.Applicant{
				FirstName: "Bami",
				LastName: "John",
				Age: 11,
			},
			output: 8,
			word: "Admit Student 6",
		},
		{
			input: students.Applicant{
				FirstName: "Bami",
				LastName: "John",
				Age: 13,
			},
			output: 9,
			word: "Admit Student 7",
		},
	}

	for _, applicant := range applicants {
		t.Run(applicant.word, func(t *testing.T) {
			result, _, _ := principal.Admit(applicant.input, school.Students)
			if result != applicant.output{
				t.Errorf("Expected %v, Got %v\n", applicant.output, result)
			}
		})
	}
}

func TestPrincipal_Expel(t *testing.T) {
	student := []struct {
		input students.Student
		output int
		word string
	} {
		{
			input: students.Student{ID: 3},
			output: 8,
			word: "Expel Student 3",
		},
		{
			input: students.Student{ID: 20},
			output: 8,
			word: "Expel Student 3",
		},
	}

	//iPrincipal := IPrincipal(principal)
	principal.ID =  int(time.Now().Unix())
	principal.FirstName = "Mr Dada"
	principal.LastName = "O"
	principal.Role = school.PRINCIPAL

	for _, stud := range student {
		t.Run(stud.word, func(t *testing.T) {
			exist, total, _ := principal.Expel(stud.input)
			if !exist && stud.output != total {
				t.Errorf("Expected %v got %v", stud.output, total)
			}
		})
	}
}

func TestPrincipal_Promote(t *testing.T) {
	student := []struct {
		input students.Student
		output int
		word string
	} {
		{
			input: students.Student{ID: 2},
			output: 6,
			word: "Promote Student SS2 to SS3",
		},
		{
			input: students.Student{ID: 3},
			output: -1,
			word: "Promote Non-Student",
		},
		{
			input: students.Student{ID: 1},
			output: 6,
			word: "Promote Student SS3 to Alumnus",
		},
	}
	//iPrincipal := IPrincipal(principal)
	principal.ID =  int(time.Now().Unix())
	principal.FirstName = "Mr Dada"
	principal.LastName = "O"
	principal.Role = school.PRINCIPAL

	//var i school.Promotional
	//i = principal
	for _, stud := range student {
		t.Run(stud.word, func(t *testing.T) {
			class, _:= principal.Promote(stud.input)
			if stud.output != class {
				t.Errorf("Expected %v got %v", stud.output, class)
			}
		})
	}
}
