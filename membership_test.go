package onesphere

import "testing"

func TestGetMemberships(t *testing.T) {
	setup()

	_, err := client.GetMemberships("")
	if err != nil {
		t.Error(err)
	}
}

func TestGetMembershipsByProject(t *testing.T) {
	setup()

	_, err := client.GetMembershipsByProject("2")
	if err != nil {
		t.Error(err)
	}
}
