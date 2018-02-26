package scraping

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/konojunya/twblock-suspicious-account/model"
)

var (
	requestURL = "https://twitter.com/intent/user?user_id="
)

// GetUserFromTwitter idを元にスクレイピングする
func GetUserFromTwitter(id int) (model.User, error) {
	time.Sleep(1 * time.Second)
	idStr := strconv.Itoa(id)
	doc, err := goquery.NewDocument(requestURL + idStr)
	if err != nil {
		return model.User{}, err
	}

	el := doc.Find("h2 a")
	screenName, ok := el.Attr("href")
	if !ok {
		return model.User{}, errors.New("Can't read screen_name")
	}

	name := strings.TrimSpace(el.Text())
	icon, ok := el.Find("img").Attr("src")
	if !ok {
		return model.User{}, errors.New("Can't read profile image url")
	}

	description := strings.TrimSpace(doc.Find(".note").Text())

	return model.User{
		ID:                   idStr,
		ScreeName:            screenName,
		Name:                 name,
		Description:          description,
		ProfileImageURLHttps: icon,
	}, nil
}
