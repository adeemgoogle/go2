package models

type Member struct {
	MemberID int    `json:"id" db:"MemberId"`
	FullName string `json:"fullname" db:"fullname"`
}

type BorrowedBook struct {
	MemberID int `json:"memberid" db:"memberid"`
	BookID   int `json:"bookid" db:"bookid"`
}
