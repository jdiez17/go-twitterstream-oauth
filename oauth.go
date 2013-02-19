package twitterstream

import (
	"github.com/mrjones/oauth"
)

const (
	k_auth_url         = "https://api.twitter.com/oauth/authorize"
	k_token_url        = "https://api.twitter.com/oauth/request_token"
	k_access_token_url = "https://api.twitter.com/oauth/access_token"
)

type Oauth struct {
	Consumer *oauth.Consumer
}

type AuthenticationRequest struct {
	RequestToken *oauth.RequestToken
	Url          string
}

func NewOauth(ConsumerKey, ConsumerSecret string) Oauth {
	c := oauth.NewConsumer(
		ConsumerKey,
		ConsumerSecret,
		oauth.ServiceProvider{
			RequestTokenUrl:   k_token_url,
			AuthorizeTokenUrl: k_auth_url,
			AccessTokenUrl:    k_access_token_url,
		})

	return Oauth{c}
}

func (o Oauth) NewAuthenticationRequest() (*AuthenticationRequest, error) {
	requestToken, url, err := o.Consumer.GetRequestTokenAndUrl("oob")
	if err != nil {
		return nil, err
	}

	return &AuthenticationRequest{requestToken, url}, nil
}

func (o Oauth) GetAccessToken(RequestToken *oauth.RequestToken, code string) (*oauth.AccessToken, error) {
	accessToken, err := o.Consumer.AuthorizeToken(RequestToken, code)
	return accessToken, err
}
