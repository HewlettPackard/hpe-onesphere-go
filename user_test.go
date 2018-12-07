package onesphere

import "testing"

func TestGetUsers(t *testing.T) {
	setup()

	_, err := client.GetUsers("")
	if err != nil {
		t.Error(err)
	}
}

func TestGetUserByID(t *testing.T) {
	setup()

	userList, err := client.GetUsers("")
	if err != nil {
		t.Error(err)
		return
	}
	if len(userList.Members) == 0 {
		t.Error("TestGetUserByID Could not find any Users")
		return
	}

	testId := userList.Members[0].ID
	user, err := client.GetUserByID(testId)
	if err != nil {
		t.Error(err)
	}
	if user.ID == "" {
		t.Errorf("TestGetUserByID Failed to get user: ID is ''")
	}
}
