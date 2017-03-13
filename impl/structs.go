package main

type User struct {
	ID        int64  `json:"ID"`
	Login     string `json:"Login"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Age       int    `json:"Age"`
	Phone     int64  `json:"Phone"`
}

type Film struct {
	ID          int    `json:"ID"`
	Title       string `json:"Title"`
	Genres      string `json:"Genres"`
	ReleaseYear int    `json:"ReleaseYear"`
}

type Filter struct {
	Genre   *string `json:"Genre,omitempty"`
	Release *int    `json:"Release,omitempty"`
}

type Paging struct {
	Prev *string `json:"Prev"`
	Next *string `json:"Next"`
}

type FilmResponce struct {
	Filter Filter `json:"filter"`
	Films  []Film `json:"films"`
	Paging Paging `json:"paging"`
	Count  int    `json:"count"`
}

type FilmQuery struct {
	Filter Filter `json:"filter"`
	Page   int    `json:"page"`
	Limit  int    `json:"limit"`
}
