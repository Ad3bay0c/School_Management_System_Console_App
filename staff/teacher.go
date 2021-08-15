package staff

import "github.com/Ad3bay0c/SchoolManagement/school"

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
