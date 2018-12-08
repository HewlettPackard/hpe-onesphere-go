package onesphere

import "testing"

func TestGetZoneTypes(t *testing.T) {
	setup()

	_, err := client.GetZoneTypes()
	if err != nil {
		t.Error(err)
	}
}


func TestGetZoneTypeResourceProfiles(t *testing.T) {
	setup()

	_, err := client.GetZoneTypeResourceProfiles("2")
	if err != nil {
		t.Error(err)
	}
}

