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
	Ids           []string `json:"ids"`
	NextCursorStr string   `json:"next_cursor_str"`
}

// health check struct
type HealthCheck struct {
	RateLimitContext RateLimitJson `json:"rate_limit_context"`
	Resources        ResourcesJson `json:"resources"`
}

type RateLimitJson struct {
	AccessToken string `json:"access_token"`
}

type ResourcesJson struct {
	Blocks    BlocksJson    `json:"blocks"`
	Followers FollowersJson `json:"followers"`
}

type BlocksJson struct {
	List LimitsInfo `json:"/blocks/list"`
	Ids  LimitsInfo `json:"/blocks/ids"`
}

type FollowersJson struct {
	List LimitsInfo `json:"followersList"`
	Ids  LimitsInfo `json:"followersIds"`
}

type LimitsInfo struct {
	Limit     int `json:"limit"`
	Remaining int `json:"remaining"`
	Reset     int `json:"reset"`
}