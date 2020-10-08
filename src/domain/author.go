package domain

type Author struct {
	Name string `json:"name"`
	Age  int64  `json:"age"`
	City string `json:"city"`
}

type AuthorRepository interface {
	SaveAuthor(author Author) error
}
