package routes

import (
	"database/sql"
	"net/http"

	"github.com/redis/go-redis/v9"
)

func InitMux(rdb *redis.Client, db *sql.DB) *http.ServeMux {
	//Initializing a MUX to handle authorization related requests
	mux := http.NewServeMux()

	mux.Handle("GET /authorize", &AuthorizeRoute{Rdb: rdb, Db: db})
	mux.Handle("POST /register/{id}", &RegisterRoute{Rdb: rdb, Db: db})

	return mux
}
