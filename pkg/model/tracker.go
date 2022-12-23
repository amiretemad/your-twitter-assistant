package model

import (
	"time"
)

const (
	TrackerFetchFollower = "fetch_follower"
)

type Tracker struct {
	ID          uint      `json:"id"`
	Source      string    `json:"source"`
	SourceValue string    `json:"source_value"`
	CreatedAt   time.Time `json:"created_at"`
}
