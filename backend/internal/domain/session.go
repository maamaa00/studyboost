package domain

import "time"

type Session struct {
    ID        string
    Title     string
    StartedAt time.Time
    EndedAt   time.Time
    CreatedAt time.Time
}