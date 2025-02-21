package models

import "time"

type Task struct {
	ID          string
	Title       string
	Description string
	CreationAt  time.Time
}
