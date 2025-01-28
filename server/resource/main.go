package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type ResourceResponse struct {
	UserId  int    `json:"user_id"`
	Address string `json:"address"`
	Phone   int    `json:"phone"`
	Country string `json:"country"`
	State   string `json:"state"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Remarks string `json:"remarks"`
}

func main() {
	db, err := ConnectToPG()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	rdb := ConnectToRedis()

	http.Handle("GET /api/access", &GetRoute{Rdb: rdb, Db: db})

	fmt.Println("Listening at 5545...")
	http.ListenAndServe(":5545", nil)
}

func WriteError(w *http.ResponseWriter, err string) {
	resp := ErrorResponse{Error: err}

	data, _ := json.Marshal(resp)

	io.Writer.Write(*w, data)
}
