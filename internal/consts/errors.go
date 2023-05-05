package consts

import "fmt"

type CustomError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Detail  string `json:"detail"`
}

func (e *CustomError) Error() string {
	return fmt.Sprintf(`%d: %s\n%s`, e.Code, e.Message, e.Detail)
}

var URL_PARSE_ERROR = &CustomError{Message: "URL parse failed!", Code: 500}
