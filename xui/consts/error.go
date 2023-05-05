package consts

import "fmt"

type CustomError struct {
	Message string
	Code    int
	Detail  string
}

func (e *CustomError) Error() string {
	return fmt.Sprintf(`%d: %s\n%s`, e.Code, e.Message, e.Detail)
}

var (
	JSON_MARSHAL_ERROR = &CustomError{Message: "Json marshal mathers failed!", Code: 500}
	BIND_JSON_ERROR    = &CustomError{Message: "Bind json failed!", Code: 500}

	ADD_SUCCESS    = &CustomError{Message: "Add succeed.", Code: 200}
	DELETE_SUCCESS = &CustomError{Message: "Delete succeed.", Code: 200}

	XUI_API_ERROR             = &CustomError{Message: "XUI API error!", Code: 500}
	CREATE_HTTP_REQUEST_ERROR = &CustomError{
		Message: "Create http.NewRequest{} error!",
		Code:    500,
	}
	CLIENT_DO_ERROR = &CustomError{Message: "Client.Do error!", Code: 500}
)
