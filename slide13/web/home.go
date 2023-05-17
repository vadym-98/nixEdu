package web

import (
	"errors"
	"github.com/vadym-98/playground/slide13/domain"
	"github.com/vadym-98/playground/slide13/response"
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(
		"slide12/templates/home.gohtml",
		"slide12/templates/header.gohtml",
		"slide12/templates/footer.gohtml",
	)
	if err != nil {
		response.SendServerError(w, errors.New("couldn't parse templates"))
		return
	}

	t.ExecuteTemplate(w, "home", r.Context().Value(domain.AuthKey))
}
