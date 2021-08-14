package school

const (
	PRINCIPAL = iota
	TEACHER
)

type Promotional interface {
	Promote()
}
