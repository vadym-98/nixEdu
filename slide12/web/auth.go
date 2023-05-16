package web

import (
	"errors"
	"github.com/vadym-98/playground/slide12/response"
	"html/template"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(
		"slide12/templates/login.gohtml",
		"slide12/templates/header.gohtml",
		"slide12/templates/footer.gohtml",
	)
	if err != nil {
		response.SendServerError(w, errors.New("couldn't parse templates"))
		return
	}

	isAuthFailed := false
	if a := r.URL.Query().Get("error"); a != "" {
		isAuthFailed = true
	}

	t.ExecuteTemplate(w, "login", isAuthFailed)
}
