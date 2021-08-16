package school

const (
	PRINCIPAL = iota
	TEACHER
)

type Promotional interface {
	Promote(s Student) (int, int)
}