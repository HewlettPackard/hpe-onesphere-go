package onesphere

import "testing"

func TestGetZoneTypes(t *testing.T) {
	setup()

	_, err := client.GetZoneTypes()
	if err != nil {
		t.Error(err)
	}
}

