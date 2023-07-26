package custom

type CustomError interface {
	Error() string
	GetHTTPCode() int
}
