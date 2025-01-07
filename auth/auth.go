package auth

import (
	"net/http"
)

type Session struct {
	Id        int32
	ReturnUrl string
}

func InitMux(sessions *[]*Session) *http.ServeMux {
	//Initializing a MUX to handle authorization related requests
	mux := http.NewServeMux()

	mux.Handle("GET /authorize", &AuthorizeRoute{})
	return mux
}
