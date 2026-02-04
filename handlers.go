package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

var books = []Book{
	{ID: 1, Title: "Dune", Author: "Frank Herbert", Year: 1965, Price: 24.90, IsRead: true},
	{ID: 2, Title: "1984", Author: "George Orvel", Year: 1949, Price: 15.00, IsRead: true},
	{ID: 3, Title: "Alchemist", Author: "Paolo Coelho", Year: 1988, Price: 21.50, IsRead: false},
}

func booksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(books)
	case http.MethodPost:
		var newBook Book
		err := json.NewDecoder(r.Body).Decode(&newBook)
		if err != nil {
			http.Error(w, "Los Json", http.StatusBadRequest)
			return
		}

		if newBook.Title == "" || newBook.Author == "" {
			http.Error(w, "Nedostaju podaci o knjizi", http.StatusBadRequest)
			return
		}

		maxID := 0
		for _, book := range books {
			if book.ID > maxID {
				maxID = book.ID
			}
		}
		newBook.ID = maxID + 1

		books = append(books, newBook)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newBook)

	default:
		http.Error(w, "Metod nije dozvoljen", http.StatusMethodNotAllowed)
	}
}

func booksByIdHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	//Find book by id
	case http.MethodGet:
		path := r.URL.Path
		idStr := strings.TrimPrefix(path, "/books/")
		id, err := strconv.Atoi(idStr) //ovaj path iza books pretvaram da je broj i pamti u id
		if idStr == path || idStr == "" || err != nil {
			http.Error(w, "Nevalidan ID", http.StatusBadRequest)
			return
		}
		var foundBook *Book
		for i := range books {
			if books[i].ID == id {
				foundBook = &books[i]
				break
			}
		}

		if foundBook == nil {
			http.Error(w, "Knjiga nije pronadjena", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(foundBook)

		//Update book
	case http.MethodPut:

		path := r.URL.Path
		idStr := strings.TrimPrefix(path, "/books/")
		id, err := strconv.Atoi(idStr)
		if idStr == path || idStr == "" || err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}
		//look for book
		//here I copy book with that id to foundBook so later when I update foundBook I actually update book in books slice
		var foundBook *Book
		for i := range books {
			if books[i].ID == id {
				foundBook = &books[i]
				break
			}
		}

		if foundBook == nil {
			http.Error(w, "Book is not found", http.StatusNotFound)
			return
		}

		var updatedBook Book
		err = json.NewDecoder(r.Body).Decode(&updatedBook)

		if err != nil {
			http.Error(w, "Invalid JSON FORMAT", http.StatusBadRequest)
			return
		}

		if updatedBook.Title == "" || updatedBook.Author == "" {
			http.Error(w, "Missing data about the book", http.StatusBadRequest)
			return
		}

		//Update fields
		if updatedBook.Title != "" {
			foundBook.Title = updatedBook.Title
		}
		if updatedBook.Author != "" {
			foundBook.Author = updatedBook.Author
		}
		if updatedBook.Year != 0 {
			foundBook.Year = updatedBook.Year
		}
		if updatedBook.Price != 0 {
			foundBook.Price = updatedBook.Price
		}

		foundBook.IsRead = updatedBook.IsRead

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(foundBook)
		//Return updated book as result

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)

	}
}
