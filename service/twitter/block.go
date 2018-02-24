package twitter

import (
	"net/url"
)

// BlockUsers ユーザーをブロックする
func (api *Client) BlockUsers(screenName string) error {
	v := url.Values{}
	v.Set("screen_name", screenName)
	res, err := oauthClient.Post(nil, api.credentials, "https://api.twitter.com/1.1/blocks/create.json", v)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return nil
}
