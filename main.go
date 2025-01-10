package main

import (
	"fmt"
	"lapisoauth/auth"
	"lapisoauth/cache"
	"lapisoauth/db"
	"net/http"
)

func main() {
	//Connecting to Remote Redis Server
	rdb := cache.ConnectToRedis()

	//Connecting to SQL Database
	db, err := db.ConnectToDB()
	if err != nil {
		fmt.Println(err.Error())
	}

	//Initializing the root Mux Server
	mux := auth.InitMux(rdb, db)

	fmt.Println("Listening at :5544...")
	http.ListenAndServe(":5544", mux)
}
