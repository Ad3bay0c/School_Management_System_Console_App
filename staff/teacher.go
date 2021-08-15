package staff

import "github.com/Ad3bay0c/SchoolManagement/school"

type Teacher struct {
	ID int
	FirstName, LastName string
	Role                int
}

func (teacher Teacher) GradeStudent(student *school.Student) {
	student.Grade = map[string]string{
		"English" : "A",
		"Mathematics" : "B",
		"Chemistry" : "A1",
	}
	school.Students[student.ID] = *student
}
