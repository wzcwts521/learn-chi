package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// user function to get all users.
func user(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world!"))
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	r.Get("/user", user)

	http.ListenAndServe(":2589", r)
}
