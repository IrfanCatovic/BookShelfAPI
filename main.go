package main

import (
	"log"
	"net/http"
)

// type Book struct {
// 	ID     int     `json:"id"`
// 	Title  string  `json:"title"`
// 	Author string  `json:"author"`
// 	Year   int     `json:"year"`
// 	Price  float64 `json:"price"`
// 	IsRead bool    `json:"isread"`
// }

var books = []Book{
	{ID: 1, Title: "Dune", Author: "Frank Herbert", Year: 1965, Price: 24.90, IsRead: true},
	{ID: 2, Title: "1984", Author: "George Orvel", Year: 1949, Price: 15.00, IsRead: true},
	{ID: 3, Title: "Alchemist", Author: "Paolo Coelho", Year: 1988, Price: 21.50, IsRead: false},
}

func main() {
	http.HandleFunc("/books", booksHandler)      //One route and in handler i use POST and GET
	http.HandleFunc("/books/", booksByIdHandler) //Get specific book by ID

	log.Fatal(http.ListenAndServe(":8080", nil))
}
