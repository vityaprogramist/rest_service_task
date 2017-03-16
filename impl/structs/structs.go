package structs

type User struct {
	ID        *int64  `json:"ID,omitempty"`
	Login     string  `json:"Login"`
	FirstName string  `json:"FirstName"`
	LastName  string  `json:"LastName"`
	Age       int     `json:"Age"`
	Phone     int64   `json:"Phone"`
	PassInfo  *string `json:"PassInfo,omitempty"`
}

type AuthUser struct {
	Login    string `json:"Login,required"`
	Password string `json:"Password,required"`
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

type FilmResponce struct {
	Films []Film `json:"films"`
	Count int    `json:"count"`
}
