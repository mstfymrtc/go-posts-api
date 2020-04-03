package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mstfymrtc/go-posts-api/app"
	"net/http"
	"os"
)

func main() {
	router := mux.NewRouter()
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
