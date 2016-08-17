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
}

// GroupDetails - contains the group information and the owner & members
type GroupDetails struct {
	Group
	Owner   User   `json:"owner"`
	Members []User `json:"members"`
}

// Create - Creates a group
func (gr *GroupResource) Create(owner string, name string) (g GroupDetails, err error) {
	path := fmt.Sprintf("/groups/%s/", owner)
	values := url.Values{}
	values.Add("name", name)
	err = gr.client.do("POST", path, nil, values, &g)
	return
}

// Update - Update group
func (gr *GroupResource) Update(owner string, name string) (g *GroupDetails, err error) {
	path := fmt.Sprintf("/groups/%s/%s", owner, name)
	err = gr.client.do("PUT", path, nil, nil, g)
	return
}

// Delete - Deletes a group from an account
func (gr *GroupResource) Delete(owner string, name string) error {
	path := fmt.Sprintf("/groups/%s/%s", owner, name)
	return gr.client.do("DELETE", path, nil, nil, nil)
}

// List - Lists groups by owner
func (gr *GroupResource) List(owner string) (g []*GroupDetails, err error) {
	path := fmt.Sprintf("/groups/%s/", owner)
	err = gr.client.do("GET", path, nil, nil, &g)
	return g, nil
}

// Find - Retrieve a group by owner and slug
func (gr *GroupResource) Find(owner string, slug string) (*GroupDetails, error) {
	filter := fmt.Sprintf("%s/%s", owner, slug)
	params := url.Values{
		"group": {filter},
	}
	var groups []GroupDetails
	err := gr.client.do("GET", "/groups", params, nil, &groups)
	if err != nil {
		return nil, err
	}

	if len(groups) > 1 {
		return nil, fmt.Errorf("Expected 1 group, found %d for %s", len(groups), filter)
	} else if len(groups) == 0 {
		return nil, nil
	} else {
		return &groups[0], nil
	}
}
