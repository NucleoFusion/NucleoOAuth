package auth

import (
	"time"
)

type Session struct {
	Id        string
	ReturnUrl string
}

var ValidTime, _ = time.ParseDuration("1h")
var AccessExpiry, _ = time.ParseDuration("6h")
