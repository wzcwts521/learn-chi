package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func requestTime(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "The current time is: %s\n", time.Now())
}

func requestSay(w http.ResponseWriter, r *http.Request) {
	val := r.FormValue("name")
	if val != "" {
		fmt.Fprintf(w, "Hello %s!", val)
	} else {
		fmt.Fprintf(w, "Hello ... you.")
	}
}

func main() {
	fmt.Println("Starting server on port :3000")

	// http.HandleFunc("/time", requestTime)
	// http.HandleFunc("/say", requestSay)
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/time", requestTime)
	r.Route("/say", func(r chi.Router) {
		r.Get("/{name}", requestSay)
		r.Get("/", requestSay)
	})
	err := http.ListenAndServe(":3001", r)
	if err != nil {
		fmt.Println("ListenAndServe:", err)
	}
}
