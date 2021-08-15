package staff

import (
	"github.com/Ad3bay0c/SchoolManagement/school"
	s "github.com/Ad3bay0c/SchoolManagement/students"
)

type Teacher struct {
	ID int
	FirstName, LastName string
	Subjects			[]string
	Role                int
}

func (teacher Teacher) GradeStudent(student *school.Student) {
	student.Grade = map[string]string{
		teacher.Subjects[0] : "A",
		teacher.Subjects[1] : "B",
		teacher.Subjects[2] : "A1",
	}
	school.Students[student.ID] = *student
}

func (teacher Teacher) Promote(student s.Student) (int, int) {
	if stud, exist := school.Students[student.ID]; exist {
		school.Class[stud.Class]--
		if stud.Class == 6 {
			school.Alumni++
		} else {
			stud.Class++
		}
		return stud.Class, teacher.ID
	}
	return -1, teacher.ID
}
