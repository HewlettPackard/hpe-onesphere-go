package onesphere

import "testing"

func TestGetUsers(t *testing.T) {
	setup()

	_, err := client.GetUsers("")
	if err != nil {
		t.Error(err)
	}
}
