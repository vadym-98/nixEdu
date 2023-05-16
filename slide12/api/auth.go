package api

import (
	"github.com/vadym-98/playground/slide12/domain"
	"github.com/vadym-98/playground/slide12/request"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	lr := request.LoginRequest{}

	err := r.ParseForm()
	if err != nil {
		http.Redirect(w, r, "/login?error=true", http.StatusFound)
		return
	}

	lr.Email = r.FormValue("email")
	lr.Password = r.FormValue("password")

	err = request.Validate(&lr)
	if err != nil {
		http.Redirect(w, r, "/login?error=true", http.StatusFound)
		return
	}

	user := domain.User{Email: lr.Email, Password: lr.Password}
	isUserExists := domain.Authenticate(&user)
	if !isUserExists {
		http.Redirect(w, r, "/login?error=true", http.StatusFound)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}
