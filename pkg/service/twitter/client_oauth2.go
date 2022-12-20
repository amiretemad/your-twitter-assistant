package twitter

import (
	"github.com/dghubble/go-twitter/twitter"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

type OAuth2Service struct {
	Credential
}

func (r *OAuth2Service) GetClient() *twitter.Client {

	config := &clientcredentials.Config{
		ClientID:     r.ConsumerKey,
		ClientSecret: r.ConsumerSecret,
		TokenURL:     "https://api.twitter.com/oauth2/token",
	}

	httpClient := config.Client(oauth2.NoContext)

	return twitter.NewClient(httpClient)
}
