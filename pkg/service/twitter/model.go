package twitter

import "github.com/dghubble/go-twitter/twitter"

type GetClient interface {
	GetClient() *twitter.Client
}

type Credential struct {
	ConsumerKey    string
	ConsumerSecret string
	Token          string
	TokenSecret    string
}
