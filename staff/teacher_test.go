package staff

import (
	"github.com/Ad3bay0c/SchoolManagement/school"
	"github.com/Ad3bay0c/SchoolManagement/students"
	"testing"
	"time"
)

var teacher = new(Teacher)

func TestTeacher_GradeStudent(t *testing.T) {
	students := []struct{
		input		school.Student
		output		int
		word 		string
	} {
		{
			input: school.Students[1],
			output: 3,
			word: "Grade Student 1",
		},
	}
	teacher := Teacher{
		ID:        1,
		FirstName: "Mr Ojo",
		LastName:  "L",
		Role:      school.TEACHER,
		Subjects: []string{"Yoruba", "Physics", "Biology"},
	}
	for _, student := range students {
		t.Run(student.word, func(t *testing.T) {
			teacher.GradeStudent(&student.input)
			if len(student.input.Grade) != student.output {
				t.Errorf("Expected %v; Got %v", student.output, len(student.input.Grade))
			}
		})
	}
}

func TestTeacher_Promote(t *testing.T) {
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
	teacher.ID =  int(time.Now().Unix())
	teacher.FirstName = "Mr Dada"
	teacher.LastName = "O"
	teacher.Role = school.TEACHER

	//var i school.Promotional
	//i = principal
	for _, stud := range student {
		t.Run(stud.word, func(t *testing.T) {
			class, _:= teacher.Promote(stud.input)
			if stud.output != class {
				t.Errorf("Expected %v got %v", stud.output, class)
			}
		})
	}
}
