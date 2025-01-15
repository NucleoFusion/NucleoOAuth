package routes

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"lapisoauth/auth"
	"net/http"
	"strconv"

	"github.com/redis/go-redis/v9"
)

type RefreshRoute struct {
	Rdb *redis.Client
	Db  *sql.DB
}

type RefreshResponse struct {
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
	Type         string `json:"type"`
	ExpiresIn    int32  `json:"expires_in"`
}

func (s *RefreshRoute) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	r.ParseForm()
	body := r.Form

	// Getting Refresh Token and verifying its existence
	val, ok := body["refresh_token"]
	if !ok {
		io.WriteString(w, "ERROR: missing and/or invalid parameters provided")
	}

	refresh_token := val[0]

	//Declaring Channels
	userExists := make(chan bool, 1)
	userDataAccess := make(chan string) // Communicate between WriteAccessToCacheRefresh and UserExistsRefresh for user data

	userDataRefresh := make(chan string) // Communicate between CreateNewRefreshToken and UserExistsRefresh for user data
	getNewRefresh := make(chan string, 1)

	cacheError := make(chan error, 1)
	getAccessToken := make(chan string, 1) // Get generated access token in main thread

	go UserExistsRefresh(s.Db, refresh_token, userExists, userDataAccess, userDataRefresh)
	go WriteAccessToCacheRefresh(s.Rdb, cacheError, userDataAccess, getAccessToken)
	go CreateNewRefreshToken(s.Db, userDataRefresh, getNewRefresh)

	if !(<-userExists) {
		io.WriteString(w, "ERROR: no user with refresh_token found, the token may have been regenerated please authenticate again")
		return
	}

	err := <-cacheError
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}

	newAccess := <-getAccessToken
	newRefresh := <-getNewRefresh

	response := RefreshResponse{
		AccessToken:  newAccess,
		RefreshToken: newRefresh,
		Type:         "Bearer",
		ExpiresIn:    int32(auth.AccessExpiry.Seconds()),
	}

	data, _ := json.Marshal(response)
	io.Writer.Write(w, data)
}

func UserExistsRefresh(db *sql.DB, refresh_token string, userExists chan bool, giveDataAccess chan string, giveDataRefresh chan string) {
	var (
		id    int
		name  string
		email string
	)

	res := db.QueryRow("SELECT id, name, email FROM users WHERE refresh_token = $1", refresh_token)
	err := res.Scan(&id, &name, &email)
	if err == sql.ErrNoRows {
		userExists <- false
	} else if err != nil {
		fmt.Println(err.Error())
		userExists <- false
	}

	// Give data to WriteAccessToCacheRefresh goroutine
	giveDataAccess <- strconv.Itoa(id)
	giveDataAccess <- name
	giveDataAccess <- email

	// Give data to CreateNewRefreshToken goroutine
	giveDataRefresh <- strconv.Itoa(id)
	giveDataRefresh <- name
	giveDataRefresh <- email

	userExists <- true
}

func WriteAccessToCacheRefresh(rdb *redis.Client, cacheError chan error, getUserData chan string, giveAccessToken chan string) {
	id_str := <-getUserData
	name := <-getUserData
	email := <-getUserData

	newToken := auth.GenerateToken(name + email)

	_, err := rdb.Set(context.Background(), fmt.Sprintf("ACCESS:%s", id_str), newToken, auth.AccessExpiry).Result()
	cacheError <- err

	giveAccessToken <- newToken
}

func CreateNewRefreshToken(db *sql.DB, getUserData chan string, giveNewRefresh chan string) {
	id_str := <-getUserData
	name := <-getUserData
	email := <-getUserData

	id, _ := strconv.Atoi(id_str)

	newRefresh := auth.GenerateToken(name + email + name)

	db.Exec("UPDATE users SET refresh_token = $1 WHERE id = $2", newRefresh, id)

	giveNewRefresh <- newRefresh
}
