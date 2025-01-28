package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/redis/go-redis/v9"
)

type GetRoute struct {
	Rdb *redis.Client
	Db  *sql.DB
}

func (s *GetRoute) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	token := r.URL.Query().Get("access_token")
	if token == "" {
		WriteError(&w, "ERROR: invalid/missing access_token")
		return
	}

	user_id := r.URL.Query().Get("user_id")
	if user_id == "" {
		WriteError(&w, "ERROR: invalid/missing user_id")
		return
	}

	_, err := MatchAccessToken(s.Rdb, token, user_id)
	if err != nil {
		WriteError(&w, err.Error())
		return
	}

	resp, err := GetDataFromSQL(s.Db, user_id)
	if err != nil {
		WriteError(&w, err.Error())
		return
	}

	data, _ := json.Marshal(resp)

	io.Writer.Write(w, data)
}

func GetDataFromSQL(db *sql.DB, user_id string) (*ResourceResponse, error) {
	response := ResourceResponse{}
	var id int

	int_id, _ := strconv.Atoi(user_id)

	res := db.QueryRow("SELECT * FROM user_data WHERE user_id = $1", int_id)

	err := res.Scan(&id, &response.UserId, &response.Name, &response.Email, &response.Address, &response.Country, &response.State, &response.Phone, &response.Remarks)

	if err != nil {
		return &response, err
	}

	return &response, nil
}

// Matches given access_token and user_id with the redis cache to check validity and/or tampering
func MatchAccessToken(rdb *redis.Client, access_token string, id string) (bool, error) {
	res, err := rdb.Get(context.Background(), "ACCESS:"+id).Result()
	if err != nil {
		return false, errors.New("ERROR: no access token found for given user_id")
	}

	if res != access_token {
		return false, errors.New("ERROR: access token does not match for given user")
	}

	return true, nil
}
