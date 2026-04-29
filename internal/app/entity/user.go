package entity

import "time"

type User struct {
	ID        int64
	FirstName string
	LastName  string
	Email     string
	Phone     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserFilter struct {
	ID *int64
}
