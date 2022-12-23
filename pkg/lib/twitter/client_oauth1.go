package twitter

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

type OAuth1Service struct {
	Credential
}

func (r *OAuth1Service) GetClient() *twitter.Client {
	config := oauth1.NewConfig(r.ConsumerKey, r.ConsumerSecret)
	token := oauth1.NewToken(r.Token, r.TokenSecret)

	httpClient := config.Client(oauth1.NoContext, token)

	return twitter.NewClient(httpClient)
}
