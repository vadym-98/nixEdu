package main

import (
	"fmt"
	"github.com/vadym-98/playground/slide12/api"
	"github.com/vadym-98/playground/slide12/handler"
	"github.com/vadym-98/playground/slide12/middleware"
	"github.com/vadym-98/playground/slide12/web"
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

	mux.Get("/posts", api.GetAllPosts, middleware.RequireAuthentication)
	mux.Get("/posts/\\d+", api.ShowPost, middleware.RequireAuthentication)
	mux.Post("/posts", api.StorePost, middleware.RequireAuthentication)
	mux.Put("/posts", api.UpdatePost, middleware.RequireAuthentication)
	mux.Delete("/posts/\\d+", api.DeletePost, middleware.RequireAuthentication)

	//todo create your middleware for unauthorized users
	mux.Get("/login", web.Login, nil)
	mux.Post("/login", api.Login, nil)

	//todo create logout endpoint

	mux.Get("/", web.Home, middleware.RequireAuthentication)

	fmt.Printf("Server started on port: %s\n", port)
	log.Fatal(http.ListenAndServeTLS(
		fmt.Sprintf(":%s", port),
		"localhost.pem",
		"localhost-key.pem",
		mux,
	))
}
