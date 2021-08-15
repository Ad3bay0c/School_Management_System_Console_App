package students

import "github.com/Ad3bay0c/SchoolManagement/school"

type Applicant struct {
	ID 					int
	FirstName, LastName string
	Age 				int
}

func(applicant Applicant) CheckStatus() bool {
	_, err := school.Students[applicant.ID]

	return err
}