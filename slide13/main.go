package main

import (
	"context"
	"fmt"
	"github.com/vadym-98/playground/slide13/api"
	"github.com/vadym-98/playground/slide13/handler"
	"github.com/vadym-98/playground/slide13/middleware"
	"github.com/vadym-98/playground/slide13/web"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"path"
	"time"
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

	srv := &http.Server{Handler: mux}
	ln, err := net.Listen("tcp", "localhost:"+port)
	if err != nil {
		panic("failed init listener")
	}

	go func() {
		fmt.Printf("Server started on port: %s\n", port)

		log.Fatal(srv.ServeTLS(ln,
			"localhost.pem",
			"localhost-key.pem",
		))
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	fmt.Println("Got interruption signal. Gracefully stopping the server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
