package entities

import "strings"

//ErrHTTP ... Error to write in a file like a template
type ErrHTTP struct {
	Content string
	Num     int
}

//ErrorErrHTTP ... Handle errors when creating ErrHTTP
type ErrorErrHTTP struct {
	ErrorString string
}

const (
	//MaxStatuCode ... Max status of the response
	MaxStatuCode = 600
	//MinStatusCode ... Min status of the response
	MinStatusCode = 100
	//ErrNotFound ... Error to handle 404 error
	ErrNotFound = "Not found"
	//ErrInternalServerError ... Error to handle 500 error
	ErrInternalServerError = "Internal server error"
	//ErrForbidden ... Error to handle 302
	ErrForbidden = "Forbidden, you cant enter here"
)

var (
	//ErrInvalidStatusCode ... Error getting status code
	ErrInvalidStatusCode = &ErrorErrHTTP{ErrorString: "Error parsing status code because is an invalid type"}
	//ErrInvalidContent ... Error getting content
	ErrInvalidContent = &ErrorErrHTTP{ErrorString: "Error parsing content of the status because is empty(probably)"}
)

//NewErrHTTP ... Return a fresh instance of an err type http
func NewErrHTTP(c string, n int) (ErrHTTP, error) {
	if n < MinStatusCode || n > MaxStatuCode {
		return ErrHTTP{}, ErrInvalidStatusCode
	}
	if len(strings.TrimSpace(c)) == 0 {
		return ErrHTTP{}, ErrInvalidStatusCode
	}
	return ErrHTTP{
		Content: c,
		Num:     n,
	}, nil
}

func (eeh *ErrorErrHTTP) Error() string {
	return eeh.ErrorString
}
