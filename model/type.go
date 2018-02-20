package model

// User Twitter User Model
type User struct {
	ID                   string `json:"id_str"`
	Description          string `json:"description"`
	Name                 string `json:"name"`
	ScreeName            string `json:"screen_name"`
	ProfileImageURLHttps string `json:"profile_image_url_https"`
}

// UsersResponse Twitter response
type UsersResponse struct {
	Users []User `json:"users"`
}
