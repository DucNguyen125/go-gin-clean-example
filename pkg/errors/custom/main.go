package custom

type CustomError interface {
	Error() string
	GetStatus() int
}
