package main

import (
	"golang.org/x/crypto/acme/autocert"
	"net/http"
)

func main() {
	http.Serve(autocert.NewListener("example.com"), nil)

	// more advanced server configuration
	m := &autocert.Manager{
		Cache:      autocert.DirCache("golang-autocert"),
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("example.org", "www.example.org"),
	}

	server := &http.Server{
		Addr:      ":443",
		TLSConfig: m.TLSConfig(),
	}

	server.ListenAndServeTLS("", "")
}
