package routes

import (
	"encoding/json"
	"io"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func WriteError(w *http.ResponseWriter, err string) {
	resp := ErrorResponse{Error: err}

	data, _ := json.Marshal(resp)

	io.Writer.Write(*w, data)
}
