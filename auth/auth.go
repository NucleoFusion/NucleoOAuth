package auth

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
)

type Session struct {
	Id        string
	ReturnUrl string
}

var ValidTime, _ = time.ParseDuration("1h")
var AccessExpiry, _ = time.ParseDuration("6h")

func InitMux(rdb *redis.Client, db *sql.DB) *http.ServeMux {
	//Initializing a MUX to handle authorization related requests
	mux := http.NewServeMux()

	mux.Handle("GET /authorize", &AuthorizeRoute{Rdb: rdb, Db: db})
	mux.Handle("POST /register/{id}", &RegisterRoute{Rdb: rdb, Db: db})

	return mux
}
