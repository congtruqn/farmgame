package lib

type Error struct {
	StatusCode int      `json:"status_code"`
	ErrorsCode []string `json:"errors_code"`
}

func NewError(status_code int) *Error {

	err := &Error{
		StatusCode: status_code,
		ErrorsCode: make([]string, 0),
	}

	return err
}
