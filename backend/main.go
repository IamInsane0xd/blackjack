package main

import (
	"net/http"

	"backend/handlers"
	"backend/models"
)

var (
	userStore = models.NewUserStorage()
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/login", handlers.NewLoginHandler(userStore))
	mux.Handle("/register", handlers.NewRegisterHandler(userStore))

	if err := http.ListenAndServe("localhost:3001", mux); err != nil {
		panic(err)
	}
}
