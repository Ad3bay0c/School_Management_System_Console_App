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
	Role                int
}

//PrincipalProfile.ID =  int(time.Now().Unix())
//principal.FirstName = "Mr Dada"
//principal.LastName = "O"
//principal.Role = school.PRINCIPAL

func (p Principal) Admit(applicant s.Applicant) (sid int,err error, pid int) {
	student := school.Student{
		FirstName: applicant.FirstName,
		LastName: applicant.LastName,
		Age: applicant.Age,
	}
	switch {
	case applicant.Age > 6 && applicant.Age <= 8:
		student.Class = school.JSS1
	case applicant.Age > 8 && applicant.Age <= 9:
		student.Class = school.JSS2
	case applicant.Age > 9 && applicant.Age <= 10:
		student.Class = school.JSS3
	case applicant.Age > 10 && applicant.Age <= 12:
		student.Class = school.SSS1
	case applicant.Age > 12 && applicant.Age <= 14:
		student.Class = school.SSS2
	case applicant.Age > 14:
		student.Class = school.SSS3
	default:
		return 0, errors.New("age too low for Admission to Our School"), p.ID
	}
	student.ID = len(school.Students) + 1
	id := len(school.Students) + 1
	school.Class[student.Class]++
	school.Students[id] = student

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
		if stud.Class == 6 {
			school.Alumni++
		} else {
			stud.Class++
		}
		return stud.Class, p.ID
	}
	return -1, p.ID
}


