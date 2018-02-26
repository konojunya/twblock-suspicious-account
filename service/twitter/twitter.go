package twitter

import (
	"io/ioutil"
	"net/url"

	"github.com/garyburd/go-oauth/oauth"
	"github.com/konojunya/twblock-suspicious-account/auth"
)

type Client struct {
	credentials *oauth.Credentials
}

var (
	oauthClient *oauth.Client
	api         *Client
)

func init() {
	oauthClient = auth.GetOauthClient()
}

func requestClient(credentials *oauth.Credentials, url string, params url.Values) ([]byte, error) {
	res, err := oauthClient.Get(nil, credentials, url, params)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, err
}

// SetupClient モデルを設定する
func SetupClient(oauthToken, oauthVerifier string) error {
	at, err := auth.GetAccessToken(&oauth.Credentials{
		Token: oauthToken,
	}, oauthVerifier)
	if err != nil {
		return err
	}
	api = &Client{
		credentials: at,
	}
	return nil
}

// GetClient モデルを取得する
func GetClient() *Client {
	return api
}

// CanUse twitter APIが使えるかどうか
func CanUse() bool {
	return api != nil
}
