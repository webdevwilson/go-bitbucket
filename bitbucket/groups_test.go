package bitbucket

import (
	"os"
	"testing"
)

var testGroup = "bitbucket_go_test"

func TestMain(m *testing.M) {
	setup()
	result := m.Run()
	teardown()
	os.Exit(result)
}

func setup() {
	client.Groups.Create(testUser, testGroup)
}

func teardown() {
	client.Groups.Delete(testUser, testGroup)
}

func Test_List(t *testing.T) {
	groups, err := client.Groups.List(testUser)
	if err != nil {
		t.Error(err)
	}

	if len(groups) == 0 {
		t.Error("no groups found")
	}

	group := groups[0]

	// find group
	found, err := client.Groups.Find(testUser, group.Slug)
	if err != nil {
		t.Error(err)
	}

	if found.Slug != group.Slug {
		t.Errorf("Error retrieving group %s/%s, got %s/%s", testUser, group.Slug, found.AccountName, found.Slug)
	}
}

func Test_GroupFind(t *testing.T) {
	group, err := client.Groups.Find(testUser, testGroup)
	if err != nil {
		t.Error(err)
	}

	if group.Name != testGroup {
		t.Errorf("Error retrieving group %s/%s, got %s/%s", testUser, testGroup, group.AccountName, group.Slug)
	}
}
