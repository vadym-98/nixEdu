package main

import (
	"fmt"
	"github.com/vadym-98/playground/slide8/handler"
	"net/http"
)

const port = "8080"

func main() {
	mux := handler.NewRegexMuxer()

	mux.Get("/posts", handler.GetAllPosts)
	mux.Get("/posts/\\d+", handler.ShowPost)
	mux.Post("/posts", handler.StorePost)
	mux.Put("/posts", handler.UpdatePost)
	mux.Delete("/posts/\\d+", handler.DeletePost)

	fmt.Printf("Server started on port: %s\n", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), mux)
}
