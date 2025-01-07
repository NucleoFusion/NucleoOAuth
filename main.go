package main

import (
	"fmt"
	"lapisoauth/auth"
	"net/http"
)

func main() {
	// Initializing Sessions Array, that will hold all current sessions
	sessions := make([]*auth.Session, 100)

	//Initializing the root Mux Server
	mux := auth.InitMux(&sessions)

	fmt.Println("Listening at :5544...")
	http.ListenAndServe(":5544", mux)
}
