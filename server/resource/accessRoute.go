package main

import (
	"database/sql"
	"net/http"

	"github.com/redis/go-redis/v9"
)

type AccessRoute struct {
	Rdb *redis.Client
	Db  *sql.DB
}

func (s *AccessRoute) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
