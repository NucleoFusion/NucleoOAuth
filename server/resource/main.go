package main

import (
	"fmt"
	"net/http"
)

func main() {
	db, err := ConnectToPG()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	rdb := ConnectToRedis()

	http.Handle("POST /access", &AccessRoute{Rdb: rdb, Db: db})

	fmt.Println("Listening at 5545...")
	http.ListenAndServe(":5545", nil)
}
