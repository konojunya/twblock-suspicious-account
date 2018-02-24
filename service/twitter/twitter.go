package twitter

import (
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
