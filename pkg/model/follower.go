package model

import "time"

type Follower struct {
	ID        uint      `json:"id"`
	TrackId   uint      `json:"track_id"`
	TwitterId string    `json:"twitterId"`
	Username  string    `json:"username"`
	Name      string    `json:"name"`
	Location  string    `json:"location"`
	Following bool      `json:"following"`
	CreatedAt time.Time `json:"created_at"`
}
