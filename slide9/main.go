package main

import (
	"fmt"
	"github.com/vadym-98/playground/slide9/handler"
	"log"
	"net/http"
	"path"
)

const port = "8080"

const RedirectAddress = "https://localhost:8080"

func main() {
	go func() {
		log.Fatal(http.ListenAndServe(":8090", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			redirectURL := path.Join(RedirectAddress, r.RequestURI)
			log.Println(redirectURL)
			http.Redirect(w, r, redirectURL, http.StatusMovedPermanently)
		})))
	}()

	mux := handler.NewRegexMuxer()

	mux.Get("/posts", handler.GetAllPosts)
	mux.Get("/posts/\\d+", handler.ShowPost)
	mux.Post("/posts", handler.StorePost)
	mux.Put("/posts", handler.UpdatePost)
	mux.Delete("/posts/\\d+", handler.DeletePost)

	fmt.Printf("Server started on port: %s\n", port)
	log.Fatal(http.ListenAndServeTLS(fmt.Sprintf(":%s", port), "localhost.pem", "localhost-key.pem", mux))
}
