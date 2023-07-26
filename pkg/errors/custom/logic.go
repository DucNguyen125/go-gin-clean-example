package custom

type LogicError struct {
	HTTPCode int    `json:"-"`
	Code     string `json:"code,omitempty"`
}

func NewLogicError(httpCode int, errorMessage string) CustomError {
	return &LogicError{
		HTTPCode: httpCode,
		Code:     errorMessage,
	}
}

func (le *LogicError) Error() string {
	return le.Code
}

func (le *LogicError) GetHTTPCode() int {
	return le.HTTPCode
}
