package students

import "github.com/Ad3bay0c/SchoolManagement/school"

type Student struct {
	ID 					int
	FirstName, LastName string
	Age 				int
	Class				int
	Subject				[]string
}

var Students = map[int]Student{
	1: {
		ID: 1,
		FirstName: "Clin",
		LastName:  "Ade",
		Age:       15,
		Class:     school.SSS3,
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

func (s Student) TakeCourse() bool {
	s.Subject = school.Subject[s.Class]
	//if s.Subject[1] != "" {
	//	return true
	//}
	return true
}

