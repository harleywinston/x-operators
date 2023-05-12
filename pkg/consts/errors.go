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

	MANAGER_CONNECTION_ERROR = &CustomError{Message: "Couldn't connect to manager!", Code: 500}
	MANAGER_USER_COUNT_ERROR = &CustomError{
		Message: "There is no user assigned to this server!",
		Code:    500,
	}
	INBOUND_COUNT_ERROR = &CustomError{
		Message: "There is no inbound to create clients on!",
		Code:    500,
	}
	BIND_JSON_ERROR = &CustomError{Message: "Bind json failed!", Code: 500}

	ADD_SUCCESS    = &CustomError{Message: "Add succeed.", Code: 200}
	DELETE_SUCCESS = &CustomError{Message: "Delete succeed.", Code: 200}

	XUI_API_ERROR             = &CustomError{Message: "XUI API error!", Code: 500}
	CREATE_HTTP_REQUEST_ERROR = &CustomError{
		Message: "Create http.NewRequest{} error!",
		Code:    500,
	}
	CLIENT_DO_ERROR = &CustomError{Message: "Client.Do error!", Code: 500}
)
