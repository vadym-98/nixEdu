package response

import "net/http"

type Basic struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func SendOK(w http.ResponseWriter, data any) {
	sendJson(w, http.StatusOK, data)
}
