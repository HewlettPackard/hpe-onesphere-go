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

func TestCreateUser(t *testing.T) {
	setup()

	userRequest := UserRequest{}

	if _, err := client.CreateUser(userRequest); err != nil {
		t.Error(err)
	}
}

func TestUpdateUser(t *testing.T) {
	setup()

	user := User{ID: "2"}
	updates := UserRequest{}

	if _, err := client.UpdateUser(user, updates); err != nil {
		t.Error(err)
	}
}

