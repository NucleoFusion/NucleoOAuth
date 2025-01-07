package auth

import (
	"fmt"
	"io"
	"net/http"
)

type AuthorizeRoute struct {
	sessions *[]*Session
}

func (s *AuthorizeRoute) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//Getting Return URL
	returnURL := r.URL.Query()["returnURL"][0]

	//Generating a random ID for current session
	id := generateRandomID()

	//Opening a new window to authenticate user
	code := fmt.Sprintf("<script>window.open('http://localhost:5173/%s');</script>", id)
	io.WriteString(w, code)
}

func generateRandomID() string {
	return "1234"
}
