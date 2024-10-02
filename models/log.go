package models

import "time"

type UserLog struct {
    UserID    uint      `json:"user_id"`
    Event     string    `json:"event"`
    Data      string    `json:"data"`
    CreatedAt time.Time `json:"created_at"`
}
