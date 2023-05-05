package consts

import (
	"net/http"
	"net/url"
	"os"
)

var (
	BaseURL    *url.URL
	HTTPClient *http.Client
)

func InitAPISession() error {
	baseURL, err := url.Parse(os.Getenv("MANAGER_BASE_URL"))
	if err != nil {
		return &CustomError{
			Message: URL_PARSE_ERROR.Message,
			Code:    URL_PARSE_ERROR.Code,
			Detail:  err.Error(),
		}
	}
	BaseURL = baseURL

	HTTPClient = &http.Client{}
	return nil
}
