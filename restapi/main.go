package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Article{Title: "Test Title", Desc: "Description", Content: "Content"}
	fmt.Println("Endpoint Hit. All articles.")
	json.NewEncoder(w).Encode(articles)

}

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage hit endpoint")
}

func handleRequests() {
	http.HandleFunc("/", homepage)
	http.HandleFunc("/articles", allArticles)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
func main() {
	handleRequests()
}
