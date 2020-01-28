package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type article struct {
	Title       string `json:"Title"`
	Description string `json:"description"`
	Content     string `json:"content"`
}

type articles []article

// Calculate function adds x + 2
func Calculate(x int) (result int) {
	result = x + 2
	return result
}

// GetArticles returns articles
func GetArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: All Articles Endpoint")

	articles := articles{
		article{Title: "Test Title", Description: "Test Description", Content: "Test Content"},
	}
	json.NewEncoder(w).Encode(articles)
}

func addArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: Post Articles Endpoint")
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

func handleRequest(port string) {

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", GetArticles).Methods("GET")
	myRouter.HandleFunc("/articles", addArticles).Methods("POST")
	log.Fatal(http.ListenAndServe(":"+port, myRouter))
	err := http.ListenAndServe(":"+port, myRouter)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func main() {
	port := os.Getenv("LISTENING_PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Listening port: %s", port)
	handleRequest(port)
}
