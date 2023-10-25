package response

import "fmt"

type Error struct {
	ErrorType    string `json:"errorType"`
	ErrorMessage string `json:"errorMessage"`
}

func Wrap(err error, errorType string, message string) Error {
	return Error{
		ErrorType:    errorType,
		ErrorMessage: fmt.Sprintf("%s; %s", message, err),
	}
}
