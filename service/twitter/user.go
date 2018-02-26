package twitter

import (
	"encoding/json"
	"net/url"

	"github.com/konojunya/twblock-suspicious-account/model"
	"github.com/konojunya/twblock-suspicious-account/service/scraping"
)

// GetUsers 怪しいアカウント一覧を取得する
func (api *Client) GetUsers() ([]model.User, error) {
	users := make([]model.User, 0)
	userCh := make(chan model.UserWithErr)
	count := 0

	v := url.Values{}
	user, err := api.GetMe()
	if err != nil {
		return users, err
	}

	v.Set("screen_name", user.ScreeName)
	v.Set("count", "200")

	nextCursor := "-1"

	for {
		v.Set("cursor", nextCursor)

		body, err := requestClient(api.credentials, "https://api.twitter.com/1.1/followers/ids.json", v)
		if err != nil {
			return users, err
		}

		usersRes := &model.UsersResponse{}
		json.Unmarshal(body, &usersRes)

		count += len(usersRes.Ids)

		for _, id := range usersRes.Ids {
			go func(id int, userCh chan model.UserWithErr) {
				user, err := scraping.GetUserFromTwitter(id)
				userCh <- model.UserWithErr{
					User: user,
					Err:  err,
				}
			}(id, userCh)
		}

		nextCursor = usersRes.NextCursorStr

		if nextCursor == "" || nextCursor == "0" {
			break
		}

	}

	for i := 0; i < count; i++ {
		r := <-userCh
		if r.Err == nil {
			users = append(users, r.User)
		}
	}

	return users, nil
}

// GetMe 認証したユーザーの情報を取得する
func (api *Client) GetMe() (model.User, error) {
	body, err := requestClient(api.credentials, "https://api.twitter.com/1.1/account/verify_credentials.json", nil)
	if err != nil {
		return model.User{}, err
	}

	var user model.User
	json.Unmarshal(body, &user)

	return user, nil
}
