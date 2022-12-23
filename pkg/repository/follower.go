package repository

import (
	"gorm.io/gorm"
	"tweak_twitter/pkg/model"
)

type UserInterface interface {
	SaveList(follower []model.Follower) (bool, error)
}

type FollowerRepository struct {
	DB *gorm.DB
}

func (u *FollowerRepository) SaveList(follower []model.Follower) (bool, error) {
	u.DB.Create(follower)
	return true, nil
}

func (u *FollowerRepository) GetDiff(tracker model.Tracker) ([]model.Follower, []model.Follower) {
	var followers []model.Follower
	var unfollows []model.Follower

	u.DB.
		Raw("select * from followers where track_id = ? and twitter_id not in (select twitter_id from followers where track_id <> ?)", tracker.ID, tracker.ID).
		Scan(&followers)

	u.DB.
		Raw("select * from followers where track_id <> ? and twitter_id not in (select twitter_id from followers where track_id = ?)", tracker.ID, tracker.ID).
		Scan(&unfollows)

	return followers, unfollows
}
