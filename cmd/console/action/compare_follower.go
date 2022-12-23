package action

import (
	"errors"
	"fmt"
	"github.com/urfave/cli"
	"gorm.io/gorm"
	"tweak_twitter/pkg/model"
	"tweak_twitter/pkg/repository"
)

func CompareFollower(c *cli.Context, db *gorm.DB) error {
	if c.String("username") == "" {
		return errors.New("please provide valid username")
	}

	username := c.String("username")

	trackerRepository := repository.TrackerRepository{DB: db}
	lastTracker, err := trackerRepository.GetBySourceAndValue(model.TrackerFetchFollower, username)
	if err != nil {
		return err
	}

	if lastTracker.ID <= 0 {
		panic(fmt.Sprintf("There is not any record for %s", username))
	}

	followerRepository := repository.FollowerRepository{DB: db}
	follows, unfollows := followerRepository.GetDiff(lastTracker)

	for _, unfollow := range unfollows {
		fmt.Println(fmt.Sprintf("user (%s) unfollowed you!", unfollow.Username))
	}

	for _, follow := range follows {
		fmt.Println(fmt.Sprintf("user (%s) start to follow you!", follow.Username))
	}

	return nil
}
