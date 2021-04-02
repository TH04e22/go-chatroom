package entity

import "time"

type User struct {
	ID        int64
	Name      string
	Password  string
	CreateAt  time.Time
	UpdatedAt time.Time
}
