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
	Name                    string `json:"name"`
	Slug                    string `json:"slug"`
	AutoAdd                 bool   `json:"auto_add"`
	Permission              string `json:"permission"`
	ResourceURI             string `json:"resource_uri"`
	EmailForwardingDisabled bool   `json:"email_forwarding_disabled"`
}

// GroupDetails - contains the group information and the owner & members
type GroupDetails struct {
	Group
	Owner   User   `json:"owner"`
	Members []User `json:"members"`
}

// Create - Creates a group
func (gr *GroupResource) Create(owner string, name string) (g *GroupDetails, err error) {
	ownerOrCurrentUser(gr, &owner)

	path := fmt.Sprintf("/groups/%s/", owner)
	values := url.Values{}
	values.Set("name", name)
	err = gr.client.do("POST", path, nil, values, &g)

	return
}

// Update - Update group
func (gr *GroupResource) Update(group *GroupDetails) (g *GroupDetails, err error) {
	owner := group.Owner.Username
	ownerOrCurrentUser(gr, &owner)

	path := fmt.Sprintf("/groups/%s/%s", owner, group.Slug)
	// body := url.Values{}
	// body.Set("group", group)
	err = gr.client.do("PUT", path, nil, nil, &g)
	return
}

// Delete - Deletes a group from an account
func (gr *GroupResource) Delete(owner string, slug string) error {
	ownerOrCurrentUser(gr, &owner)

	path := fmt.Sprintf("/groups/%s/%s", owner, slug)
	return gr.client.do("DELETE", path, nil, nil, nil)
}

// List - Lists groups by owner
func (gr *GroupResource) List(owner string) (g []*GroupDetails, err error) {
	ownerOrCurrentUser(gr, &owner)

	path := fmt.Sprintf("/groups/%s/", owner)
	err = gr.client.do("GET", path, nil, nil, &g)
	return g, nil
}

// Find - Retrieve a group by owner and slug
func (gr *GroupResource) Find(owner string, slug string) (*GroupDetails, error) {
	ownerOrCurrentUser(gr, &owner)

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

// AddMember - Add a member to an existing group
func (gr *GroupResource) AddMember(owner string, group string, member string) (user *User, err error) {
	ownerOrCurrentUser(gr, &owner)

	path := fmt.Sprintf("/groups/%s/%s/members/%s", owner, group, member)
	err = gr.client.do("PUT", path, nil, nil, &user)
	return
}

// RemoveMember - Remove a member from an existing group
func (gr *GroupResource) RemoveMember(owner string, group string, member string) (user *User, err error) {
	ownerOrCurrentUser(gr, &owner)

	path := fmt.Sprintf("/groups/%s/%s/members/%s", owner, group, member)
	err = gr.client.do("DELETE", path, nil, nil, &user)
	return
}

// ownerOrCurrentUser sets the owner string to the current user when empty
func ownerOrCurrentUser(gr *GroupResource, owner *string) error {
	if owner == nil || (*owner) == "" {
		current, err := gr.client.Users.Current()
		if err == nil {
			return err
		}
		*owner = current.User.Username
	}
	return nil
}
