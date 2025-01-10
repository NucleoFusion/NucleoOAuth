package auth

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/redis/go-redis/v9"
)

//Directly Route Related

type RegisterRoute struct {
	Rdb *redis.Client
	Db  *sql.DB
}

func (s *RegisterRoute) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	//Getting Session ID
	id := r.PathValue("id")

	//Parsing Form and collecting Data
	err := r.ParseForm()
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}

	body := r.PostForm
	user, err := DecodeBody(&body)
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}

	user.SessionID = id
	user.TokenType = "Bearer"
	user.ExpiresIn = 6 * 60 * 60

	//Declaring Channels
	matchError := make(chan error, 1) // Checking Errors in MatchIDWithSessions goroutine
	matchFound := make(chan bool, 1)  // If there is a match for provided with current sessions
	delSession := make(chan bool, 1)  // Tell MatchIDWithSessions goroutine to remove that sessions from cache

	dbError := make(chan error, 1)
	tokenWritten := make(chan bool, 1) // Tell main thread that token was written

	writeToken := make(chan bool, 1) // Tell WriteAccessToken goroutine to Write the AccessToken to cache
	GetToken := make(chan string, 1) // Give the AccessToken created to the WriteToDB goroutine

	//Matching ID with current Session and Writing To DB
	go MatchIDWithSessions(s.Rdb, id, matchFound, matchError, delSession)
	go WriteToDB(s.Db, user, delSession, GetToken, writeToken, dbError, tokenWritten)
	go WriteAccessToken(s.Rdb, user, writeToken, GetToken)

	// if sessionids was not found
	if !(<-matchFound) {
		io.WriteString(w, "ERROR: no session with this id found, create a new session")
		return
	}

	// if error occured while matching session ids
	err = <-matchError
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}

	// if error occured while Writiing to DB
	err = <-dbError
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}

	delSession <- true
	<-tokenWritten

	data, _ := json.Marshal(user)
	io.Writer.Write(w, data)
}

//Indirectly Route Related

type RegisterBody struct {
	UserID       int    `json:"user_id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Password     string `json:"-"`
	AccessToken  string `json:"access_token"`
	SessionID    string `json:"-"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int32  `json:"expires_in"`
}

func DecodeBody(body *url.Values) (*RegisterBody, error) {
	user := RegisterBody{}

	for k, v := range *body {
		switch k {
		case "name":
			if len(v) == 0 && v[0] == "" {
				return &user, errors.New("ERROR: invalid parameters provided")
			}

			user.Name = v[0]

		case "email":
			if len(v) == 0 && v[0] == "" {
				return &user, errors.New("ERROR: invalid parameters provided")
			}

			user.Email = v[0]
		case "password":
			if len(v) == 0 && v[0] == "" {
				return &user, errors.New("ERROR: invalid parameters provided")
			}

			hashed, err := HashPassword(v[0])
			if err != nil {
				return &user, err
			}

			user.Password = hashed
		}
	}

	return &user, nil
}

func WriteToDB(db *sql.DB, user *RegisterBody, delKey chan bool, GetAccessToken chan string, WriteToken chan bool, dbError chan error, tokenWritten chan bool) {
	user.RefreshToken = GenerateToken(user.Name + user.Email + user.Name)

	fmt.Println(user)

	res := db.QueryRow("INSERT INTO users(name, email, password, refresh_token) VALUES ($1, $2, $3, $4) RETURNING id", user.Name, user.Email, user.Password, user.RefreshToken)

	err := res.Scan(&user.UserID)
	if err != nil {
		dbError <- err
	}
	dbError <- nil

	// Tell WriteAccessToken goroutine to write to cache the AccessToken
	WriteToken <- true

	// Get the Access Token from the goroutine
	user.AccessToken = <-GetAccessToken
	tokenWritten <- true
}

func MatchIDWithSessions(rdb *redis.Client, id string, matchFound chan bool, errorChan chan error, delSession chan bool) {

	// Seeing if there is a record for the given session
	_, err := rdb.Get(context.Background(), id).Result()

	// handling if session is not found or an error occurs
	if err == redis.Nil {
		matchFound <- false
	} else if err != nil {
		matchFound <- false
		errorChan <- err
	}

	// session is found and no error occurred
	matchFound <- true
	errorChan <- nil

	// Checking whether authentication was done to delete current session
	if <-delSession {
		rdb.Del(context.Background(), id)
	}
}

func WriteAccessToken(rdb *redis.Client, user *RegisterBody, writeToken chan bool, SendToken chan string) {
	token := GenerateToken("ACCESS:" + user.Name + user.Email)

	if <-writeToken {
		rdb.Set(context.Background(), "ACCESS:"+strconv.Itoa(user.UserID), token, AccessExpiry)
		SendToken <- token
	}
}
