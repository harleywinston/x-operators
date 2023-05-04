package xui

import (
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/harleywinston/x-operators/xui/internal/transport"
)

var (
	Clinet    *http.Client
	CookieJar *cookiejar.Jar
	BaseURL   *url.URL
)

func initApiSession() error {
	BaseURL, err := url.Parse(os.Getenv("BASE_URL"))
	if err != nil {
		return err
	}

	CookieJar, err = cookiejar.New(nil)
	if err != nil {
		return err
	}

	Clinet = &http.Client{
		Jar: CookieJar,
	}

	loginURL := BaseURL.ResolveReference(&url.URL{Path: "login"})
	_, err = Clinet.PostForm(
		loginURL.String(),
		url.Values{"username": {os.Getenv("USERNAME")}, "password": {os.Getenv("PASSWORD")}},
	)
	if err != nil {
		return err
	}

	return nil
}

func registerHandlers() error {
	r := gin.Default()

	setupHandlers := transport.SetupHandlers{}
	r.GET("/add", setupHandlers.AddClientHandler)
	r.GET("/delete", setupHandlers.DeleteClientHandler)
	return nil
}

func InitApp() error {
	if err := initApiSession(); err != nil {
		return err
	}
	if err := registerHandlers(); err != nil {
		return err
	}
	return nil
}
