package api

import (
	"gopkg.in/resty.v1"
	"obas/src/config"
)

const BASE_URL string = "https://localhost:9000"

func Rest() *resty.Request {
	return resty.R().SetAuthToken("").
		SetHeader("Accept", "application/json").
		SetHeader("email", "email").
		SetHeader("site", "site").
		SetHeader("Content-Type", "application/json")
}

var JSON = config.ConfigWithCustomTimeFormat
