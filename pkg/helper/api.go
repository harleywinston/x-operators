package helper

import (
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/harleywinston/x-operators/pkg/consts"
)

var (
	BaseURL    *url.URL
	HTTPClient *http.Client
)

func InitAPISession() error {
	baseURL, err := url.Parse(os.Getenv("MANAGER_BASE_URL"))
	if err != nil {
		return &consts.CustomError{
			Message: consts.URL_PARSE_ERROR.Message,
			Code:    consts.URL_PARSE_ERROR.Code,
			Detail:  err.Error(),
		}
	}
	BaseURL = baseURL

	HTTPClient = &http.Client{}

	log.Println(BaseURL.String())
	// req, err := http.NewRequest(http.MethodGet, BaseURL.String(), nil)
	// if err != nil {
	// 	return &CustomError{
	// 		Message: MANAGER_CONNECTION_ERROR.Message,
	// 		Code:    MANAGER_CONNECTION_ERROR.Code,
	// 		Detail:  err.Error(),
	// 	}
	// }
	// resp, err := HTTPClient.Do(req)
	// if err != nil {
	// 	return &CustomError{
	// 		Message: MANAGER_CONNECTION_ERROR.Message,
	// 		Code:    MANAGER_CONNECTION_ERROR.Code,
	// 		Detail:  err.Error(),
	// 	}
	// }
	// var respData string
	// err = json.NewDecoder(resp.Body).Decode(&respData)
	// if err != nil {
	// 	return &CustomError{
	// 		Message: JSON_UNMARSHAL_ERROR.Message,
	// 		Code:    JSON_UNMARSHAL_ERROR.Code,
	// 		Detail:  err.Error(),
	// 	}
	// }
	//
	// log.Println(respData)

	return nil
}
