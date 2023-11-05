package models

type Book struct {
	ID       int    `json:"id" db:"id"`
	Title    string `json:"title" db:"title"`
	Genre    string `json:"genre" db:"genre"`
	ISBN     string `json:"isbn" db:"isbn"`
	AuthorID int    `json:"author_id" db:"author_id"`
}
