package bitbucket

import (
	"fmt"
	"net/url"
)

// GroupResource returns the schema for the group resource
type GroupResource struct {
	client *Client
}

// Group - represents a group json structure in bitbucket
type Group struct {
	AccountName string `json:"accountname"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	AutoAdd     bool   `json:"auto_add"`
	Permission  string `json:"permission"`
	Owner       UserResource
}

// Create - Creates a group
func (g *GroupResource) Create(owner string, name string) error {
	path := fmt.Sprintf("/groups/%s/", owner)
	values := url.Values{}
	values.Add("name", name)
	err := g.client.do("POST", path, nil, values, nil)
	return err
}

// Delete - Deletes a group from an account
func (g *GroupResource) Delete(owner string, name string) error {
	path := fmt.Sprintf("/groups/%s/%s", owner, name)
	err := g.client.do("DELETE", path, nil, nil, nil)
	return err
}

// List - Lists groups by owner
func (g *GroupResource) List(owner string) ([]*Group, error) {
	path := fmt.Sprintf("/groups/%s/", owner)
	groups := []*Group{}
	err := g.client.do("GET", path, nil, nil, &groups)
	if err != nil {
		return nil, err
	}
	return groups, nil
}

// Find - Retrieve a group by owner and slug
func (g *GroupResource) Find(owner string, slug string) (*Group, error) {
	filter := fmt.Sprintf("%s/%s", owner, slug)
	params := url.Values{
		"group": {filter},
	}
	var groups []Group
	err := g.client.do("GET", "/groups", params, nil, &groups)
	if err != nil {
		return nil, err
	}

	if len(groups) != 1 {
		return nil, fmt.Errorf("Expected 1 group, found %d for %s", len(groups), filter)
	}
	return &groups[0], nil
}
