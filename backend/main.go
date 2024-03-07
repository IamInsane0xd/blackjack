package main

import (
	"net/http"

	"backend/handlers"
	"backend/models"
)

var (
	userStore = models.NewUserStorage()
	game      = models.NewGame()
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/login", handlers.NewLoginHandler(userStore))
	mux.Handle("/register", handlers.NewRegisterHandler(userStore))
	mux.Handle("/forgotPassword", handlers.NewForgotPasswordHandler(userStore))
	mux.Handle("/game/start", handlers.NewStartHandler(game))

	if err := http.ListenAndServe("localhost:3001", mux); err != nil {
		panic(err)
	}
}
