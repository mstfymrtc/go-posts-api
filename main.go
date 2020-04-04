package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mstfymrtc/go-posts-api/app"
	"github.com/mstfymrtc/go-posts-api/controllers"
	"net/http"
	"os"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/register", controllers.Register).Methods("POST")
	router.HandleFunc("/api/authenticate", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/api/posts", controllers.GetPosts).Methods("GET")
	router.HandleFunc("/api/posts/{id}", controllers.GetPost).Methods("GET")
	router.HandleFunc("/api/posts", controllers.CreatePost).Methods("POST")

	router.Use(app.JwtAuthentication)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Println(err)
	}
}
