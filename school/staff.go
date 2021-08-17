package school

const (
	PRINCIPAL = iota
	TEACHER
)

type StudentLists interface {
	CheckStudentList() map[int]Student
}