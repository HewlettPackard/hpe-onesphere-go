package onesphere

import "testing"

func TestGetMembershipRoles(t *testing.T) {
	setup()

	_, err := client.GetMembershipRoles()
	if err != nil {
		t.Error(err)
	}
}

func TestGetMembershipRoleByName(t *testing.T) {
	setup()

	_, err := client.GetMembershipRoleByName("2")
	if err != nil {
		t.Error(err)
	}
}
