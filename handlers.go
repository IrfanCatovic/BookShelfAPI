package main

import (
	"encoding/json"
	"net/http"
)

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

}
