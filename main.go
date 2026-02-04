package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/books", booksHandler)      //One route and in handler i use POST and GET
	http.HandleFunc("/books/", booksByIdHandler) //Get specific book by ID

	log.Fatal(http.ListenAndServe(":8080", nil))
}
