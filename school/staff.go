package school

const (
	PRINCIPAL = iota
	TEACHER
)

type Promotional interface {
	CheckStudentList() map[int]Student
}