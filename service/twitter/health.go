package twitter

import (
	"encoding/json"
	"net/url"

	"github.com/konojunya/twblock-suspicious-account/model"
)

func (api *Client) HealthCheck() (model.HealthCheck, error) {
	v := url.Values{}
	v.Set("resources", "followers,blocks")

	body, err := requestClient(api.credentials, "https://api.twitter.com/1.1/application/rate_limit_status.json", v)
	if err != nil {
		return model.HealthCheck{}, err
	}

	var hc model.HealthCheck
	err = json.Unmarshal(body, &hc)
	if err != nil {
		return model.HealthCheck{}, nil
	}

	return hc, nil
}
