package students

import "github.com/Ad3bay0c/SchoolManagement/school"

type Student struct {
	ID 					int
	FirstName, LastName string
	Age 				int
	Class				int
	Subject				[]string
}

func (s Student) TakeCourse() bool {
	s.Subject = school.Subject[s.Class]
	if s.Subject[1] != "" {
		return true
	}
	return false
}

