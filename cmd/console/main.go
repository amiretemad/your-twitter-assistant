package main

import (
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"github.com/urfave/cli"
	"log"
	"os"
	"tweak_twitter/cmd/console/action"
	database "tweak_twitter/db"
	TwitterService "tweak_twitter/pkg/lib/twitter"
)

func main() {
	// Setup Database Connection
	db, err := database.DB(fmt.Sprintf("%s.db", os.Getenv("DATABASE_NAME")))
	if err != nil {
		log.Fatal("failed to connect database")
	}

	// Setup Twitter Credentials
	twitterCredentials := TwitterService.Credential{
		ConsumerKey:    os.Getenv("CONSUMER_KEY"),
		ConsumerSecret: os.Getenv("CONSUMER_SECRET"),
		Token:          os.Getenv("TOKEN"),
		TokenSecret:    os.Getenv("TOKEN_SECRET"),
	}

	app := cli.NewApp()
	app.Name = "Set of commands to have more control over your twitter account!"

	app.Commands = []cli.Command{
		{
			Name:     "fetch_followers",
			HelpName: "fetch_followers",
			Action: func(c *cli.Context) {
				err := action.FetchFollowers(c, db, twitterCredentials)
				if err != nil {
					return
				}
			},
			ArgsUsage:   ` `,
			Usage:       `Fetch followers for provided username.`,
			Description: `Fetch followers.`,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "username",
					Usage:    "Fetch followers for this username.",
					Required: true,
				},
			},
		},
		{
			Name:     "compare_followers",
			HelpName: "compare_followers",
			Action: func(c *cli.Context) {
				err := action.CompareFollower(c, db)
				if err != nil {
					return
				}
			},
			ArgsUsage:   ` `,
			Usage:       `Compare your followers with previous iterations`,
			Description: `this command returns list of users who started to follow/unfollow you.`,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "username",
					Usage:    "Fetch followers for this username.",
					Required: true,
				},
			},
		},
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
