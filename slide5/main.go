package main

import (
	"fmt"
	"github.com/vadym-98/playground/slide5/handler"
	"github.com/vadym-98/playground/slide5/response"
	"net/http"
	"regexp"
	"strconv"
)

const port = "8080"

type post struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

var posts = []post{
	{1, "sunt aut facere repellat provident occaecati excepturi optio reprehenderit", "quia et suscipit\nsuscipit recusandae consequuntur expedita et cum\nreprehenderit molestiae ut ut quas totam\nnostrum rerum est autem sunt rem eveniet architecto"},
	{2, "qui est esse", "est rerum tempore vitae\nsequi sint nihil reprehenderit dolor beatae ea dolores neque\nfugiat blanditiis voluptate porro vel nihil molestiae ut reiciendis\nqui aperiam non debitis possimus qui neque nisi nulla"},
	{3, "ea molestias quasi exercitationem repellat qui ipsa sit aut", "et iusto sed quo iure\nvoluptatem occaecati omnis eligendi aut ad\nvoluptatem doloribus vel accusantium quis pariatur\nmolestiae porro eius odio et labore et velit aut"},
}

func main() {
	mux := handler.NewRegexMuxer()

	mux.Get("/posts", func(w http.ResponseWriter, r *http.Request) {
		response.SendOK(w, posts)
	})

	mux.Get("/posts/\\d+", func(w http.ResponseWriter, r *http.Request) {
		var res *post

		pattern := regexp.MustCompile("\\d+")

		match := pattern.FindString(r.URL.Path)

		postID, err := strconv.Atoi(match)
		if err != nil {
			response.SendBadRequestError(w, fmt.Errorf("couldn't transform string %s to int", match))
			return
		}

		for _, p := range posts {
			if p.Id == postID {
				res = &p
				break
			}
		}

		if res == nil {
			response.SendBadRequestError(w, fmt.Errorf("there is not post with id: %d", postID))
			return
		}

		response.SendOK(w, res)
	})

	fmt.Printf("Server started on port: %s\n", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), mux)
}

//func main() {
//	//mux := http.NewServeMux()
//	mux := http.DefaultServeMux
//
//	mux.HandleFunc("/hello/golang", func(w http.ResponseWriter, r *http.Request) {
//		fmt.Fprint(w, "Hello Golang")
//	})
//
//	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
//		fmt.Fprint(w, "Hello World")
//	})
//
//	http.ListenAndServe(":8080", mux)
//}
