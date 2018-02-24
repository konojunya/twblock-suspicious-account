package service

import (
	"github.com/konojunya/twblock-suspicious-account/model"
	"github.com/konojunya/twblock-suspicious-account/service/twitter"
)

var api *twitter.Client

// SetupTwitterClient twitterのclientを任意のタイミングで使えるようにする
func SetupTwitterClient() {
	api = twitter.GetClient()
}

// GetUsers 怪しいユーザー一覧
func GetUsers() ([]model.User, error) {
	return suspiciousFilter(api.GetUsers())
}

// HealthCheck api limitを確認する
func HealthCheck() (model.HealthCheck, error) {
	return api.HealthCheck()
}

// BlockUser ユーザーをブロックする
func BlockUser(screenName string) error {
	return api.BlockUsers(screenName)
}
