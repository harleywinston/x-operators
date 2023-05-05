package configs

import (
	"fmt"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
)

var (
	Clinet    *http.Client
	CookieJar *cookiejar.Jar
	BaseURL   *url.URL
)

func InitApiSession() error {
	baseURL, err := url.Parse(os.Getenv("XUI_BASE_URL"))
	if err != nil {
		return err
	}
	BaseURL = baseURL

	CookieJar, err = cookiejar.New(nil)
	if err != nil {
		return err
	}

	Clinet = &http.Client{
		Jar: CookieJar,
	}

	loginURL := BaseURL.ResolveReference(&url.URL{Path: "login"})
	resp, err := Clinet.PostForm(
		loginURL.String(),
		url.Values{
			"username": {os.Getenv("XUI_USERNAME")},
			"password": {os.Getenv("XUI_PASSWORD")},
		},
	)
	if err != nil {
		return err
	}

	log.Println(loginURL)
	cookies := Clinet.Jar.Cookies(resp.Request.URL)
	for _, cookie := range cookies {
		fmt.Println(cookie.Name, cookie.Value)
	}

	return nil
}
