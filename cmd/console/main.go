package main

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	_ "github.com/joho/godotenv/autoload"
	"log"
	"os"
	"strconv"
	"tweak_twitter/pkg/model"
	"tweak_twitter/pkg/repository"
	TwitterService "tweak_twitter/pkg/service/twitter"
)

func main() {

	followerFile, err := os.Create("followers.json")
	if err != nil {
		log.Fatal(err)
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(followerFile)

	twitterCredentials := TwitterService.Credential{
		ConsumerKey:    os.Getenv("CONSUMER_KEY"),
		ConsumerSecret: os.Getenv("CONSUMER_SECRET"),
		Token:          os.Getenv("TOKEN"),
		TokenSecret:    os.Getenv("TOKEN_SECRET"),
	}

	service := TwitterService.OAuth1Service{Credential: twitterCredentials}
	client := service.GetClient()

	var userList []model.User
	userRepository := repository.UserFileRepository{File: followerFile}

	searchParams := &twitter.FollowerListParams{Cursor: -1, Count: 40}
	for {
		list, _, err := client.Followers.List(searchParams)

		if err != nil {
			if len(userList) > 0 {
				userRepository.SaveList(userList)
			}
			log.Fatal(err)
			return
		}

		for _, user := range list.Users {
			u := model.User{
				TwitterId: strconv.FormatInt(user.ID, 10),
				Username:  user.ScreenName,
				Name:      user.Name,
				Location:  user.Location,
				Following: user.Following,
			}
			userList = append(userList, u)
		}

		if list.NextCursor == 0 {
			fmt.Println("End of the list!")
			break
		}

		searchParams.Cursor = list.NextCursor
	}

	userRepository.SaveList(userList)
}
