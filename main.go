package main

import (
	"encoding/json"
	"log"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     int    `json:id`
	Title  string `json:title`
	Author string `json:author`
	Year   string `json:year`
}

var books []Book

func main() {
	router := mux.NewRouter()

	books = append(books, Book{ID: 1, Title: "Golang Pointers", Author: "Mr. Golang", Year: "2010"},
		Book{ID: 2, Title: "GoRoutines", Author: "Mr. GoRoutines", Year: "2011"},
		Book{ID: 3, Title: "Golang Routers", Author: "Mr. Router", Year: "2012"},
		Book{ID: 4, Title: "Golang Concurrency", Author: "Mr. Currency", Year: "2013"},
		Book{ID: 5, Title: "Golang Good parts", Author: "Mr. Good", Year: "2014"})

	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", removeBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	// log.Println("Get all books is called")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Get book is called")
	params := mux.Vars(r)

	i, _ := strconv.Atoi(params["id"])

	log.Println(reflect.TypeOf(i))
	for _, book := range books {
		if book.ID == i {
			json.NewEncoder(w).Encode(&book)
		}
	}

}
func addBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Add book is called")
	var book Book
	log.Println(book)
	json.NewDecoder(r.Body).Decode(&book)

	books = append(books, book)
	json.NewEncoder(w).Encode(books)
}
func updateBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Update Books is called")
	var book Book
	json.NewDecoder(r.Body).Decode(&book)

	for i, item := range books {
		if item.ID == book.ID {
			books[i] = book
		}
	}
	json.NewEncoder(w).Encode(books)
}
func removeBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Remove book is called")
	params := mux.Vars(r)

	id, _ := strconv.Atoi(params["id"])
	for index, item := range books {
		if item.ID == id {
			books = append(books[:index], books[index+1:]...)
		}
	}
	json.NewEncoder(w).Encode(books)
}
