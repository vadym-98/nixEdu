package request

import "github.com/vadym-98/playground/slide12/domain"

const (
	titleMaxLength    = 100
	titleMaxLengthErr = "too long title. Max allowed length is 100"

	nonExistingIDErr = "there is no post with such ID"
)

type CreatePostRequest struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func (cpr *CreatePostRequest) Validate() []string {
	var errors []string

	if len(cpr.Title) > titleMaxLength {
		errors = append(errors, titleMaxLengthErr)
	}

	return errors
}

type UpdatePostRequest struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func (upr *UpdatePostRequest) Validate() []string {
	var errors []string

	targetPost := domain.FindPostByID(upr.Id)
	if targetPost == nil {
		errors = append(errors, nonExistingIDErr)
	}

	if len(upr.Title) > titleMaxLength {
		errors = append(errors, titleMaxLengthErr)
	}

	return errors
}
