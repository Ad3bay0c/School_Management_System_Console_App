package students

type Student struct {
	ID 					int
	FirstName, LastName string
	age 				int
	class				int
}

const (
	JSS = iota
	JSS1 = iota
	JSS3
	SSS1
	SSS2
	SSS3
)