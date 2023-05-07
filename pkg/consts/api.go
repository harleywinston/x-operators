package consts

import (
	"log"
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

	log.Println(BaseURL.String())
	req, err := http.NewRequest(http.MethodGet, BaseURL.String(), nil)
	log.Println(HTTPClient.Do(req))
	return nil
}
