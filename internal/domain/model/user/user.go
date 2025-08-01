package user

import "time"

type User struct {
	ID                  int
	TotalRequests       int64
	CreatedAt           time.Time
	GenerationsPerMonth int64
}

func New(id int) *User {
	return &User{
		ID:                  id,
		TotalRequests:       0,
		CreatedAt:           time.Now(),
		GenerationsPerMonth: 0,
	}
}

type Users []User
