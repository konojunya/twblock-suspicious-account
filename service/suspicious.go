package service

import (
	"regexp"

	"github.com/konojunya/twblock-suspicious-account/model"
)

func suspiciousFilter(orgUsers []model.User, err error) ([]model.User, error) {
	users := make([]model.User, 0)

	if err != nil {
		return users, err
	}

	for _, user := range orgUsers {
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
