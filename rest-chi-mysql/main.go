package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/go-sql-driver/mysql"
)

var router *chi.Mux
var db *sql.DB

const (
	dbName = "go-mysql-crud"
	dbPass = "12345"
	dbHost = "localhost"
	dbPort = "33066"
)

func routers() *chi.Mux {
	// router.Get("posts", AllPosts)
	// router.Get("/posts/{id}", DetailPost)
	router.Post("/posts", CreatePost)
	router.Put("/posts/{id}", UpdatePost)
	router.Delete("/posts/{id}", DeletePost)
	return router
}

func init() {
	router = chi.NewRouter()
	router.Use(middleware.Recoverer)

	dbSource := fmt.Sprintf("root:%s@tcp(%s:%s)/%s?charset=utf8", dbPass, dbHost, dbPort, dbName)

	var err error
	db, err = sql.Open("mysql", dbSource)

	catch(err)
}

// Post struct with JSON return.
type Post struct {
	ID      string `json: "id"`
	Title   string `json: "title"`
	Content string `json: "content"`
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	var post Post
	json.NewDecoder(r.Body).Decode(&post)

	query, er := db.Prepare("Insert posts SET title=?, content=?")
	catch(er)

	defer query.Close()

	respondwithJSON(w, http.StatusCreated, map[string]string{"message": "successfully created"})
}

// UpdatePost updates a post..
func UpdatePost(w http.ResponseWriter, r *http.Request) {
	var post Post
	id := chi.URLParam(r, "id")
	json.NewDecoder(r.Body).Decode(&post)

	query, err := db.Prepare("Update posts set title=?, content=? where id=?")
	catch(err)
	_, er := query.Exec(post.Title, post.Content, id)
	catch(er)

	defer query.Close()
	respondwithJSON(w, http.StatusOK, map[string]string{"message": "Update successfully!"})
}

// DeletePost delete a post..
func DeletePost(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	query, err := db.Prepare("delete post where id=?")
	catch(err)
	query.Close()

	_, er := query.Exec(id)
	catch(er)
	query.Close()

	respondwithJSON(w, http.StatusOK, map[string]string{"message": "successfully deleted"})
}

func main() {
	routers()
	http.ListenAndServe(":8005", Logger())
}
