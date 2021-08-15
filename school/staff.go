package school

const (
	PRINCIPAL = iota
	TEACHER
	BURSARY
)

type Promotional interface {
	Promote(s Student) (int, int)
}