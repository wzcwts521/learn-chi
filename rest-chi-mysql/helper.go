package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func catch(err error) {
	if err != nil {
		panic(err)
	}
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondwithJSON(w, code, map[string]string{"message": msg})
}

func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	respose, _ := json.Marshal(payload)
	fmt.Println(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(respose)
}

// Logger return log message
func Logger() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(time.Now(), r.Method, r.URL)
		router.ServeHTTP(w, r)
	})
}
