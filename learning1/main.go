package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// users definition.
type users struct {
	ID        int    `json: "id"`
	FirstName string `json: "firstname"`
	LastName  string `json: "lastname"`
}

// user function to get all users.
func user(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	profile := users{001, "zicheng", "wang"}
	js, err := json.Marshal(profile)

	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("SmartGuy", "zwang")
	w.Write(js)
}

func main() {
	r := chi.NewRouter()
	// r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	r.Get("/user", user)

	http.ListenAndServe(":2589", r)
}
