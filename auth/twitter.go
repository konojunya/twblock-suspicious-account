package auth

import (
	"log"
	"os"

	"github.com/garyburd/go-oauth/oauth"
	"github.com/joho/godotenv"
)

var (
	ck string
	cs string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ck = os.Getenv("consumerKey")
	cs = os.Getenv("consumerSecret")
}

// GetOauthClient OAuthクライアントを取得する
func GetOauthClient() *oauth.Client {
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

// GetAccessToken アクセストークンを取得する
func GetAccessToken(credentials *oauth.Credentials, oauthVerifier string) (*oauth.Credentials, error) {
	client := GetOauthClient()
	at, _, err := client.RequestToken(nil, credentials, oauthVerifier)
	return at, err
}
