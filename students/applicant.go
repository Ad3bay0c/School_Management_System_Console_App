package students

import "github.com/Ad3bay0c/SchoolManagement/school"

type Applicant struct {
	ID 					int
	FirstName, LastName string
	Age 				int
}

// CheckStatus 'applicant to check Admission Status'
func(applicant Applicant) CheckStatus() bool {
	_, err := school.Students[applicant.ID] // if present in list of students

	return err
}