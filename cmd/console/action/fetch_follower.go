package action

import (
	"errors"
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/urfave/cli"
	"gorm.io/gorm"
	"log"
	"strconv"
	TwitterService "tweak_twitter/pkg/lib/twitter"
	"tweak_twitter/pkg/model"
	"tweak_twitter/pkg/repository"
)

func FetchFollowers(c *cli.Context, db *gorm.DB, credential TwitterService.Credential) error {
	if c.String("username") == "" {
		return errors.New("please provide valid username")
	}

	username := c.String("username")

	service := TwitterService.OAuth1Service{Credential: credential}
	client := service.GetClient()

	var followerList []model.Follower

	trackerRepository := repository.TrackerRepository{DB: db}
	trackId, err := trackerRepository.AddBySourceAndValue(model.TrackerFetchFollower, username)
	if err != nil {
		return err
	}

	followerRepository := repository.FollowerRepository{DB: db}
	searchParams := &twitter.FollowerListParams{Cursor: -1, Count: 40}
	for {
		list, _, err := client.Followers.List(searchParams)

		if err != nil {
			log.Print(err)
			break
		}

		for _, user := range list.Users {
			u := model.Follower{
				TwitterId: strconv.FormatInt(user.ID, 10),
				TrackId:   trackId,
				Username:  user.ScreenName,
				Name:      user.Name,
				Location:  user.Location,
				Following: user.Following,
			}
			followerList = append(followerList, u)
		}

		if list.NextCursor == 0 {
			fmt.Println("End of the list!")
			break
		}

		searchParams.Cursor = list.NextCursor
	}

	if len(followerList) > 0 {
		_, err = followerRepository.SaveList(followerList)
		if err != nil {
			log.Fatal(err)
			return err
		}
	}

	return nil
}
