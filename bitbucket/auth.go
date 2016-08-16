package bitbucket

import (
	"net/http"

	"github.com/webdevwilson/go-bitbucket/oauth1"
)

// Auth represents an authentication strategy
type Auth interface {
	authenticate(r *http.Request) error
}

// OAuth for doing OAuth authentication
type OAuth struct {
	ConsumerKey    string
	ConsumerSecret string
	AccessToken    string
	TokenSecret    string
}

func (auth *OAuth) authenticate(r *http.Request) error {
	client := oauth1.Consumer{
		ConsumerKey:    auth.ConsumerKey,
		ConsumerSecret: auth.ConsumerSecret,
	}
	token := oauth1.NewAccessToken(auth.AccessToken, auth.TokenSecret, nil)

	// sign the request
	if err := client.Sign(r, token); err != nil {
		return err
	}

	return nil
}

// BasicAuth BASIC authentication strategy
type BasicAuth struct {
	Username string
	Password string
}

// authenticate adds BASIC Auth header
func (auth *BasicAuth) authenticate(r *http.Request) error {
	r.SetBasicAuth(auth.Username, auth.Password)
	return nil
}

// Anonymous requests do no authentication
type Anonymous struct{}

func (auth *Anonymous) authenticate(r *http.Request) error {
	return nil
}
