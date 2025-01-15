package main

import (
	"fmt"
	"lapisoauth/cache"
	"lapisoauth/db"
	"lapisoauth/routes"
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
	mux := routes.InitMux(rdb, db)

	fmt.Println("Listening at :5544...")
	http.ListenAndServe(":5544", mux)
}
