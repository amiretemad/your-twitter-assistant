package model

type User struct {
	ID        int64  `json:"id"`
	TwitterId string `json:"twitterId"`
	Username  string `json:"username"`
	Name      string `json:"name"`
	Location  string `json:"location"`
	Follower  bool   `json:"follower"`
	Following bool   `json:"following"`
}
