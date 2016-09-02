package bitbucket

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupsUnmarshal(t *testing.T) {
	jsonString, err := ioutil.ReadFile("groups_test_GroupDetails.json")
	assert.NoError(t, err)
	group := &GroupDetails{}
	err = json.Unmarshal(jsonString, group)
	assert.NoError(t, err)
	assert.Equal(t, "Rebel Alliance", group.Name)
	assert.NotNil(t, group.Owner)
	assert.Equal(t, "rebel_alliance", group.Slug)
	assert.Equal(t, false, group.AutoAdd)
	assert.Equal(t, false, group.EmailForwardingDisabled)
}

func TestGroupsCreate(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping in short mode")
	}
	testGroup := "TestGroupsCreate"
	group, err := client.Groups.Create(testUser, testGroup)
	defer client.Groups.Delete(testUser, group.Slug)

	// create
	assert.NoError(t, err)
	assert.NotNil(t, group)
	assert.NotNil(t, group.Owner)
	assert.Equal(t, "testgroupscreate", group.Slug)
	assert.Equal(t, false, group.EmailForwardingDisabled)
	assert.Equal(t, false, group.AutoAdd)
	assert.Equal(t, testUser, group.Owner.Username)
	assert.Equal(t, testGroup, group.Name)
}

func TestGroupsDelete(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping in short mode")
	}
	testGroup := "TestGroupsDelete"
	group, err := client.Groups.Create(testUser, testGroup)
	defer client.Groups.Delete(testUser, group.Slug)

	assert.NoError(t, err)
	assert.NotNil(t, group)

	err = client.Groups.Delete(testUser, group.Slug)
	assert.NoError(t, err)
}

func TestGroupsList(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping in short mode")
	}
	testGroup := "TestGroupsList"
	group, err := client.Groups.Create(testUser, testGroup)
	defer client.Groups.Delete(testUser, group.Slug)

	groups, err := client.Groups.List(testUser)
	if err != nil {
		t.Error(err)
		return
	}
	assert.NotEmpty(t, groups, "no groups found")
}

func TestGroupsAddMembers(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping in short mode")
	}
	testGroup := "TestGroupsAddMember"
	group, err := client.Groups.Create(testUser, testGroup)
	defer client.Groups.Delete(testUser, group.Slug)

	_, err = client.Groups.AddMember(testUser, group.Slug, testUser)
	assert.NoError(t, err)

	members, err := client.Groups.GetMembers(testUser, group.Slug)
	assert.NoError(t, err)
	assert.NotNil(t, members)
	assert.True(t, len(*members) == 1)
	assert.Equal(t, testUser, (*members)[0].Username)
}

func TestGroupsRemoveMember(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping in short mode")
	}
	testGroup := "TestGroupsRemoveMember"
	created, err := client.Groups.Create(testUser, testGroup)
	defer client.Groups.Delete(testUser, created.Slug)

	assert.NoError(t, err)
	assert.NotNil(t, created)

	_, err = client.Groups.AddMember(testUser, created.Slug, testUser)
	assert.NoError(t, err)

	members, err := client.Groups.GetMembers(testUser, created.Slug)
	assert.Equal(t, 1, len(*members))

	err = client.Groups.RemoveMember(testUser, created.Slug, testUser)
	assert.NoError(t, err)

	members, err = client.Groups.GetMembers(testUser, created.Slug)
	assert.NoError(t, err)
	assert.Equal(t, 0, len(*members))
}

func TestGroupsGet(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping in short mode")
	}
	testGroup := "TestGroupsFind"
	created, err := client.Groups.Create(testUser, testGroup)
	defer client.Groups.Delete(testUser, created.Slug)

	group, err := client.Groups.Get(testUser, created.Slug)
	assert.NotNil(t, group)
	assert.NoError(t, err)
	assert.Equal(t, testGroup, group.Name)
	assert.Equal(t, testUser, created.Owner.Username)
}
