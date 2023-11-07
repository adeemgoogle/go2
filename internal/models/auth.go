package models

type Author struct {
	ID             int    `json:"id" db:"id"`
	FullName       string `json:"fullname" db:"full_name"`
	Pseudonym      string `json:"pseudonym" db:"pseudonym"`
	Specialization string `json:"specialization" db:"specialization"`
}
