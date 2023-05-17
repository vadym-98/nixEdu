package api

import (
	"fmt"
	"github.com/vadym-98/playground/slide13/domain"
	"github.com/vadym-98/playground/slide13/request"
	"github.com/vadym-98/playground/slide13/response"
	"net/http"
	"regexp"
	"strconv"
)

func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	response.SendOK(w, domain.GetAllPosts())
}

func ShowPost(w http.ResponseWriter, r *http.Request) {
	pattern := regexp.MustCompile("\\d+")

	match := pattern.FindString(r.URL.Path)

	postID, err := strconv.Atoi(match)
	if err != nil {
		response.SendBadRequestError(w, fmt.Errorf("couldn't transform string %s to int", match))
		return
	}

	post := domain.FindPostByID(postID)

	if post == nil {
		response.SendBadRequestError(w, fmt.Errorf("there is not post with id: %d", postID))
		return
	}

	response.SendOK(w, post)
}

func StorePost(w http.ResponseWriter, r *http.Request) {
	rp := request.CreatePostRequest{}

	err := request.ParseJsonRequest(r, &rp)
	if err != nil {
		response.SendBadRequestError(w, err)
		return
	}

	post := domain.Post{Title: rp.Title, Body: rp.Body}
	domain.AddPost(&post)

	response.SendOK(w, post)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	rp := request.UpdatePostRequest{}

	err := request.ParseJsonRequest(r, &rp)
	if err != nil {
		response.SendBadRequestError(w, err)
		return
	}

	post := domain.Post{Id: rp.Id, Title: rp.Title, Body: rp.Body}
	domain.UpdatePost(&post)

	response.SendOK(w, post)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	pattern := regexp.MustCompile("\\d+")

	match := pattern.FindString(r.URL.Path)

	postID, err := strconv.Atoi(match)
	if err != nil {
		response.SendBadRequestError(w, fmt.Errorf("couldn't transform string %s to int", match))
		return
	}

	post := domain.FindPostByID(postID)

	if post == nil {
		response.SendBadRequestError(w, fmt.Errorf("there is not post with id: %d", postID))
		return
	}

	domain.RemovePost(post)

	response.SendOK(w, post)
}
