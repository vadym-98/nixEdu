package handler

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/vadym-98/playground/slide8_updated/domain"
	"github.com/vadym-98/playground/slide8_updated/request"
	"github.com/vadym-98/playground/slide8_updated/response"
	"net/http"
	"strconv"
)

func GetAllPosts(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	response.SendOK(w, domain.GetAllPosts())
}

func ShowPost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	postWildcard := ps.ByName("post")

	postID, err := strconv.Atoi(postWildcard)
	if err != nil {
		response.SendBadRequestError(w, fmt.Errorf("couldn't transform string %s to int", postWildcard))
		return
	}

	post := domain.FindPostByID(postID)

	if post == nil {
		response.SendBadRequestError(w, fmt.Errorf("there is not post with id: %d", postID))
		return
	}

	response.SendOK(w, post)
}

func StorePost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	rp := request.CreatePostRequest{}

	err := request.ParseRequest(r, &rp)
	if err != nil {
		response.SendBadRequestError(w, err)
		return
	}

	post := domain.Post{Title: rp.Title, Body: rp.Body}
	domain.AddPost(&post)

	response.SendOK(w, post)
}

func UpdatePost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	rp := request.UpdatePostRequest{}

	err := request.ParseRequest(r, &rp)
	if err != nil {
		response.SendBadRequestError(w, err)
		return
	}

	post := domain.Post{Id: rp.Id, Title: rp.Title, Body: rp.Body}
	domain.UpdatePost(&post)

	response.SendOK(w, post)
}

func DeletePost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	postWildcard := ps.ByName("post")

	postID, err := strconv.Atoi(postWildcard)
	if err != nil {
		response.SendBadRequestError(w, fmt.Errorf("couldn't transform string %s to int", postWildcard))
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
