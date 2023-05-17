package domain

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

const AuthKey = "authenticatedUser"

var (
	users = []User{
		{"john.doe", "john.doe@gmail.com", "12345678"},
	}
	authenticatedUser *User
)

func GetAuthUser() *User {
	return authenticatedUser
}

func Authenticate(u *User) bool {
	for _, user := range users {
		if user.Email == u.Email && user.Password == u.Password {
			authenticatedUser = &user
			return true
		}
	}

	return false
}
