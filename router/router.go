package router

import (
	"log"
	"net/http"

	"github.com/konojunya/twblock-suspicious-account/auth"
	"github.com/konojunya/twblock-suspicious-account/service"

	"github.com/gin-gonic/gin"
	"github.com/konojunya/twblock-suspicious-account/middleware"
	"github.com/konojunya/twblock-suspicious-account/service/twitter"
)

func getRedirectURL() string {
	config := auth.GetOauthClient()
	rt, err := config.RequestTemporaryCredentials(nil, "http://127.0.0.1:8080/oauth", nil)
	if err != nil {
		panic(err)
	}
	url := config.AuthorizationURL(rt, nil)

	return url
}

func GetRouter() *gin.Engine {
	r := gin.Default()

	r.Static("/js", "./public/js")
	r.Static("/image", "./public/image")
	r.Static("/css", "./public/css")

	r.LoadHTMLGlob("view/*")

	r.GET("/", func(c *gin.Context) {
		if !twitter.CanUse() {
			c.HTML(http.StatusOK, "login.html", nil)
		} else {
			c.HTML(http.StatusOK, "index.html", nil)
		}
	})
	r.GET("/auth", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, getRedirectURL())
	})
	r.GET("/oauth", func(c *gin.Context) {
		oauthToken := c.Query("oauth_token")
		oauthVerifier := c.Query("oauth_verifier")

		err := twitter.SetupClient(oauthToken, oauthVerifier)
		if err != nil {
			log.Fatal(err)
		}
		service.SetupTwitterClient()

		c.Redirect(http.StatusFound, "/")
	})

	api := r.Group("/api")
	api.Use(middleware.IsAuthenticated())
	apiRouter(api)

	return r
}
