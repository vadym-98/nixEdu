package request

const (
	minPasswordLength    = 8
	minPasswordLengthErr = "password must have 8 characters"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (lr *LoginRequest) Validate() []string {
	var errors []string

	if len(lr.Password) < minPasswordLength {
		errors = append(errors, minPasswordLengthErr)
	}

	return errors
}
