package staff

type Profile struct {
	ID string
	FirstName, LastName string
	Position int
}

const (
	PRINCIPAL = iota
	TEACHER
)
