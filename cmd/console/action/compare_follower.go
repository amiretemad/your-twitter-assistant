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
	follows, unfollows := followerRepository.GetDiff(username, lastTracker)

	fmt.Println(fmt.Sprintf("%d users has been unfollowed %s!", len(unfollows), username))
	fmt.Println(fmt.Sprintf("%d new users has been started to follow %s!", len(follows), username))

	return nil
}
