package service

import (
	"encoding/json"
	"io/ioutil"
	"net/url"
	"regexp"

	"github.com/garyburd/go-oauth/oauth"
	"github.com/konojunya/twblock-suspicious-account/model"
)

type TwitterClient struct {
	Credentials *oauth.Credentials
}

var (
	ck  string
	cs  string
	api *TwitterClient
)

// SetKeys Twitter APIを使うためのキーをセットする
func SetKeys(consumerKey, consumerSecret string) {
	ck = consumerKey
	cs = consumerSecret
}

// GetUsers 怪しいユーザー一覧
func GetUsers() ([]model.User, error) {
	users, err := suspiciousFilter(api.GetUsers())
	if err != nil {
		return nil, err
	}

	return users, nil
}

func suspiciousFilter(orgUsers model.UsersResponse, err error) ([]model.User, error) {
	if err != nil {
		return nil, err
	}

	var users []model.User

	for _, user := range orgUsers.Users {
		if isSuspicious(user.Description) {
			users = append(users, user)
		}
	}

	return users, nil
}

func isSuspicious(description string) bool {
	r := regexp.MustCompile(`投資`)
	return r.MatchString(description)
}

// BlockUser ユーザーをブロックする
func BlockUser(screenName string) error {
	return api.BlockUsers(screenName)
}

// GetClient OAuthクライアントを取得する
func GetClient() *oauth.Client {
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
	client := GetClient()
	at, _, err := client.RequestToken(nil, credentials, oauthVerifier)
	return at, err
}

func (api *TwitterClient) getMe() (model.User, error) {
	client := GetClient()

	res, err := client.Get(nil, api.Credentials, "https://api.twitter.com/1.1/account/verify_credentials.json", nil)
	if err != nil {
		return model.User{}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return model.User{}, err
	}

	var user model.User
	json.Unmarshal(body, &user)

	return user, nil
}

// GetUsers 怪しいアカウント一覧を取得する
func (api *TwitterClient) GetUsers() (model.UsersResponse, error) {
	client := GetClient()
	v := url.Values{}
	user, err := api.getMe()
	if err != nil {
		return model.UsersResponse{}, err
	}

	v.Set("screen_name", user.ScreeName)
	v.Set("count", "200")
	res, err := client.Get(nil, api.Credentials, "https://api.twitter.com/1.1/followers/list.json", v)
	if err != nil {
		return model.UsersResponse{}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return model.UsersResponse{}, err
	}

	var users model.UsersResponse
	json.Unmarshal(body, &users)

	return users, nil
}

// BlockUsers ユーザーをブロックする
func (api *TwitterClient) BlockUsers(screenName string) error {
	client := GetClient()
	v := url.Values{}
	v.Set("screen_name", screenName)
	res, err := client.Post(nil, api.Credentials, "https://api.twitter.com/1.1/blocks/create.json", v)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return nil
}

// SetTwitterClient モデルを設定する
func SetTwitterClient(c *TwitterClient) {
	api = c
}

// GetTwitterClient モデルを取得する
func GetTwitterClient() *TwitterClient {
	return api
}
