package custom

type InternalServerError struct {
	HTTPCode int    `json:"-"`
	Code     string `json:"code,omitempty"`
	Message  string `json:"message,omitempty"`
}

func (ise *InternalServerError) Error() string {
	return ise.Code
}

func (ise *InternalServerError) GetHTTPCode() int {
	return ise.HTTPCode
}
