package main

import (
	"github.com/konojunya/twblock-suspicious-account/router"
)

func main() {
	r := router.GetRouter()
	r.Run(":8080")
}
