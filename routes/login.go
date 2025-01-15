package routes

//TODO: Change Error management and sending

import (
	"database/sql"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
)

type LoginRuote struct {
	Db  *sql.DB
	Rdb *redis.Client
}

func (s *LoginRuote) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	//Getting Session ID
	id := r.PathValue("id")

	//Parsing Form and collecting Data
	err := r.ParseForm()
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}

	body := r.PostForm
	user, err := DecodeLoginBody(&body)
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}

	// Declaring channels
	matchFound := make(chan bool, 1)  // checking whether id matches with current sessions
	matchError := make(chan error, 1) // error channel for MatchIDWithSessions

	delSession := make(chan bool, 1) // Telling the MatchIDWithSessions to delete current session from active

	go MatchIDWithSessions(s.Rdb, id, matchFound, matchError, delSession)

	err = <-matchError
	if err != nil {
		fmt.Println(err.Error())
		io.WriteString(w, err.Error())
		return
	}

	if !(<-matchFound) {
		io.WriteString(w, "ERROR: no current session with given id found")
		return
	}

	matches, err := FindingUserAndMatchPasswords(s.Db, &user)
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}

	if !matches {
		io.WriteString(w, "ERROR: invalid password")
		return
	}
}

func FindingUserAndMatchPasswords(db *sql.DB, user *RegisterBody) (bool, error) {
	res := db.QueryRow("SELECT id, password FROM users WHERE email = $1")

	var (
		id              int
		hashed_password string
	)

	err := res.Scan(&id, &hashed_password)
	if err == sql.ErrNoRows {
		return false, errors.New("ERROR: no user with given email found")
	} else if err != nil {
		return false, err
	}

	result := bcrypt.CompareHashAndPassword([]byte(hashed_password), []byte(user.Password))
	if result != nil {
		return false, nil
	}

	return true, nil
}

func DecodeLoginBody(body *url.Values) (RegisterBody, error) {
	user := RegisterBody{}

	email, ok := (*body)["email"]
	if !ok {
		return user, errors.New("ERROR: email not provided")
	}

	pass, ok := (*body)["pass"]
	if !ok {
		return user, errors.New("ERROR: password not provided")
	}

	if pass[0] == "" || email[0] == "" {
		return user, errors.New("ERROR: invalid parameters")
	}

	user.Password = pass[0]
	user.Email = email[0]

	return user, nil
}
