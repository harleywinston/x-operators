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

var (
	URL_PARSE_ERROR      = &CustomError{Message: "URL parse failed!", Code: 500}
	JSON_MARSHAL_ERROR   = &CustomError{Message: "Json marshal mathers failed!", Code: 500}
	JSON_UNMARSHAL_ERROR = &CustomError{Message: "Json unmarshal unmathers failed!", Code: 500}
	XUI_DRIVER_ERROR     = &CustomError{Message: "XUI DRIVER ERROR!", Code: 500}
)
