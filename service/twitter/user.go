package twitter

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/url"

	"github.com/konojunya/twblock-suspicious-account/model"
)

// GetUsers 怪しいアカウント一覧を取得する
func (api *Client) GetUsers() ([]model.User, error) {
	users := make([]model.User, 0)

	v := url.Values{}
	user, err := api.GetMe()
	if err != nil {
		return users, err
	}

	v.Set("screen_name", user.ScreeName)
	v.Set("count", "200")

	// nextCursor := "-1"

	// for {
	// 	v.Set("cursor", nextCursor)

	// 	res, err := oauthClient.Get(nil, api.Credentials, "https://api.twitter.com/1.1/followers/list.json", v)
	// 	if err != nil {
	// 		return users, err
	// 	}

	// 	defer res.Body.Close()

	// 	body, err := ioutil.ReadAll(res.Body)
	// 	if err != nil {
	// 		return users, err
	// 	}

	// 	usersRes := &model.UsersResponse{}
	// 	json.Unmarshal(body, &usersRes)

	// 	nextCursor = usersRes.NextCursorStr

	// 	if nextCursor == "" || nextCursor == "0" {
	// 		break
	// 	}

	// }

	return users, nil
}

func (api *Client) GetMe() (model.User, error) {
	log.Printf("api: %v\n", api)
	res, err := oauthClient.Get(nil, api.credentials, "https://api.twitter.com/1.1/account/verify_credentials.json", nil)
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
