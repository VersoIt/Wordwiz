package model

import "time"

type User struct {
	ID                  int
	TotalRequests       int64
	CreatedAt           time.Time
	GenerationsPerMonth int64
}

type Users []User
