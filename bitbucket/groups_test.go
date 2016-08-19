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
	testGroup := "Test_Create"
	group, err := client.Groups.Create(testUser, testGroup)
	defer client.Groups.Delete(testUser, group.Slug)

	// create
	assert.NoError(t, err)
	assert.NotNil(t, group)
	assert.NotNil(t, group.Owner)
	assert.Equal(t, "test_create", group.Slug)
	assert.Equal(t, false, group.EmailForwardingDisabled)
	assert.Equal(t, false, group.AutoAdd)
	assert.Equal(t, testUser, group.Owner.Username)
	assert.Equal(t, testGroup, group.Name)
}

func TestGroupsDelete(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping in short mode")
	}
	testGroup := "Test_Delete"
	group, err := client.Groups.Create(testUser, testGroup)
	err = client.Groups.Delete(testUser, group.Slug)
	assert.NoError(t, err)
}

//
// func Test_AddMember(t *testing.T) {
// 	testGroup := "Test_AddMember"
// 	group, err := client.Groups.Create(testUser, testGroup)
// 	defer client.Groups.Delete(testUser, testGroup)
//
// 	_, err = client.Groups.AddMember(testUser, group.Slug, testUser)
// 	assert.NoError(t, err)
//
// 	_, err = client.Groups.Find(testUser, group.Slug)
// 	assert.NoError(t, err)
// }
//
// func Test_List(t *testing.T) {
// 	testGroup := "Test_List"
// 	_, err := client.Groups.Create(testUser, testGroup)
// 	defer client.Groups.Delete(testUser, testGroup)
//
// 	groups, err := client.Groups.List(testUser)
// 	if err != nil {
// 		t.Error(err)
// 		return
// 	}
// 	assert.NotEmpty(t, groups, "no groups found")
// }
//
func TestGroupsFind(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping in short mode")
	}
	testGroup := "Test_Find"
	group, err := client.Groups.Create(testUser, testGroup)
	defer client.Groups.Delete(testUser, group.Slug)

	group, err = client.Groups.Find(testUser, group.Slug)
	assert.NotNil(t, group)
	assert.NoError(t, err)
	assert.Equal(t, testGroup, group.Name)
	assert.Equal(t, testUser, group.Owner.Username)
}
