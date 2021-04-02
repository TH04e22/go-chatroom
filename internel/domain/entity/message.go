package entity

import "time"

type Message struct {
	ID        int64
	UserID    int64
	UserName  string
	RoomID    int64
	RoomName  string
	Message   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
