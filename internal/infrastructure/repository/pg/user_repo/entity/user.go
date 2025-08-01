package entity

import (
	"time"
	"wordwiz/internal/domain/model/user"
)

type User struct {
	ID                  int       `db:"id"`
	TotalRequests       int64     `db:"total_requests"`
	CreatedAt           time.Time `db:"created_at"`
	GenerationsPerMonth int64     `db:"generations_per_month"`
}

func (u *User) FromDomain(user user.User) {
	u.ID = user.ID
	u.TotalRequests = user.TotalRequests
	u.CreatedAt = user.CreatedAt
	u.GenerationsPerMonth = user.GenerationsPerMonth
}

func (u *User) ToDomain() user.User {
	return user.User{
		ID:                  u.ID,
		TotalRequests:       u.TotalRequests,
		CreatedAt:           u.CreatedAt,
		GenerationsPerMonth: u.GenerationsPerMonth,
	}
}

type Users []User
