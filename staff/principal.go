package staff

import (
	"errors"
	"github.com/Ad3bay0c/SchoolManagement/school"
	s "github.com/Ad3bay0c/SchoolManagement/students"
)

//type IPrincipal interface {
//	Admit(student s.Student, stud map[int]s.Student) (int, error, int)
//	Expel(student s.Student) (bool, int, int)
//}

type Principal struct {
	ID int
	FirstName, LastName string
	Position int
}

//PrincipalProfile.ID =  int(time.Now().Unix())
//principal.FirstName = "Mr Dada"
//principal.LastName = "O"
//principal.Position = school.PRINCIPAL

func (p Principal) Admit(student s.Student, students map[int]s.Student) (sid int,err error, pid int) {
	switch {
	case student.Age > 6 && student.Age <= 8:
		student.Class = school.JSS1
	case student.Age > 8 && student.Age <= 9:
		student.Class = school.JSS2
	case student.Age > 9 && student.Age <= 10:
		student.Class = school.JSS3
	case student.Age > 10 && student.Age <= 12:
		student.Class = school.SSS1
	case student.Age > 12 && student.Age <= 14:
		student.Class = school.SSS2
	case student.Age > 14:
		student.Class = school.SSS3
	default:
		return 0, errors.New("age too low for Admission to Our School"), p.ID
	}
	student.ID = len(students) + 1
	id := len(students) + 1
	school.Class[student.Class]++
	students[id] = student

	return id, nil, p.ID
}

func (p Principal) Expel(student s.Student) (bool, int, int) {
	if stud, exist := school.Students[student.ID]; exist {
		delete(school.Students, stud.ID)
		school.Class[stud.Class]--
		return true, len(school.Students), p.ID
	}

	return false, len(school.Students), p.ID
}

func (p Principal) Promote(student s.Student) (int, int) {
	if stud, exist := school.Students[student.ID]; exist {
		school.Class[stud.Class]--
		stud.Class++
		school.Class[stud.Class]--
		return stud.Class, p.ID
	}
	return student.Class, p.ID
}


