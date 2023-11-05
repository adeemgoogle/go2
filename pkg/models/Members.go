package models

type Member struct {
	ID       int    `json:"id" db:"id"`
	FullName string `json:"full_name" db:"full_name"`
}

type BorrowedBook struct {
	MemberID int `json:"member_id" db:"member_id"`
	BookID   int `json:"book_id" db:"book_id"`
}
