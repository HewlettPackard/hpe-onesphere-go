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

func TestGetUserByName(t *testing.T) {
	setup()

	_, err := client.GetUserByName("2")
	if err != nil {
		t.Error(err)
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

	updates := UserRequest{}

	if _, err := client.UpdateUser("2", updates); err != nil {
		t.Error(err)
	}
}

func TestDeleteUser(t *testing.T) {
	setup()

	if err := client.DeleteUser("2"); err != nil {
		t.Skip(err)
	}
}
