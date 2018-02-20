package main

import (
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/garyburd/go-oauth/oauth"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	ck string
	cs string
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	ck = os.Getenv("consumerKey")
	cs = os.Getenv("consumerSecret")
}

func main() {

	config := getClient()
	rt, err := config.RequestTemporaryCredentials(nil, "http://127.0.0.1:8080/oauth", nil)
	if err != nil {
		log.Fatal(err)
	}

	url := config.AuthorizationURL(rt, nil)

	r := gin.Default()
	r.LoadHTMLGlob("view/*")
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, url)
	})
	r.GET("/oauth", func(c *gin.Context) {
		oauthToken := c.Query("oauth_token")
		oauthVerifier := c.Query("oauth_verifier")

		at, err := getAccessToken(&oauth.Credentials{
			Token: oauthToken,
		}, oauthVerifier)
		if err != nil {
			log.Fatal(err)
		}

		tweet("hello", at)

		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.Run(":8080")
}

func tweet(text string, at *oauth.Credentials) {
	client := getClient()
	v := url.Values{}
	v.Set("status", text)

	resp, err := client.Post(nil, at, "https://api.twitter.com/1.1/statuses/update.json", v)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
}

func getClient() *oauth.Client {
	return &oauth.Client{
		TemporaryCredentialRequestURI: "https://api.twitter.com/oauth/request_token",
		ResourceOwnerAuthorizationURI: "https://api.twitter.com/oauth/authorize",
		TokenRequestURI:               "https://api.twitter.com/oauth/access_token",
		Credentials: oauth.Credentials{
			Token:  ck,
			Secret: cs,
		},
	}
}

func getAccessToken(credentials *oauth.Credentials, oauthVerifier string) (*oauth.Credentials, error) {
	client := getClient()
	at, _, err := client.RequestToken(nil, credentials, oauthVerifier)
	return at, err
}
