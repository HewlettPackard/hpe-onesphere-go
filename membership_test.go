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

func TestGetMembershipByUser(t *testing.T) {
	setup()

	_, err := client.GetMembershipsByUser("2")
	if err != nil {
		t.Error(err)
	}
}

func TestGetMembershipsByUserGroup(t *testing.T) {
	setup()

	_, err := client.GetMembershipsByUserGroup("2")
	if err != nil {
		t.Error(err)
	}
}

func TestGetMembershipsByRole(t *testing.T) {
	setup()

	_, err := client.GetMembershipsByRole("2")
	if err != nil {
		t.Error(err)
	}
}

func TestGetMembershipByID(t *testing.T) {
	setup()

	_, err := client.GetMembershipByID("2")
	if err != nil {
		t.Error(err)
	}
}

func TestCreateMembership(t *testing.T) {
	setup()

	membershipRequest := MembershipRequest{}

	if _, err := client.CreateMembership(membershipRequest); err != nil {
		t.Error(err)
	}
}

func TestDeleteMembership(t *testing.T) {
	setup()

	membership := Membership{
		GroupURI:          "2",
		ProjectURI:        "2",
		UserURI:           "2",
		MembershipRoleURI: "2",
	}

	if err := client.DeleteMembership(membership); err != nil {
		t.Skip(err)
	}
}
