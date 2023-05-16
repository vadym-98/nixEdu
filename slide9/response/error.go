package response

import "net/http"

func SendServerError(w http.ResponseWriter, err error) {
	sendJson(w, http.StatusInternalServerError, Basic{
		Status:  http.StatusInternalServerError,
		Message: err.Error(),
	})
}

func SendBadRequestError(w http.ResponseWriter, err error) {
	sendJson(w, http.StatusBadRequest, Basic{
		Status:  http.StatusBadRequest,
		Message: err.Error(),
	})
}
