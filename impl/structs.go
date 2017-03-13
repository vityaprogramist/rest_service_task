package main

type User struct {
	ID        int64
	Login     string
	FirstName string
	LastName  string
	Age       int
	Phone     int64
}

type Film struct {
	ID          int
	Title       string
	Genres      string
	ReleaseYear int
}
