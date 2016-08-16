package bitbucket

import (
	"errors"
)

var (
	ErrNilClient = errors.New("client is nil")
)

// New creates an instance of the Bitbucket Client
func New(auth Auth) *Client {
	c := &Client{auth: auth}

	c.Keys = &KeyResource{c}
	c.Repos = &RepoResource{c}
	c.Users = &UserResource{c}
	c.Emails = &EmailResource{c}
	c.Brokers = &BrokerResource{c}
	c.Teams = &TeamResource{c}
	c.RepoKeys = &RepoKeyResource{c}
	c.Sources = &SourceResource{c}
	c.Groups = &GroupResource{c}
	return c
}

// Client the Bitbucket client
type Client struct {
	auth Auth

	Repos    *RepoResource
	Users    *UserResource
	Emails   *EmailResource
	Keys     *KeyResource
	Brokers  *BrokerResource
	Teams    *TeamResource
	Sources  *SourceResource
	RepoKeys *RepoKeyResource
	Groups   *GroupResource
}

// Guest Client that can be used to access
// public APIs that do not require authentication.
var Guest = New(&Anonymous{})
