package staff

import (
	"github.com/Ad3bay0c/SchoolManagement/school"
	"testing"
)

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
