package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/vadym-98/playground/slide8_updated/handler"
	"net/http"
)

const port = "8080"

func main() {
	router := httprouter.New()

	router.GET("/posts", handler.GetAllPosts)
	router.GET("/posts/:post", handler.ShowPost)
	router.POST("/posts", handler.StorePost)
	router.PUT("/posts", handler.UpdatePost)
	router.DELETE("/posts/:post", handler.DeletePost)

	fmt.Printf("Server started on port: %s\n", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), router)
}
