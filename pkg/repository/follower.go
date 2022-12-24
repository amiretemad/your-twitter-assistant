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

func (u *FollowerRepository) GetDiff(username string, tracker model.Tracker) ([]model.Follower, []model.Follower) {
	var followers []model.Follower
	var unfollows []model.Follower

	u.DB.
		Raw("SELECT * FROM followers f INNER JOIN trackers t ON f.track_id = t.id and t.source_value = ? and t.source = ? where f.track_id = ? and f.twitter_id not in (SELECT twitter_id FROM followers where track_id IN (select id from trackers where source_value = ? and source = ? and id < ?))",
			username, model.TrackerFetchFollower, tracker.ID, username, model.TrackerFetchFollower, tracker.ID).
		Scan(&followers)

	u.DB.
		Raw("SELECT * FROM followers f INNER JOIN trackers t ON f.track_id = t.id and t.source_value = ? and t.source = ? where f.track_id < ? and f.twitter_id not in (SELECT twitter_id FROM followers where track_id IN (select id from trackers where source_value = ? and source = ? and id = ?))",
			username, model.TrackerFetchFollower, tracker.ID, username, model.TrackerFetchFollower, tracker.ID).
		Scan(&unfollows)

	return followers, unfollows
}
