package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

// Book structs (Model)
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author Struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Init Books var as slice book struct
var books []Book

// Get All books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(books)
}

// Get Single book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Get Params
	params := mux.Vars(r)
	// looping through the boosk to find specific id
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

// create book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(1000000)) //Not a safe way generating
}

// update book
func updateBooks(w http.ResponseWriter, r *http.Request) {

}

// delete book
func deleteBooks(w http.ResponseWriter, r *http.Request) {

}

func main() {
	// init Router
	r := mux.NewRouter()

	books = append(books, Book{ID: "1", Isbn: "43543", Title: "Book one", Author: &Author{Firstname: "Damian", Lastname: "Kelechi"}})

	books = append(books, Book{ID: "2", Isbn: "43544", Title: "Book two", Author: &Author{Firstname: "Emma", Lastname: "Odoh"}})

	books = append(books, Book{ID: "3", Isbn: "43545", Title: "Book three", Author: &Author{Firstname: "Diamond", Lastname: "Chidera"}})

	// Router Handlers / Endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBooks).Methods("PUT")
	r.HandleFunc("/api/books", deleteBooks).Methods("DELETE")

	// server the Request
	log.Fatal(http.ListenAndServe(":8000", r))
}
