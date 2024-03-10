package main

import (
	auth "auth/pkg"
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/login", auth.Login)
	mux.HandleFunc("/register", auth.Register)
	mux.HandleFunc("/update", auth.Update)
	mux.HandleFunc("/CreateFilm", auth.CreateFilm)

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}

	log.Println("Server is listening on :8080")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
		return
	}
}
