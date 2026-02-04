package main

type Book struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Year   int     `json:"year"`
	Price  float64 `json:"price"`
	IsRead bool    `json:"isread"`
}
