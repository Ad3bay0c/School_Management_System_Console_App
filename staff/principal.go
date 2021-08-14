package staff

import (
	"errors"
	"github.com/Ad3bay0c/SchoolManagement/school"
	s "github.com/Ad3bay0c/SchoolManagement/students"
	"time"
)

type Principal struct {
	school.Profile
}

func (p Principal) Admit(student *s.Student, students []*s.Student) (int, error) {
	switch {
	case student.Age < 7:
		student.Class = school.JSS1
	case student.Age < 9:
		student.Class = school.JSS2
	case student.Age < 10:
		student.Class = school.JSS3
	case student.Age < 12:
		student.Class = school.SSS1
	case student.Age < 14:
		student.Class = school.SSS2
	case student.Age > 14:
		student.Class = school.SSS3
	default:
		return 0, errors.New("age too low for Admission to Our School")
	}
	student.ID = int(time.Now().Unix())
	students = append(students, student)
	return student.ID, nil
}


